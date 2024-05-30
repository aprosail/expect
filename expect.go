package expect

import (
	"fmt"
	"testing"
)

type TestHandler struct {
	tester *testing.T
}

func Init(t *testing.T) *TestHandler {
	return &TestHandler{tester: t}
}

type ExpectHandler struct {
	tester  *testing.T
	handler any
}

func (t *TestHandler) Expect(actual any) *ExpectHandler {
	return &ExpectHandler{tester: t.tester, handler: actual}
}

func (e *ExpectHandler) ToBe(expect any) {
	if e.handler != expect {
		e.tester.Error(
			dim("Expect:"), green(fmt.Sprint(expect)),
			dim("but get:"), red(fmt.Sprint(e.handler)),
		)
	}
}
