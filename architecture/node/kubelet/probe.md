# 一 概述:
## (1)概述:
- pkg/kubelet/prober

# 二 prober_manager:
## (1)概述:
- 管理pod的探测, 为每个指定prober的容器创建probe worker, worker按周期探测它分配的容器并cache探测结果, 每个probe对应一个worker.
- Manager根据探测结果来设置相应的pod status.

## (2)manager属性:
- workers: 当前active的workers.
- statusManager: 提供探测的Pod IP和容器IDs.
- livenessManager: 管理liveness探测的结果.
- readinessManager: 管理readiness探测的结果.
- startupManager: 管理startup探测的结果.
- prober: 执行探测(prober.go).

## (3)manager方法:
- Start
- AddPod: 在每个pod创建时候被调用, 为每个容器probe创建新的probe worker.
- RemovePod
- CleanupPods
- UpdatePodStatus

# 三 worker

# 四 prober
