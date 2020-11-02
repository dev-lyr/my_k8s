# 一 概述
## (1)概述:

## (2)配置:
- --cidr-allocator-type: 默认RangeAllocator.
- --allocate-node-cidrs: Should CIDRs for Pods be allocated and set on the cloud provider. 
- --cluster-cidr: CIDR Range for Pods in cluster, 需要allocate-node-cidrs为true.
- --node-cidr-mask-size: Mask size for node cidr in cluster, 默认针对ipv4是24,ipv6是64.
- --node-cidr-mask-size-ipv4
- --node-cidr-mask-size-ipv6
- --service-cluster-ip-range: CIDR Range for Services in cluster, 要求allocate-node-cidrs为true.
- 等等.

## (3)备注:
- pkg/controller/nodeipam
