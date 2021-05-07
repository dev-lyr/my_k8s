# 一 概述:
## (1)概述:
- evictionManager是kubelet的一个属性, 在kubelet启动时候启动.
- evictionManager作为Kubelet admitHandlers中的一个handler.

# 二 managerImpl:
## (1)概述:
- 实现Manager接口.

## (2)属性:
- killPodFunc
- imageGC
- containerGC
- nodeConditions: node当前的conditions.
- signalToRankFunc: 每类资源对驱逐pod时候使用的排序函数.

## (3)方法:
- Start: starts the control loop to observe and response to low compute resources.
- Admit: rejects a pod if its not safe to admit for node stability.
- synchronize: 核心处理流程.
- IsUnderDiskPressure
- IsUnderMemoryPressure
- IsUnderPIDPressure

# 三 synchronize:
## (1)概述:
- 功能: main control loop that enforces eviction thresholds.
- 执行间隔: 10s执行一次, 若有pod需要销毁则等待(最长30s)执行一次.

## (2)执行流程
