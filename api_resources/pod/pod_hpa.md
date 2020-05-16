# 一 概述:
## (1)Horizontal Pod Autoscaler:
- 横向Pod自动伸缩通过Horizontal控制器来对Pod控制器(replicaSet等)管理的pod的副本数据进行自动伸缩, 基于观察到的CPU利用率或者借助custom metrics支持基于应用级别的metrics.
- 通过**HorizontalpodAutoscaler(HPA)资源**和**Horizontal控制器**来实现HPA.
- Horizontal控制器**周期**性检查Pod性能指标, 计算满足HPA资源所配置指标所需的副本数量, 从而调整对应Pod控制器的replicas字段.

## (2)相关配置:
- --horizontal-pod-autoscaler-sync-period: 默认15s, 控制器的检查周期.
- --horizontal-pod-autoscaler-downscale-stabilization: 默认5min.

## (3)备注:
- kubernetes/pkg/controller/podautoscaler
- https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/
- 当前稳定版本只支持cpu autoscaling, beta版本支持内存和自定义metrics的autoscaling.
- Vertical Pod Autoscaler(VPA).
