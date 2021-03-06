package go_hello_world

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func TestService(t *testing.T) {
	// 串行
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	// 无缓冲，必须有人消费了才能释放阻塞
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	asyncService := AsyncService()
	otherTask()
	fmt.Println(<-asyncService)
	time.Sleep(time.Second * 1)
}

func BufferedAsyncService() chan string {
	// 加入缓冲
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestBufferedAsyncService(t *testing.T) {
	asyncService := BufferedAsyncService()
	otherTask()
	fmt.Println(<-asyncService)
	time.Sleep(time.Second * 1)
}
