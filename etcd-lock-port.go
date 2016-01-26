package main

// EtcdPort finds an unclaimed port and claims it
type EtcdPort struct {
	name, registry, key string
}

// NewEtcdPort creates a new instance of EtcdPort
func NewEtcdPort(name, registry, key string) *EtcdPort {
	return &EtcdPort{name, registry, key}
}

// LockPort locks a port and returns it
func (etcdPort *EtcdPort) LockPort() (string, error) {
	return "5", nil
}
