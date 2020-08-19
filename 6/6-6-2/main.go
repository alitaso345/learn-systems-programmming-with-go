package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	sendMessage := []string{"ASCII", "PROGRAMMING", "PLUS"}
	current := 0
	var conn net.Conn = nil
	for {
		var err error
		if conn == nil {
			conn, _ = net.Dial("tcp", "localhost:8888")
			fmt.Printf("Access: %d\n", current)
		}

		request, _ := http.NewRequest("POST", "http://localhost:8888", strings.NewReader(sendMessage[current]))
		request.Write(conn)
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		dump, _ := httputil.DumpResponse(response, true)
		fmt.Println(string(dump))
		current++
		if current == len(sendMessage) {
			break
		}
	}
	conn.Close()
	fmt.Println("closed")
}
