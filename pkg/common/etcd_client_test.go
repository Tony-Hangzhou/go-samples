package common

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/pkg/testutil"
	"golang.org/x/net/context"
	"testing"
)

func TestOpsForGet(t *testing.T) {

	key, val := OpsForGet()("op/limit/cache")
	testutil.AssertNotNil(t, key)
	testutil.AssertNotNil(t, val)

}

func TestOpsForFun(t *testing.T) {

	OpsForFun(func(cli *clientv3.Client, ctx context.Context) {

		resp, _ := cli.Get(ctx, "op/limit/cache")
		fmt.Printf("Kvs is %s", resp.Kvs)
		testutil.AssertNotNil(t, resp.Kvs)
	})
}
