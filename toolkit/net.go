package toolkit

import (
	"fmt"
	"net"
)

func GetOutboundIP() (err error, ip net.IP) {
	conn, err := net.Dial("udp", "223.5.5.5:53")
	if err != nil {
		fmt.Println("获取IP地址失败", err)
		return
	}
	defer conn.Close()

	//_, err = conn.Write([]byte("Hello, server!"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = localAddr.IP
	return
}

func GetOutboundIpString() (ip string, err error) {
	conn, err := net.Dial("udp", "223.5.5.5:53")
	if err != nil {
		fmt.Println("获取IP地址失败", err)
		return
	}
	defer conn.Close()

	//_, err = conn.Write([]byte("Hello, server!"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = localAddr.IP.String()
	return
}
