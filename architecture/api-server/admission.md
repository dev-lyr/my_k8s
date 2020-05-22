# 一 概述:
## (1)概述
- admission模块是一个软件模块, 用于修改或拒绝请求, 还可以set complex defaults for fields.
- 可配置运行多个admission模块, 每个依序调度; 与认证和授权不同, 若任一admission模块拒绝请求, 则请求立马被拒绝.
- Once a request passes all admission controllers, it is validated using the validation routines for the corresponding API object, and then written to the object store.
- 除了编译进去的admission插件, admission插件已扩展形式开发并以**webhooks**方式运行, 能够在运行时配置.

## (2)备注:
- https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/
- https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/

# 二 admission控制器:
## (1)概述:
- An **admission controller** is a piece of code that intercepts requests to the Kubernetes API server **prior to** persistence of the object, but **after** the request is authenticated and authorized. 
- The admission control process proceeds in **two phases**. In the **first** phase, **mutating** admission controllers are run. In the **second** phase, **validating** admission controllers are run. Note again that some of the controllers are both.

## (2)类型:
- validating
- mutating
- validating and mutating both.

## (3)admission控制器种类:
- AlwaysPullImages: 修改每个pod强制image pull策略为Always.
- DefaultStorageClass
- DefaultTolerationSeconds
- EventRateLimit
- **ValidatingAdmissionWebhook**: calls any validating webhooks which match the request. Matching webhooks are called in parallel; if any of them rejects the request, the request fails.
- **MutatingAdmissionWebhook**: calls any mutating webhooks which match the request. Matching webhooks are called in serial; each one may modify the object if it desires.
- 备注: https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#what-does-each-admission-controller-do

## (4)备注:
- 代码: kubernetes/plugin/pkg/admission

# 三 admission webhooks:
## (1)概述:
- admission webhooks是接收到admission请求的**HTTP callback**.
- 可以定义两种类型admission webhooks: **validating admission webhooks**和**mutating admission webhooks(先被调用)**.

## (2)API资源:
- MutatingWebhookConfiguration: 重要属性webhooks(webhook列表, 类型为MutatingWebhook).
- ValidatingWebhookConfiguration: 重要属性webhooks(webhook列表, 类型为ValidatingWebhook).

## (3)MutatingWebhook:
- clientConfig: 定义与hook的通信方式, 必选, 类型为WebhookClientConfig.
- 等等.

## (4)ValidatingWebhook:
- clientConfig: 定义与hook的通信方式, 必选, 类型为WebhookClientConfig.
- 等等.

## (5)WebhookClientConfig:
- url: webhook的地址, 标准URL格式(scheme://host:port/path).
- service: webhook的服务的引用, url或service必须制定一个,若webhook运行在集群内则使用service.
- 备注: 包含于hook进行TLS连接的信息.
