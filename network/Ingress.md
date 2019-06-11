# 一 概述:
## (1)功能:
- Ingress是一个API对象, 暴露集群外部到集群内部服务的HTTP和HTTPS路由, 流量路由被定义在Ingress资源中的规则控制.
- Ingress可以为服务提供一个外部可到达的URLS, **load balancing**, **SSL termination**和 **name-based virtual hosting**.
- 流量方向: Internet->Ingress->服务.

## (2)对外暴露服务的方式:
- Ingress: Http/Https方式.
- 若以非HTTP和HTTPS方式暴露服务到外部, 通常使用服务类型Service.Type=Node.Port和Service.Type=LoadBalancer.

## (3)备注:
- https://kubernetes.io/docs/concepts/services-networking/ingress/

# 二 Ingress控制器:
## (1)概述:
- 为了Ingress资源能够work, 集群必须有一个ingress控制器运行, Ingress控制器不是自动启动的, 需选择集群最佳的ingress控制器实现方案.
- k8s当前支持和维护GCE和nginx controller.
- 可以在一个集群中部署多个Ingress控制器.

## (2)其它控制器:
- istio
- haproxy
- 等等.

## (3)备注:
- nginx controller: https://github.com/kubernetes/ingress-nginx/blob/master/README.md
- https://kubernetes.io/docs/concepts/services-networking/ingress-controllers

# 三 Ingress资源:
## (1)概述:
- Ingress通常使用注解来配置一些选项, 不同Ingress控制器支持不同的配置.
- Ingress的spec有配置loadbalancer或proxy服务器所需的所有信息, 还包含一些**rules**列表, 用来匹配进来的请求, Ingress资源支持持redirect HTTP流量的规则.

## (2)Rules信息:
- 一个可选的host.
- 一个paths的列表: 例如:/testpath.
- 一个**backend**: 是一个服务和端口名字的组合, HTTP(和HTTPS)请求匹配了该rule的host和path, 则请求会被发送到backend列表.

## (3)默认backend:
- 默认backend通常在Ingress控制器中配置, 那些没有match到spec中host和path的请求会被路由到default backend.

# 四 Ingress的类型:
## (1)单服务Ingress:
- 使用一个指定默认backend且无rules的Ingress.

## (2)Simple fanout:
- fanount噢诶之路由**单个Ip地址**的流量到多个服务, 基于请求HTTP URI.

## (3)Name based virtual hosting:
- 基于名字的虚拟host支持路由HTTP流量到多个多个host名字(一样的IP).

## (4)TLS:
- 可以通过指定一个包含TLS私有key和certificate的**secret**来使Ingress变安全.

## (5)Loadbalancing.

# 五 ingress-nginx:
## (1)概述:
## (2)备注:
- https://kubernetes.github.io/ingress-nginx/
- https://github.com/kubernetes/ingress-nginx
