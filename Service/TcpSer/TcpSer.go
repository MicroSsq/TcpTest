package TcpSer

import (
	"net"
	"sync"
	"fmt"
)

type TcpCon struct {
	con net.Conn
}


type TOnReciverData func()
type TOnSendData func()

type MsConnect struct {
	con net.Conn
}

type MsTcpSer struct {
	Listener  net.Listener
	OnReciverData TOnReciverData
	OnSendData TOnSendData
}

var (
	poor sync.Pool
)



func handleConnection(conn net.Conn)  {

	fmt.Println("有新的连接：")

}

func (ser *MsTcpSer)InitSer()error{
	ln, err := net.Listen("tcp", ":3535")
	if err != nil {
		fmt.Println(err)
		return err
	}
	ser.Listener = ln
	return nil
}

func (ser *MsTcpSer)Run(){

	for {
		conn, err := ser.Listener.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}


}

