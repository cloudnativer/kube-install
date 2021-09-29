二进制方式一键离线安装高可用的kubernetes集群，同时支持添加kubernetes节点、删除kubernetes节点、销毁kubernetes主机、重建kubernetes主机、卸载集群等。
<br>
(不需要在目标主机上安装任何软件，只需要有纯净的裸机即可离线完成高可用kubernetes集群的部署！)
<br>

![kube-install](docs/images/kube-install-logo.jpg)

<br>

切换语言： <a href="README0.7.md">English Documents</a> | <a href="README0.7-zh.md">中文文档</a>

<br>

# [1] 兼容性

<br>
兼容性说明:

<table>
<tr><td><b>kube-install版本</b></td><td><b>支持的Kubernetes版本</b></td><td><b>支持的操作系统</b></td><td><b>相关文档</b></td></tr>
<tr><td> kube-install v0.7.* </td><td> kubernetes v1.23, v1.22, v1.20, v1.19, v1.18, v1.17 </td><td> CentOS 7 , RHEL 7 , CentOS 8 , RHEL 8 , SUSE Linux 15 , Ubuntu 20 </td><td><a href="README0.7.md">README0.7.md</a></td></tr>
<tr><td> kube-install v0.6.* </td><td> kubernetes v1.22, v1.21, v1.20, v1.19, v1.18, v1.17, v1.16, v1.15, v1.14 </td><td> CentOS 7 , RHEL 7 , CentOS 8 , RHEL 8 , SUSE Linux 15 </td><td><a href="README0.6.md">README0.6.md</a></td></tr>
<tr><td> kube-install v0.5.* </td><td> kubernetes v1.21, v1.20, v1.19, v1.18, v1.17, v1.16, v1.15, v1.14 </td><td> CentOS 7 , RHEL 7 </td><td><a href="README0.5.md">README0.5.md</a></td></tr>
<tr><td> kube-install v0.4.* </td><td> kubernetes v1.21, v1.20, v1.19, v1.18, v1.17, v1.16, v1.15, v1.14 </td><td> CentOS 7 , RHEL 7 </td><td><a href="README0.4.md">README0.4.md</a></td></tr>
<tr><td> kube-install v0.3.* </td><td> kubernetes v1.18, v1.17, v1.16, v1.15, v1.14 </td><td>CentOS 7</td><td><a href="README0.3.md">README0.3.md</a></td></tr>
<tr><td> kube-install v0.2.* </td><td> kubernetes v1.14 </td><td> CentOS 7 </td><td><a href="README0.2.md">README0.2.md</a></td></tr>
<tr><td> kube-install v0.1.* </td><td> kubernetes v1.14 </td><td> CentOS 7 </td><td><a href="README0.1.md">README0.1.md</a></td></tr>
</table>

<br>
注意：kube-install支持CentOS 7、CentOS 8、SUSE 15、RHEL 7和RHEL 8操作系统环境，<a href="docs/os-support.md">点击这里查看kube-install所支持的操作系统发行版的列表</a>。
<br>
<br>
<br>

# [2] 快速安装kubernetes集群

<br>

如果你有四台服务器，k8s-master安装在三台服务器（192.168.1.11、192.168.1.12、192.168.1.13）上，k8s-node安装在四台服务器（192.168.1.11、192.168.1.12、192.168.1.13、192.168.1.14）上。服务器的操作系统是纯净的CentOS Linux或RHEL（RedHat Enterprise Linux），具体如下表所示：
<table>
<tr><td><b>IP地址</b></td><td><b>需要安装的组件</b></td><td><b>操作系统版本</b></td><td><b>root密码</b></td></tr>
<tr><td>192.168.1.11</td><td>k8s-master,k8s-node,kube-install</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.12</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.13</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.14</td><td>k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
</table>
安装后的部署架构如下图所示：

![kube-install-arch](docs/images/kube-install-arch-1.jpg)

<br>
注意：这里使用192.168.1.11作为kube-install源安装机。事实上，您可以将任何主机(包括kubernetes集群之外的任何主机)用来作为kube-install源安装机！
<br>

## 2.1 下载kube-install软件包

<br>

