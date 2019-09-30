# 一 概述:
## (1)概述:
- 负责管理和运行**控制器**, 逻辑上每个控制器是独立的进程, 但为了减少复杂性, 将它们编译到一个binary并且运行在一个single进程.
- 控制器: 一个控制循环, 通过API服务器来监听集群中资源的状态变化, 并执行相关操作, 尝试将资源当前状态向期望状态收敛.
- 除了常用控制器, 也可以自定义控制器, 参考自定义资源文章.

## (2)controller manager:
- kubernetes/cmd/kube-controller-manager

## (3)控制器源码:
- 目录: kubernetes/pkg/controller
- 每个控制器一般都有一个构造器, 内部会创建一个Informer(本质是个监听器)， 每次API对象有更新时候会被调用.
- 备注: client-go可用来帮助实现控制器.

## (4)常用控制器:
- Replication控制器
- ReplicaSet, DaemonSet和Job控制器
- Deployment控制器
- StatefulSet控制器
- Node控制器
- Service控制器
- Endpoint控制器
- Namespace控制器
- PersistentVolume控制器
- 等等.

# 二 kube-controller-manager
## (1)概述:
- kube-controller-manager是一个包含kube核心控制循环的daemon.
- 用法: kube-controller

## (2)常用

## (3)controller并发相关:
- --concurrent-endpoint-syncs: 默认5.
- --concurrent-deployment-syncs: 默认5.
- 等等.
