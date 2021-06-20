# 一 概述:
## (1)概述：
- kruise由一些扩展和补充kubernetes核心控制器的控制器组成, 用于workload管理.

## (2)主要特性:
- 本地更新(in-place update): 相对于deployment/statefulset的recreate update, 本地更新速度更快, 且几乎对pod内其它运行容器没有影响.
- sidecar容器管理
- 多fault domains部署

## (3)workload控制器:
- CloneSet.
- Advanced StatefuleSet.
- SidecarSet
- UnitedDeployment
- BroadcastJob

## (4)备注:
- https://openkruise.io/
- https://github.com/openkruise/kruise

# 二 组件:
## (1)概述:
- crds: clonesets等.
- **kruise-manager**:运行controllers和webhooks的控制面板, 以statefulset部署在**kruise-system**命名空间.
- **kruise-daemon(将要实现)**: 以daemonset部署, 运行在每个node并管理镜像资源(预拉镜像).
