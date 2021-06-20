# 一 概述:
## (1)概述:
- Node Problem Detector以ds(或standalone)方式运行在每个node, 检测node问题并上报给apiserver.
- 当前k8s不会针对npd生成的NodeCondition和Event采取任何操作, 未来会接入自愈系统(remedy system)来处理node相关问题.

## (2)上报方式:
- NodeCondition: 一些使得Node不可用的持续问题(Permanent problem).
- Event: 一些有限影响pod的临时问题(temporary problem).

## (3)自愈系统:
- descheduler: https://github.com/kubernetes-sigs/descheduler

## (4)自定义监控插件:
- https://github.com/kubernetes/node-problem-detector/tree/master/pkg/custompluginmonitor

## (5)备注:
- https://kubernetes.io/docs/tasks/debug-application-cluster/monitor-node-health/
- https://github.com/kubernetes/node-problem-detector

