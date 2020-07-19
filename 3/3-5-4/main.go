package main

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	newFile, err := os.Create("Lenna2.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer newFile.Close()

	chunks := readChunks(file)
	// シグネチャ書き込み
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")

	// 先頭に必要なIHDRチャンクを書き込み
	io.Copy(newFile, chunks[0])

	// テキストチャンクを追加
	io.Copy(newFile, textChunk("ASCII PROGRAMING++"))

	// 残りのチャンクを追加
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}
}

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	buffer.Write(byteData)

	// CRCを計算して追加
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	crc.Write(byteData)
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())

	return &buffer
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader

	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}
