package data

import (
	"chat_with_me/common/model/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strings"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Create(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}
func (u *UserRepo) Update(ctx context.Context, id int64, user map[string]any) error {
	return u.db.Model(&entity.User{}).WithContext(ctx).Where("id = ?", id).Updates(user).Error
}
func (u *UserRepo) UpdateAttr(ctx context.Context, id int64, column string, pathArr []string, val any) error {
	return u.db.Model(&entity.User{}).WithContext(ctx).Where("id = ?", id).Transaction(func(db *gorm.DB) error {
		if err := BuildPostgresJsonbMissObject(db, column, pathArr); err != nil {
			return err
		}
		path := JoinPostgresJsonbPath(pathArr)
		return db.Update("attr", datatypes.JSONSet(column).Set(path, val)).Error
	})
}

// 构建jsonb缺失的中间路径(不会覆盖原有路径)，注意传入的dbWithModelAndWhere 必须是db.Model(表结构).WithContext(ctx).Where(条件)这样的
func BuildPostgresJsonbMissObject(dbWithModelAndWhere *gorm.DB, targetColumn string, pathArr []string) error {
	lenPath := len(pathArr)
	//只有一层就不需要
	if lenPath <= 1 {
		return nil
	}
	for i := 0; i < lenPath-1; i++ {
		//截取的路径
		cutPath := pathArr[:i+1]
		checkPath, err := joinPostgresJsonbPathChain(targetColumn, cutPath)
		if err != nil {
			return err
		}
		var checkResult sql.NullString
		err = dbWithModelAndWhere.Select(checkPath + " as val").Scan(&checkResult).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("record not found")
			}
			return errors.New("failed to find record,err=" + err.Error())
		}
		//存在字段
		if checkResult.Valid {
			//值为空字符串
			if checkResult.String == "" {
				return errors.New("the path value is not an object,path=" + checkPath)
			}
			//值不为空且非对象
			if string(checkResult.String[0]) != "{" || string(checkResult.String[len(checkResult.String)-1]) != "}" {
				return errors.New("the path is not an object，path=" + checkPath)
			}
			//是对象就跳过下一层
			continue
		}
		//截止到这里断了，一次性创建除最后一个节点的对象
		missObj := buildAllMissPath(pathArr, i+1)
		path := JoinPostgresJsonbPath(cutPath)
		//创建空对象
		err = dbWithModelAndWhere.Update(targetColumn, gorm.Expr("JSONB_SET("+targetColumn+",?,?,?)", path, missObj, true)).Error
		if err != nil {
			return errors.New(fmt.Sprintf("创建中间路径失败，path=%s,err=%s", checkPath, err.Error()))
		}
		//结束循环
		break
	}
	return nil
}

// jsonb路径{a,b,c}
func JoinPostgresJsonbPath(pathArr []string) string {
	return "{" + strings.Join(pathArr, ",") + "}"
}

// 连接jsonb路径链路 targetColumn->a->>b
func joinPostgresJsonbPathChain(column string, pathArr []string) (string, error) {
	lenPath := len(pathArr)
	if lenPath == 0 {
		return "", errors.New("jsonb path is empty")
	}
	if lenPath == 1 {
		return column + "->>'" + pathArr[0] + "'", nil
	}
	pathBuilder := strings.Builder{}
	pathBuilder.WriteString(column)
	for k, v := range pathArr {
		if k < lenPath-1 {
			pathBuilder.WriteString("->")
		} else {
			pathBuilder.WriteString("->>")
		}
		pathBuilder.WriteString("'")
		pathBuilder.WriteString(v)
		pathBuilder.WriteString("'")
	}
	return pathBuilder.String(), nil
}

// 构建自缺失的所有路径
func buildAllMissPath(pathArr []string, beginIndex int) string {
	if beginIndex >= len(pathArr)-1 {
		return "{}"
	}
	newObj := strings.Builder{}
	newObj.WriteString("{")
	newObj.WriteString("\"")
	newObj.WriteString(pathArr[beginIndex])
	newObj.WriteString("\":")
	newObj.WriteString(buildAllMissPath(pathArr, beginIndex+1))
	newObj.WriteString("}")
	return newObj.String()
}
