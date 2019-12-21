# 一概述:
## (1)功能:
- DaemonSet确保**一些或者全部**Node上运行指定pod, 当node添加到集群时, 该pod会被自动创建, 当node从集群中移除时, pod会被垃圾回收.
- DaemonSet被删除时, 它创建的Pod会被清除.
- 最简单情况, daemonSet在所有节点运行每种类型daemon; 复杂情况是一种类型使用多个daemonSet来运行, 但是不同daemonSet使用不同flags或者对不同硬件类型使用不同的cpu或内存request.

## (2)典型使用场景:
- 在所有node上运行一个node monitoring daemon, 例如: Prometheus Node Exporter.
- 在所有Node上运行一个日志收集daemon, 例如:fluentd或logstash.
- 在每个node上允许一个集群存储daemon, 例如: glusterd或cepth.

# 二 DaemonSetSpec
