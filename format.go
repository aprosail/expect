package expect

import (
	"fmt"
	"reflect"
	"strconv"
)

func format(value any) string {
	t := reflect.TypeOf(value)
	return fmt.Sprint("<"+t.String()+">", formatValue(t.Kind(), value))
}

func formatValue(kind reflect.Kind, value any) string {
	switch kind {
	case reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(int64(value.(int)), 10)
	case reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(uint64(value.(uint)), 10)
	}
	return fmt.Sprint(value)
}
