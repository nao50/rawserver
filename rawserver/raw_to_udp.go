package main

import (
	"fmt"
	"net"
	"syscall"
)

func main() {
	const proto = (syscall.ETH_P_ALL<<8)&0xff00 | syscall.ETH_P_ALL>>8
	/////////////////////////////////////////////////////////////////////
	// recv
	fmt.Println("\n===== syscall.Socket() =====")
	recvFd, _ := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_DGRAM, proto)
	defer syscall.Close(recvFd)

	recvIf, _ := net.InterfaceByName("eth1")

	var recvHaddr [8]byte
	copy(recvHaddr[0:7], recvIf.HardwareAddr[0:7])
	fmt.Println("\n===== syscall.SockaddrLinklayer() =====")
	recvAddr := syscall.SockaddrLinklayer{
		Protocol: proto,
		Ifindex:  recvIf.Index,
		Halen:    uint8(len(recvIf.HardwareAddr)),
		Addr:     recvHaddr,
	}

	fmt.Println("\n===== syscall.Bind() =====")
	if err := syscall.Bind(recvFd, &recvAddr); err != nil {
		fmt.Println("bind: ", err)
	}

	recvBuf := make([]byte, 8214)
	/////////////////////////////////////////////////////////////////////
	// send
	conn, err := net.Dial("udp", "10.0.11.20:2152")
	if err != nil {
		fmt.Println("conn: ", err)
	}
	defer conn.Close()
	/////////////////////////////////////////////////////////////////////
	// main loop
	fmt.Println("Starting raw server...")
	for {
		fmt.Println("\n===== syscall.Recvfrom() =====")
		n, _, _ := syscall.Recvfrom(recvFd, recvBuf, 0)

		fmt.Println("recieved size: ", n)
		fmt.Printf("Receive Packet: %02v", recvBuf[:n])

		go func() {
			_, err = conn.Write(recvBuf[:n])
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

}
