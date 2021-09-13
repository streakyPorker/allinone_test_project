package basic

import (
	"fmt"
	"reflect"
)

func init() {
	//TypeFunc()

}

func TypeFunc() {
	a := "hello，世界！"
	b := []byte(a)
	c := []rune(a)
	println(b, c, reflect.TypeOf(b), reflect.TypeOf(c))
	descStr := "go语言的复合类型如下：\n" +
		"* pointerType ：指针类型，" +
		"[n] eleType ：定长数组类型，" +
		"[] eleType ：切片类型，" +
		"map [keyType]valueType ：映射类型，" +
		"chan eleType ：通道类型\n" +
		"struct {" +
		"feildName feildType" +
		"}：结构类型，" +
		"interface{" +
		"methodName(argTypes)(outputTypes)" +
		"}：接口类型\n"
	println(descStr)

	arr := [...]int{1, 2, 3} // 数组一旦创建就是定长的
	arr2 := [10]int{1, 2, 3} // 可以显式指定长度
	slice := []int{1, 2, 3}  // 切片类型在打印时不显示长度，因其是不定长的
	fmt.Println(arr, slice, reflect.TypeOf(arr), reflect.TypeOf(arr2), reflect.TypeOf(slice))
	for i, v := range slice /*arr2*/ {
		fmt.Println(i, v)
	}
	fmt.Println(arr2[:2], reflect.TypeOf(arr2[:3]))
	a_s1 := make([]int, 10, 15)
	fmt.Println(a_s1, reflect.TypeOf(a_s1))

}

func MapFunc() {
	m1 := map[string]int{}
	fmt.Println(m1, reflect.TypeOf(m1))

	m2 := make(map[string]int, 10)
	m2["a"] = 10
	fmt.Println(m2, reflect.TypeOf(m2))

	delete(m2, "a")

	userMap := make(map[int]User)
	userMap[0] = User{age: 10, name: "asd"}
	fmt.Println(userMap[0])
	userRef := userMap[0] // 是一次深拷贝，并不能修改
	userRef.age = 12      // 无法修改原map数据
	fmt.Println(userMap[0])

	alterableUserMap := map[int]*User{
		0: {age: 10, name: "asd"},
	}
	alterUser(*alterableUserMap[0])
	fmt.Println(alterableUserMap[0])

}

func SwitchFunc() {
	switch i := 1; i {
	case 1, 2:
		fmt.Println(i)
		fallthrough
	case 3, 4:
		fmt.Println("fallthrough in 34")
		fallthrough
	default:
		fmt.Println("here in default")
		fallthrough
	case 5, 6:
		fmt.Println("case after default")
	}

}
func FuncFunc() {
	slice := []int{1, 2, 3, 4}
	array := [...]int{1, 2, 3, 4}
	defer func() {
		println("defer 1")
	}()
	multiArg(slice...)
	defer func(arg int) {
		fmt.Println("defer 2 with arg :", arg)
	}(123)
	defer multiArg(array[:]...)

}

func alterUser(user User) {
	user.age += 10
}

func multiArg(arr ...int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

type User struct {
	name string
	age  int
}
