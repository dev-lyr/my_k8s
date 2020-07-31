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

# 二 主要特性:
## (1)本地更新:
- 流程: 当pod在进行本地更新时, controller首先使用readinessGate来更新pod status使得pod变为not-ready; 然后更新pod spec中的镜像来触发kubelet在node上重建容器.
- gracePeriodSeconds: 控制器更新pod status和更新pod image的间隔时间, 避免其它控制器(例如:endpoints控制器)还没感知到pod not-ready就更新了镜像, 从而导致请求失败.
