# 一 概述:
## (1)概述:
- In Kubernetes, a **controller** is a **control loop** that watches the shared state of the cluster through the apiserver and makes changes attempting to move the current state towards the desired state.
- The Kubernetes **controller manager** is a daemon that embeds the core control loops shipped with Kubernetes. In applications of robotics and automation, a control loop is a non-terminating loop that regulates the state of the system.

## (2)类型:
- **Node Controller**: Responsible for noticing and responding when nodes go down.
- **Replication Controller**: Responsible for maintaining the correct number of pods for every replication controller object in the system.
- **Endpoints Controller**: Populates the Endpoints object (that is, joins Services & Pods).
- **Service Account & Token Controllers**: Create default accounts and API access tokens for new namespaces.
- 备注: kube-controller-manager组件.

## (3)备注:
- **支持自定义控制器**: 参考自定义资源文章.
