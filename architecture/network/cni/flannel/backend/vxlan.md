# 一 概述:
## (1)概述:
- 使用内核的VXLAN来封装包.

## (2)选项:
- Type: vxlan
- VNI: VXLAN Identifier(VNI), 在linux上默认为1.
- Port: 用来发送被封装包的**UDP**端口,在linux上,默认是内核默认的(8472).
- GBP: 默认是false, 相关: https://github.com/torvalds/linux/commit/3511494ce2f3d3b77544c79b87511a4ddb61dc89.
- DirectRouting: 默认为false.
- MacPrefix: 只适用于windows.

# 二 VXLANBackend:
## (1)概述:
- 实现Backend接口.

## (2)属性:
- subnetMgr
- extIface

## (3)RegisterNetwork

# 三 network:
## (1)概述:
- vxlan_network.go

## (2)属性:
- backend.SimpleNetwork
- dev
- subnetMgr

## (3)方法:
- Run
- MTU
- Lease

# 四 Run:
## (1)概述:
- 循环处理subnet事件.

## (2)EventAdded:
- 若是DirectRouting: 添加direct route.
- 若非DirectRouting: 添加ARP,添加FDB,添加vlanRoute.

## (3)EventRemoved:
- 删除(2)添加的信息.
