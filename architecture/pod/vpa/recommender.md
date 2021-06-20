# 一 概述:
## (1)概述:
- Recommender是vpa系统内的核心binary, 基于pod资源历史和当前的使用情况计算pod推荐的资源请求, 推荐结果放到VPA资源的status.

## (2)flag:
- recommender-interval: fetch metrics的频率, 默认1分钟.
- storage: prometheus或checkpoint(默认).
- address: expose prometheus metrics的地址, 默认为:8942.
- kubeconfig
- prometheus-address
- prometheus-cadvisor-job-name

# 二 recommender:
## (1)概述:
- 基于定时从metrics api获得的利用率, 为指定容器计算推荐资源.

## (2)方法:
- RunOnce: performs one iteration of recommender duties followed by update of recommendations in VPA objects, 根据recommender-interval定时执行.
- UpdateVPAs: computes recommendations and sends VPAs status updates to API Server.
- GetClusterState
- GetClusterStateFeeder
- MaintainCheckpoints
- GarbageCollect: 删除旧的AggregateCollectionStates.

## (3)RunOnce:
- clusterStateFeeder.LoadVPAs
- clusterStateFeeder.LoadPods
- clusterStateFeeder.LoadRealTimeMetrics
- UpdateVPAs
- MaintainCheckpoints
- GarbageCollect

## (4)UpdateVPAs:

# 三 ClusterState:
## (1)概述:
- ClusterState包含VPA操作需要的所有运行时信息,包括:资源配置(pod,容器,vpa对象),计算资源(cpu和内存)的aggregated utilization和事件(容器oom).
- VPA推荐算法依赖的所有输入都在该结构中.

## (2)属性:
- Pods
- Vpas
- EmptyVPAs
- ObservedVpas
- aggregateStateMap: All container aggregations where the usage samples are stored.
- labelSetMap

# 四 ClusterStateFeeder:
## (1)概述:
- ClusterStateFeeder用于更新ClusterState对象的状态.

## (2)方法：
- InitFromHistoryProvider
- InitFromCheckpoints
- LoadVPAs
- LoadPods
- LoadRealTimeMetrics
- GarbageCollectCheckpoints
