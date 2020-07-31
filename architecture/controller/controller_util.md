# 一 概述:
## (1)概述:
- controller_util.go

# 二 ControllerExpectations:
## (1)概述:
- ControllerExpectations is a cache mapping controllers to what they expect to see before being woken up for a sync.
- 实现: ControllerExpectationsInterface接口, 允许用户设置和等待expectations.

## (2)属性:
- store cache.Store

## (3)方法:
- SetExpectations: 为给定controller注册一个新的expectations(ControlleeExpectations).
- DeleteExpectations: 从TTLStore中删除给定controller的expectations.
- ExpectCreations: 调用SetExpectations.
- ExpectDeletions: 调用SetExpectations.
- GetExpectations
- LowerExpectations: 降低给定controller的expect counts(add和del).
- RaiseExpectations: 增加给定controller的expect counts(add和del).
- CreationObserved: 调用LowerExpectations, 给add减1.
- DeletionObserved: 调用LowerExpectations, 给del减1.
- SatisfiedExpectations

# 三 ControlleeExpectations:
## (1)概述:
- 功能: track controllee的create和delete.

## (2)属性:
- add
- del
- key
- timestamp

## (3)方法:
- Add
- IsExpired
- Fulfilled
- GetExpectations

# 四 UIDTrackingControllerExpectations
## (1)概述:
- 扩展ControllerExpectations, track 需删除pod的UID.

## (2)属性:
- ControllerExpectationsInterface
- uidStoreLock
- uidStore: 元素类型为UIDSet.

## (3)UIDSet:
- sets.String
- key string

