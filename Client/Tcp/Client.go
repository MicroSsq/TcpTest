package Tcp

import (
	"net"
	"bytes"
	"fmt"
	"time"
)

type MsClient struct {
	conn net.Conn
}


func (cli *MsClient)OnReciver()error {
	bt := make([]byte,1024)
	bf := bytes.NewBuffer(bt)
	for {
		_,err :=  cli.conn.Read(bt)
		if err != nil{
			fmt.Println(err)
			cli.conn.Close()
			return err
		}
		fmt.Println(string(bt[:]))
		bf.Reset()
		bf.WriteString("Client:")
		bf.WriteString(time.Now().Format("2006-01-02 15:04:05"))
		cli.conn.Write(bf.Bytes())


	}
}

func (cli *MsClient)Run()error{
	conn,err := net.Dial("tcp",":3535")
	if err != nil {
		fmt.Println(err)
		return err
	}
	cli.conn = conn
	cli.OnReciver()
	return nil
}