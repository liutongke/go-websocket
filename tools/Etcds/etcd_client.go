package etcds

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
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

// Put 添加一个key
func Put(key, value string) error {
	_, err := GetInstance().Put(context.Background(), key, value)
	return err
}

// Get 获取一个key
func Get(key string) (resp *clientv3.GetResponse, err error) {
	resp, err = GetInstance().Get(context.Background(), key)
	return resp, err
}

// GetPrefix 通过前缀获取key
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

// Del 删除key
func Del(key string) (int, error) {
	kv := clientv3.NewKV(GetInstance())
	resp, err := kv.Delete(context.Background(), key)
	if err != nil {
		return 0, fmt.Errorf("del error")
	}
	return int(resp.Deleted), nil
}
