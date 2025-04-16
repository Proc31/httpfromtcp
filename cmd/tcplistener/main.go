package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

func catch(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:42069")
	catch(err)
	defer listener.Close()
	fmt.Println("Listening on:", listener.Addr())

	for {
		conn, err := listener.Accept()
		catch(err)
		fmt.Println("Connection acccepted from", conn.RemoteAddr())
		linesChan := getLinesChannel(conn)
		for lines := range linesChan {
			fmt.Println(lines)
		}
		fmt.Println("Connection to", conn.RemoteAddr(), "closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {

	messages := make(chan string)

	go func() {
		defer f.Close()
		defer close(messages)
		currentLine := ""
		for {
			buffer := make([]byte, 8)
			n, err := f.Read(buffer)
			if err != nil {
				if currentLine != "" {
					messages <- currentLine
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			str := string(buffer[:n])
			parts := strings.Split(str, "\n")
			for i := range len(parts) - 1 {
				messages <- currentLine + parts[i]
				currentLine = ""
			}
			currentLine += parts[len(parts)-1]
		}
	}()

	return messages
}
