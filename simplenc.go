package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", os.Args[1]+":"+os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	for {
		go copy(conn, os.Stdout)
		copy(os.Stdin, conn)
	}
}

func copy(src io.Reader, dst io.Writer) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
