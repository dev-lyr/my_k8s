# 一 概述:
## (1)功能:

## (2)类型:
- EventRecorder
- EventBroadcaster

## (3)k8s Event资源(core/v1):
- action
- apiVersion
- count: 事件发生的次数.
- eventTime: 事件首次被observed到的时间.
- firstTimestamp: 事件首次recorded的时间.
- involvedObject
- lastTimestamp
- message: 可读的操作的状态的描述.
- **reason**: 简短可读描述状态改变的原因.
- related
- **sources**: 上报事件的模块.
- **type**: 事件类型, 目前有Normal和Warning.
