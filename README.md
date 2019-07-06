# Myriad-Dreamin's Test

使用方法

```go
package mtest

import "testing"
import qwq "github.com/Myriad-Dreamin/mydrest"

type MyTestHelper struct {
	qwq.TestHelper
}

var s MyTestHelper

func TestA(t *testing.T) {
	s.AssertEqual(t, 1, 1)
}

func TestB(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			return
		} else {
			t.Error("does not fetch error...")
		}
	}()
	s.AssertEqual(t, 1, 2)
}

```

## Assertion

传入一个`testing.T`或者`testing.B`，即可开启Assertion

目前支持的方法：

#### `Assert(t, i, f)`

如果$f(i)$不成立，则向`t`报告失败，并panic

## `AssertEqual(t, a, b)`

如果$a=b$不成立，则向`t`报告失败，并panic

#### AssertNoErr(t, e)

如果`e`不为空，则向`t`报告错误，并panic

## CheckFunc

#### `IsNil(i)`

如果`i`为`nil`，则返回true，否则返回false

#### `IsNotNil(i)`

如果`i`不为`nil`，则返回true，否则返回false
