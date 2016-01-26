package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/client"
)

// EtcdLockPort finds an unclaimed port and claims it
type EtcdLockPort struct {
	name, registry, key string
	etcd                client.Client
}

// New creates a new instance of EtcdLockPort
func New(name, registry, key string) *EtcdLockPort {
	var etcd client.Client
	return &EtcdLockPort{name, registry, key, etcd}
}

// Connect connects the EtcdLockPort instance to Etcd
func (etcdLockPort *EtcdLockPort) Connect() error {
	etcd, err := client.New(client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
	})
	etcdLockPort.etcd = etcd
	return err
}

// LockPort locks a port and returns it
func (etcdLockPort *EtcdLockPort) LockPort() (string, error) {
	rand.Seed(time.Now().UnixNano())
	port := fmt.Sprintf("%v", (20000 + rand.Intn(45000)))
	registryKey := fmt.Sprintf("%v/%v", etcdLockPort.registry, port)

	api := client.NewKeysAPI(etcdLockPort.etcd)
	api.Set(context.Background(), registryKey, etcdLockPort.name, nil)

	whoGotTheLock, err := api.Get(context.Background(), registryKey, nil)
	if err != nil {
		return "", err
	}

	if whoGotTheLock.Node.Value == etcdLockPort.name {
		return port, nil
	}

	return etcdLockPort.LockPort()
}
