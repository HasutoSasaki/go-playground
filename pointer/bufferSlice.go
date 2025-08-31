package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const fileName = "data/sample.txt"

func main() {
	err := readFile()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func readFile() error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	data := make([]byte, 100)
	for {
		count, err := file.Read(data)
		process(data[:count]) // 読み込んだデータ量
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
	}
}

func process(data []byte) {
	fmt.Printf("読み込んだデータ: %s", string(data))
}
