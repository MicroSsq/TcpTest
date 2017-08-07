package main

import (
	"fmt"
	"TcpTest/Client/Tcp"
)

func main(){
	client := new(Tcp.MsClient)
	fmt.Println("hello Client")
	client.Run()
	//net.Dial("tcp",":3535")
}