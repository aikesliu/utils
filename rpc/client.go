package rpc

import (
	"context"
	"github.com/docker/libkv/store"
	"github.com/pkg/errors"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"time"
)

var (
	dRpc = &defaultRpc{
		rpcClients: make(map[string]client.XClient),
	}
)

type defaultRpc struct {
	rpcClients map[string]client.XClient
	basePath   string
	etcdAddr   string
}

func Init(basePath, etcdAddr string) {
	dRpc.basePath = basePath
	dRpc.etcdAddr = etcdAddr
}

func getRpc(serviceName string) (client.XClient, error) {
	if v, ok := dRpc.rpcClients[serviceName]; ok {
		return v, nil
	}
	optionD := new(store.Config)
	optionD.ConnectionTimeout = 15 * time.Second
	d := client.NewEtcdDiscovery(dRpc.basePath, serviceName, []string{dRpc.etcdAddr}, optionD)
	option := client.DefaultOption
	option.Breaker = nil
	option.SerializeType = protocol.ProtoBuffer
	c := client.NewXClient(serviceName, client.Failfast, client.RandomSelect, d, option)
	if c == nil {
		return nil, errors.Errorf("create rpc: %v failed", serviceName)
	}
	dRpc.rpcClients[serviceName] = c
	return c, nil
}

func Call(service, method string, args, resp interface{}) error {
	c, err := getRpc(service)
	if err != nil {
		return err
	}
	return c.Call(context.Background(), method, args, resp)
}
