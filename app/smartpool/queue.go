package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("====== Start transaction pool queue ========")
	sev, err := net.Listen("tcp", ":7550")

	if err != nil {
		log.Fatal(err)
	}

	cnn, _ := sev.Accept()

	for {
		msg, _ := bufio.NewReader(cnn).ReadString('\n')
		fmt.Println(msg)
	}

}
