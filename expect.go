package expect

import (
	"testing"
)

type TestHandler struct {
	tester *testing.T
	rules  *FormatRules
}

func Init(tester *testing.T, rules *FormatRules) *TestHandler {
	return &TestHandler{
		tester: tester,
		rules:  rules,
	}
}

type ExpectHandler struct {
	tester  *testing.T
	rules   *FormatRules
	handler any
}

func (t *TestHandler) Expect(actual any) *ExpectHandler {
	return &ExpectHandler{
		tester:  t.tester,
		rules:   t.rules,
		handler: actual,
	}
}

func (e *ExpectHandler) ToBe(expect any) {
	position := Trace(2)
	if e.handler != expect {
		e.tester.Error(
			yellow(position.String())+dim(":"),
			dim("Expect:"), green(format(expect, e.rules)),
			dim("but get:"), red(format(e.handler, e.rules)),
		)
	}
}
