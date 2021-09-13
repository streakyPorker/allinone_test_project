package fileProcess

import (
	"fmt"
	"io/ioutil"
)

func init() {
	file, err := ioutil.ReadFile("./asd.txt")
	if err != nil {
		return
	}
	fmt.Println(file)
	dir, err := ioutil.ReadDir("./")
	if err != nil {
		return
	}
	for i, info := range dir {
		fmt.Println(i,info)
	}


}
