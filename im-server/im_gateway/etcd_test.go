package main

import (
	"context"
	"fmt"
	"im-server/core"
	"testing"
)

func TestReadEtcd(t *testing.T) {
	kv := core.InitEtcd()

	kv.Put(context.TODO(), "yzy", "9999")

	resp, _ := kv.Get(context.TODO(), "yzy")
	val := string(resp.Kvs[0].Value)
	fmt.Println(val)
}

func TestWriteEtcd(t *testing.T) {

}
