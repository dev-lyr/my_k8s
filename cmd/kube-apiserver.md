# 一kube-apiserver命令:
## (1)格式:
- kube-apiserver [flags]

## (2)常用选项:
- --advertise-address ip: The IP address on which to advertise the apiserver to members of the cluster, 该地址必须是集群内可达的, 若不指定, 则使用--bind-address的地址, 若--bind-address没有指定, 则使用host的默认接口.
- --bind-address: 监听--secure-port的IP地址, 该接口必须是集群内可达, 以及CLI和web用户可达, 若为空, 则所有接口会被使用(ipv4的0.0.0.0和ipv6的::).
- --secure-port: 默认**6443**, 用来提供带authentication和authorization的HTTPS服务.
- --apiserver-count: 集群中apiserver的运行数量, 默认为1.
- --etcd-servers: 可连接的etcd服务器的地址, 格式:scheme://ip:port, 逗号分隔.
- --service-cluster-ip-range: 一个分配服务cluster ips的CIDR IP范围, 不能和分配给Pod的ip overlap, 默认为:10.0.0.0/24.

## (3)日志相关:
- --log-dir: 若非空, 写入该目录下日志文件.
- --log-file: 若非空, 使用该日志文件.
- --log-flush-frequency: 默认5s, 日志刷新的最大时间.
- --logtostderr: 默认为true, 将日志写入stderr而不是文件, 默认为true.
- --alsologtostderr: 将日志同时写入stderr和文件.
