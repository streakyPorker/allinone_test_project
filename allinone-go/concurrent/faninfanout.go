package concurrent

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("in fan in&out test")
	simpleTest()
}

type TestStruct struct {

}

func getGenerator(done chan interface{}) chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				goto Label
			}
		}
	Label:
		close(ch)
	}()
	return ch
}

func simpleTest() {
	done := make(chan interface{})
	ch := getGenerator(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)


}

func fanIOTest() {
}
