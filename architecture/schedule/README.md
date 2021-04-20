# 一 概述:
## (1)功能:
- watch新创建的但未分配node的pods, 并选择一个node去执行它们.
- 调度决定考虑的因素: individual and collective resource requirements, hardware/software/policy constraints, affinity and anti-affinity specifications, data locality, inter-workload interference and deadlines.
- k8s自带默认的调度器(kube-scheduler), 若自带调度器满足不了需求, 可以自定义调度器, 甚至可以伴随默认调度器同时运行多个调度器并告诉kubernetes pod使用的调度器.

## (2)备注:
- kube-scheduler命令: https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/
- https://kubernetes.io/docs/tasks/administer-cluster/topology-manager/

# 二 scheduler扩展:
## (1)实现方式:
- SchedulerExtender
- Multiple schedulers
- Scheduler Framework(优先)

## (2)SchedulerExtender:
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/scheduling/scheduler_extender.md

## (3)多个调度器:
- 使用podSpec的schedulerName来指定使用哪个调度器, 默认是默认调度器.

## (4)备注:
- 运行多个调度器: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
- 实现调度器: https://github.com/kubernetes/kubernetes/tree/master/pkg/scheduler

