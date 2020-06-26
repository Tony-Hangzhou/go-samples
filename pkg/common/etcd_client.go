package common

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func InstanceClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return nil, err
	}
	fmt.Println("connect to etcd success")
	return cli, err
}

/**
通过 func 入参，提供简化、灵活的Etcd操作
*/
func OpsForFun(f func(cli *clientv3.Client, ctx context.Context)) {
	cli, _ := InstanceClient()
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	f(cli, ctx)
	defer cancel()
}

/**
封装Get操作，通过返回func
*/
func OpsForGet() func(key string) (*string, *string) {
	return func(key string) (*string, *string) {
		cli, err := InstanceClient()
		defer cli.Close()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, key)
		defer cancel()
		if err != nil {
			fmt.Printf("get from etcd failed, err:%v\n", err)
			return nil, nil
		}
		for _, ev := range resp.Kvs {
			//ev.String();
			fmt.Printf("KeyValue: %s\n", ev.String())
			key := string(ev.Key[:])
			value := string(ev.Value[:])
			return &key, &value
		}
		return nil, nil
	}
}
