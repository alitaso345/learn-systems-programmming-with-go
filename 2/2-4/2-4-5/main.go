package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "www.google.co.jp:80")
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: www.google.co.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
}
