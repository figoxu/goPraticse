docker exec $(docker ps -a --filter name=etcd -q) etcdctl put root/service/svc1/key1 val1
docker exec $(docker ps -a --filter name=etcd -q) etcdctl get root/service/svc1/key1