package utils

import (
	"errors"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strings"
)

// 设置postgres jsonb属性
func SetPgJsonbValue(db *gorm.DB, model any, rowId int64, dbColumn string, objectPath []string, value any) error {
	lenPath := len(objectPath)
	//空就不需要处理
	if lenPath == 0 {
		return nil
	}
	//只有一层直接赋值
	if lenPath == 1 {
		missPath := joinPostgresJsonbPath(objectPath)
		err := setValue(db, model, rowId, dbColumn, missPath, value)
		if err != nil {
			return errors.New(fmt.Sprintf("set jsonb err，path:{%s},err=%s", missPath, err.Error()))
		}
		return nil
	}
	//检查前面的路径是否对象
	selectStr := buildGetJsonbPathTypeSelectStr(dbColumn, objectPath)
	var checkResult string
	if err := db.Model(model).Where("id = ?", rowId).Select(selectStr + " as result").Scan(&checkResult).Error; err != nil {
		return err
	}
	pathAttrTypes := strings.Split(checkResult, ",")
	for index, attrType := range pathAttrTypes {
		//最后一个，直接赋值
		if index == lenPath-1 {
			missObjVal := buildMissPathVal(objectPath, lenPath, value)
			missPath := joinPostgresJsonbPath(objectPath)
			err := setValue(db, model, rowId, dbColumn, missPath, missObjVal)
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
			err := setValue(db, model, rowId, dbColumn, missPath, missObjVal)
			if err != nil {
				return errors.New(fmt.Sprintf("set jsonb err，path:{%s},err=%s", missPath, err.Error()))
			}
			return nil
		case "array": //数组
			return errors.New(fmt.Sprintf("the attribute type of path:{%s} is not an object but %s}", tipPath, attrType))
		}
	}
	return nil
}

// 设置
func setValue(db *gorm.DB, model any, rowId int64, dbColumn string, missPath string, missObjVal any) error {
	return db.Model(model).Where("id = ?", rowId).Update(dbColumn, datatypes.JSONSet(dbColumn).Set(missPath, missObjVal)).Error
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
