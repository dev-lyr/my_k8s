# 一 概述:
## (1)概述:
- 功能: kubernetes包管理器.
- 语法: helm [command]

## (2)环境变量:
- HELM_CACHE_HOME: 设置存储cache文件的目录.
- HELM_CONFIG_HOME: 设置存储config文件的目录.
- HELM_DATA_HOME: 设置存储data文件的目录.
- HELM_DRIVER: 设置backend storage driver, 可选: configmap,secret,memory,postgres.
- HELM_DRIVER_SQL_CONNECTION_STRING: 设置SQL存储driver使用的连接字符串.
- HELM_NO_PLUGINS: disable插件.
- KUBECONFIG: 设置kubernetes配置文件目录(默认:~/.kube/config).
- 备注: helm env

## (3)相关目录:
- helm存放cache,配置和数据的目录顺序: 优先使用对应环境变量; 否则系统支持XDG base目录规范则使用XDG变量;否则使用系统默认.
- linux上默认路径: $HOME/.cache/helm,$HOME/.config/helm,$HOME/.local/share/helm.

## (4)flags:
- --kube-apiserver: apiserver的地址和port.
- -n,--namespace: 请求的namespace(k8s的namespace)范围.

# 二 命令:
## (1)repo:
- 功能: 对chart仓库的增删改查.

## (2)chart相关:
- search: 从用户添加的hub和repo中搜索heml charts.
- show: 查询一个chart的信息.
- create: 创建一个新的chart.
- pull: 从repo下载和unpack一个char到本地目录.
- install: 安装一个chart.
- package: 将一个chart目录打包为chart archive.
- dependency: 管理chart的依赖.

## (3)release相关:
- list: 返回指定namespace的releases, 若没指定则使用当前namespace.
- get
- status
- test
- upgrade
- uninstall
- rollback
- history
