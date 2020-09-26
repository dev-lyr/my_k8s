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

## (3)方法:
- Start
- Admit
- IsUnderDiskPressure
- IsUnderMemoryPressure
- IsUnderPIDPressure
