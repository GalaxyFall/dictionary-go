package db

import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	config  clientv3.Config
	client  *clientv3.Client
	err     error
	kv      clientv3.KV
	getResp *clientv3.GetResponse
)
