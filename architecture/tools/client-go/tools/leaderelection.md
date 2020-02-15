# 一 概述:
## (1)概述:
- leaderelection目录

## (2)LeaderElectionConfig(struct):
- Lock: 用来当锁的资源.
- LeaseDuration
- RenewDeadline
- RetryPeriod
- Callbacks
- WatchDog
- ReleaseOnCancel
- Name

## (3)函数:
- RunOrDie

# 二 LeaderElector:
## (1)属性:
- config LeaderElectionConfig
- name

## (2)方法:
- Run: 启动leader election循环.
- GetLeader
- IsLeader
- acquire: 尝试获得leader lease, 获得时立即返回, 否则定时轮询.
- renew: 获得leader lease后调用.
- tryAcquireOrRenew
- release
