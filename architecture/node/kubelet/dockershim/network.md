# 一 概述:

# 二 NetworkPlugin:
## (1)概述:
- kubelet网络插件要实现的接口.
- kubenet(kubenetNetworkPlugin)和cni(cniNetworkPlugin)都实现该接口.

## (2)方法:
- Init: 初始化插件.
- Event: 响应各类事件.
- Name: 插件名字.
- Capabilities: 插件能力.
- SetUpPod: 在infra容器被创建后但其他容器启动前调用.
- TearDownPod: 在infra容器被删除前调用.
- GetPodNetworkStatus: 获取容器网络状态.
- Status: 当创建处于error状态时返回error.

## (3)备注:
- https://github.com/kubernetes/kubernetes/blob/v1.18.0/pkg/kubelet/dockershim/network/plugins.go

# 三 cniNetworkPlugin:
## (1)概述:
- 实现NetworkPlugin接口的方法.
- kubernetes/pkg/kubelet/dockershim/network/cni

## (2)方法:
- NetworkPlugin接口的方法.
- addToNetwork: 被SetUpPod方法调用.
- deleteFromNetwork: 被TearDownPod方法调用.

## (3)属性:
- network.NoopNetworkPlugin
- loNetwork: cniNetwork类型指针.
- defaultNetwork: cniNetwork类型指针.
- host
- execer
- nsenterPath
- confDir
- binDirs
- cacheDir
- podCidr

## (4)cniNetwork结构:
- name
- NetworkConfig: libcni.NetworkConfigList指针.
- CNIConfig: libcni.CNI类型.
- Capabilities
