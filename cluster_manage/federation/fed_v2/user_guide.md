# 一 概述:
## (1)概述:
- https://github.com/kubernetes-sigs/kubefed/blob/master/docs/userguide.md

## (2)安装:
- https://github.com/kubernetes-sigs/kubefed/blob/master/docs/installation.md

# 二 集群注册:
## (1)概述:
- https://github.com/kubernetes-sigs/kubefed/blob/master/docs/cluster-registration.md

# 三 federated API types:
## (1)概述:
- 可以通过kubefedctl命令来开启任意kubernetes API类型(包括CRD)的federation.
- kubefedctl会为联邦类型创建一个CRD, 命名为**Federated<kind>**; 在KubeFed system namespace中创建一个**FederatedTypeConfig**.

## (2)语法:
- kubefedctl enable <target kubernetes API type>
- API type的形式: kind(例如:Deployment), 复数名字(例如:deployments), 组限制的复数名字(例如:deployments.apps)或短名字(例如:deploy).

## (3)关闭API类型的propagation:
- 方法1: 修改**FederatedTypeConfig**资源的spec.propagation为Disabled.
- 方法2: kubefedctl disable <FederatedTypeConfig name>.

# 四 Federating a target resource:
## (1)概述:
- **kubefedctl federate** creates a federated resource from a kubernetes resource. The federated resource will embed the kubernetes resource as its template and its placement will select all clusters.

# 五 Propagation status:
## (1)概述:
- When the **sync controller** reconciles a federated resource with member clusters, propagation status will be written to the resource.

# 六 Overrides:
## (1)概述:
- **Overrides** can be specified for any federated resource and allow varying resource content from the **template** on a per-cluster basis.

# 七 Using Cluster Selector

