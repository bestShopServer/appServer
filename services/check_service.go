package services

import (
	"reflect"
	"shop/pkg/logging"
)

// Check All Params From The Request
func CheckParams(v interface{}, cat string, err error, optionalFields []string) bool {
	check := true
	if err != nil {
		logging.Error("Parsing Param Err:" + err.Error())
		check = false
	} else if !CheckField(v, optionalFields) {
		check = false
	}
	return check
}

// Check All Field In A Struct
func CheckField(v interface{}, optionalFields []string) bool {
	//ref := reflect.ValueOf(v)
	ref := reflect.TypeOf(v)
	check := true
	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)

		var optional bool
		for j := 0; j < len(optionalFields); j++ {
			if field.Name == optionalFields[j] {
				optional = true
				break
			}
		}

		if optional {
			continue
		}

		switch field.Type.String() {
		case "int":
			if reflect.ValueOf(v).Field(i).Int() == 0 {
				check = false
			}
		case "string":
			if reflect.ValueOf(v).Field(i).String() == "" {
				check = false
			}
		case "struct":
			check = CheckField(reflect.ValueOf(v).Field(i).Interface(), []string{})
		}

	}
	return check
}
