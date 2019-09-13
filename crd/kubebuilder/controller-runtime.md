# 一 概述:
## (1)功能:
- https://godoc.org/sigs.k8s.io/controller-runtime

## (2)Managers
- 每个controller和webhook最后都通过一个Manager(pkg/manager)来运行.
- manager: 追踪所有controller的运行, 同时建立共享caches和到api server的clients; 启动manager时, 它会接着启动所有controller和webhook, manager一直运行直至接收一个graceful shutdown信号.

## (3)Controllers:
- Controller(pkg/controller)使用事件(pkg/events)来trigger reconcile请求.

## (4)Reconcilers:
- 控制器的逻辑以Reconcilers(pkg/reconcile)的形式实现.
- Reconciler实现一个函数, 该函数使用一个包含需要被调和对象的namespace和name的reconcile请求, 调和该对象, 然后返回一个Response或这个错误表示是否重新入队列进行第二轮处理.

## (5)Client和Caches:
- Reconcilers使用Clients(pkg/client)来访问API对象.
- manager默认提供的client从local shared cache(pkg/cache)中读取对象且直接写入API Server, 但是client也可以构造成直接访问API Server, 不使用cache.
- The Cache will auto-populate with watched objects, as well as when other structured objects are requested.

## (6)Schemes:
- Clients, Caches, and many other things in Kubernetes use Schemes (pkg/scheme) to associate Go types to Kubernetes API Kinds (Group-Version-Kinds, to be specific).
- GroupVersionKind(GVK): kind in a particular group-verison.

## (7)webhooks

## (8)Logging和Metrics

## (9)Testing

# 二 Manager
## (1)概述:
- Package manager is required to create Controllers and provides shared dependencies such as clients, caches, schemes, etc. Controllers must be started by calling Manager.Start.
- 创建方法: manager.New, 参数为rest.Config和Options.

## (2)Manager结构:
- Start
- GetConfig
- GetScheme
- GetClient: 返回一个config配置好的client, client可能会从cache读, 由Options.NewClient决定.
- GetCache
- GetApiReader: 返回一个Reader, 该Reader使用API server, 该对象应该节约使用并且在client不满足场景时使用.
- GetWebhookServer
- 等等.

## (3)rest.Config:
- config.GetConfig: 创建一个和kubernete API server交互的rest.Config, 若--kubeconfig设置, 则会使用在该目录下的kubeconfig文件; 否则假设运行在集群中并使用集群提供的kubeconfig.

## (4)Options:
- Scheme: 用来解析runtime.Object到GVK或资源, 默认是kubernetes/client-go scheme.Scheme, 最好传递自己的scheme.
- MapperProvider: 提供rest mapper用来map go types到k8s APIs.
- SyncPeriod: 设置被watch资源被调和的最小frequency, 默认是10小时.
- LeaderElection: 是否使用领导者选举.
- LeaderElectionName: 领导竞选configMap被创建的namespace.
- LeaderElectionID: 用于竞选表示hold leader lock的configMap的name
- LeaseDuration: 等待强制获得leadership的时间, 默认是15s.
- RenewDeadline: acting master尝试refreshing leadership的持续时间(放弃前), 默认10s.
- RetryPeriod: Leader竞选者等待重试的间隔, 默认为2s.
- Namespace: 限制manager的cache去watch指定namespace的对象, 默认是所有namespace.
- MetricsBindAddress: controller用来服务permetheus metrics需绑定的TCP地址.
- Host和Port: webhook服务器需要绑定的host和port.
- NewCache cache.NewCacheFunc
- NewClient NewClientFunc: 创建一个manager使用的client, 若不设置, 则会创建一个DelegatingClient(使用cache来读和client来写).

## (5)备注:
- https://godoc.org/sigs.k8s.io/controller-runtime/pkg/manager

# 三 Controller:
## (1)概述:
- pkg/controller提供构建controller使用的类型和函数.
- A Controller manages a **work queue fed reconcile.Requests** from source.Sources. Work is performed through the **reconcile.Reconciler** for each enqueued item. Work typically is reads and writes Kubernetes objects to make the system state match the state specified in the object Spec.

## (2)创建:
- 为了创建新的controller, 首先要创建一个manager.Manager并将它传递给**controller.New**函数, controller必须通过调用Manager.Start来启动.
- 格式: func New(name string, mgr manager.Manager, options Options)(Controller, error), 创建一个注册到Manager的新controller, Manager保证在Controller启动前已sync共享cache.
- Options: MaxConcurrentReconciles(默认为1)表示最大并发调和的数量; Reconciler: 对象对于的Reconciler对象.

## (3)Controller结构:
- reconcile.Reconciler
- Watch: watch的资源以及对于eventHandler, 还可以指定Predicate来filter事件, 一个controller可以watch多个资源.
- Start: 启动controller, start会block直至stop被closed或有一个启动错误.

## (4)handler

# 四 reconciler:
## (1)功能:
- Package reconcile defines the Reconciler interface to implement Kubernetes APIs. Reconciler is provided to Controllers at creation time as the API implementation.

## (2)Reconciler接口:
- Reconcile(Request)(Result, err): 执行调和的函数.

## (3)Request:
- types.NamespacedName: 需要调和对象的name和namespace.
- Request包含调和对象所必须的信息(可以唯一标记一个对象), 但是它不包含任何事件和对象自己的内容.

## (4)Result:
- Requeue bool: 告诉controller重新入队, 默认为false.
- RequeueAfter time.Duration

