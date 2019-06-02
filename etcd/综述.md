# 一 概述:
## (1)概述:
- etcd是一个可靠和快速的分布式kv存储, 高可用和持久化数据存储和获取.
- etcd使用go实现, 使用Raft一致性算法来实现一个高可用的复制日志.
- etcd可用于实现各种分布式协调: 分布式锁,领导者选举,写barriers等.

## (2)读和写:
- **Raft is leader-based**; the leader handles all client requests which need cluster consensus. However, the client does not need to know which node is the leader. Any request that requires **consensus** sent to a follower is automatically forwarded to the leader. Requests that do not require consensus (e.g., serialized reads) can be processed by any cluster member.

## (3)备注:
- https://github.com/etcd-io/etcd/tree/v3.2.17/Documentation
- https://coreos.com/etcd/docs/latest/

# 二 数据模型:
## (1)概述:
- etcd是一个持久的, 多版本, 并发控制数据模型.
- etcd被用来可靠存储不频繁更新的数据,并提供可靠的watch查询.
- etcd暴露历史版本的k-v对来支持快照和watch历史事件.
- k-v存储是不可变的, 针对它的操作不会原地更新其结构, 而是产生一条新的更新后的结构, 更新后, key的所有历史版本是可访问和可watch的.

## (2)逻辑视图:
- 存储从逻辑上是一个flat binary key空间, 该key空间有一个lexically sorted index on byte string keys, 所以区间查询很方便.
- key空间是多**revisions**, 每个原子的变化操作创建一个新的revision, 先前revision的所有数据保持不变且可以访问.
- 若为了节省空间, 存储会被压缩(compacted), 在compact revision前的revisions会被删除.
- 创建一个key会增加key的**version**, 若key当前不存在则version为1; 删除一个key会将version设置为0; 每一个修改操作会增加它的version.

## (3)物理视图:
- etcd在一个持久的b+树内以键值对形式存储物理数据, etcd还保持一个secondary in-momory b树索引来加速区间查询.
- 键值对的key是一个3元组(major,sub,type).

## (4)备注:
- https://github.com/etcd-io/etcd/blob/master/Documentation/learning/data_model.md

# 三 API:
## (1)概述:
- 所有etcd3的API以**gRPC**服务方式定义, 每个AP发送到etcd服务器的请求是一个gRPC rpc.

## (2)API种类:
- 处理etcd key空间: KV(创建,更新,fetch和删除键值对); watch(监控keys的变化); Lease.
- 管理etcd集群: Auth, Cluster和Maintenance.

## (3)备注:
- https://github.com/etcd-io/etcd/blob/master/Documentation/learning/api.md
