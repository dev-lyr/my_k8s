# 一 概述:
## (1)两种方式:
- stacked控制面板节点(默认): etcd成员和控制面板在同一个节点, 需要更少节点.
- 使用外部etcd集群: etcd成员和控制面板节点分离, 需要更多节点.

## (2)比较:
- https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/ha-topology/

## (3)前提:
- 三个满足kubeadm最小要求的master节点.
- 三个满足kubeadm最小要求的worker节点.
- 集群中所有网络全联调(公网或私网).
- 所有机器的sudo权限.
- 所有节点的ssh访问.
- kubeadm和kubelet在所有节点安装, kubectl是可选的.
- 三个额外的集群用于etcd成员(使用外部etcd集群时).

## (4)备注:
- https://github.com/kubernetes/kubeadm/blob/master/docs/ha-considerations.md

# 二 为kube-apiserver创建load balancer:
## (1)概述:
- 适用于两种HA方式.

## (2)步骤:
- 创建一个kube-apiserver load balancer.
- 添加控制面板节点到load balancer并测试连通性.
- 将剩余的控制面板节点添加到load balancer目标组.

