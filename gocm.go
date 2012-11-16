package main

import (
	"flag"
	"fmt"
	"net"
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

	err := startServer(*portFlag)
	if err != nil {
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

}
