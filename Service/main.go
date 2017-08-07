package main

import (
	"fmt"
	"TcpTest/Service/TcpSer"
)

var
  srv *TcpSer.MsTcpSer
func main(){
	fmt.Println("hello world")
	srv = new(TcpSer.MsTcpSer)
	srv.InitSer()
	srv.Run()
}
