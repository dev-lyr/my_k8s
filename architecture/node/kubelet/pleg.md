# 一 概述:
## (1)概述:
- pod生命周期事件生成器(pod lifecycle event generator(pleg))
- 代码: kubelet/pleg

## (2)PodLifeCycleEventType:
- ContainerStarted: 容器的新状态是running.
- ContainerChanged: 容器的新状态时unknown.
- ContainerDied: 容器新状态是exited.
- ContainerRemoved: 容器旧状态是exited.
- PodSync: 触发pod的sync, 用于观察到的事件不归属于上述几类的情况.
