# 一 概述:
## (1)概述:
- https://kubernetes.io/docs/concepts/containers/images/

# 二 镜像:
## (1)概述:
- 在k8s Pod关联镜像之前需构建镜像并推送到仓库.
- 默认的pull策略是**IfNotPresent**: 若镜像已存在则kubelet跳过pull镜像.

## (2)强制每次pull的方法:
- 设置imagePullPolicy为**Always**.
- 忽略imagePullPolicy并且使用:latest(不建议使用).
- 忽略imagePullPolicy和使用tag.
- enable **AlwaysPullImages** admission控制器.

# 三 容器环境变量:
## (1)概述:
- K8s容器环境为容器提供一些重要资源: 一个文件系统(结合一个镜像和一个或多个volumes); 容器自己相关的信息; 集群中其它对象的信息.

## (2)容器信息:
- hostname: is the name of the Pod in which the Container is running, 可通过hostname命令或gethostname函数获取.
- Pod的名字和namespace作为环境变量通过downward API获取.
- Pod定义中用户自定义的环境变量, 以及Docker镜像中静态指定的环境变量.

## (3)集群信息:
- 容器创建时所有的运行中的服务的列表以环境变量的形式对容器可见.
- 服务的显式IP地址通过DNS对容器可见(需enable DNS).

# 四 Runtime类:
## (1)概述:
- RuntimeClass是一种用于选择容器运行时配置(container runtime configuration)的特性, 容器运行时配置用于运行一个Pod的容器.

## (2)setUp:
- **RuntimeClass**特性必在apiservers和kubelets上开启: 在Node上配置CRI实现; 创建对应的RuntimeClass资源.
- 通过RuntimeClass配置依赖容器运行时接口(CRI)的实现, 不同CRI实现不一样, 配置(Configuration)有一个对应的handler名字, 被RuntimeClass引用.
- 对于每个handler, 创建一个相应的RuntimeClass对象, RuntimeClass资源有两个属性: RuntimeClass名字(metadata.name)和handler(handler).

## (3)使用

# 五 容器生命周期Hook:
## (1)概述:
- K8s为容器提供lifecycle hooks, 这些hook可以让容器感知它们生命周期的事件, 并且在hook被执行时运行handler中的代码.
- 容器可以通过实现和注册一个handler到hook方式来访问一个hook.

## (2)export给容器的hook类型:
- **PostStart**: 当容器被创建时该hook立即执行, 但是不保证hook会在容器ENTRYPOINT之前执行, 没有参数传递给handler.
- **PreStop**: 当容器被终止之前调用, 终止方式包括: API请求或管理事件(liveness probe失败. preemption, 资源紧张或其他). 当容器已经是终止或完成状态, 则该Hook调用失败.

## (3)Hook Handler实现类型:
- **Exec**: 执行一个特定的命令, 在容器的cgroups和namespaces中执行.
- **HTTP**: Executes an HTTP request against a specific endpoint on the Container.

## (4)Handler的执行:
- 当容器lifecycle管理hook被调用时, k8s在注册该hook的容器中执行handler.
- Hook handler的调用是同步的, 若hooks执行或hung太久, 则容器不能到达running状态; PreStop Hook类似, 若hung太久, 则Pod处于Terminating状态.
- 若PostStart或PreStop hook失败, 则会kill容器.
- 用户应该让hook handler越轻量越好.

## (5)Hook传递保证:
- Hook传递是至少一次, 即一个hook对于一个给定事件可能会被调用多次, 需要hook的实现来处理正确性.

## (6)debug
