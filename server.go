package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// tcpの接続アドレスを作成する
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":3030")
	if err != nil {
		log.Fatal(err)
	}

	// リスナーを作成する
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start TCP Server...")
	for {
		// クライアントからのコネクション情報を受け取る
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}

		// ハンドラーに接続情報を渡す
		echoHandler(conn)
	}
}

func echoHandler(conn *net.TCPConn) {
	defer conn.Close()
	for {
		// リクエストを受け付けたらサーバー側に「response from server」を返す
		_, err := io.WriteString(conn, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
