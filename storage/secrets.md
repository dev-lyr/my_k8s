# 一 概述:
## (1)概述:
- **secret**对象用来存储和管理一些敏感信息, 例如: 密码, Oauth token和ssh keys等.
- 将信息存放在secret比存放在pod的定义或者容器镜像中更加灵活和安全.
- 用户和系统都可以创建secret, 为了使用secret, pod需要引用secret.
- 使用secret的两种方式: 以卷中文件的方式挂载到一个或多个容器; 被kubelet在pulling pod镜像时候使用.

## (2)built-in secrets:
- Service Accounts Automatically Create and Attach Secrets with API Credentials.

## (3)创建:
- kubectl create secret
- 构造一个yaml/json, 然后使用kubectl创建.

## (4)secret属性:
- apiVersion
- kind
- metadata
- data
- stringData
- type

## (5)type:
- SecretTypeOpaque SecretType = "Opaque"
- SecretTypeServiceAccountToken SecretType = "kubernetes.io/service-account-token"
- SecretTypeDockercfg SecretType = "kubernetes.io/dockercfg"
- SecretTypeDockerConfigJSON SecretType = "kubernetes.io/dockerconfigjson"
- SecretTypeBasicAuth SecretType = "kubernetes.io/basic-auth"
- SecretTypeSSHAuth SecretType = "kubernetes.io/ssh-auth"
- SecretTypeTLS SecretType = "kubernetes.io/tls"
- SecretTypeBootstrapToken SecretType = "bootstrap.kubernetes.io/token"

## (6)备注:
- https://kubernetes.io/docs/concepts/configuration/secret/
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/auth/secrets.md

# 二 使用secrets
## (1)概述:
- secrets可以以**数据卷**方式挂载到pod的容器中, 或者作为**环境变量**被容器使用.