你可以从 https://github.com/cloudnativer/kube-install/releases 这里下载`kube-install-*.tgz`软件包。 <br>
举例，下载`kube-install-allinone-v0.7.0.tgz`软件包进行安装：<br>

```
# cd /root/
# curl -O https://github.com/cloudnativer/kube-install/releases/download/v0.7.0-beta2/kube-install-allinone-v0.7.0.tgz
# tar -zxvf kube-install-allinone-v0.7.0.tgz
# cd /root/kube-install/
```

注意：由于软件包在Github上，如果你本地的网络环境不是太好的话，建议你使用支持断点续传的下载软件进行软件包下载，这样可以获得更好的下载体验。

<br>

## 2.2 初始化系统环境

<br>
首先你需要使用root用户对kube-install源安装机本地环境进行初始化操作，进入解压后的软件目录执行`-exec init`命令：<br>

```
# cd /root/kube-install/
# ./kube-install -exec init -ostype "centos7"
```

注意：kube-install软件支持`rhel7`, `rhel8`, `centos7`, `centos8`, `ubuntu20`, `suse15`等版本的操作系统，在做初始化操作的时候，请确保`-ostype`参数设置正确。<br>

<br>

## 2.3 打通到目标主机的SSH通道

<br>
在你开始给目标主机安装kubernetes集群之前，请先打通kube-install源安装机本地到目标主机的SSH免密通道。
你可以自己手工打通到目标主机的SSH通道，也可以通过`kube-install -exec sshcontrol`命令来打通：<br>

```
# cd /root/kube-install/
# ./kube-install -exec sshcontrol -sship "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -sshpass "cloudnativer"
```

你也可以通过kube-install的Web管理平台来打通到目标主机的SSH通道，<a href="docs/webssh0.7.md">点击这里查看使用Web管理平台打通SSH通道的方法</a>！<br>

<br>

## 2.4 一键安装部署kubernetes集群

<br>
在kube-install源安装机上使用root用户执行下面这台命令即可：

```
# cd /root/kube-install/
# ./kube-install -exec install -master "192.168.1.11,192.168.1.12,192.168.1.13" -node "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -k8sver "1.22" -ostype "centos7" -label "192168001011"
```

注意：kube-install软件支持`rhel7`, `rhel8`, `centos7`, `centos8`, `ubuntu20`, `suse15`等版本的操作系统，在做安装部署操作的时候，请确保`-ostype`参数设置正确。<br>
另外，如果你需要自定义制定Kubernetes集群安装在目标主机上的目录路径的话，可以带上`-softdir`参数来设置。

<br>


## 2.5 登录kubernetes dashboard界面

<br>
通过查看loginkey.txt文件可以获取kube-dashboard的登录地址和密钥。<br>

```
# cat /opt/kube-install/loginkey.txt
```


![loginkey](docs/images/loginkey2.jpg)

如下面的截图所示为kube-dashboard的登录地址和密钥：

![kube-dashboard](docs/images/kube-dashboard3.jpg)


![kube-dashboard](docs/images/kube-dashboard4.jpg)

<br>

# [3] 通过Web平台安装kubernetes集群

<br>
你也可以通过kube-install的Web管理平台来安装kubernetes集群。kube-install的Web管理平台具备SSH打通、定时安装部署、Node扩容、Master修复、集群卸载等强大的功能，你可以在Web管理平台上获得更好的安装体验。
<br>

## 运行kube-install的Web管理服务

首先，你需要执行`kube-install -exec init`命令初始化系统环境(如果前面已经初始化过了可以跳过)，然后执行`systemctl start kube-install`命令来运行kube-install的Web管理平台服务。

```
# cd /root/kube-install/
# ./kube-install -exec init -ostype "centos7"
#
# systemctl start kube-install.service
#
# systemctl status kube-install.service
  ● kube-install.service - kube-install One click fast installation of highly available kubernetes cluster.
     Loaded: loaded (/etc/systemd/system/kube-install.service; disabled; vendor preset: disabled)
     Active: active (running) since Fri 2021-08-20 14:30:55 CST; 21min ago
       Docs: https://cloudnativer.github.io/
   Main PID: 2768 (kube-install)
     CGroup: /system.slice/kube-install.service
             └─2768 /go/src/kube-install/kube-install -daemon
   ...

```

