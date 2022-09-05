package web

import (
	"net"
	"github.com/rn0l485/core/setting"
)


func Broadcast(senderIP []byte, port int) error {
	if len(senderIP) != 4 {
		 return setting.Err_NotSupport
	}

	senderAddr := net.UDPAddr{
		IP: net.IPv4( senderIP[0], senderIP[1], senderIP[2], senderIP[3]),
		Port: port,
	}

	receiverAddr := net.UDPAddr{
		IP: net.IPv4(255, 255, 255, 255),
		Port: port,
	}

	conn, err := net.DialUDP("udp", &senderAddr, &receiverAddr)
	if err != nil {
		return err
	}

	conn.Write([]byte(`Hi all`))
	conn.Close()

	return nil
}