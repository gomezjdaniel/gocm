package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"
)

var apiKeyFlag = flag.String("apikey", "", "google api key for gcm")
var portFlag = flag.Int("port", 0, "listening port")

func main() {
	flag.Parse()

	if *apiKeyFlag == "" || *portFlag == 0 {
		fmt.Println("argument required")
		flag.Usage()
		return
	}

	if *portFlag < 8000 {
		fmt.Println("can't start server. port must be >= 8000")
		return
	}

	if err := startServer(*portFlag); err != nil {
		fmt.Println("can't start server")
		return
	}

}

func startServer(port int) error {
	portString := fmt.Sprintf(":%d", *portFlag)
	l, err := net.Listen("tcp", portString)
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go msgReceiver(conn)
	}
	return nil
}

func msgReceiver(conn net.Conn) {

	buf := bufio.NewReader(conn)
	line, isPrefix, err := buf.ReadLine()

	if err != nil {
		fmt.Println(err)
		return
	}

	if isPrefix != false {
		fmt.Println("line too long")
		return
	}

	body := string(line)
	parts := strings.Split(body, ":")

	_ = parts
}
