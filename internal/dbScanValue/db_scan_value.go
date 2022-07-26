package dbScanValue

import (
	"fmt"
	"reflect"
)

// This can be used from golang 1.18 on
func GetValueFromDB[E](dbValue interface{}, mapping map[string]E) (*E, error) {
	if v, ok := mapping[GetStringFromDB(dbValue)]; ok {
		return &v, nil
	}
	return nil, fmt.Errorf("could not read %s", GetStringFromDB(dbValue))
}

func GetStringFromDB(dbValue interface{}) string {
	var value string
	if reflect.ValueOf(dbValue).Kind() == reflect.Ptr {
		if dbValue == nil || dbValue.(*string) == nil {
			return ""
		}
		value = *dbValue.(*string)
	} else {
		value = dbValue.(string)
	}
	return value
}
