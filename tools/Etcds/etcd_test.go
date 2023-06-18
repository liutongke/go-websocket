package Etcds

import (
	"fmt"
	"go-websocket/tools/Timer"
	"log"
	"testing"
	"time"
)

func TestEtcdClient(t *testing.T) {
	log.Println("EtcdConfig:", EtcdConfig())

	key := fmt.Sprintf("/TestEtcd/%d", Timer.GetNowUnix())
	log.Println("Put:", Put(key, Timer.GetNowStr()))

	resp, _ := Get(key)
	for _, item := range resp.Kvs {
		log.Printf("Get:Key: %s, Value: %s\n", item.Key, item.Value)
	}

	log.Println("Del:", Del(key))

	prefix, _ := GetPrefix("/TestEtcd")
	for _, item := range prefix.Kvs {
		log.Printf("GetPrefix:Key: %s, Value: %s\n", item.Key, item.Value)
	}
}
func TestEtcdDiscovery(t *testing.T) {
	//config.InitTest()
	go NewEtcdDiscovery().EtcdStartDiscovery()

}

func TestEtcdRegister(t *testing.T) {
	//config.InitTest()
	go NewEtcdRegister().EtcdStartRegister()

	time.Sleep(1 * time.Second)
	log.Printf("LeaseID:%d", GetEtcdRegister().LeaseID)
	key, err := GetEtcdRegister().PutKey("GetEtcdRegister", "GetEtcdRegister")
	log.Println("GetEtcdRegister", key, err)
	log.Println("GetEtcdRegister", GetEtcdRegister().DelKey("GetEtcdRegister"))

}
