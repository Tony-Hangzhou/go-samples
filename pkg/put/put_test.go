package put

import (
	"go.etcd.io/etcd/pkg/testutil"
	"os"
	"testing"
)

/**
TestMain名字 和 testing.M 必须匹配，其他测试用例Test开头配合使用testing.T；
优先被执行，类似 @before 方法
*/

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestPut(t *testing.T) {

	key, value := PutAndGet("k1", "v2")
	testutil.AssertNotNil(t, key)
	testutil.AssertNotNil(t, value)

}
