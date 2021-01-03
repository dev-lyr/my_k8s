# 一 概述:
## (1)概述:
- kubernetes使用kubeconfig文件来组织有关集群,用户,命令空间和身份认证机制的信息.
- kubectl使用kubeconfig文件来选择集群所需信息并与集群API服务器通信.
- kubectl config: 修改和查询kubeconfig文件.

## (2)kubeconfig文件加载顺序:
- 若--kubeconfig选项指定, 则只有指定文件会被加载.
- 若KUBECONFIG环境变量设置, 则指定kubeconfig文件列表,这些kubeconfig文件会被merge.
- 最后,使用${HOME}/.kube/config文件, 不会进行merge.

## (3)生成方式:
- kubectl config
- kubeadm alpha kubeconfig

## (4)备注:
- https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/
- https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/

# 二 上下文(context):
## (1)概述:
- 使用kubeconfig文件的context约束可以对访问参数进行分组.
- kubectl config use-context: 选择上下文.

## (2)子元素:
- cluster
- namespace
- user

