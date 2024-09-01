// 并发执行
package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个channel 用于协程通信 这里是主线程与协程之间的通信
	// 创建一个channel，通信的数据类型为字符串类型
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "sheep"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "cow"
			time.Sleep(time.Millisecond * 2000)
		}
	}()
	 
	// 从channel中收听信息 使用select解决阻塞问题
	// 当有多个channel发送消息的时候，使用select保证第一时间接收channel中的数据，而不需要等待其他channel的顺序执行阻塞时间
	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}
