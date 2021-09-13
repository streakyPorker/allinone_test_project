package userpcx

import (
	"fmt"
	"reflect"
)

type TS struct {
	A, B int
}

func (t *TS) Mul() int {
	return t.A * t.B
}

type Muller interface {
	Mul() int
}

func BrokeStruct(int interface{}) interface{} {
	t := reflect.TypeOf(int)
	fmt.Println("t is", reflect.TypeOf(int), "input type has", t.NumMethod(), "methods")

	callBroker := make(map[string]interface{})
	fmt.Println(callBroker)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println("name:", method.Name)
		fmt.Println("type:", method.Type)
	}
	return int
}
