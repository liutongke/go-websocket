package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Pool struct {
	pool []chan *SendData
}

var (
	MaxPool    = 10 //消费者最大数量
	capacity   = 2  //队列容量
	Wg         sync.WaitGroup
	WgSendData sync.WaitGroup
)

func NewPool() *Pool {
	return &Pool{pool: make([]chan *SendData, MaxPool)}
}

// 生成工作work
func (p *Pool) startPool() {
	for i := 0; i < MaxPool; i++ {
		p.pool[i] = make(chan *SendData, capacity)
		go p.startOneWork(i, p.pool[i])
	}
}

// 创建工作work
func (p *Pool) startOneWork(workerID int, taskQueue chan *SendData) {
	fmt.Println("Worker ID = ", workerID, " is started.")
	Wg.Done()
	//不断的等待队列中的消息
	for {
		select {
		//有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			fmt.Printf("接收到的任务信息：workdId:%d,数据：%d\n", workerID, request.Data)
			time.Sleep(1 * time.Second)
			WgSendData.Done()
		}
	}
}

type SendData struct {
	Data int
}

// 生产者
func (p *Pool) SendToWork(data *SendData) {
	i := rand.Intn(9)
	p.pool[i] <- data
}

func main() {
	Wg.Add(MaxPool)
	pool := NewPool()
	pool.startPool()

	Wg.Wait()

	for i := 0; i <= 100; i++ {
		WgSendData.Add(1)

		//生产者，投递任务
		pool.SendToWork(&SendData{
			Data: i,
		})
	}

	WgSendData.Wait()
}
