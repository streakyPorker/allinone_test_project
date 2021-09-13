package userpcx

import "flag"

func Test(){
	addr := flag.String("addr", "localhost:8972", "server address")
	println(addr)
}