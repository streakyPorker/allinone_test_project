package concurrent

import (
	"context"
	"fmt"
	"time"
)

func init() {
}

// context库的核心功能：退出通知、元数据传递
// context维护了调用树
// 第一个创建Context的goroutine称为root节点

type otherContext struct {
	context.Context
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "get msg to cancel")
			return
		default:
			fmt.Println(name, "is running")
			time.Sleep(1 * time.Second)
		}
	}
}

func workWithVal(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "get msg to cancel")
			return
		default:
			fmt.Println(name, "is running, value is", ctx.Value("key").(string))
			time.Sleep(1 * time.Second)
		}
	}
}

func UseContext() {
	ctxa, cancel := context.WithCancel(context.Background())

	go work(ctxa, "work1")

	tm := time.Now().Add(3 * time.Second)
	ctxb, _ := context.WithDeadline(ctxa, tm)

	go work(ctxb, "work2")

	oc := otherContext{ctxb}
	ctxc := context.WithValue(oc, "key", "RISEUPLZY")

	go workWithVal(ctxc, "work3")

	time.Sleep(10 * time.Second)
	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println("main stops")
}
