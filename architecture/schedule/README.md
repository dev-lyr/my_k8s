# 一 概述:
## (1)功能:
- watch新创建的但未分配node的pods, 并选择一个node去执行它们.
- 调度决定考虑的因素: individual and collective resource requirements, hardware/software/policy constraints, affinity and anti-affinity specifications, data locality, inter-workload interference and deadlines.
- k8s自带默认的调度器, 若自带调度器满足不了需求, 可以自定义调度器, 并且可以伴随默认调度器同时运行多个调度器, 还可以你的Pods中的每一个使用哪个调度器.

## (2)备注:
- 默认调度器: https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/
- 运行多个调度器: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
- 实现调度器: https://github.com/kubernetes/kubernetes/tree/master/pkg/scheduler
- kube-scheduler命令: https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/
