package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
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

	// 选一个服务机（这里选最后一个）
	var addr string
	for k, v := range serviceMap {
		fmt.Printf("k: %s\nv: %#v\n", k, v)
		addr = v.Address + ":" + strconv.Itoa(v.Port)
		fmt.Printf("addr: %s\n", addr)
	}

	// 建立RPC连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()

	// 发送RPC请求
	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req := &pb.HelloRequest{
		Name: "World",
	}
	resp, err := client.SayHello(ctx, req)
	if err != nil {
		fmt.Printf("client.YourMethod failed, err: %v", err)
		return
	}
	fmt.Printf("Response: %v", resp)
}
