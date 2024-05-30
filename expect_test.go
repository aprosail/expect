package expect_test

import (
	"testing"

	"github.com/aprosail/expect"
)

func TestExpect(t *testing.T) {
	// Single test chain.
	expect.Init(t, nil).Expect(1 + 1).ToBe(2)

	// Multi test with same expect handler.
	e := expect.Init(t, nil)
	e.Expect(1 + 1).ToBe(2)
	e.Expect(1 + 2).ToBe(3)
}
