# 一 概述:
## (1)功能:
- watch新创建的但未分配node的pods, 并选择一个node去执行它们.
- 调度决定考虑的因素: individual and collective resource requirements, hardware/software/policy constraints, affinity and anti-affinity specifications, data locality, inter-workload interference and deadlines.
- k8s自带默认的调度器(kube-scheduler), 若自带调度器满足不了需求, 可以自定义调度器, 甚至可以伴随默认调度器同时运行多个调度器并告诉kubernetes pod使用的调度器.

## (2)scheduler扩展机制:
- SchdulerExtender
- Multiple schedulers
- Scheduler Framework(优先)

## (3)备注:
- 默认调度器: https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/
- 运行多个调度器: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
- 实现调度器: https://github.com/kubernetes/kubernetes/tree/master/pkg/scheduler
- kube-scheduler命令: https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/
- https://kubernetes.io/docs/tasks/administer-cluster/topology-manager/

# 二 kube-scheduler命令:
## (1)格式:
- kube-scheduler [flags]

## (2)网络相关:
- --bind-address: 用于监听--secure-port端口的IP地址,该IP地址可被集群内其他部分,外部CLI/WEB用户访问, 若为空, 则所有接口会被使用. 用于替换--address.
- --secure-port: 默认10259, 服务HTTPS.
- --port: 默认10251, 服务HTTP.
- --master: API server的地址.

## (3)配置相关:
- --config
- --policy-config-file
- --policy-configmap
- --policy-configmap-namespace: 默认kube-system.

