package validators

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func unique(src []interface{}) []interface{} {
	keys := make(map[interface{}]bool)
	list := []interface{}{}
	for _, entry := range src {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (cv *cValidator) ExistMultiple() validator.Func {
	return func(fl validator.FieldLevel) bool {

		arr := strings.Split(fl.Param(), ".")

		interfaceValue := reflect.ValueOf(fl.Field().Interface())

		if interfaceValue.Kind() != reflect.Slice && interfaceValue.Kind() != reflect.Array {
			return false
		}

		interfaceSlice := make([]interface{}, interfaceValue.Len())

		for i := 0; i < interfaceValue.Len(); i++ {
			interfaceSlice[i] = interfaceValue.Index(i).Interface()
		}

		var count int64
		err := cv.db.
			Table(arr[0]).
			Distinct(arr[1]).
			Where(arr[1]+" IN (?)", fl.Field().Interface()).
			Count(&count).Error

		if err != nil {
			return false
		}

		return (int64)(len(unique(interfaceSlice))) == count
	}
}
