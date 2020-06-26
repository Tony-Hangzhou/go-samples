package watch

import (
	"context"
	"fmt"
	"go-samples/pkg/common"
	"go.etcd.io/etcd/clientv3"
)

func Watch(k string) {

	common.OpsForFun(func(cli *clientv3.Client, ctx context.Context) {
		ch := cli.Watch(context.Background(), k)
		// 监视通道
		for wresp := range ch {
			for _, evt := range wresp.Events {
				fmt.Printf("Type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
			}
		}
	})

}
