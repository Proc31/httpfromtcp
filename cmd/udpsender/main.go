package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func catch(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	server, err := net.ResolveUDPAddr("udp4", "localhost:42069")
	catch(err)
	conn, err := net.DialUDP("udp4", nil, server)
	catch(err)
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("UDP Sender started on", conn.RemoteAddr().String())
	for {
		fmt.Print(">")
		str, err := reader.ReadString('\n')
		catch(err)
		conn.Write([]byte(str))
	}
}
