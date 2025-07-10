package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/datatypes"
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
	for i := lenPath - 2; i >= 0; i-- {
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
		var (
			missObjVal map[string]any
			missPath   string
		)
		//存在字段
		if checkResult.Valid {
			//值为空字符串
			if checkResult.String == "" {
				return errors.New("the path value is not an object,path=" + checkPath)
			}
			//值不为空且非对象
			if string(checkResult.String[0]) != "{" || string(checkResult.String[len(checkResult.String)-1]) != "}" {
				return errors.New("the path value is not an object，path=" + checkPath)
			}
			//是对象
			//倒数第二路径直接是对象就不用继续探查
			if i == lenPath-2 {
				return nil
			}
			//否则创建所有中间对象（除最后节点）
			missObjVal = buildMissPathVal(objectPath, i+2)
			missPath = JoinPostgresJsonbPath(objectPath[:i+2])
		} else { //不存在字段
			//如果不是到根部都不存在则跳过
			if i != 0 {
				continue
			}
			//到了根部一次性创建所有中间对象（除最后节点）
			missObjVal = buildMissPathVal(objectPath, i+1)
			missPath = JoinPostgresJsonbPath(cutPath)
		}
		//创建对象
		err = dbWithModelAndWhere.Update(dbColumn, datatypes.JSONSet(dbColumn).Set(missPath, missObjVal)).Error
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

// 构建自缺失的中间对象
func buildMissPathVal(objectPath []string, beginIndex int) map[string]any {
	if beginIndex >= len(objectPath)-1 {
		return map[string]any{}
	}
	return map[string]any{
		objectPath[beginIndex]: buildMissPathVal(objectPath, beginIndex+1),
	}
}
