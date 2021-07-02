package test

import (
	"go-websocket/app/services/socket"
	"go-websocket/utils/Timer"
	"math/rand"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	p := socket.NewPool()
	p.StartPool()
	for i := 0; i < 10; i++ {
		p.SendData(Timer.NowStr())
	}
	time.Sleep(10)
}
