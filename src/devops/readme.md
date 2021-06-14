# devops

- 使用 podman 取代 docker

```sh
alias docker=podman
podman run -d -p 6379:6379 -v /root/devops/redis.conf:/etc/redis/redis.conf --name redis redis:alpine
podman run -d -p 3306:3306 -e TZ='Asia/Shanghai' -e MYSQL_ROOT_PASSWORD=root --name mysql mysql
podman run -d -p 2379:2379 -e ETCD_ADVERTISE_CLIENT_URLS='http://0.0.0.0:2379' -e ETCD_LISTEN_CLIENT_URLS='http://0.0.0.0:2379' -e ETCDCTL_API='3' --name etcd quay.io/coreos/etcd
```