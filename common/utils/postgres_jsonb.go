package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

// 构建jsonb缺失的中间路径(不会覆盖原有路径)，注意传入的dbWithModelAndWhere 必须是db.Model(表结构).WithContext(ctx).Where(条件)这样的
func BuildPostgresJsonbMissObject(dbWithModelAndWhere *gorm.DB, dbColumn string, objectPath []string) error {
	lenPath := len(objectPath)
	//只有一层就不需要
	if lenPath <= 1 {
		return nil
	}
	for i := 0; i < lenPath-1; i++ {
		//截取的路径
		cutPath := objectPath[:i+1]
		checkPath, err := joinPostgresJsonbPathChain(dbColumn, cutPath)
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
		missObj := buildAllMissPath(objectPath, i+1)
		path := JoinPostgresJsonbPath(cutPath)
		//创建空对象
		err = dbWithModelAndWhere.Update(dbColumn, gorm.Expr("JSONB_SET("+dbColumn+",?,?,?)", path, missObj, true)).Error
		if err != nil {
			return errors.New(fmt.Sprintf("创建中间路径失败，path=%s,err=%s", checkPath, err.Error()))
		}
		//结束循环
		break
	}
	return nil
}

// jsonb路径{a,b,c}
func JoinPostgresJsonbPath(objectPath []string) string {
	return "{" + strings.Join(objectPath, ",") + "}"
}

// 连接jsonb路径链路 targetColumn->a->>b
func joinPostgresJsonbPathChain(dbColumn string, objectPath []string) (string, error) {
	lenPath := len(objectPath)
	if lenPath == 0 {
		return "", errors.New("jsonb path is empty")
	}
	if lenPath == 1 {
		return dbColumn + "->>'" + objectPath[0] + "'", nil
	}
	pathBuilder := strings.Builder{}
	pathBuilder.WriteString(dbColumn)
	for k, v := range objectPath {
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
func buildAllMissPath(objectPath []string, beginIndex int) string {
	if beginIndex >= len(objectPath)-1 {
		return "{}"
	}
	newObj := strings.Builder{}
	newObj.WriteString("{")
	newObj.WriteString("\"")
	newObj.WriteString(objectPath[beginIndex])
	newObj.WriteString("\":")
	newObj.WriteString(buildAllMissPath(objectPath, beginIndex+1))
	newObj.WriteString("}")
	return newObj.String()
}
