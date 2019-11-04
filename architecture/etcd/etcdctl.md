# 一 概述:
## (1)功能:
- etcd的客户端命令行工具.
- 默认情况下, ETCDCTL_API未设置, 使用etcdctl v2, 需设置ETCDCTL_API=3来使用etcdctl v3 API.

## (2)用法:
- etcdctl [global options] command [command options] [arguments...]

## (3)全局选项:
- --endpoints=[127.0.0.1:2379]: gRPC endpoints.
- -w,--write-out="simple": 设置输出格式(fields, json, protobuf, simple, table), 使用非simple信息更丰富, 可查看reversion等信息.

# 二 基本命令:
## (1)get

## (2)put

## (3)del

## (4)txn
