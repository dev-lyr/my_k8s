# 一 概述:
## (1)功能:
- Helm是kubernetes的包管理器.
- Helm是一个管理kube charts的工具, charts是预先配置的kube资源的packages.
- Help简化了Kube应用的安装和管理, 类似apt/yum等工具.

## (2)组件:
- helm client: 是一个命令行工具.
- helm library

## (3)备注:
- https://helm.sh/
- https://github.com/helm/helm

# 二 相关概念:
## (1)chart:
- helm使用的包格式, 一个chart是创建一个kube应用的实例所必要的信息的一个bundle.
- chart有特定的目录树, 可以被打包成版本化的archives进行部署.

## (2)config:
- 包含可以被merge到chart中用来创建一个released对象的配置信息.

## (3)release:
- 一个chart的运行实例, 结合特定的config.

## (4)repository:
- 是一个收集和分享charts的地方.
