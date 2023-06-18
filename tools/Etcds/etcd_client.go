package Etcds

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"log"
	"sync"
)

var (
	etcdKvClient *clientv3.Client
	mu           sync.Mutex
)

func GetInstance() *clientv3.Client {
	if etcdKvClient == nil {
		if client, err := clientv3.New(EtcdConfig()); err != nil {
			log.Println(err)
			return nil
		} else {
			//创建时才加锁
			mu.Lock()
			defer mu.Unlock()
			etcdKvClient = client
			return etcdKvClient
		}

	}
	return etcdKvClient
}

// 添加一个key
func Put(key, value string) error {
	_, err := GetInstance().Put(context.Background(), key, value)
	return err
}

func Get(key string) (resp *clientv3.GetResponse, err error) {
	resp, err = GetInstance().Get(context.Background(), key)
	return resp, err
}

func GetPrefix(prefix string) (*clientv3.GetResponse, error) {
	kv := clientv3.NewKV(GetInstance())
	// 使用 Get 方法获取指定前缀的键值对
	resp, err := kv.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	//for _, item := range resp.Kvs {
	//	fmt.Printf("Key: %s, Value: %s\n", item.Key, item.Value)
	//}
	return resp, err
}

func Del(key string) int {
	kv := clientv3.NewKV(GetInstance())
	resp, err := kv.Delete(context.Background(), key)
	if err != nil {
		return 0
	}
	return int(resp.Deleted)
}
