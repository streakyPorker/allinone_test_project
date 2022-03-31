package basic

import (
	"fmt"
	"testing" // 必须引入testing
)

// 必须以Test开头
func TestSome(t *testing.T /*必须使用testing.T作为参数*/) {
	fmt.Println("here")
}

func TestDeferFunc(t *testing.T) {
	user := &User{"asd", 10}
	fmt.Println(user.string())
}
