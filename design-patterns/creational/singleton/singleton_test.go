package creational

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
	data string
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	// 使用 once.Do 或者直接比较与 nil 比较
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			// fmt.Printf("%p\n", obj)
			fmt.Printf("%X\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
