package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	n, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer n.Close()

	for {
		conn, err := n.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connection accepted...")
		for line := range getLinesChannel(conn) {
			fmt.Println(line)
		}
		fmt.Println("...connection has been closed.")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer f.Close()
		defer close(lines)

		currentLine := ""
		for {
			buff := make([]byte, 8)
			n, err := f.Read(buff)
			if err != nil {
				if currentLine != "" {
					lines <- currentLine
				}
				if errors.Is(err, io.EOF) {
					break
				}
				log.Fatal(err)
			}

			parts := strings.Split(string(buff[:n]), "\n")
			for i := range len(parts) - 1 {
				lines <- currentLine + parts[i]
				currentLine = ""
			}
			currentLine += parts[len(parts)-1]
		}
	}()

	return lines
}
