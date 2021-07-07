# 一 镜像:
## (1)概述:
- 在kube pod中使用前, 需创建镜像并push到仓库.
- 参考: https://kubernetes.io/docs/concepts/containers/images

## (2)更新镜像:
- **imagePullPolicy**: IfNotPresent(默认), 当使用:latest tag时默认为Always; Always; Never.
- 开启**AlwaysPullImages** admission controller.
- 需避免使用:latest.

## (3)私有镜像仓库.

# 二 ContainerSpec:
## (1)基本信息:
- command: entrypoint数组, 若未指定,使用docker镜像的ENTRYPOINT.
- args: entrypoint的参数, 若未指定使用docker镜像的CMD.
- env: 在容器中设置的环境变量的列表.
- envFrom: 一个生产容器环境变量的源的列表,例如:ConfigMap,Secret等.
- image: 镜像.
- imagePullPolicy: 镜像拉取策略,可选:Always,Never,IfNotPresent;若指定:latest则默认为Always,其它情况默认为IfNotPresent.
- name
- terminationMessagePath
- terminationMessagePolicy

## (2)resources

## (3)ports:
- 功能: 列出容器expose的端口,主要是说明容器的网络连接信息,不指定ports并不会阻止端口的expose.
- containerPort: 容器IP地址expose的端口号.
- hostIP
- hostPort: 在host上expost的端口号, 若podSpec的hostNetwork为true, 则hostPort需与containerPort相等.
- name
- protocol

## (4)存储相关:
- volumeMounts: 挂载到容器文件系统的pod volumes.
- volumeDevices: 容器使用的block device列表.
- workdingDir

## (5)probe相关:
- startupProbe
- livenessProbe
- readinessProbe

## (6)lifecycle:
- postStart: 在容器创建后立即调用,若handler失败,则容器被终止并根据重启策略重启. 
- preStop: 在容器由于API请求或者管理事件(liveness/startup探测失败,抢占或资源竞争等)被终止时调用, 容器crash或exits时不会调用.
- 备注: 类型为Handler.

## (7)其它:
- stdin
- stdinOnce
- tty

# 三 containerStatus:
## (1)概述:
- Pod的status属性包含一个containerStatuses属性表示Pod每个容器的状态, initContainerStatuses表示每个init容器的状态.
- containerStatuses是一个**ContainerStatus**对象的数组, ContainerStatus对象的state表示容器的状态.

## (2)containerStatus:
- containerId
- image
- imageId
- name
- ready: 容器是否通过readiness probe.
- restartCount
- started: 容器是否通过startup probe.
- state ContainerState

## (3)ContainerState:
- Running: 表示容器正在没问题的执行中, 当容器进入Running状态时, postStart hook(若有)会被执行.
- Terminated: 表示容器已经完成执行操作且已停止运行, 当容器成功执行完成或由于一些原因失败时候会进入该状态, 在进入Terminated状态之前, PreStop hook(若有)会被执行.
- Waiting: 容器的默认状态, 若容器不在Running或Terminated状态, 则处于Waiting状态, Waiting状态的容器依旧执行它要求的操作, 例如拉镜像, 应用secrets等.

# 四 容器probes:
## (1)概述:
- kubelet定期在容器上执行诊断性探测(Probe), 为了执行诊断, kubelet调用容器实现的handler.
- 每次探测的结果: Success(容器通过诊断), Failure(容器诊断失败), Unknown(诊断失败, 不会执行action).

## (2)探测类型:
- **livenessProbe**: 表示容器是否在运行中, 若探测失败, kubelet会kill容器, 并根据restart policy来决定如何操作, 若容器不提供livenessProbe探测则默认状态为Success.
- **readinessProbe**: 表示容器是否准备好接收请求, 若探测失败, 则**endpoint控制器**会从所有服务的endpoints中删除pod的ip地址.readiness在init delay之前的默认状态为Failure, 若不提供readiness探测则默认state为Success.
- **startUpProbe**: 可在应用启动时设置.
- 备注: 对应Probe资源.

## (3)handler类型:
- ExecAction: 在容器内执行一个指定命令, 若退出码为0则表明诊断执行成功.
- TCPSocketAction: 在容器IP地址上指定Port执行一个TCP检查, 若port是打开则表示诊断成功.
- HTTPGetAction: 在容器IP地址上指定port和路径上执行一个Get请求, 若返回码>=200且<400则表示诊断成功.

