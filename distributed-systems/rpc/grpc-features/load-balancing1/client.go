package main

import (
	"context"
	"fmt"
	example "grpc-demo/proto"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type RoundRobinBalancer struct {
	servers []string
	current int
	mu      sync.Mutex
}

func (rr *RoundRobinBalancer) getNextServer() string {
	rr.mu.Lock()
	defer rr.mu.Unlock()
	server := rr.servers[rr.current]
	rr.current = (rr.current + 1) % len(rr.servers)
	return server
}

func main() {
	// 定义多个服务器地址
	servers := []string{
		"localhost:50051",
		"localhost:50052", // 这里假设有多个 gRPC 服务器在运行
	}

	// 创建负载均衡器
	balancer := &RoundRobinBalancer{
		servers: servers,
	}

	// 启动客户端请求
	for i := 0; i < 10; i++ {
		// 获取下一个服务器
		serverAddress := balancer.getNextServer()

		// 连接到选中的服务器
		conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("failed to connect to server %s: %v", serverAddress, err)
		}
		defer conn.Close()

		client := example.NewGreeterClient(conn)

		// 发送请求
		req := &example.HelloRequest{Name: fmt.Sprintf("Client %d from Server %s", i, serverAddress)}
		resp, err := client.SayHello(context.Background(), req)
		if err != nil {
			log.Printf("Error calling SayHello: %v", err)
			continue
		}
		fmt.Printf("Response from server %s: %s\n", serverAddress, resp.GetMessage())

		// 模拟延迟
		time.Sleep(500 * time.Millisecond)
	}
}
