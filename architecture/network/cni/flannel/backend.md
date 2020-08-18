# 一 概述:
## (1)概述:
- flannel可能会搭配不同的backends,一旦设置,在运行时不能改变. 

## (2)类型:
- VXLAN: 推荐方式, 使用kernel的VXLAN来封装包.
- host-gw: 适用于希望性能提升的有经验用户, 并且基础设施支持(通常不能在云环境使用).
- UDP: 用于debug或者一些不支持VXLAN的旧kernel.

## (3)其它实验性:
- AliVPC
- AWS VPC
- IPIP
- IPSec

## (4)备注:
- https://github.com/coreos/flannel/blob/master/Documentation/backends.md

# 二 公共结构:
## (1)概述:
- manager.go
- common.go

## (2)Manager接口:
- GetBackend: 根据backend类型创建Backend.

## (3)Backend接口:
- RegisterNetwork: 当backend创建或开始管理一个新网络时候调用.

## (4)Network接口:
- Run
- Lease
- MTU


# 三 VXLAN:
## (1)概述:
- 使用内核的VXLAN来封装包.

## (2)选项:
- Type: vxlan
- VNI: VXLAN Identifier(VNI), 在linux上默认为1.
- Port: 用来发送被封装包的**UDP**端口,在linux上,默认是内核默认的(8472).
- GBP: 默认是false, 相关: https://github.com/torvalds/linux/commit/3511494ce2f3d3b77544c79b87511a4ddb61dc89.
- DirectRouting: 默认为false.
- MacPrefix: 只适用于windows.