## (5)使用场景:
- 若希望容器在测探失败时被kill并且重启, 则提供一个liveness probe, 且指定restartPolicy为Always或OnFailure.
- 若希望在探测成功后再发生请求给Pod, 则提供一个readiness probe.

## (6)Probe资源属性:
- exec: ExecAction.
- httpGet: HTTPGetAction.
- tcpSocket: TCPSocketActon.
- initialDelaySeconds: 容器启动指定时间后再触发探测(三种类型都适用).
- periodSeconds: 探测间隔, 默认10s, 最小1.
- failureThreshold: 默认为3,最小为1,经过指定次数的连续probe失败才被认为是失败.
- successThreshold: 默认是1, 针对liveness和startup必须为1, 经过连续多次探测成功才算成功.
- timeoutSeconds: 探测的超时时间, 默认1s.

## (7)startup探针:
- 在容器启动时间超过**initialDelaySeconds + failureThreshold * periodSeconds**时, 可使用一个与liveness探针同样endpoint的startup探针, startup探针periodSeconds默认为30s, 可指定一个高的failureThreshold从而使得容器有足够的时间启动.
- 若指定startup探针, **在它成功完成前其他probes不会被执行**.
- 若探测失败, 则和liveness探针一样, 容器会被重启.

## (8)liveness探针经验:
- 对于生产中的Pod一定要定义一个存活探针, 没有存活探针, kubernete不会知道应用是否还活着, 只要进程还在运行, kubernete就认为Pod是健康的.
- 存活探针不应该消耗太多计算资源, 且运行时间不应该太长, 默认情况下, 探针的执行比较频繁, 必须在一秒内执行完成.
- 无需再探针中实现重试循环.
- 容器崩溃或存活探针失败, Pod所在节点的kubelet会**重启容器**, k8s的控制面板组件不参与该操作; 若节点崩溃, kubelet则无法执行相关操作, 因此需要使用各种Pod Controller.

## (9)备注:
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/

# 五 容器环境变量:
## (1)概述:
- kube容器环境为容器提供一些重要的资源: 一个文件系统(镜像和一个或多个volume的组合); 容器自身相关信息; 集群中其它对象相关信息.
- https://kubernetes.io/docs/concepts/containers/container-environment-variables/

## (2)容器信息:
- hostname: 容器的hostname是包含该容器的pod的name.
- pod的name和namespace可通过downward API作为环境变量使用. 
- 在pod定义中自定义的环境变量对容器也是可用的.

## (3)集群信息:
- services: 在容器创建时所有运行中的服务作为容器的环境变量使用.
- Services have dedicated IP addresses and are available to the Container via DNS, if DNS addon is enabled. 

# 六 运行时类(runtime class):
## (1)概述:
- RuntimeClass用于选择容器运行时的配置.

## (2)备注:
- https://kubernetes.io/docs/concepts/containers/runtime-class/

# 七 容器生命周期:
## (1)概述:
- kube为容器提供了生命周期hooks, 可以让容器感知它们管理生命周期的事件, 并且hook被执行时运行对应的handler中实现的代码.

## (2)hook类型:
- **PostStart**: 当容器被创建后立即执行, 不保证会在容器的ENTRYPOINT前执行.
- **PreStop**

## (3)hook handler实现类型:
- Exec: 执行一个指定命令, 例如: pre-stop.sh.
- HTTP: 执行一个HTTP请求.

## (4)PostStart:
- PostStart操作是个blocking调用, 容器状态会保持Waiting(pod状态为Pending)直至postStart操作完成. 
- This nature of postStart can be used to delay the startup state of the container while giving time to the main container process to initialize.
- Another use of postStart is to prevent a container from starting when the Pod does not fulfill certain preconditions. For example, when the postStart hook indicates an error by returning a nonzero exit code, the main container process gets killed by Kubernetes.

## (5)PreStop:
- Even though preStop is blocking, holding on it or returning a nonsuccessful result does not prevent the container from being deleted and the process killed. 
- PreStop is only a convenient alternative to a SIGTERM signal for graceful application shutdown

## (6)备注:
- https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/
