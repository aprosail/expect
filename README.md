# expect

Syntax sugar for easier and highlighted golang testing.

Before:

```go
package demo_test

import (
  "testing"
)

func TestDemo(t *testing.T) {
  if 1 + 1 != 2 {
    t.Fatalf("Expect %v, but got %v", 2, 1 + 1)
  }

  if 1 + 2 != 3 {
    t.Fatalf("Expect %v, but got %v", 3, 1 + 2)
  }

  if 1 + 3 != 4 {
    t.Fatalf("Expect %v, but got %v", 4, 1 + 3)
  }
}
```

After:

```go
package demo_test

import (
  "testing"
  "github.com/aprosail/expect"
)

func TestDemo(t *testing.T) {
  const expect = 2
  const actual = 1 + 1
  e := expect.Init(t, nil)
  e.Expect(1 + 1).ToBe(2)
  e.Expect(1 + 2).ToBe(3)
  e.Expect(1 + 3).ToBe(4)
}
```
