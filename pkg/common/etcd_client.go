package common

import (
	"go.etcd.io/etcd/clientv3"
	"time"
)

func InstanceClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	return cli, err
}
