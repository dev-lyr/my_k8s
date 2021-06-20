# 一 概述 
## (1)概述:
- Kubernetes支持多种授权模块, 例如: ARAC模式,RBAC模式和Webhook模式.

## (2)授权模式:
- Node 
- ABAC
- RBAC
- WebHook

## (3)备注:
- https://kubernetes.io/docs/reference/access-authn-authz/authorization/

# 二 RBAC: 
## (1)概述:
- Role-based access control: 是一种基于用户的角色要控制计算或网络等资源访问的方法.
- RBAC使用**rbac.authorization.k8s.io**API组来进行授权决定, 运行admin通过k8s API动态配置策略.

## (2)备注:
- https://kubernetes.io/docs/reference/access-authn-authz/rbac/

# 三 API:
## (1)概述:
- RBAC API声明四个类型资源: **Role**和**ClusterRole**, **RoleBinding**和**ClusterRoleBinding**.
- 角色: 包含rules, 这些rules代表一些权限(permissions)的集合;permissions是累加的(不存在deny rule); 角色可以通过**Role**定义在一个namespace内, 也可以通过**ClusterRole**定义在集群范围内.
- 角色绑定: 将定义在role中的权限绑定到一个或多个user上; 包含一个对象列表(users, groups或service accounts)和一个被授权的role的引用; 单个namespace内绑定使用RoleBinding, 集群级别使用ClusterRoleBinding.

## (2)Role和ClusterRole:
- Role: 只能在单个namespace内对资源进行访问授权.
- ClusterRole: 集群范围的访问授权, 除了Role可以干的外, 它还可以进行其他访问授权: 对集群范围资源(例如:nodes); non-resources endpoints(例如:/healthz); 跨所有namespaces的namespaced资源.

## (3)RoleBinding:
- roleRef(RoleRef): 引用一个当前namespace的Role或一个全局namespace的ClusterRole.
- subjects(Subject数组): 表示Role绑定到的对象的信息.
- Subject: kind(User,Group或ServiceAccount), apiGroup(kind为ServiceAccount时为"", User和Group时默认为"rabc.authorization.k8s.io"), name, namespaced(若对象kind是非namespace的, 例如:User或Group, 则该值需为空).
- 备注: RoleBinding的roleRef可以引用当前namespace内的Role或全局namespace内的ClusterRole; 与clusterrolebinding不同, rolebinding即使引用一个clusterrole, 也只是授权clusterrole当中包含的资源在rolebinding所在ns访问权限.
- 通过将rolebinding ref到一个clusterrole, 适用于管理员需要在整个集群中定义一些公共roles, 并且在多个namespaces中使用.

## (4)ClusterRoleBinding:
- roleRef(RoleRef): 引用一个全局namespace的ClusterRole.
- subjects: 同上.
- 备注: ClusterRoleBinding的roleRef引用全局namespace内的ClusterRole.
