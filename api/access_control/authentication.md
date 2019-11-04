# 一 概述:
## (1)概述:
- 可以运行一个或多个认证模块, 认证模块包括: Client Certificates, Paasowrd, Plain tokens, Boostrap tokens和JWT token.
- 认证插件会返回已经认证过用户的**用户名**和**组**, 这些信息被用来验证用户是否被授权执行某个操作.
- 可以同时enable多个认证模块, 但通常至少使用两种方法: 用于service accounts的service account token和至少一种用于user认证的方法.

## (2)用户类型:
- **service account**: Pod使用(确切说上面运行的应用使用), 允许集群内服务和API server交互.
- **普通用户**. 
- 备注: API请求必须关联到普通用户或者一个service account, 否则会被作为匿名请求对待.

## (3)user accounts和service accounts:
- user accounts用于人工(humans); service account用于运行于pod中的进程
- user accounts是global的, name必须在整个cluster内唯一; service accounts是namespaced.
- Typically, a cluster’s User accounts might be synced from a corporate database, where new user account creation requires special privileges and is tied to complex business processes. Service account creation is intended to be more lightweight, allowing cluster users to create service accounts for specific tasks (i.e. principle of least privilege).
- Auditing considerations for humans and service accounts may differ.
- A config bundle for a complex system may include definition of various service accounts for components of that system. Because service accounts can be created ad-hoc and have namespaced names, such config is portable.
- 参考: https://kubernetes.io/docs/reference/access-authn-authz/service-accounts-admin/

## (4)备注:
- https://kubernetes.io/docs/reference/access-authn-authz/authentication/
- https://kubernetes.io/docs/reference/access-authn-authz/service-accounts-admin/

# 二 认证策略
