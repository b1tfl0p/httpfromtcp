package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buff := make([]byte, 8)
	currentLine := ""
	for {
		n, err := file.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		parts := strings.Split(string(buff[:n]), "\n")
		currentLine += parts[0]

		if len(parts) == 2 {
			fmt.Printf("read: %s\n", currentLine)
			currentLine = parts[1]
		}
	}

	if len(currentLine) > 0 {
		fmt.Printf("read: %s\n", currentLine)
	}
}
