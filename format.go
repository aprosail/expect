package expect

import (
	"fmt"
	"reflect"
	"strconv"
)

var GlobalFormatRules = &FormatRules{
	DigitSystemBase: IntBaseDec,
}

type FormatRules struct {
	DigitSystemBase int
}

const (
	IntBaseDec = 10
	IntBaseHex = 16
	IntBaseBin = 2
	IntBaseOct = 8
)

func format(value any, rules *FormatRules) string {
	t := reflect.TypeOf(value)
	return "<" + t.String() + "> " + formatValue(t.Kind(), value, rules)
}

func formatValue(kind reflect.Kind, value any, rules *FormatRules) string {
	switch kind {
	case reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(int64(value.(int)), rules.DigitSystemBase)
	case reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(uint64(value.(uint)), rules.DigitSystemBase)
	}
	return fmt.Sprint(value)
}
