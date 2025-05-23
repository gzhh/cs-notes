# Load Balancing
- https://en.wikipedia.org/wiki/Load_balancing_(computing)
- What Is Load Balancing? https://www.nginx.com/resources/glossary/load-balancing/

## Nginx Config Load Balancing
- Nginx负载均衡配置 https://blog.csdn.net/xyang81/article/details/51702900

## Example
- https://hackernoon.com/this-150-line-go-script-is-actually-a-full-on-load-balancer

### Implement Load Balancer with Go.

```
package main

import (
	"fmt"
	"math/rand"
)

type LoadBalancer struct {
	size    int
	clients []*Client
}

func NewLoadBalancer(size int) *LoadBalancer {
	loadBalancer := &LoadBalancer{
		size:    size,
		clients: make([]*Client, size),
	}

	for i := 0; i < size; i++ {
		loadBalancer.clients[i] = &Client{
			ID:   i,
			Name: fmt.Sprintf("client_%d", i),
		}
	}

	return loadBalancer
}

func (lb *LoadBalancer) GetClient() *Client {
	return lb.clients[rand.Intn(len(lb.clients))]
}

type Client struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Client) Do() {
	fmt.Printf("client %d: doing \n", c.ID)
}

func main() {
	lb := NewLoadBalancer(4)
	for i := 0; i < 10; i++ {
		c := lb.GetClient()
		c.Do()
	}
}
```
