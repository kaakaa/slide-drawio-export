package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net/url"
	"os"
)

func parse(r io.Reader) (err error) {
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(r)
	if err != nil {
		return err
	}

	//　PNGシグネチャの読み込み
	if string(buffer.Next(8)) != "\x89PNG\r\n\x1a\n" {
		return fmt.Errorf("not a PNG")
	}

	// IDATチャンクの読み込み
	data := make([]byte, 0, 32)
	loop := true
	for loop {
		length := int(binary.BigEndian.Uint32(buffer.Next(4)))
		chunkType := string(buffer.Next(4))

		switch chunkType {
		case "tEXt":
			fmt.Println("chunk: tEXt")
			data = append(data, buffer.Next(length)...)
			_ = buffer.Next(4) // CRC
		case "IEND":
			fmt.Println("chunk: IEND")
			loop = false
		default:
			fmt.Println("chunk:", chunkType)
			_ = buffer.Next(length) // chunk data
			_ = buffer.Next(4)      // CRC
		}
	}

	text, _ := url.QueryUnescape(string(data))
	fmt.Println("tEXt chunk: " + text)

	return nil
}

func main() {
	imageFile := "test.drawio.png"
	file, err := os.Open(imageFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = parse(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Complete")
}
