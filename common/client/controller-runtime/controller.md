# 一 概述:
## (1)概述:
- pkg/controller提供构建controller使用的类型和函数.
- pkg/reconciler

# 二 controller:
- A Controller manages a **work queue fed reconcile.Requests** from source.Sources. 
- Work is performed through the **reconcile.Reconciler** for each enqueued item. Work typically is reads and writes Kubernetes objects to make the system state match the state specified in the object Spec.

## (2)创建:
- 为了创建新的controller, 首先要创建一个manager.Manager并将它传递给**controller.New**函数, controller必须通过调用Manager.Start来启动.
- 格式: func New(name string, mgr manager.Manager, options Options)(Controller, error), 创建一个注册到Manager的新controller, Manager保证在Controller启动前已sync共享cache.
- Options: MaxConcurrentReconciles(默认为1)表示最大并发调和的数量; Reconciler: 对象对于的Reconciler对象.

## (3)Controller结构:
- reconcile.Reconciler
- Watch: watch的资源以及对于eventHandler, 还可以指定Predicate来filter事件, 一个controller可以watch多个资源.
- Start: 启动controller, start会block直至stop被closed或有一个启动错误.

## (4)handler

# 三 reconciler:
## (1)功能:
- Package reconcile defines the Reconciler interface to implement Kubernetes APIs. Reconciler is provided to Controllers at creation time as the API implementation.

## (2)Reconciler接口:
- Reconcile(Request)(Result, err): 执行调和的函数.

## (3)Request:
- types.NamespacedName: 需要调和对象的name和namespace.
- Request包含调和对象所必须的信息(可以唯一标记一个对象), 但是它不包含任何事件和对象自己的内容.

## (4)Result:
- Requeue bool: 告诉controller重新入队, 默认为false.
- RequeueAfter time.Duration: 若大于0则表示告诉controller等待对应时间后重新入队,隐含表示Requeue为true(不用再指定Requeue).
