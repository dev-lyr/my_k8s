# 一 kube-scheduler:
## (1)格式:
- kube-scheduler [flags]

## (2)网络相关:
- --bind-address: 用于监听--secure-port端口的IP地址,该IP地址可被集群内其他部分, 外部CLI/WEB用户访问, 若为空, 则所有接口会被使用. 用于替换--address.
- --secure-port: 默认10259, 服务HTTPS.
- --port: 默认10251, 服务HTTP.
- --master: API server的地址.
