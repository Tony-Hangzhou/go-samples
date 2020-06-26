package put

import (
	"context"
	"go-samples/pkg/common"
	"go.etcd.io/etcd/clientv3"
)

func PutAndGet(k string, v string) (*string, *string) {

	common.OpsForFun(func(cli *clientv3.Client, ctx context.Context) {
		cli.Put(ctx, k, v)
	})

	return common.OpsForGet()(k)
}
