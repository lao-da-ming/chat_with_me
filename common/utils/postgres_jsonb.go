package utils

import (
	"errors"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strings"
)

// 构建jsonb缺失的中间路径(不会覆盖原有路径)，注意传入的dbWithModelAndWhere 必须是db.Model(表结构).WithContext(ctx).Where(条件)这样的
func BuildPostgresJsonbMissObject(dbWithModelAndWhere *gorm.DB, dbColumn string, objectPath []string) error {
	lenPath := len(objectPath)
	//只有一层或者空就不需要处理
	if lenPath <= 1 {
		return nil
	}
	//检查前面的路径是否对象（除路径最后一个）
	selectStr := buildGetJsonbPathTypeSelectStr(dbColumn, objectPath)
	var checkResult string
	if err := dbWithModelAndWhere.Select(selectStr + " as result").Scan(&checkResult).Error; err != nil {
		return err
	}
	pathAttrTypes := strings.Split(checkResult, ",")
	for index, attrType := range pathAttrTypes {
		tipPath := strings.Join(objectPath[:index+1], ",")
		switch attrType {
		case "string": //字符串
			return errors.New(fmt.Sprintf("the attribute type of path:{%s} is not an object but string}", tipPath))
		case "number": //数值
			return errors.New(fmt.Sprintf("the attribute type of path:{%s} is not an object but number}", tipPath))
		case "boolean": //布尔
			return errors.New(fmt.Sprintf("the attribute type of path:{%s} is not an object but boolean}", tipPath))
		case "object": //对象
			continue
		case "null": //不存在字段
			missObjVal := buildMissPathVal(objectPath, index+1)
			missPath := JoinPostgresJsonbPath(objectPath[:index+1])
			//创建对象
			err := dbWithModelAndWhere.Update(dbColumn, datatypes.JSONSet(dbColumn).Set(missPath, missObjVal)).Error
			if err != nil {
				return errors.New(fmt.Sprintf("创建中间路径失败，path:{%s},err=%s", missPath, err.Error()))
			}
			return nil
		case "array": //数组
			return errors.New(fmt.Sprintf("the attribute type of path:{%s} is not an object but array}", tipPath))
		}
	}
	return nil
}

// jsonb路径{a,b,c}
func JoinPostgresJsonbPath(objectPath []string) string {
	return "{" + strings.Join(objectPath, ",") + "}"
}

// 连接jsonb路径链路 targetColumn->a->>b
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
func buildMissPathVal(objectPath []string, beginIndex int) map[string]any {
	if beginIndex >= len(objectPath)-1 {
		return map[string]any{}
	}
	return map[string]any{
		objectPath[beginIndex]: buildMissPathVal(objectPath, beginIndex+1),
	}
}

// 构建检查路径用的select的值
func buildGetJsonbPathTypeSelectStr(dbColumn string, objectPath []string) string {
	lenPath := len(objectPath)
	//检查前面的路径是否对象（除路径最后一个）
	selectStr := strings.Builder{}
	for i := 0; i < lenPath-1; i++ {
		selectStr.WriteString("COALESCE(jsonb_typeof(")
		selectStr.WriteString(joinPostgresJsonbPathChain(dbColumn, objectPath[:i+1]))
		selectStr.WriteString(")::TEXT,'null')")
		if i < lenPath-2 {
			selectStr.WriteString(" || ',' || ")
		}
	}
	return selectStr.String()
}
