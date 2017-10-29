###打通Docker容器以及Docker与分布主机网络
一、准备环境：
1、安装etcd
根据不同系统下载最新的etcd二进制安装包文件
https://github.com/coreos/etcd/releases/download/v3.2.1/etcd-v3.2.1-linux-arm64.tar.gz

解压压缩包
tar -xvf etcd-v3.2.1-linux-arm64.tar.gz
将解压包文件：etcd 、etcdctl 使用root用户拷贝到/usr/bin目录下，并保证有root运行权限

配置etcd集群，参照：
http://blog.csdn.net/u010511236/article/details/52386229

目前采用静态配置、例如：
`#!/bin/bash
etcd --name infra0 --initial-advertise-peer-urls http://192.168.1.147:2380 \
--listen-peer-urls http://192.168.1.147:2380 \
--listen-client-urls http://192.168.1.147:2379,http://127.0.0.1:2379 \
--advertise-client-urls http://192.168.1.147:2379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster infra0=http://192.168.1.147:2380,infra1=http://192.168.1.182:2380,infra2=http://192.168.1.197:2380 \
--initial-cluster-state new
`
配置完后使用etcdctl客户端工具，检查：
 etcdctl  member list
 查看配置成员信息
 
 etcdctl cluster-health
 检测个个节点的健康状况
 
 确定etcd可以使用之后，我们需要设置分配给docker网络的网段
etcdctl mk /coreos.com/network/config '{"Network":"172.17.0.0/16", "SubnetMin": "172.17.1.0", "SubnetMax": "172.17.254.0"}'

使用命令： etcdctl ls /coreos.com/network/subnets
/coreos.com/network/subnets/172.17.66.0-24
/coreos.com/network/subnets/172.17.12.0-24
/coreos.com/network/subnets/172.17.32.0-24
liuhy@liuhy ~ $ etcdctl get /coreos.com/network/subnets/172.17.66.0-24
{"PublicIP":"192.168.1.147"}
liuhy@liuhy ~ $ etcdctl get /coreos.com/network/subnets/172.17.12.0-24
{"PublicIP":"192.168.1.182"}
liuhy@liuhy ~ $ etcdctl get /coreos.com/network/subnets/172.17.32.0-24
{"PublicIP":"192.168.1.197"}
可以看到加入etcd集群中分配的地址信息

2、安装flanneld
去地址下载最新的版本：flannel-v0.7.1
https://github.com/coreos/flannel/releases/download/v0.7.1/flannel-v0.7.1-linux-amd64.tar.gz

将解压后的文件放置到/usr/bin目录下
使用root启动flanneld程序
flanneld -etcd-endpoints=http://192.168.1.147:2379

其中192.168.1.147地址是本地etcd客户端监听地址
如果启动成功，会创建虚拟网卡
`flannel0  Link encap:UNSPEC  HWaddr 00-00-00-00-00-00-00-00-00-00-00-00-00-00-00-00  
          inet addr:172.17.66.0  P-t-P:172.17.66.0  Mask:255.255.0.0
          UP POINTOPOINT RUNNING NOARP MULTICAST  MTU:1472  Metric:1
          RX packets:32 errors:0 dropped:0 overruns:0 frame:0
          TX packets:32 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:500 
          RX bytes:2688 (2.6 KB)  TX bytes:2688 (2.6 KB)
`

使用flanneld脚本文件生成docker启动参数：
运行：
 /usr/bin/mk-docker-opts.sh -i -d docker_opts.env -c

3、配置docker启动参数
编辑docker启动配置文件：
docker 服务配置文件：
vi /lib/systemd/system/docker.service
修改：
EnvironmentFile=/home/liuhy/docker_opts.env

使配置生效：
systemctl daemon-reload
重启docker
systemctl restart docker

如果配置成功docker0会配置172.17.66网段的一个地址
`docker0   Link encap:Ethernet  HWaddr 02:42:f8:24:9f:c7  
          inet addr:172.17.66.1  Bcast:0.0.0.0  Mask:255.255.255.0
          inet6 addr: fe80::42:f8ff:fe24:9fc7/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1472  Metric:1
          RX packets:95 errors:0 dropped:0 overruns:0 frame:0
          TX packets:103 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:7172 (7.1 KB)  TX bytes:10523 (10.5 KB)
`

二 、参照网络资源：
http://www.jianshu.com/p/a2039a8855ec
http://www.linuxidc.com/Linux/2016-01/127784.htm
http://blog.csdn.net/qq_32971807/article/details/54693254
http://www.linuxidc.com/Linux/2016-01/127784.htm
http://blog.csdn.net/qq_32971807/article/details/54693254
https://my.oschina.net/dxqr/blog/607854
http://www.dockerinfo.net/%E9%AB%98%E7%BA%A7%E7%BD%91%E7%BB%9C%E9%85%8D%E7%BD%AE
https://mritd.me/2016/09/03/Dokcer-%E4%BD%BF%E7%94%A8-Flannel-%E8%B7%A8%E4%B8%BB%E6%9C%BA%E9%80%9A%E8%AE%AF/


三、
最后

三台机器都配置好了之后，我们在三台机器上分别开启一个docker容器，测试它们的网络是不是通的。

docker run -ti centos bash
一次查看容器IP
cat /etc/hosts
172.17.97.2     334cec104721
测试连通性，都成功就OK了
到跨物理机容器
ping -c 1 172.16.164.7
ping -c 1 172.17.67.1
到宿主机
ping -c 1 172.16.164.7
到别的物理机
ping -c 1 172.16.164.6










