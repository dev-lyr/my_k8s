# 一 概述:
## (1)概述:
- 可以运行一个或多个认证模块,按照顺序一一尝试, 直至任一成功,apiserver不保证认证顺序.
- 认证插件会返回已经认证过用户的**用户名**和**组**, 这些信息被用来验证用户是否被授权执行某个操作.
- 可以同时enable多个认证模块, 但通常至少使用两种方法: 用于service accounts的service account token和至少一种用于user认证的方法.
- **system:authenticated组**: 表示所有验证成功的用户.
- 可通过authenticating proxy或authentication webhook和其它authentication协议集成.

## (2)用户类型:
- **service account**: Pod使用(确切说上面运行的应用使用), 允许集群内服务和API server交互.
- **普通用户**. 
- 备注: API请求必须关联到普通用户或者一个service account, 否则会被作为匿名请求对待.

## (3)认证方式:
- Client Certificates: 通过apiserver的--client-ca-file=file来开启.
- Static Token file: apiserver通过--token-auth-file=file来读取bearer tokens.
- Bootstrap Tokens: --enable-bootstrap-token-auth和TokenCleaner控制器.
- Static Passwork File: --basic-auth-file=file.
- Service Account Tokens
- Webhook Token Authentication.
- Authenticating Proxy

## (4)备注:
- https://kubernetes.io/docs/reference/access-authn-authz/authentication/
- https://kubernetes.io/docs/reference/access-authn-authz/service-accounts-admin/
- https://kubernetes.io/docs/concepts/cluster-administration/certificates/
