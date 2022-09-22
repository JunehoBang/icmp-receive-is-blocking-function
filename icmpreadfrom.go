package main

import (
	"log"
	"net"

	// "log"
	"fmt"

	"golang.org/x/net/icmp"
	// "golang.org/x/sys/unix"
	// "golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func main() {
	var num int
	num = 1
	// log.Prinitln(num)
	fmt.Println(num)
	conn, err := net.ListenIP("ip4:icmp", &net.IPAddr{IP: net.ParseIP("192.168.9.34")})
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	buf := make([]byte, 65536)
	n, addr, err := conn.ReadFrom(buf)
	// Blocks the software. In case the machine that this code runs receives an ICMP message,
	// it proceeds and arrives at the end of this code
	if err != nil {
		panic(err)
	}
	rm, _ := icmp.ParseMessage(1, buf[:n])
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		log.Println("icmp replied by: ", addr.String())
	case ipv4.ICMPTypeEcho:
		log.Println("icmp requested by: ", addr.String())

	}

}
