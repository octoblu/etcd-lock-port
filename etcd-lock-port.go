package main

import (
	"fmt"
	"math/rand"
	"time"
)

// EtcdLockPort finds an unclaimed port and claims it
type EtcdLockPort struct {
	name, registry, key string
	etcd                *EtcdClient
}

// New creates a new instance of EtcdLockPort
func New(name, registry, key string) (*EtcdLockPort, error) {
	etcd, err := NewEtcdClient()
	if err != nil {
		return nil, err
	}
	return &EtcdLockPort{name, registry, key, etcd}, nil
}

// LockPort locks a port and returns it
func (etcdLockPort *EtcdLockPort) LockPort() (string, error) {
	port, err := etcdLockPort.getExistingLock()
	if err != nil {
		return "", err
	}
	if port != "" {
		return port, nil
	}
	return etcdLockPort.lockNewPort()
}

func (etcdLockPort *EtcdLockPort) getExistingLock() (string, error) {
	etcd := etcdLockPort.etcd
	port, err := etcd.Get(etcdLockPort.key)
	if err != nil {
		return "", err
	}

	if port == "" {
		return "", nil
	}

	registryKey := fmt.Sprintf("%v/%v", etcdLockPort.registry, port)
	whoGotTheLock, err := etcd.Get(registryKey)
	if err != nil {
		return "", err
	}
	if whoGotTheLock != etcdLockPort.name {
		return "", nil
	}
	return port, nil
}

func (etcdLockPort *EtcdLockPort) lockNewPort() (string, error) {
	port := randomPort()
	registryKey := fmt.Sprintf("%v/%v", etcdLockPort.registry, port)

	etcd := etcdLockPort.etcd
	err := etcd.Set(registryKey, etcdLockPort.name)
	if err != nil {
		return "", err
	}

	whoGotTheLock, err := etcd.Get(registryKey)
	if err != nil {
		return "", err
	}

	if whoGotTheLock != etcdLockPort.name {
		return etcdLockPort.lockNewPort()
	}

	etcd.Set(etcdLockPort.key, port)
	return port, nil
}

func randomPort() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v", (20000 + rand.Intn(45000)))
}
