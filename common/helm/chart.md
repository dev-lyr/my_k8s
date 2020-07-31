# 一 概述:
## (1)概述:
- helm使用的包格式被称作chart, chart有特定的目录结构.
- 一个chart是一个描述相关k8s资源集的文件的集合.
- 单个chart可用于部署简单的事情, 也可以用来部署复杂的事情.

## (2)helm相关命令:
- search: 从用户添加的hub和repo中搜索heml charts.
- show: 查询一个chart的信息.
- create: 创建一个新的chart.
- pull: 从repo下载和unpack一个char到本地目录.
- install: 安装一个chart.
- package: 将一个chart目录打包为chart archive.
- dependency: 管理chart的依赖.

# 二 chart文件结构:
## (1)概述:
- 目录名是chart的名称(不带版本), 目录下charts,crds和templates三个目录是helm保留的.

## (2)包含:
- charts目录: 包含该chart依赖的charts.
- crds目录: crd目录.
- templates目录: 模板的目录, 集合values.yaml可以生产合法的kubernetes manifest文件.
- Charts.yaml: 一个包含chart相关信息的yaml文件.
- values.yaml: 默认的chart配置值.
- values.schema.json(可选)
- README.md(可选)
- LICENSE(可选)
