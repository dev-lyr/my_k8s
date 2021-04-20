# 一 概述:
## (1)概述:
- vertical pod autoscaler(VPA): frees the users from necessity of setting up-to-date resource limits and requests for the containers in their pods. 
- VPA will set the requests automatically based on usage and thus allow proper scheduling onto nodes so that appropriate resource amount is available for each pod. 
- It can both **down-scale** pods that are over-requesting resources, and also **up-scale** pods that are under-requesting resources based on their usage over time.

## (2)VPA组件:
- **Recommender**: monitor当前和过去的资源使用情况, 并基于此提供容器的cpu和内存request的推荐值.
- **Updater**: it checks which of the managed pods have correct resources set and, if not, kills them so that they can be recreated by their controllers with the updated requests.
- **Admission Plugin**: 为新创建pod设置准确的资源requests.

## (3)使用:
- 通过配置一个**VerticalPodAutoscaler** CRD来使用.

## (4)限制:
- 更新已运行pod还是实验性feature.
- VPA不会驱逐不被controller管理的pod.
- VPA不应该与HPA在CPU或mem上共同使用, 其它自定义或external metrics可以.
- VPA推荐可能会超过可用资源从而导致pod pending, 部分情况可结合ClusterAutoScaler解决.
- 多个VPA资源match到同一个pod的行为不确定.

## (5)备注:
- https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler
- https://cloud.google.com/architecture/best-practices-for-running-cost-effective-kubernetes-applications-on-gke
- examples: https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler/examples

# 二 VerticalPodAutoscalerSpec:
## (1)概述:
- TargetRef: 指向pod控制器, 例如: deployment,statefuleset等.
- UpdatePolicy(PodUpdatePolicy): 控制changes apply到pod的规则.
- ResourcePolicy(PodResourcePolicy): 控制autoscaler如何计算推荐资源, 是一个ContainerResourcePolicy数组.

## (3)PodUpdatePolicy:
- Auto: 默认, 目前等同于Recreate.
- Recreate: Pod创建时候VPA分配资源请求, 同时通过驱逐重建来给更新已运行pod的资源请求.
- Initial: VPA只在Pod创建时候分配资源请求, 后续不做处理.
- Off: VPA不自动修改pod的资源请求, 只是在VPA对象上显示推荐值.

## (4)ContainerResourcePolicy:
- ContainerName: 容器的名字或者默认容器资源策略(DefaultContainerResourcePolicy).
- Mode: 控制针对特定容器是否开启autoscaler.
- MinAllowed
- MaxAllowed
- ControllerdResources
- ControlledValues

# 三 VerticalPodAutoscalerStatus:
## (1)概述:
- Recommendation(RecommendedContainerResources): 最近一次autoscaler计算的资源推荐.
- Conditions(VerticalPodAutoscalerCondition数组): Conditions is the set of conditions required for this autoscaler to scale its target and indicates whether or not those conditions are met.

## (2)RecommendedContainerResources

## (3)VerticalPodAutoscalerCondition

# 四 Recommender:
## (1)概述:
- Recommender是vpc系统内核心的binary, 基于pod资源历史和当前的使用情况为pod计算推荐资源requests.
