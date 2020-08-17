package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8888")
	request, _ := http.NewRequest("GET", "http://localhost:8888", nil)
	request.Write(conn)
	response, _ := http.ReadResponse(bufio.NewReader(conn), request)
	dump, _ := httputil.DumpResponse(response, true)
	fmt.Println(string(dump))
}
