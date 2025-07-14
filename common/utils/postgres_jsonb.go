package utils

import (
	"errors"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strings"
)

// 构建jsonb缺失的中间路径(不会覆盖原有路径)，dbWithCtxModelWhere传类似org.WithContext(ctx).Model(&entity.User).Where(条件)
func SetPgJsonbValue(dbWithCtxModelWhere *gorm.DB, dbColumn string, objectPath []string, value any) error {
	lenPath := len(objectPath)
	//空就不需要处理
	if lenPath == 0 {
		return nil
	}
	//只有一层直接赋值
	if lenPath == 1 {
		err := dbWithCtxModelWhere.Update(dbColumn, datatypes.JSONSet(dbColumn).Set(joinPostgresJsonbPath(objectPath), value)).Error
		if err != nil {
			return errors.New(fmt.Sprintf("failed to set jsonb value，path:{%s},err=%s", objectPath[0], err.Error()))
		}
		return nil
	}
	//检查前面的路径是否对象
	selectStr := buildGetJsonbPathTypeSelectStr(dbColumn, objectPath)
	var checkResult string
	if err := dbWithCtxModelWhere.Select(selectStr + " as result").Scan(&checkResult).Error; err != nil {
		return err
	}
	pathAttrTypes := strings.Split(checkResult, ",")
	for index, attrType := range pathAttrTypes {
		//最后一个，直接赋值
		if index == lenPath-1 {
			missObjVal := buildMissPathVal(objectPath, lenPath, value)
			missPath := joinPostgresJsonbPath(objectPath)
			err := dbWithCtxModelWhere.Update(dbColumn, datatypes.JSONSet(dbColumn).Set(missPath, missObjVal)).Error
			if err != nil {
				return errors.New(fmt.Sprintf("set jsonb err，path:{%s},err=%s", missPath, err.Error()))
			}
			return nil
		}
		cutPath := objectPath[:index+1]
		tipPath := strings.Join(cutPath, ".")
		switch attrType {
		case "string": //字符串
			return errors.New(fmt.Sprintf("the attribute type of path:{%s} is not an object but %s}", tipPath, attrType))
		case "number": //数值
			return errors.New(fmt.Sprintf("the attribute type of path:{%s} is not an object but %s}", tipPath, attrType))
		case "boolean": //布尔
			return errors.New(fmt.Sprintf("the attribute type of path:{%s} is not an object but %s}", tipPath, attrType))
		case "object": //对象
			continue
		case "null": //不存在字段
			missObjVal := buildMissPathVal(objectPath, index+1, value)
			missPath := joinPostgresJsonbPath(cutPath)
			//设置
			err := dbWithCtxModelWhere.Update(dbColumn, datatypes.JSONSet(dbColumn).Set(missPath, missObjVal)).Error
			if err != nil {
				return errors.New(fmt.Sprintf("set jsonb err，path:{%s},err=%s", missPath, err.Error()))
			}
			return nil
		case "array": //数组
			return errors.New(fmt.Sprintf("the attribute type of path:{%s} is not an object but %s}", tipPath, attrType))
		default:
			return nil
		}
	}
	return nil
}

// jsonb路径{a,b,c}
func joinPostgresJsonbPath(objectPath []string) string {
	return "{" + strings.Join(objectPath, ",") + "}"
}

// 连接jsonb路径链路 targetColumn->a->b
func joinPostgresJsonbPathChain(dbColumn string, objectPath []string) string {
	lenPath := len(objectPath)
	if lenPath == 0 {
		return ""
	}
	if lenPath == 1 {
		return dbColumn + "->'" + objectPath[0] + "'"
	}
	pathBuilder := strings.Builder{}
	pathBuilder.WriteString(dbColumn)
	for _, v := range objectPath {
		pathBuilder.WriteString("->")
		pathBuilder.WriteString("'")
		pathBuilder.WriteString(v)
		pathBuilder.WriteString("'")
	}
	return pathBuilder.String()
}

// 构建自缺失的中间对象
func buildMissPathVal(objectPath []string, beginIndex int, value any) any {
	if beginIndex >= len(objectPath) {
		return value
	}
	return map[string]any{
		objectPath[beginIndex]: buildMissPathVal(objectPath, beginIndex+1, value),
	}
}

// 构建检查路径用的select的值
func buildGetJsonbPathTypeSelectStr(dbColumn string, objectPath []string) string {
	lenPath := len(objectPath)
	//检查前面的路径是否对象
	selectStr := strings.Builder{}
	for i := 0; i < lenPath; i++ {
		selectStr.WriteString("COALESCE(jsonb_typeof(")
		selectStr.WriteString(joinPostgresJsonbPathChain(dbColumn, objectPath[:i+1]))
		selectStr.WriteString(")::TEXT,'null')")
		if i < lenPath-1 {
			selectStr.WriteString(" || ',' || ")
		}
	}
	return selectStr.String()
}
