# 一 概述:
## (1)功能:
- kubelet是每个node上运行的"node agent".
- Kubelet和PodSpec一起工作, 保证PodSpec描述的容器健康运行, kubelet不管理非kebernetes创建的容器.

## (2)PodSpec的来源:
- apiserver(主要).
- File: 通过命令行flag传递文件路径, 该路径下的文件会被周期性监控更新, 默认监控周期为20s.
- HTTP endpoint
- HTTP server

## (3)用法:
- kubelet [flags]

# 二 kubelet选项:
## (1)网络相关:
- --address: 默认为0.0.0.0, kubelet提供服务的地址.
- --port: 默认10250, kubelet服务的端口地址.
- --healthz-port: 默认10248, localhost healthz endpoint的部分.
- --node-ip: 节点的IP地址, 若指定kubelet会使用该IP地址.
- --pod-cidr: 用于Pod IP地址的CIDR, 只能用于standalone模式.

## (2)配置相关:
- --kubeconfig: kubeconfig文件的位置, 指定如何连接到API Server.
- --bootstrap-kubeconfig:用于获得client certificate的kubeconfig文件的地址.

## (3)日志相关:
- --log-dir: 若非空, 写入该目录下日志文件.
- --log-file: 若非空, 使用该日志文件.
- --log-flush-frequency: 默认5s, 日志刷新的最大时间.
- --logtostderr: 默认为true, 将日志写入stderr而不是文件.
- --alsologtostderr: 将日志同时写入stderr和文件.

## (4)cgroup相关:
- --cgroup-driver: kubulet用来操作host上cgroups使用的驱动.
- --cgrup-root. 
- --cgroups-per-qos

## (5)系统限制:
- --max-open-files: 可被kubelet进程打开的文件的数量, 默认为1000000.
- --max-pods: 可运行在kubelets上的Pods的数量.
