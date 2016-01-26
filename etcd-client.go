package main

import (
	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/client"
)

// EtcdClient lets your Get/Set from Etcd
type EtcdClient struct {
	etcd client.Client
}

// NewEtcdClient constructs a new EtcdClient
func NewEtcdClient() (*EtcdClient, error) {
	etcd, err := client.New(client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
	})
	if err != nil {
		return nil, err
	}
	return &EtcdClient{etcd}, nil
}

// Get gets a value in Etcd
func (etcdClient *EtcdClient) Get(key string) (string, error) {
	api := client.NewKeysAPI(etcdClient.etcd)
	response, err := api.Get(context.Background(), key, nil)
	if err != nil {
		if client.IsKeyNotFound(err) {
			return "", nil
		}
		return "", err
	}
	return response.Node.Value, nil
}

// Set sets a value in Etcd
func (etcdClient *EtcdClient) Set(key, value string) error {
	api := client.NewKeysAPI(etcdClient.etcd)
	_, err := api.Set(context.Background(), key, value, nil)
	return err
}
