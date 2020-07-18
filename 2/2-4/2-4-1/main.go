package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	file.Write([]byte("os.FIle example"))
	file.Close()
}
