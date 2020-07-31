# 一 概述:
## (1)语法:
- kubeadm [command]

## (2)命令:
- init: 启动控制面板节点.
- join: 启动工作/控制面板节点并加入集群.
- upgrade: 升级集群到新的版本.
- token: 管理kubeadm join使用的tokens.
- 等等.

## (3)备注:
- https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm/

# 二 kubeadm init:
## (1)概述:
- 初始化和启动控制面板节点.
- kubeadm init [flags]
- kubeadm init [command]: 可用命令: phase(用于调用init workflow中指定单个phase的执行).
- https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init/

## (2)执行phases:
- preflight: 运行pre-flight check.
- kubelet-start: 写入kubelet配置并启动或重启kubelet.
- certs: 证书生成.
- kubeconfig: 生成创建控制面板需要的kubeconfig文件和admin kubeconfig文件.
- control-plane: 生成创建控制面板需要的static pod manifest文件,apiserver,scheduler和controller-manager.
- etcd: 生成local etcd的static pod manifest.
- upload-config: 将kubeadm和kubelet配置上传到configmap.
- upload-certs: 上传certificates到kubeadm-certs.
- mark-control-plane: 标记节点作为控制节点.
- bootstrap-token: 生成bootstrap token用于添加节点到集群.
- kubelet-finalize
- addon: 安装coredns和kube-proxy.

# 三 kubeadm join:
## (1)概述:
- 初始化一个工作/控制面板节点并加入集群.
- https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-join

## (2)工作节点执行phases

## (3)控制面板节点执行phases

# 四 kubeadm upgrade:
## (1)概述:
- kubeadm upgrade将负责的升级逻辑封装在一个命令里, 支持规划一个upgrade以及实际执行它.
- kubeadm upgrate [flags]
- kubeadm upgrate [command]
- https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-upgrade/


