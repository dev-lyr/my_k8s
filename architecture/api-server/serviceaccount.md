# 一 概述:
## (1)概述:
- A service account provides an identity for processes that run in a Pod.
- Pod中进程需要访问apiserver时, 通过service account来进行校验(authenticated).
- 创建sa时候会自动创建对应的secret.

## (2)属性:
- automountServiceAccountToken
- imagePullSecrets: ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount.
- secrets: Secrets is the list of secrets allowed to be used by pods running using this ServiceAccount.

## (3)使用
- 当创建pod时, 可通过spec.serviceAccountName来指定希望使用的sa; 若不指定一个service account, 则会将统一ns内的"default" service account自动分配给pod.
- 若pod引用一个sa, 但是创建是sa不存在, 则创建pod会失败.

## (4)源码:
- serviceaccounts_controller.go
- tokens_controller.go

## (5)备注:
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
- https://kubernetes.io/docs/reference/access-authn-authz/service-accounts-admin/

# 二 serviceaccount自动化操作:
## (1)概述:
- 三个单独的组件实现serviceaccount的自动化操作: **service account admission controller**, **token controller**和**service account controller**.
