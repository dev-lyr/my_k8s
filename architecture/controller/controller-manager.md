# 一 概述:
## (1)概述:
- 源码: kubernetes/cmd/kube-controller-manager
- 每个controller作为一个goroutine与kube-controller-manager进程一起运行.
- 所有controller使用同一个InformerFactory, 但有自己的queue, 调和也是基于自己的队列进行的.
- controller从cache中list和get.
- 领导者选举(HA): 是kube-controller-manager节点界别的.
- 每个controller源码中NewXXX用来配置watch event handler.

## (2)controller相关:
- controller详细列表: 参考NewControllerInitializers函数. 
- ControllersDisabledByDefault表示默认不启动的controller.

# 二 kube-controller-manager命令:
## (1)概述:
- kube-controller-manager是一个包含kube核心控制循环的daemon.
- 用法: kube-controller-manager [flags]

## (2)--controllers:
- 配置enable的控制器列表, 星号表示开启默认的controller. 其它的例如: foo表示开启,-foo表示关闭.
- 所有controller列表: attachdetach, bootstrapsigner, cloud-node-lifecycle, clusterrole-aggregation, cronjob, csrapproving, csrcleaner, csrsigning, daemonset, deployment, disruption, endpoint, endpointslice, garbagecollector, horizontalpodautoscaling, job, namespace, nodeipam, nodelifecycle, persistentvolume-binder, persistentvolume-expander, podgc, pv-protection, pvc-protection, replicaset, replicationcontroller, resourcequota, root-ca-cert-publisher, route, service, serviceaccount, serviceaccount-token, statefulset, tokencleaner, ttl, ttl-after-finished
- 默认关闭: bootstrapsigner, tokencleaner

## (3)controller并发相关:
- --concurrent-endpoint-syncs: 默认5.
- --concurrent-deployment-syncs: 默认5.
- 等等.
