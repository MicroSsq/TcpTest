package TcpSer

import (
	"net"
	"sync"
	"fmt"
	"bytes"
	"time"
)

type TcpCon struct {
	con net.Conn
}


type TOnReciverData func(msg *MsgStr)
type TOnSendData func()

type MsgStr struct {
	param map[string]interface{}
	value interface{}
	result interface{}
}

type MsConnect struct {
	con net.Conn
	OnReciverData TOnReciverData
}

type MsTcpSer struct {
	Listener  net.Listener
	OnSendData TOnSendData
}

var (
	poor sync.Pool
)

func str2Msgstr(str string)MsgStr{
	if str == ""{
		return nil
	}else{

	}
}

func onReciever(conn *MsConnect)  {
	by := make([]byte,1024)
	bf := bytes.NewBuffer(by)
	for{
		_,err := conn.con.Read(by)
		if err != nil{
			conn.con.Write(by)
			conn.con.Close()
			return
		}
		fmt.Println(string(by[:]))
		if conn.OnReciverData != nil{
			bf.Write(by)

			conn.OnrecieverMsg()
		}
		//select {
		//case <- time.After(time.Second):
		//	//bf.Reset()
		//	//bf.WriteString("Service:")
		//	//bf.WriteString(time.Now().Format("2006-01-02 15:04:05"))
		//	//conn.Write(bf.Bytes())
		//}

	}

}

func (ser *MsTcpSer)SendMsg(msg *MsgStr,con net.Conn){
	//

}

func (con *MsConnect)OnrecieverMsg(msg *MsgStr){
	//

}

func handleConnection(conn *MsConnect)  {

	fmt.Println("有新的连接：",conn.con.RemoteAddr())
	go onReciever(conn.con)
	time.Sleep(300)
	conn.con.Write([]byte("sdf"))

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
		mscon := new(MsConnect)
		mscon.con = conn
		mscon.OnReciverData = mscon.OnrecieverMsg
		go handleConnection(mscon)
	}


}

