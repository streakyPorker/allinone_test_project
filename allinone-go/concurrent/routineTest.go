package concurrent

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	//selectTest()
}

func routineTest() {
	i := 1
	runtime.GOMAXPROCS(3)
	for ; i < 10; i++ {
		iInline := i
		go func(iInline int) {
			for {
				println("in the ", iInline, "th round")
				time.Sleep(1 * time.Second)
			}

		}(iInline)
		time.Sleep(2 * time.Second)
		fmt.Println("NUmGoroutine=", runtime.NumGoroutine())
	}
}

func waitGroupTest() {
	var urls = []string{
		" https://www.golang.org",
		" https://www.google.com",
		" https://www.qq.com",
	}
	var wg sync.WaitGroup // 类似于CountDownLatch
	wg.Add(len(urls))
	for idx, url := range urls {
		go func(idx int, url string) {
			defer wg.Done()
			fmt.Println(url)
			time.Sleep(time.Duration(rand.Int31n(6)) * time.Second)
			println(idx, " routine is done")
		}(idx, url)
	}
	wg.Wait() // 如果Done的次数比Add的多，会panic
}

