# 一 概述:
## (1)概述:
- containernetworking小组维护的CNI插件: https://github.com/containernetworking/plugins
- calico CNI插件: https://github.com/projectcalico/cni-plugin

## (2)类型:
- **Main(接口创建)**: bridge,ipvlan,loopback,macvlan,ptp,vlan,host-device.
- **IPAM(IP地址分配)**: dhcp,host-local,static.
- **Meta(其它插件)**: flannel,portmap等.
