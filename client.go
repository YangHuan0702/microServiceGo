package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"microServiceGo/message"
	"time"
)

func main() {
	dial, err := grpc.Dial("localhost:8086", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	client := message.NewOrderServiceClient(dial)

	orderRequest := &message.OrderRequest{OrderId: "201907300001", TimeStamp: time.Now().Unix()}
	info, err := client.GetOrderInfo(context.Background(), orderRequest)
	if nil != info {
		fmt.Printf("rpc success! orderName:%s", info.OrderName)
	}
}
