package main

import (
	"fmt"
	"net"
	"syscall"
)

const (
	EPOLLET        = 1 << 31
	MaxEpollEvents = 32
)

func main() {
	const proto = (syscall.ETH_P_ALL<<8)&0xff00 | syscall.ETH_P_ALL>>8
	/////////////////////////////////////////////////////////////////////
	// recv
	fmt.Println("\n===== syscall.Socket() =====")
	// Set nonblock & cloexec
	recvFd, _ := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_DGRAM|syscall.O_NONBLOCK|syscall.SOCK_CLOEXEC, proto)
	defer syscall.Close(recvFd)

	// Set nonblock
	if err := syscall.SetNonblock(recvFd, true); err != nil {
		fmt.Println("setnonblock recvFd: ", err)
	}

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
	fmt.Println("\n===== syscall.Socket() =====")
	sendFd, _ := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_DGRAM|syscall.O_NONBLOCK|syscall.SOCK_CLOEXEC, proto)
	defer syscall.Close(sendFd)

	// Set nonblock
	if err := syscall.SetNonblock(sendFd, true); err != nil {
		fmt.Println("setnonblock sendFd: ", err)
	}

	sendIf, _ := net.InterfaceByName("eth0")

	var sendHaddr [8]byte
	copy(sendHaddr[0:7], sendIf.HardwareAddr[0:7])
	fmt.Println("\n===== syscall.SockaddrLinklayer() =====")
	sendAddr := syscall.SockaddrLinklayer{
		Protocol: proto,
		Ifindex:  sendIf.Index,
		Halen:    uint8(len(sendIf.HardwareAddr)),
		Addr:     sendHaddr,
	}

	fmt.Println("\n===== syscall.Bind() =====")
	if err := syscall.Bind(sendFd, &sendAddr); err != nil {
		fmt.Println("bind: ", err)
	}

	// sendBuf := make([]byte, 8214)

	/////////////////////////////////////////////////////////////////////
	// main loop
	fmt.Println("Starting raw server...")
	for {
		// fmt.Println("\n===== syscall.Recvfrom() =====")
		n, _, err := syscall.Recvfrom(recvFd, recvBuf, 0)
		switch err {
		case nil:
			fmt.Printf("Receive Packet: %02x\n", recvBuf[:n])
			go func() {
				if err := syscall.Sendto(sendFd, recvBuf[:n], 0, &sendAddr); err != nil {
					fmt.Println("Sendto:", err)
				}
			}()
		case syscall.EAGAIN:
			// ????
			continue
		default:
			// UNKNOWN ERROR
			// break
		}

		// fmt.Printf("Receive Packet: %02x\n", recvBuf[:n])

		// go func() {
		// 	if err := syscall.Sendto(sendFd, recvBuf[:n], 0, &sendAddr); err != nil {
		// 		fmt.Println("Sendto:", err)
		// 	}
		// }()
	}
}
