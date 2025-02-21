package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver" // 这个很重要，用于注册 consul resolver
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {
	// 连接consul
	cc, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("api.NewClient failed, err:%v\n", err)
		return
	}
	// 返回的是一个 map[string]*api.AgentService
	// 其中key是服务ID，值是注册的服务信息
	serviceMap, err := cc.Agent().ServicesWithFilter("Service matches `hello-*`")
	if err != nil {
		fmt.Printf("query service from consul failed, err:%v\n", err)
		return
	}

	fmt.Printf("serviceMap %+v\n", serviceMap)

	// 建立RPC连接
	conn, err := grpc.Dial(
		// consul服务
		"consul://127.0.0.1:8500/hello?healthy=true",
		// 指定round_robin策略
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Printf("c.SayHello failed, err:%v\n", err)
		return
	}
	defer conn.Close()

	// 发送RPC请求
	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req := &pb.HelloRequest{
		Name: "gzhh",
	}
	resp, err := client.SayHello(ctx, req)
	if err != nil {
		fmt.Printf("client.YourMethod failed, err: %v", err)
		return
	}
	fmt.Printf("Response: %v", resp)
}