此时，你就可以使用网页浏览器打开`http://kube-install源安装机IP:9080`，访问kube-install的Web管理平台了。
<br>
注意：kube-install的Web管理平台服务默认监听`TCP 9080`。如果你想修改这个监听地址的话，可以通过修改`/etc/systemd/system/kube-install.service`文件中的`kube-install -daemon -listen ip:port`参数来进行设置，<a href="docs/systemd0.7.md">点击这里可以查看详细文档</a>！<br>

## 在Web界面上安装kubernetes

然后，点击Web界面右上角的的`Install Kubernetes`按钮开始kubernetes集群的安装。

![kube-dashboard](docs/images/webinstall001.jpg)

注意：在你开始给目标主机安装kubernetes集群之前，请先打通kube-install源安装机本地到目标主机的SSH免密通道。
你可以自己手工打通到目标主机的SSH通道，也可以点击右上角的`Open SSH Channel of Host`按钮来进行打通，<a href="docs/webssh0.7.md">点击这里可以查看更加详细的文档</a>。

![kube-dashboard](docs/images/webinstall002.jpg)

你可以<a href="docs/webinstall0.7-zh.md">点击这里可以查看更多通过kube-install的Web管理平台安装部署的详细信息</a>。
<br>
<br>
<br>

# [4] 扩容与销毁Node|修复Master|卸载集群

<br>

Kube-install不仅可以很方便的安装单机和高可用的kubernetes集群，还可以支持k8s-node的扩容与销毁、k8s-master的销毁与修复、kubernetes集群的卸载等。<br>

举例，现在需要给第[2]章节中安装好的kubernetets集群，增加2个k8s-node节点(192.168.1.15 and 192.168.1.16)，相关信息如下：

<table>
<tr><td><b>IP地址</b></td><td><b>需要安装的组件</b></td><td><b>操作系统版本</b></td><td><b>root密码</b></td></tr>
<tr><td>192.168.1.11</td><td>k8s-master,k8s-node,kube-install</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.12</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.13</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.14</td><td>k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td><b>192.168.1.15</b></td><td>k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td><b>192.168.1.16</b></td><td>k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
</table>

在kube-install源主机上使用root用户执行如下命令：<br>

```
# kube-install -exec addnode -node "192.168.1.15,192.168.1.16" -k8sver "1.22" -ostype "centos7" -label "192168001011"
```

注意：kube-install软件支持`rhel7`, `rhel8`, `centos7`, `centos8`, `ubuntu20`, `suse15`等版本的操作系统，在做安装部署操作的时候，请确保`-ostype`参数设置正确。<br>
另外，如果你需要自定义制定Kubernetes集群安装在目标主机上的目录路径的话，可以带上`-softdir`参数来设置。

<br>

安装完毕之后的部署架构如下图所示：

![kube-install-arch](docs/images/kube-install-arch-2.jpg)

<br>

除了使用`kube-install -exec addnode`命令进行k8s-node节点扩容外，你也同样可以使用kube-install的Web管理平台来对k8s-node节点进行扩容。<a href="docs/webinstall0.7-zh.md">点击这里可以查看使用kube-install的Web管理平台来扩容k8s-node节点的方法</a>。

![kube-dashboard](docs/images/webnodeadd001.jpg)

<br>

注意：你可以<a href="docs/operation0.7-zh.md">点击这里查看更多关于销毁k8s-node和k8s-master、修复k8s-master、卸载集群的操作</a>。

<br>
<br>


# [5] kube-install命令行帮助

<br>

你可以执行`kube-install -help`命令查看kube-install的使用帮助文档，<a href="docs/parameters0.7.md">你也可以点击这里查看更加详细的命令行帮助文档</a>。

<br>
<br>


# [6] 欢迎提交Issues和PR

如果你在使用过程中遇到问题，可以点击https://github.com/cloudnativer/kube-install/issues向我们提交Issues，也可以Fork源代码，然后尝试修复BUG之后，向我们提交PR。<br>
<br>

```
# git clone your-fork-code
# git checkout -b your-new-branch
# git commit -am "Fix bug or add some feature"
# git push origin your-new-branch
```

<br>
欢迎给我们提交Issues和PR。
<br>
谢谢每一位贡献者！

<br>
<br>
<br>



