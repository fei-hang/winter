package utils

import (
	"reflect"
	"strconv"
)

func StructToParm(s interface{}) string {

	stype := reflect.TypeOf(s)
	svalue := reflect.ValueOf(s)
	if stype.Kind() == reflect.Ptr {
		stype = stype.Elem()
	}
	if svalue.Kind() == reflect.Ptr {
		svalue = svalue.Elem()
	}
	numFileld := stype.NumField()
	var pram string
	for i := 0; i < numFileld; i++ {
		filed := stype.Field(i)
		if filed.IsExported() {
			filedJson := filed.Tag.Get("json")
			if filedJson != "" {
				if pram == "" {
					pram = "?" + filedJson + "=" + svalue.FieldByName(filed.Name).String()
				} else {
					pram = pram + "&" + filedJson + "=" + svalue.FieldByName(filed.Name).String()
				}
			} else {
				var value string
				if pram == "" {
					pram = "?" + filed.Name + "=" + svalue.Field(i).String()
				} else {
					filedValue := svalue.FieldByName(filed.Name)
					if filedValue.Kind() == reflect.Ptr {
						filedValue = filedValue.Elem()
					}
					switch filedValue.Kind().String() {
					case "int", "int64":
						value = strconv.FormatInt(filedValue.Int(), 10)
					default:
						value = filedValue.String()
					}
					pram = pram + "&" + filed.Name + "=" + value
				}
			}
		}
	}
	return pram
}
