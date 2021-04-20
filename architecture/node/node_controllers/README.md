# 一 概述:
## (1)概述:
- node控制器是一个管理node许多方面的控制面板组件, 在node的生命周期内有很多作用.

## (2)种类:
- node_lifecycle_controller
- node_ipam_controller

## (3)功能:
- 当node注册时给它**分配CIDR块**.
- 监控node的健康并在Node不可达时平滑驱逐pod.

## (4)备注:
- https://kubernetes.io/docs/concepts/architecture/nodes/#node-controller
- pkg/controller/nodelifecycle
- pkg/controller/nodeipam
- 配置: https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager/
