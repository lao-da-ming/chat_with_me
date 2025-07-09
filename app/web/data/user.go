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
	//检查路径并构建完整路径
	dbWithModelAndWhere := u.db.Model(&entity.User{}).WithContext(ctx).Where("id = ?", id)
	if err := BuildJsonbMissPath(ctx, dbWithModelAndWhere, column, pathArr); err != nil {
		return err
	}
	path := JoinJsonbPathToObj(pathArr)
	return u.db.Model(&entity.User{}).WithContext(ctx).Where("id = ?", id).Update("attr", datatypes.JSONSet(column).Set(path, val)).Error
}

func BuildJsonbMissPath(ctx context.Context, dbWithModelAndWhere *gorm.DB, targetColumn string, pathArr []string) error {
	lenPath := len(pathArr)
	//只有一层就不需要
	if lenPath <= 1 {
		return nil
	}
	return dbWithModelAndWhere.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < lenPath-1; i++ {
			//截取的路径
			cutPath := pathArr[:i+1]
			checkPath, err := JoinJsonbPath(targetColumn, cutPath)
			if err != nil {
				return err
			}
			var checkResult sql.NullString
			err = tx.Select(checkPath + " as val").Scan(&checkResult).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return errors.New("record not found")
				}
				return errors.New("failed to find record,err=" + err.Error())
			}
			if checkResult.Valid {
				switch checkResult.String {
				case "": //空字符串，非对象
					return errors.New("the path value is not an object but string,path=" + checkPath)
				default:
					//非对象
					if string(checkResult.String[0]) != "{" || string(checkResult.String[len(checkResult.String)-1]) != "}" {
						return errors.New("the path is not an object，path=" + checkPath)
					}
					//是对象就跳过下一层
					continue
				}
			}
			path := JoinJsonbPathToObj(cutPath)
			//创建空对象
			err = tx.WithContext(ctx).Update(targetColumn, gorm.Expr("JSONB_SET("+targetColumn+",?,?,?)", path, "{}", true)).Error
			if err != nil {
				return errors.New(fmt.Sprintf("创建中间路径失败，path=%s,err=%s", checkPath, err.Error()))
			}
		}
		return nil
	})
}

// jsonb路径连接成{a,b,c}
func JoinJsonbPathToObj(pathArr []string) string {
	return "{" + strings.Join(pathArr, ",") + "}"
}

// 连接jsonb路径 targetColumn->a->>b
func JoinJsonbPath(column string, pathArr []string) (string, error) {
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
