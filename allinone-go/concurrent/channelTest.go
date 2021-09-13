package concurrent

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	println("init channelTest")
	//selectTest2()
}

func channelTest() {
	c := make(chan interface{})
	c2 := make(chan interface{}, 20)
	go func(c chan interface{}, c2 chan interface{}) {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
			c2 <- sum
		}
		close(c2)
		println(sum)
		time.Sleep(2 * time.Second)
		c <- sum
	}(c, c2)
	println(runtime.NumGoroutine())
	//
	for v := range c2 {
		fmt.Print(v, " ")
	}
	some := <-c
	fmt.Println(some, "outside")

}

func channelTest2() {
	c := make(chan int, 10)
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)

	for i := 0; i < 20; i++ {
		fmt.Print(<-c, " ")
		time.Sleep(200 * time.Millisecond)
	}
	/*
		channel被关闭会使得读取全部读到默认值，但并不会阻塞
	*/

	c2 := make(chan int, 1)
	go func() {
		for i := 0; i < 20; i++ {
			c2 <- i
			time.Sleep(time.Duration(rand.Int31n(600)) * time.Millisecond)
		}
		close(c2)
	}()
	// range操作并不是取快照，而是能够阻塞取，但channel若不关闭会panic
	for v := range c2 { // 对channel进行range操作，只能接到元素，不能接到index
		fmt.Print(v, " ")
	}

}

func selectTest() {
	ch := make(chan int, 1)
	done := make(chan interface{})
	go func(chan int) {
		for {
			// 如果没有channel可读就会阻塞，有则随机选择
			select {
			case ch <- 0:
				fmt.Println("write 0")
			case ch <- 1:
				fmt.Println("write 1")
			case <-done: // 这种case用于监控channel是否关闭
				fmt.Println("done channel is done")
				return
			}
		}
	}(ch)
	for i := 0; i < 10; i++ {
		println(<-ch)
	}
	close(done)
}

func selectTest2() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5*time.Second)
		close(done)

	}()

	for {
		select {
		case <-done: // 这种case用于监控channel是否关闭,但是写入channel时，也会引起这一事件
		// 所以往往需要把通知channel设为只读
			fmt.Println("done channel is done")
			<-done
		default:
			fmt.Println("nothing happens")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
