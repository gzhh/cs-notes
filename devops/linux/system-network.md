## 网络

### centos 配置IP网络

1.检查网络配置

ifconfig

查看网卡：通常有两个网卡：lo 和 eth0 (或其他名称的网卡)

cd /etc/sysconfig/network-scripts/

ls

然后到 cd 到网络配置目录下查看网络对应 eth0(或其他名称网卡)的配置文件，配置文件通常为 ifcfg- 开头

vim ifcfg-

然后用 vim 命令打开该文件修改IP配置，如下(省略了一部分)

BOOTPROTO=dhcp

...

ONBOOT=yes

IPADDR0=本机IP

PREFIX0=24

GATEWAY0=255.255.255.0

DNS1=本机DNS

重启网络服务

service network restart

查看网络是否 ping 的通

ping www.163.com

2.虚拟机网络模式配置

VMnet0 桥接模式 公网、局域网都可以访问(通过网卡来通讯,有线网卡和无线)

VMnet1 仅主机模式

VMnet8 NAT模式 只可以访问公网
