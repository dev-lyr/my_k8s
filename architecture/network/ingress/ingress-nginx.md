# 一 概述:
## (1)概述:
- ingress-nginx是一个Ingress controller, 使用nginx作为反向代理和负载平衡.

## (2)部署:
- bare-metal.
- 各云厂商
- https://kubernetes.github.io/ingress-nginx/deploy/

## (3)备注:
- https://github.com/kubernetes/ingress-nginx
- https://github.com/nginxinc/kubernetes-ingress
- 备注: 注意不同.

# 二 裸机部署:
## (1)概述:
- 部署: kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-0.32.0/deploy/static/provider/baremetal/deploy.yaml, 使用node port来暴露nginx controller.

## (2)备注:
- https://kubernetes.github.io/ingress-nginx/deploy/baremetal/
