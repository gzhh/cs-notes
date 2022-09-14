package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func dog(dogChan chan bool, catChan chan bool) {
	for {
		select {
		case <-dogChan:
			fmt.Println("dog")
			catChan <- true
		default:
			break
		}
	}
}

func cat(catChan chan bool, fishChan chan bool) {
	for {
		select {
		case <-catChan:
			fmt.Println("cat")
			fishChan <- true
		default:
			break
		}
	}
}

func fish(fishChan chan bool, dogChan chan bool) {
	i := 1
	for {
		select {
		case <-fishChan:
			fmt.Println("fish")
			i++ // 计数，打印完之后就溜溜结束了。
			if i > 3 {
				wg.Done()
				return
			}
			dogChan <- true
		default:
			break
		}
	}
}

// 利用goroutine和channel 连续输出3次，dog，cat，fish，并且都要按照这个dog，cat，fish的顺序输出。
func main() {
	dogChan, catChan, fishChan := make(chan bool), make(chan bool), make(chan bool)
	wg.Add(1)
	go dog(dogChan, catChan)
	go cat(catChan, fishChan)
	go fish(fishChan, dogChan)
	dogChan <- true // 记得这里进行启动条件，不然就没法启动了。
	wg.Wait()
}
