package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func receiveTCPConn(ln *net.TCPListener) {
	for {
		err := ln.SetDeadline(time.Now().Add(time.Second * 100))
		if err != nil {
			log.Fatal(err)
		}
		conn, err := ln.AcceptTCP() // ここで次の着信を待ち、新しいコネクションを返すため、一回止まる。
		if err != nil {
			log.Fatal(err)
		}
		go echoHandler(conn)
	}
}

func echoHandler(conn *net.TCPConn) {
	defer conn.Close()
	for {
		b := make([]byte, 2048)
		i, err := conn.Read(b)
		if err != nil {
			log.Println(333, err)
		}
		if i != 0 {
			fmt.Println(string(b))
		}
		time.Sleep(time.Second * 10)
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatal("1: ", err)
	}

	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal("2: ", err)
	}
	receiveTCPConn(ln)

}
