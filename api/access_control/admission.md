# 一 概述:
## (1)概述
- admission模块是一个软件模块, 用于修改或拒绝请求, 还可以set complex defaults for fields.
- 可配置运行多个admission模块, 每个依序调度; 与认证和授权不同, 若任一admission模块拒绝请求, 则请求立马被拒绝.
- Once a request passes all admission controllers, it is validated using the validation routines for the corresponding API object, and then written to the object store.
- 除了编译进去的admission插件, admission插件已扩展形式开发并以webhooks方式在运行时配置.

## (2)备注:
- https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/
- https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/

# 二 admission controller
## (1)概述:
- An **admission controller** is a piece of code that intercepts requests to the Kubernetes API server **prior to** persistence of the object, but **after** the request is authenticated and authorized. 
- The admission control process proceeds in **two phases**. In the **first** phase, **mutating** admission controllers are run. In the **second** phase, **validating** admission controllers are run. Note again that some of the controllers are both.

## (2)类型:
- validating
- mutating
- validating and mutating both.
- 备注: Admission controllers may be “validating”, “mutating”, or both. Mutating controllers may modify the objects they admit; validating controllers may not.

## (3)admission controller种类:
- AlwaysPullImages: 修改每个pod强制image pull策略为Always.
- DefaultStorageClass
- DefaultTolerationSeconds
- EventRateLimit
- ValidatingAdmissionWebhook: calls any validating webhooks which match the request. Matching webhooks are called in parallel; if any of them rejects the request, the request fails.
- MutatingAdmissionWebhook: calls any mutating webhooks which match the request. Matching webhooks are called in serial; each one may modify the object if it desires.
- 备注: https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#what-does-each-admission-controller-do

# 三 admission webhooks:
## (1)概述:
- admission webhooks是接收到admission请求的**HTTP callback**.
- 可以定义两种类型admission webhooks: **validating admission webhooks**和**mutating admission webhooks**.
- Mutating admission Webhooks are invoked **first**, and can modify objects sent to the API server to enforce custom defaults. After all object modifications are complete, and after the incoming object is validated by the API server, validating admission webhooks are invoked and can reject requests to enforce custom policies.
