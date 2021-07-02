package socket

import (
	"fmt"
	"math/rand"
)

type SocketPool struct {
	pools []chan string
}

func NewPool() *SocketPool {
	return &SocketPool{
		pools: make([]chan string, 10),
	}
}

func (s *SocketPool) StartPool() {
	for i := 0; i < 10; i++ {
		s.pools[i] = make(chan string, 2)
		go s.StartOneWorker(i, s.pools[i])
	}
}

//StartOneWorker 启动一个Worker工作流程
func (s *SocketPool) StartOneWorker(workerID int, taskQueue chan string) {
	fmt.Println("Worker ID = ", workerID, " is started.")
	//不断的等待队列中的消息
	for {
		select {
		//有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			fmt.Println("接收到的任务信息：", workerID, request)
		}
	}
}

func (s *SocketPool) SendData(data string) {
	i := rand.Intn(9)
	s.pools[i] <- data
}
