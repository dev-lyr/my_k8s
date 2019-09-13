# 一 概述:
## (1)概述:
- Operator pattern是自定义资源和自定义控制器的结合.
- Operators are clients of the Kubernetes API that act as controllers for a Custom Resource.
- Operator是一种打包,部署和管理kubernetes应用的方法, kubernetes应用指部署在kubernetes中且通过kubernetes APIs或kubuctl工具管理的应用.

## (2)实现一个operator:
- 使用kubebuider
- 使用Operator Framework
- 等等.

## (3)部署operator:
- 常用部署operator方法是添加一个CRD和它关联的controller到集群中, controller通常运行控制面板之外, 例如可以在集群中的一个deployment中运行operator.

## (4)备注:
- https://kubernetes.io/docs/concepts/extend-kubernetes/operator/
- https://coreos.com/operators/
- https://github.com/operator-framework
- https://github.com/operator-framework/awesome-operators
- https://github.com/operator-framework/operator-sdk
