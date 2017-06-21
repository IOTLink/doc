package main

import (
	"log"
	"google.golang.org/grpc"
	pb "streamserver/protocol"
	"golang.org/x/net/context"
	"encoding/json"
)

const (
	serveraddr = "127.0.0.1:50055"
)

func main() {
	conn, err := grpc.Dial(serveraddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStreamServerClient(conn)

	regInfo := &pb.RegisterReq{User:"client", Pwd:"123456"}
	msg, err := client.RegisterClient(context.Background(), regInfo)
	if err != nil {
		log.Fatalf("msg:%s,%v",msg.Message, err)
	}
	if msg != nil {
		appInfo, _ := json.Marshal(msg)
		log.Printf("appInfo: %s", appInfo)
	}

}
