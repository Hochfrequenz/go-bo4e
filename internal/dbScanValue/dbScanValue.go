package dbScanValue

import "reflect"

// GetStringFromDB converts the interface from the SQL framework to a string that can be mapped to an Enum subsequently
func GetStringFromDB(dbValue interface{}) (string, bool) {
	var value string
	if reflect.ValueOf(dbValue).Kind() == reflect.Ptr {
		if dbValue == nil || dbValue.(*string) == nil {
			return "", false
		}
		value = *dbValue.(*string)
	} else {
		value = dbValue.(string)
	}
	return value, true
}

// GetIntFromDB converts the interface from the SQL framework to an int that can be interpreted an Enum subsequently
func GetIntFromDB(dbValue interface{}) (int, bool) {
	var value int
	if reflect.ValueOf(dbValue).Kind() == reflect.Ptr {
		if dbValue == nil || dbValue.(*int) == nil {
			return 0, false
		}
		value = *dbValue.(*int)
	} else {
		value = dbValue.(int)
	}
	return value, true
}
