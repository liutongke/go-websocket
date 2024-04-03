package etcds

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

type Locker struct {
	client  *clientv3.Client
	session *concurrency.Session
	mutex   *concurrency.Mutex
}

func NewEtcdLocker(key string) (*Locker, error) {
	client, err := clientv3.New(EtcdConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to create etcd client: %v", err)
	}

	session, err := concurrency.NewSession(client)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}

	mutex := concurrency.NewMutex(session, key)

	return &Locker{
		client:  client,
		session: session,
		mutex:   mutex,
	}, nil
}

func (l *Locker) AcquireLock() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := l.mutex.Lock(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire lock: %v", err)
	}

	return nil
}

func (l *Locker) ReleaseLock() error {
	err := l.mutex.Unlock(context.TODO())
	if err != nil {
		return fmt.Errorf("failed to release lock: %v", err)
	}

	return nil
}

func (l *Locker) Close() {
	l.session.Close()
	l.client.Close()
}
