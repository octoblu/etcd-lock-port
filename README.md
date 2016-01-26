# etcd-lock-port
establish a lock on a port using etcd

## Usage

```bash
etcd-lock-port \
  --service-name my-service-1 \
  --port-registry /cluster/ports \
  --service-key /octoblu/my-service/instances/1/port
```
