package main

import (
	"context"
	"google.golang.org/grpc"
	"microServiceGo/message"
	"net"
)

type OrderServiceImpl struct {
}

func (order *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error) {
	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}
	info := orderMap[request.GetOrderId()]
	return &info, nil
}

func main() {
	server := grpc.NewServer()
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))
	listen, err := net.Listen("tcp", ":8086")
	if err != nil {
		panic(err)
	}
	server.Serve(listen)
}
