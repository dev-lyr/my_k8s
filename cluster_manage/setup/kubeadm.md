# 一 概述:
## (1)功能:
- 部署多节点kube集群.

## (2)参考:
- https://v1-13.docs.kubernetes.io/docs/setup/independent/install-kubeadm/
- https://juejin.im/post/5b8a4536e51d4538c545645c#heading-15

# 二 步骤:
## (1)先安装docker-ce:
- https://docs.docker.com/install/linux/docker-ce/centos/

## (2)安装kubeadm,kubectl和kubelet:
- 使用阿里云的源.
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
        http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF
        
- setenforce 0
- yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes
- systemctl enable kubelet && systemctl start kubelet`
- 备注: 注意对齐版本, 可使用yum list kubelet --showduplicates | sort -r, 查找版本对齐.

## (3)镜像替换:
- 寻找合适国内源重命名为google的镜像名.

## (4)启动前一些配置设置:
- export KUBECONFIG=/etc/kubernetes/admin.conf

## (5)启动master

## (6)配置网络插件

## (7)启动worker
