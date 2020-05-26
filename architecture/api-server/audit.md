# 一 概述:
## (1)概述:
- Kubernetes auditing provides a security-relevant chronological set of records documenting the sequence of activities that have affected system by individual users, administrators or other components of the system.

## (2)备注:
- https://kubernetes.io/docs/tasks/debug-application-cluster/audit/

# 二 工作方式:
## (1)概述:
- api-server来执行审计功能.
- 每个请求在执行的各个stage会产生一个event, 该事件根据特定的policy进行预处理和写入backend.
- 当前backend实现包括:日志文件和webhooks.

## (2)stage:
- RequestReceived: 只要audit handler收到请求就生成事件, 在传递给handler chain之前.
- RequestStarted: 响应header被发送, 在响应body发送之前.
- RequestComplete: 响应body已完成且没有更多数据需发送.
- Panic: 当panic发送时生成的事件.

# 三 audit policy:
## (1)概述:
- 审计策略定义了应记录哪些事件以及包含哪些数据的规则.
- kube-apiserver通过--audit-policy-file来查找审计策略文件, 若不设置,则不记录事件.

## (2)审计级别:
- None: 符合这表规则的事件不会记录.
- Metadata: 记录metadata(请求user,timestamp,resources,verb等), 不会记录请求/响应的body.
- Request: 记录事件metadata和请求body但是不记录响应body.
- RequestResponse: 记录事件的metadta,请求和响应体.

## (3)备注:
- https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/apis/audit/v1beta1/types.go


# 四 audit backend:
## (1)概述:
- 审计后端将审计事件持久化到外部存储.

## (2)类型:
- log后端(将日志写入磁盘): 将审计事件以json格式写入文件.
- webhook后端(发送事件到外部api)
- Dynamic后端(通过AuditSink API对象来配置webhook后端)

## (3)log后端:
- --audit-log-path: 文件路径, 不指定则disable log后端, -表示标准输出.
- --audit-log-maxage
- --audit-log-maxbackup
- --audit-log-maxsize
