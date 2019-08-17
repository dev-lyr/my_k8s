# 一 概述:
## (1)功能:
- API服务器是管理组件, 已RESTful API的形式提供的查询和修改集群状态的接口, 它将状态存储在etcd中.
- 可通过kubectl,第三方client库等形式调用api服务器.

## (2)通知资源变更:
- 控制面板组件, 各种控制器和各种客户端可以通过订阅资源来接收到资源被创建, 修改和删除的通知.
- 客户端可以通过创建到API服务器的HTTP连接来监听资源变更.

## (3)访问方式:
- 各种client库.
- kubectl proxy命令.
- Pod内访问: 使用每个容器内挂载的secret目录下的ca.crt,token和namespace来访问(例如:curl); 使用ambassador容器, 在该容器内执行kubectl proxy命令.
