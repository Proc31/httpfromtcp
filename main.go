package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func catch(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./messages.txt")
	catch(err)
	defer f.Close()

	for msg := range getLinesChannel(f) {
		fmt.Println("read: " + msg)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {

	messages := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		wg.Wait()
		close(messages)
	}()

	go readByBytes(messages, &wg, f)

	return messages
}

func readByBytes(ch chan string, wg *sync.WaitGroup, f io.ReadCloser) {
	defer wg.Done()

	currentLine := ""

	for {
		buffer := make([]byte, 8)
		n, err := f.Read(buffer)
		if err != nil {
			if currentLine != "" {
				ch <- currentLine
				currentLine = ""
			}
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}
		str := string(buffer[:n])
		parts := strings.Split(str, "\n")
		for i := range len(parts) - 1 {
			ch <- currentLine + parts[i]
			currentLine = ""
		}
		currentLine += parts[len(parts)-1]
	}
}
