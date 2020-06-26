package common

import (
	"go.etcd.io/etcd/pkg/testutil"
	"testing"
)

func TestOpsForGet(t *testing.T) {

	key, val := OpsForGet()("op/limit/cache")
	testutil.AssertNotNil(t, key)
	testutil.AssertNotNil(t, val)

}
