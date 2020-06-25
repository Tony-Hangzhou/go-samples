package put

import (
	"context"
	"fmt"
	"go-samples/pkg/common"
	"time"
)

func PutAndGet() (*string, *string) {
	cli, err := common.InstanceClient()
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		// nil只能赋值给指针、channel、func、interface、map或slice类型的变量。如果未遵循这个规则，则会引发panic。
		return nil, nil
	}
	fmt.Println("connect to etcd success")

	defer cli.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := "100"
	_, err = cli.Put(ctx, "op/limit/cache", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return nil, nil
	}

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "op/limit/cache")
	cancel()
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
