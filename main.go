package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func catch(e error) {
	if e != nil {
		panic(e)
	}
}

func main(){
	f, err := os.Open("./messages.txt")
	catch(err)
	defer f.Close()
	getLinesChannel(f)
}

func getLinesChannel(f io.ReadCloser) <-chan string {

	messages := make(chan string)

	currentLineContents := ""
	for {
		buffer := make([]byte, 8)
		n, err := f.Read(buffer)



		if err != nil {
            		if currentLineContents != "" {
				fmt.Printf("read: %s\n", currentLineContents)
				currentLineContents = ""
		}
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}
		str := string(buffer[:n])
		parts := strings.Split(str, "\n")
		for i := 0; i < len(parts)-1; i++ {
			fmt.Printf("read: %s%s\n", currentLineContents, parts[i])
			currentLineContents = ""
		}
		currentLineContents += parts[len(parts)-1]

	}
	return messages
}
