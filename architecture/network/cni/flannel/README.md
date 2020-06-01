# 一 概述:
## (1)概述:
- flannel是一个为kubernetes设计的可简单轻松配置三层网络的工具, 在k8s外也有广泛应用.
- flannel在每个host运行一个单独的agent(**flanneld**), 该agent根据预先配置的地址范围给每个host分配subnet.
- flannel使用k8s API或直接使用etcd存储网络配置,分配的子网和其它辅助网络信息(例如:host的公网IP).
- 包通过多种backend机制的一种进行转发,包括:VXLAN和其它多种云集成.

## (2)网络细节:
- flannel负责在集群中多个node间提供一个3层IPv4网络.
- flannel不控制容器怎么连接到host, 只管理host间的流量如何转发.

## (3)备注:
- https://github.com/coreos/flannel

# 二 backends:
## (1)概述:
- flannel可能会搭配不同的backends,一旦设置,在运行时不能改变. 

## (2)类型:
- VXLAN: 推荐方式.
- host-gw: 适用于希望性能提升的有经验用户, 并且基础设施支持(通常不能在云环境使用).
- UDP: 用于debug或者一些不支持VXLAN的旧kernel.

## (3)其它实验性:
- AliVPC
- AWS VPC
- IPIP
- IPSec

## (4)备注:
- https://github.com/coreos/flannel/blob/master/Documentation/backends.md

# 三 subnet
