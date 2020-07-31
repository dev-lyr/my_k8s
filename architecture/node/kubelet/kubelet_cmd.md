# 一 kubelet命令:
## (1)概述:
- 用法: kubelet [flags]
- 选项: kubernetes/cmd/kubelet/app/options

## (2)配置相关:
- --config string: kubelet的配置文件.
- --dynamic-config-dir: 开启动态kubelet配置.
- --kubeconfig: kubeconfig文件的位置, 指定如何连接到API Server, 若不指定则是standalone模式, 方便调试.
- --bootstrap-kubeconfig:用于获得client certificate的kubeconfig文件的地址.
- --feature-gates: A set of key=value pairs that describe feature gates for alpha/experimental features.

## (3)网络相关:
- --address: 默认为0.0.0.0, kubelet提供服务的地址.
- --port: 默认10250, kubelet服务的端口地址.
- --read-only-port: 只读端口,用于为未验证和授权的请求提供服务, 0表示disable.
- --healthz-bind-address: 默认127.0.0.1, 0.0.0.0(所有IPv4), ::(所有IPv6).
- --healthz-port: 默认10248, localhost healthz endpoint的部分.
- --node-ip: 节点的IP地址, 若指定kubelet会使用该IP地址.
- --pod-cidr: 用于Pod IP地址的CIDR, 只能用于standalone模式; 集群模式, 从master获取, 参考controller-manager配置.
- --node-status-update-frequency: Specifies how often kubelet posts node status to master, 默认10s, Note: be cautious when changing the constant, it must work with nodeMonitorGracePeriod in nodecontroller(controller-manager配置).

## (4)日志相关:
- --log-dir: 若非空, 写入该目录下日志文件.
- --log-file: 若非空, 使用该日志文件.
- --log-flush-frequency: 默认5s, 日志刷新的最大时间.
- --logtostderr: 默认为true, 将日志写入stderr而不是文件.
- --alsologtostderr: 将日志同时写入stderr和文件.

## (5)cgroup相关:
- --cgroup-driver: kubulet用来操作host上cgroups使用, 可选: cgroupfs(默认), systemd.
- --cgroup-root. 
- --cgroups-per-qos
- --kube-reserved-cgroup
- --kubelet-cgroups
- --runtime-cgroups
- --system-cgroups
- --system-reserved-cgroup

## (6)系统限制:
- --max-open-files: 可被kubelet进程打开的文件的数量, 默认为1000000.
- --max-pods: 可运行在kubelets上的Pods的数量.

## (7)apiserver相关:
- --register-node: 将node注册到apiserver, 若--kubeconfig未提供, 则该flag不重要, 默认为true.
- --kube-api-burst: 默认10, 和apiserver交互的burst, 废弃(需通过kubelet --config指定的配置文件来设置).
- --kube-api-qps: 默认5, 和apiserver交互的qps, 废弃(需通过kubelet --config指定的配置文件来设置).
- --kube-api-content-type: 默认application/vnd.kubernetes.protobuf, 发送到apiserver的请求的内容类型, 废弃(需通过kubelet --config指定的配置文件来设置).
- --hostname-override: 若非空会使用该字符串来替代真实的hostname, 若--cloud-provider配置, 则该cloud provider决定node的name.
- --cloud-provider: 云服务的provider, 若为空字符串则没有cloud provider, 若指定, 则cloud provider决定node的name.
- --cloud-config: cloud provider的配置文件路径.

## (8)容器相关:
- --container-runtime: 使用的运行时, 可能值: docker(默认), remote, rkt(废弃).
- --container-runtime-endpoint: 远程运行时服务的endpoint.
- --container-log-max-files: 一个容器最大保存的日志文件数, 值必须>=2, 只适用于运行时为remote.
- --container-log-max-size: 默认10Mi, 设置日志文件最大size, 超过会被rotate, 只适用于运行时remote.
- --docker-endpoint: 和docker交互的endpoint, 默认为unix:///var/run/docker.sock.
- --redirect-container-streaming
- --pod-infra-container-image: The image whose network/ipc namespaces containers in each pod will use,默认为:k8s.gcr.io/pause:3.2.

## (9)CNI相关:
- --network-plugin
- --cni-bin-dir
- --cni-conf-dir
- 备注: 前三个只适用于运行时为docker.

## (10)image相关:
- image-gc-high-threshold: 镜像gc执行的磁盘使用水位, 可选值0-100, 默认85, 设置100则会关闭镜像gc.
- image-gc-low-threshold: 磁盘使用比例小于该值时, 不会运行image的gc, 默认80, 不应高于image-gc-high-threshold.
- image-pull-progress-deadline: 默认1m, 只适用于docker. 
- image-service-endpoint

## (11)eviction相关

## (12)cpu:
- --cpu-cfs-quota: 默认true.
- --cpu-cfs-quota-period: cpu.cfs_period_us, 默认为linux内核默认(100ms).
- --cpu-manager-policy: none或static.
- --cpu-manager-reconcile-period
- 参考: https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/

## (13)streaming相关:
- --redirect-container-streaming: 开启容器streaming重定向.若为false,则kubelet会在api-server和容器运行时间proxy streaming数据; 若为true,kubelet会返回一个http重定向给api-server. proxy方式更安全,但是带来性能开销; 直接方式性能更好, 但是安全性较低(因为apiserver和容器运行时间的连接可能没有被验证).
- --streaming-connection-idle-timeout 

# 二 配置文件:
## (1)概述:
- kubelet配置参数的部分可以通过磁盘上的配置文件来指定, 替代通过命令行参数.
- 基于配置文件是推荐的方式, 因为简化了node的部署以及配置文件的管理.
- 命令行参数若指定和配置文件相同的flag, 则会覆盖掉配置文件的值.

## (2)配置文件:
- kubelet --config来指定, 配置文件需是JSON/YAML格式.
- 若使用--config, 则未指定的flag的默认值是KubeletConfiguration的默认值.

## (3)备注:
- https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/
- https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/kubelet/config/v1beta1/types.go

# 三 动态kubelet配置
## (1)概述:
- 功能: 允许通过部署一个ConfigMap并配置每个Node使用它来动态改变在线k8s集群的kubelet的配置.
- 参考: https://kubernetes.io/docs/tasks/administer-cluster/reconfigure-kubelet/

# 四 认证和授权:
## (1)概述:

## (2)备注:
- https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet-authentication-authorization/
