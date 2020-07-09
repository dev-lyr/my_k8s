# 一 概述:
## (1)概述:
- 功能: SDK for building Kubernetes APIs using CRDs.

## (2)相关目录:
- go.mod
- Makefile: build和deploy自定义controller.
- PROJECT
- config: 包含启动配置,CRD,RBAC配置,WebHook配置等.
- main.go
- 等等.

## (3)备注:
- https://github.com/kubernetes-sigs/kubebuilder
- https://book.kubebuilder.io/introduction.html
- 代理: https://goproxy.io/(需使用Go module).

# 二 main.go
## (1)功能:
- 为metrics设置一些基本flags.
- 实例化一个manager.
- 运行manager, manager来运行所有controller和webhooks.

## (2)manager:
- 创建方法: manager.New, 参数为rest.Config和Options.
- Start: 启动所有注册的controller并block直至Stop channel被close, 当有controller启动失败时返回一个error.

## (3)备注:
- https://godoc.org/sigs.k8s.io/controller-runtime/pkg/manager

# 三 API:
## (1)添加:
- 执行cmd: kubebuilder create api --group batch --version v1 --kind CronJob
- 产出: api/v1/cronjob_types.go.

## (2)设计规则:
- 所有被序列化属性必须是驼峰(camelCase), 使用JSON struct tags来指定.
- 属性可使用大多数基本类型, 但是数字例外: 为了API兼容, 只支持三类数字: int32, int64和resource.Quantity表示小数.

# 四 controller:
## (1)概述:
- controller主要确保对象的实际状态match期望状态, 该过程称为调和(reconciling).
- controller里实现调和逻辑的是**XXXReconciler**.

## (2)XXXReconciler(struct):
- client.Client
- Log
- Scheme

## (3)Reconcile函数:
- reconcile发生错误尽早返回,成功是返回空的result和error为nil, 失败时返回error不为nil, 会自动入队列重试.

# 五 metrics

# 六 日志

# 七 event
