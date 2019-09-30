# 一 概述:
## (1)功能:
- Helm是一个管理kube charts的工具, charts是预先配置的kube资源的packages.
- Help简化了Kube应用的安装和管理, 类似apt/yum等工具.

## (2)相关概念:
- chart: 是创建一个kube应用的实例所必要的信息的一个bundle.
- config: 包含可以被merge到chart中用来创建一个released对象的配置信息.
- release: 一个chart的运行实例, 结合特定的config.
- repository: 是一个收集和分享charts的地方.

## (3)组件:
- helm client: 是一个命令行工具.
- tiller server: 是一个在集群内的server, 和helm client交互, 和kube API server交互.

## (4)备注:
- https://helm.sh/
- https://github.com/helm/helm
