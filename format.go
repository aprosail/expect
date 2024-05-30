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

// Parse parameter for strconv.FormatFloat from a FormatRules instance.
func (r *FormatRules) floatFormat() byte {
	if r.FloatScientificNotation {
		return 'e'
	}
	return 'f'
}

// Format a value of any type with given FormatRules.
func (r *FormatRules) Format(value any) string {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(int64(value.(int)), r.IntBase)

	case reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(uint64(value.(uint)), r.IntBase)

	case reflect.Float32:
		return strconv.FormatFloat(
			float64(value.(float32)),
			r.floatFormat(), r.FloatPrecision, 32,
		)

	case reflect.Float64:
		return strconv.FormatFloat(
			value.(float64),
			r.floatFormat(), r.FloatPrecision, 64,
		)
	}
	return fmt.Sprint(value)
}

func format(value any, rules *FormatRules) string {
	return "<" + reflect.TypeOf(value).String() + "> " + rules.Format(value)
}
