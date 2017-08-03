package TcpSer

import (
	"net"
	"sync"
)

type MsTcpSer struct {
	Lis net.Listener
}

var (
	poor sync.Pool
)

func GetSer() (* MsTcpSer){
	var srv *MsTcpSer
	m := poor.Get()
	if m == nil{
		srv = new(MsTcpSer)
	}else{
		srv = m.(*MsTcpSer)
	}
	return srv
}

func (ser *MsTcpSer)Run(){
	ser.Lis,_ = net.Listen("tcp",":3535")


}

