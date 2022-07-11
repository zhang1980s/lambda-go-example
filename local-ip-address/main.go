package main 

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"

	"log"
	"net"
	"strings"
)

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

func main() {
	lambda.Start(GetOutboundIP)
}