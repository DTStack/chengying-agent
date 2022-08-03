package enums

import "reflect"

/*
 @Author: zhijian
 @Date: 2021/5/28 10:29
 @Description: 运行状态
*/

type execStatusType struct {
	Running EnumValueType
	Success EnumValueType
	Failed  EnumValueType
}

func (c execStatusType) List() (enumValues []EnumValueType) {
	v := reflect.ValueOf(c)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		enumValues = append(enumValues, field.Interface().(EnumValueType))
	}
	return enumValues
}

var ExecStatusType = execStatusType{
	Running: EnumValueType{
		Code: 1,
		Desc: "进行中",
	},
	Success: EnumValueType{
		Code: 2,
		Desc: "正常",
	},
	Failed: EnumValueType{
		Code: 3,
		Desc: "失败",
	},
}
