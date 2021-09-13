package rpc

// from go语言高级编程

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type HelloService struct {
}

// Hello 方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法。
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func SimpleRpcServer() {
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		log.Fatal(err, "can`t register")
		return
	}

	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	log.Println(listener)

	for {
		conn, err := listener.Accept() // 阻塞了
		// 只有在收到rpc请求的时候才会返回连接
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn) // 仅仅服务一次这个连接,只在调用端close之后，才会停止阻塞
		log.Println("successfully called")
	}
}

func SimpleHttpRpcServer() {
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		log.Fatal(err, " can`t register")
		return
	}
	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer:     writer,
		}
		err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			return
		}
	})
	err = http.ListenAndServe(":9999", nil)
	if err != nil {
		return
	}
}

func SimpleRpcClient(sleep int) {
	client, err := rpc.Dial("tcp", "localhost:9999")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string


	err = client.Call("HelloService.Hello", "hi,i`m lzy ", &reply)

	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Duration(sleep) * time.Second)
	fmt.Println(reply)
	err = client.Close() // close后才会唤醒阻塞方法
	log.Println(sleep,"th call closed")
	if err != nil {
		return
	}
}

func RunSimpleTest() {
	go func() {
		go SimpleRpcServer()
		//go SimpleHttpRpcServer() // 仅仅可以通过http请求来调用rpc
		time.Sleep(2 * time.Second)
		for i := 0; i < 10; i++ {
			go SimpleRpcClient(i)
		}

	}()
	// 等待结果打印完成
	time.Sleep(100 * time.Second)
}
