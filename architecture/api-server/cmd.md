# 一 概述:
## (1)格式:
- kube-apiserver [flags]

## (2)常用选项:
- --advertise-address ip: The IP address on which to advertise the apiserver to members of the cluster, 该地址必须是集群内可达的, 若不指定, 则使用--bind-address的地址, 若--bind-address没有指定, 则使用host的默认接口.
- --bind-address: 监听--secure-port的IP地址, 该接口必须是集群内可达, 以及CLI和web用户可达, 若为空, 则所有接口会被使用(ipv4的0.0.0.0和ipv6的::).
- --secure-port: 默认**6443**, 用来提供带authentication和authorization的HTTPS服务.
- --apiserver-count: 集群中apiserver的运行数量, 默认为1.
- --service-cluster-ip-range: 一个分配服务cluster ips的CIDR IP范围, 不能和分配给Pod的ip overlap, 默认为:10.0.0.0/24.

## (3)日志相关:
- --log-dir: 若非空, 写入该目录下日志文件.
- --log-file: 若非空, 使用该日志文件.
- --log-flush-frequency: 默认5s, 日志刷新的最大时间.
- --logtostderr: 默认为true, 将日志写入stderr而不是文件, 默认为true.
- --alsologtostderr: 将日志同时写入stderr和文件.

## (4)认证相关
- --enable-bootstrap-token-auth
- --token-auth-file
- --authentication-token-webhook-config-file: token认证的webhook配置文件, API server会查询远程服务来认证.

## (5)授权相关:
- --authorization-mode: 默认AlwaysAllow, 可选: AlwaysAllow, AlwaysDeny, ABAC, Webhook, RBAC, Node.
- --authorization-webhook-config-file: 当mode为webhook时使用的webhook配置(kubeconfig format), API server会查询远程服务来进行判断.

## (6)admission相关:
- --admission-control-config-file
- --enable-admission-plugins stringSlice: 除了默认开启的admission插件外需开启admission插件.
- --disable-admission-plugins stringSlice: 需关闭admission插件, 即使在默认开启的列表.
- 备注: 查询默认开启的admission controller: kube-apiserver -h | grep enable-admission-plugins

## (7)审计(audit)相关:
- --audit-webhook-config-file

## (8)storage相关:
- --etcd-servers: 可连接的etcd服务器的地址, 格式:scheme://ip:port, 逗号分隔.
- --etcd-prefix: 默认/registry, 在etcd中存储所有资源的prefix.
- --etcd-compaction-interval: The interval of compaction requests, 默认5m.
- -–storage-backend: 持久化存储的backend, 默认etcd3.
- --storage-media-type: 默认application/vnd.kubernetes.protobuf.
- --default-watch-cache-size: 默认watch的cache大小, 默认为100; 0表示disable watch cache.
- --watch-cache: 默认为true.
- --watch-cache-sizes: 指定资源的watch size的配置.

## (9)请求相关:
- --min-request-timeout: 默认1800, 目前只适用于watch handler.
- --request-timeout

