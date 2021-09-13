package util

import (
	"fmt"
	"reflect"
	"runtime"
)

func SimpleDeco(noArgFunction func()) func(){
	return func(){
		fmt.Println("---------当前运行的方法是",runtime.FuncForPC(reflect.ValueOf(noArgFunction).Pointer()).Name(),"---------")
		noArgFunction()
	}
}