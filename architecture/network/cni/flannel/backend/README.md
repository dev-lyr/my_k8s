# 一 概述:
## (1)概述:
- flannel可能会搭配不同的backends,一旦设置,在运行时不能改变. 

## (2)类型:
- VXLAN: 推荐方式, 使用kernel的VXLAN来封装包.
- host-gw: 适用于希望性能提升的有经验用户, 要求host间二层网络是连通的.
- UDP: 用于debug或者一些不支持VXLAN的旧kernel.

## (3)其它实验性:
- AliVPC
- AWS VPC
- IPIP
- IPSec

## (4)备注:
- https://github.com/coreos/flannel/blob/master/Documentation/backends.md
- flannel/backend

# 二 公共结构:
## (1)概述:
- manager.go
- common.go

## (2)Manager接口:
- GetBackend: 根据backend类型创建Backend.

## (3)Backend接口:
- RegisterNetwork: 当backend创建或开始管理一个新网络时候调用.

## (4)ExternalInterface:
- Iface
- IfaceAddr
- ExtAddr
- 备注: 对外出口的设备, 若没指定publicIP以及iface等, 则默认默认gateway接口(GetDefaultGatewayIface).

## (5)Network接口:
- Run
- Lease
- MTU
