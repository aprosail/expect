package expect

import (
	"fmt"
	"reflect"
	"strconv"
)

var GlobalFormatRules = &FormatRules{
	IntBase:                 IntBaseDec,
	FloatPrecision:          FloatPrecisionPercent,
	FloatScientificNotation: false,
}

type FormatRules struct {
	IntBase                 int
	FloatPrecision          int
	FloatScientificNotation bool
}

const (
	IntBaseDec = 10
	IntBaseHex = 16
	IntBaseBin = 2
	IntBaseOct = 8
)

const (
	FloatPrecisionPercent = 2
	FloatPrecisionMillis  = 3
	FloatPrecisionMicro   = 6
)

func format(value any, rules *FormatRules) string {
	t := reflect.TypeOf(value)
	return "<" + t.String() + "> " + formatValue(t.Kind(), value, rules)
}

func formatValue(kind reflect.Kind, value any, rules *FormatRules) string {
	switch kind {
	case reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(int64(value.(int)), rules.IntBase)

	case reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(uint64(value.(uint)), rules.IntBase)

	case reflect.Float32:
		return strconv.FormatFloat(
			float64(value.(float32)),
			resolveFloatFormat(rules.FloatScientificNotation),
			rules.FloatPrecision, 32,
		)

	case reflect.Float64:
		return strconv.FormatFloat(
			value.(float64),
			resolveFloatFormat(rules.FloatScientificNotation),
			rules.FloatPrecision, 64,
		)
	}
	return fmt.Sprint(value)
}

func resolveFloatFormat(scientificNotation bool) byte {
	if scientificNotation {
		return 'e'
	}
	return 'f'
}
