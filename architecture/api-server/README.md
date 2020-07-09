# 一 概述:
## (1)功能:
- 暴露Kubernetes API, 是Kubernetes的**控制面板**的前端.
- k8s系统组件间只能通过API服务器通信, 它们之间不会直接通信.
- API服务器是和etcd通信的唯一组件, 其它组件都通过api服务器来修改集群状态.
- API server提供一致性的对象访问, 也对对象进行校验(避免存入非法对象), 同时提供支持并发更新的乐观锁.

## (2)访问方式:
- 各种client库.
- kubectl proxy命令.
- Pod内访问: 使用每个容器内挂载的secret目录下的ca.crt,token和namespace来访问(例如:curl); 使用ambassador容器, 在该容器内执行kubectl proxy命令.

## (3)优化点:
- 自定义list-watch的store来减少存储数据的大小来避免内存占用过多.
- 分片list-watch, 实现参考: https://github.com/kubernetes/kube-state-metrics

# 二 kube-apiserver命令:
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
- --enable-admission-plugins stringSlice: 开启admission controller.
- --disable-admission-plugins stringSlice: 关闭admission controller.
- 备注: 查询默认开启的admission controller: kube-apiserver -h | grep enable-admission-plugins

## (7)审计(audit)相关:
- --audit-webhook-config-file

## (8)storage相关:
- --etcd-servers: 可连接的etcd服务器的地址, 格式:scheme://ip:port, 逗号分隔.
- --etcd-prefix: 默认/registry, 在etcd中存储所有资源的prefix.
- -–storage-backend: 持久化存储的backend, 默认etcd3.
- --storage-media-type: 默认application/vnd.kubernetes.protobuf.
- --default-watch-cache-size: 默认watch的cache大小, 默认为100; 0表示disable watch cache.
- --watch-cache: 默认为true.
- --watch-cache-sizes: 指定资源的watch size的配置.
