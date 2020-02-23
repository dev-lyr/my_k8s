# 一 概述:
## (1)概述:
- 自定义资源(Custom Resources)是k8s API的扩展.
- 自定义资源在运行中集群中通过动态注册来出现和消失, 一旦自定义资源安装, 用户可以使用kubectl来创建和访问它的对象.
- The **Operator pattern** 是自定义资源和自定义控制器的结合.

## (3)添加自定义资源的方式:
- CRD: **CustomResourceDefinition** API资源允许用户定义自定义资源, 简单易用, 若需扩展可以自己实现控制器.
- API Aggregation(AA): 允许用户提供自定义资源的定制化实现, 需要实现和部署用户自己的API Server, 更加灵活.

## (3)自定义控制器:
- A declarative API allows you to declare or specify the desired state of your resource and tries to keep the current state of Kubernetes objects in sync with the desired state. The **controller** interprets the structured data as a record of the user’s desired state, and continually maintains this state.
- 自定义属性只是允许简单的存放和获取结构化数据, 当和**自定义控制器**结合使用时, 自定义资源提供了一个真正的声明式API.
- 自定义控制器可以与任何类型资源一起工作, 但通常和自定义资源结合更高效.
- You can use custom controllers to encode domain knowledge for specific applications into an extension of the Kubernetes API.

## (4)备注:
- https://github.com/kubernetes-sigs/kubebuilder
- https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/
- https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definition-versioning/
- https://github.com/kubernetes/sample-controller/blob/master/docs/controller-client-go.md

# 二 CRD：
## (1)概述:
- CustomResourceDefinition(CRD) API资源允许用户定义自定义资源, 定义一个CRD对象创建一个自定义资源(name和scheme), kubernetes API提供和处理自定义资源的存储.

## (2)CustomResourceDefinitionSpec:
- conversion: defines conversion settings for the CRD.
- names
- group
- scope: 默认是namespaced.
- validation: 描述CustomResource的validation方法.
- versions: versions is the list of all API versions of the defined custom resource.

## (3)备注:
- https://github.com/kubernetes/sample-controller
- https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/
- https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definition-versioning/
