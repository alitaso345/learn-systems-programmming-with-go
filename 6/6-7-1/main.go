package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
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
		request.Header.Set("Accept-Encoding", "gzip")
		request.Write(conn)
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		dump, _ := httputil.DumpResponse(response, false)
		fmt.Println(string(dump))

		defer response.Body.Close()

		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, _ := gzip.NewReader(response.Body)
			io.Copy(os.Stdout, reader)
		} else {
			io.Copy(os.Stdout, response.Body)
		}
		current++
		if current == len(sendMessage) {
			break
		}
	}
	conn.Close()
	fmt.Println("closed")
}
