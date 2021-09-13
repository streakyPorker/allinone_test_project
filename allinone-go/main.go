package main

import (
	"allinone/rpc/userpcx"
	"fmt"

	//_ "allinone/fileProcess"
)

func main() {

	//util.SimpleDeco(basic.TypeFunc)()
	//util.SimpleDeco(basic.MapFunc)()
	//util.SimpleDeco(basic.SwitchFunc)()
	//util.SimpleDeco(basic.FuncFunc)()
	//webGin.SimpleServer()
	//oop.TestFunc()

	//rpc.RunSimpleTest()

	fmt.Println(userpcx.BrokeStruct(struct {
		userpcx.Muller
	}{
	}).(userpcx.Muller))

}
