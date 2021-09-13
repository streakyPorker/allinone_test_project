package userpcx

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	_ "github.com/ugorji/go/codec"
	"log"
	"reflect"
	"time"
)

type Args struct {
	A int
	B int
}

func (a *Args) Mul() int {
	return a.A * a.B
}

type Reply struct {
	C int
}

type Arith int

type MyServer struct {
	*server.Server
}

func NewServer() *MyServer {
	return &MyServer{
		server.NewServer(),
	}
}

func (s *MyServer) MyRegisterName(rcvr interface{}) {
	err := s.RegisterName("Arith", rcvr, "metadata")
	fmt.Println(reflect.TypeOf(rcvr).String())
	if err != nil {
		return
	}
}

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}



func SimpleServer() {
	s := NewServer()
	s.MyRegisterName(new(Arith))
	err := s.Serve("tcp", ":8972") // 一次使用永久有效
	if err != nil {
		return
	}
}

func SimpleClient() {
	addr := "localhost:8972"
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	xclient := client.NewXClient("Arith",
		client.Failtry,
		client.RandomSelect,
		d, client.DefaultOption)

	args := &Args{
		A: 10, B: 20,
	}
	reply := &Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}

func AsyncSimpleClient() {
	addr := "localhost:8972"
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	xclient := client.NewXClient("Arith",
		client.Failtry,
		client.RandomSelect,
		d, client.DefaultOption)

	args := &Args{
		A: 10, B: 20,
	}
	reply := &Reply{}
	call, err := xclient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		return
	}
	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	} else {
		log.Printf("async result : %d * %d = %d", args.A, args.B, reply.C)
	}

}



func foo(a, b int) {
	fmt.Println(a, " along with ", b)
}

func SimpleRun() {
	//	// 使用反射调用方法
	//	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}
	//	for i, v := range reflect.ValueOf(foo).Call(args) {
	//		fmt.Println(i, v.String())
	//	}

	go SimpleServer()
	time.Sleep(1 * time.Second)
	AsyncSimpleClient()
	time.Sleep(20 * time.Second)
}
