# 一 概述:
## (1)概述:
- 横向Pod自动伸缩(HPA)通过Horizontal控制器来对Pod控制器(replicaSet等)管理的pod的副本数据进行自动伸缩, 基于观察到的CPU利用率或者借助custom metrics支持基于应用级别的metrics.
- 通过**HorizontalpodAutoscaler(HPA)资源**和**Horizontal控制器**来实现HPA.
- Horizontal控制器**周期**性检查Pod性能指标, 计算满足HPA资源所配置指标所需的副本数量, 通过pod控制器的scale子资源来调整对应副本.

## (2)相关配置:
- --horizontal-pod-autoscaler-sync-period: 默认15s, 控制器的检查周期.
- --horizontal-pod-autoscaler-downscale-stabilization: 默认5min.

## (3)kubectl相关:
- 查询: kubectl get/descibe
- 创建: kubectl create hpa或kubectl autoscale

## (4)备注:
- kubernetes/pkg/controller/podautoscaler
- https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/
- 当前稳定版本只支持cpu autoscaling, beta版本支持内存和自定义metrics的autoscaling.
- Vertical Pod Autoscaler(VPA).

# 二 HorizontalPodAutoscaler:
## (1)概述:
- 当前存在v1,v2beta1和v2beta2版本.

## (2)HorizontalPodAutoscalerSpec:
- **scaleTargetRef**
- **minReplicas**
- **maxReplicas**
- **targetCPUUtilizationPercentage**: target average CPU utilization (represented as a percentage of requested CPU) over all the pods, 默认80, 只在v1版本存在.
- **metrics(MetricSpec数组)**: metrics contains the specifications for which to use to calculate the desired replica count, 默认为80% average CPU utilization.
- **behavior**: behavior configures the scaling behavior of the target in both Up and Down directions.

## (3)MetricSpec
- **type**: metric源的类型,可选值为:Object,Pods或Resource.
- **resource**: resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory).
- **pods**: pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
- **external**: external refers to a global metric that is not associated with any Kubernetes object. 
- **object**: object refers to a metric describing a single kubernetes object.

## (4)HorizontalPodAutoscalerBehavior:
- **scaleDown(HPAScalingRules)**: scaleDown is scaling policy for scaling Down.
- **scaleUp(HPAScalingRules)**: scaleUp is scaling policy for scaling Up.

## (5)HorizontalPodAutoscalerStatus:
- conditions
- currentMetrics
- currentReplicas
- desiredReplicas
- lastScaleTime
- observedGeneration

# 三 支持metrics:
## (1)概述:
- 当前HPA支持多种metrics, 可通过使用autoscaling/v2beta2版本来指定.
- HPA控制器会计算每个metrics并提出一个新的scale, **最大**的scales会作为新的scale.
- HPA控制器通过一系列API来获得metrics, 为了能访问这些API, API aggregation layer必须开启.

## (2)metrics API:
- **metrics.k8s.io**: 用于**resource** metrics, 通常由metrics-server提供.
- **custom.metrics.k8s.io**: 用于**custom** metrics, 通过一些metrics解决方案vendors提供的**adapter** API servers来提供数据.
- **external.metrics.k8s.io**: 用于**external** metrics, 也可以通过custom metrics adapters来提供.

## (3)备注:
- 已有的一些方案: https://github.com/kubernetes/metrics/blob/master/IMPLEMENTATIONS.md
- 自己实现metrics api server: https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/custom-metrics-api.md

# 四 计算方式:
## (1)概述:
- 公式: desiredReplicas = ceil[currentReplicas * ( currentMetricValue / desiredMetricValue )]
