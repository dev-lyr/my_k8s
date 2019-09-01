# 一 概述:
## (1)概述:
- 联邦使得管理多个集群更加容易, 联邦kubernetes集群可以是运行在不同云provider上的集群, 也可以是事先搭建好的集.
- 部署联邦集群推荐使用**kubefed**.

## (2)功能:
- **多集群间同步资源**: Federation provides the ability to keep resources in multiple clusters in sync. For example, you can ensure that the same deployment exists in multiple clusters.
- **跨集群服务发现**: Federation provides the ability to auto-configure DNS servers and load balancers with backends from all clusters.
- **高可用**: By spreading load across clusters and auto configuring DNS servers and load balancers, federation minimises the impact of cluster failure.
- **避免Provider锁定**: By making it easier to migrate applications across clusters, federation prevents cluster provider lock-in.

## (3)使用多集群的原因:
- 低延迟: 让多个region中距离用户近的集群来提供服务. 
- 错误隔离: 多个集群相对单个集群可提供错误隔离.
- 可扩展性: 单个集群的扩展有局限性(https://git.k8s.io/community/sig-scalability/goals.md).
- 混合云: 可以使用多个云provider或自己搭建的集群.

## (4)缺点:
- 增加的网络带宽和开销: 联邦控制面板需要watch所有集群来确保期望的集群状态, 若集群运行在云provider的不同region或不同provider则会导致巨大的网络开销.
- 降低跨集群隔离: 联邦控制面板里的一个bug会影响所有的集群.
- 成熟度: 联邦项目还不太成熟, 不是所有资源都可用且许多还处于alpha阶段.

## (5)相关文档:
- https://kubernetes.io/docs/concepts/cluster-administration/federation/
