package main

import (
	"net/http"
	"os"
)

func main() {
	request, _ := http.NewRequest("GET", "http://ascii.jp", nil)
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)
}
