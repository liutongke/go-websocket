package Etcds

import (
	"encoding/json"
	"fmt"
	"go-websocket/config"
	"go-websocket/tools/Timer"
	"go-websocket/tools/Tools"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"testing"
	"time"
)

func TestEtcdClient(t *testing.T) {
	config.InitTest()
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
	config.InitTest()
	go NewEtcdDiscovery(map[string]FunDiscovery{"put": EventPut, "del": EventDel}).EtcdStartDiscovery([]string{"/prefix1", "/prefix2", "/net", "go-nat-x", ETCD_SERVER_LIST, ETCD_PREFIX_ACCOUNT_INFO})

}

func EventPut(event *clientv3.Event) {
	log.Printf("watch put test---------->key:%q val:%q", event.Kv.Key, event.Kv.Value)
}
func EventDel(event *clientv3.Event) {
	log.Printf("watch del test---------->key:%q val:%q", event.Kv.Key, event.Kv.Value)
}
func TestEtcdRegister(t *testing.T) {
	config.InitTest()
	go NewEtcdRegister().EtcdStartRegister(RegisterServer)

	time.Sleep(1 * time.Second)
	log.Printf("LeaseID:%d", GetEtcdRegister().LeaseID)
	key, err := GetEtcdRegister().PutKey("GetEtcdRegister", "GetEtcdRegister")
	log.Println("GetEtcdRegister", key, err)
	log.Println("GetEtcdRegister", GetEtcdRegister().DelKey("GetEtcdRegister"))

}

type ServerInfo struct {
	ServerIp string `json:"server-ip"`
	Rpcport  string `json:"rpc-port"`
	Tm       string `json:"tm"`
}

func RegisterServer(e *EtcdRegister) {
	key := fmt.Sprintf("%s%s:nat-x", ETCD_SERVER_LIST, Tools.GetLocalIp())

	info := ServerInfo{
		ServerIp: Tools.GetLocalIp(),
		Rpcport:  "nat-x",
		Tm:       Timer.GetNowStr(),
	}
	// 将Person对象转换为JSON字符串
	val, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	e.PutKey(key, string(val))
}
