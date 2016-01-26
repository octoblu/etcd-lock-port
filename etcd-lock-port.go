package main

// EtcdLockPort finds an unclaimed port and claims it
type EtcdLockPort struct {
	name, registry, key string
}

// New creates a new instance of EtcdLockPort
func New(name, registry, key string) *EtcdLockPort {
	return &EtcdLockPort{name, registry, key}
}

// LockPort locks a port and returns it
func (etcdLockPort *EtcdLockPort) LockPort() (string, error) {
	return "5", nil
}
