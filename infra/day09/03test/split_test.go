package splitutils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplitStr(t *testing.T) {
	got := SplitStr("asdfasdas", "d")
	want := []string{"as", "fas", "as"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("error:want:%v, but got:%v\n", want, got)
	}

	fmt.Println("over!")
}

//所有Test 用例都通过时,结果才为PASS, 只有有一个Test用例失败结果都为FAIL
func Test2SplitStr(t *testing.T) {
	got := SplitStr("asdfasdas", "d")
	want := []string{"as", "fas", "as"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("error:want:%v, but got:%v\n", want, got)
	}

	fmt.Println("over!")
}
