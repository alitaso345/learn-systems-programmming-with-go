package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	file, _ := os.Create("test.text.gz")
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.writer example\n")
	writer.Close()
}
