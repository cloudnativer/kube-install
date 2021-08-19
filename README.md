One click fast installation of highly available kubernetes cluster, as well as addition of kubernetes node, deletion of kubernetes node, destruction of kubernetes master, rebuild of kubernetes master, and uninstallation of cluster in later operation and maintenance stage.
<br>

![kube-install](docs/images/kube-install-logo.jpg)


# [1] Compatibility

<br>
Compatibility matrix:
<table>
<tr><td><b>kube-install Version</b></td><td><b>Supported Kubernetes Version</b></td><td><b>Supported OS Version</b></td></tr>
<tr><td> kube-install v0.7.* </td><td> kubernetes v1.22 , kubernetes v1.21 , kubernetes v1.20 , <br> kubernetes v1.19 , kubernetes v1.18 </td><td> CentOS 7 , RHEL 7 , <br> CentOS 8 , RHEL 8 , <br> SUSE Linux 15 , <br> Ubuntu 20 </td></tr>
<tr><td> kube-install v0.6.* </td><td> kubernetes v1.22 , kubernetes v1.21 , kubernetes v1.20 , <br> kubernetes v1.19 , kubernetes v1.18 , kubernetes v1.17 , <br> kubernetes v1.16 , kubernetes v1.15 , kubernetes v1.14 </td><td> CentOS 7 , RHEL 7 , <br> CentOS 8 , RHEL 8 , <br> SUSE Linux 15 </td></tr>
<tr><td> kube-install v0.5.* </td><td> kubernetes v1.21 , kubernetes v1.20 , kubernetes v1.19 , <br> kubernetes v1.18 , kubernetes v1.17 , kubernetes v1.16 , <br> kubernetes v1.15 , kubernetes v1.14 </td><td> CentOS 7 , RHEL 7 </td></tr>
<tr><td> kube-install v0.4.* </td><td> kubernetes v1.21 , kubernetes v1.20 , kubernetes v1.19 , <br> kubernetes v1.18 , kubernetes v1.17 , kubernetes v1.16 , <br> kubernetes v1.15 , kubernetes v1.14 </td><td> CentOS 7 , RHEL 7 </td></tr>
<tr><td> kube-install v0.3.* </td><td> kubernetes v1.18 , kubernetes v1.17 , kubernetes v1.16 , <br> kubernetes v1.15 , kubernetes v1.14 </td><td>CentOS 7</td></tr>
<tr><td> kube-install v0.2.* </td><td> kubernetes v1.14 </td><td> CentOS 7 </td></tr>
<tr><td> kube-install v0.1.* </td><td> kubernetes v1.14 </td><td> CentOS 7 </td></tr>
</table>


<br>
Notice: kube-install supports CentOS 7, CentOS 8, SUSE 15, RHEL 7 and RHEL 8 operating system environments. For a list of supported operating system distributions, please refer to <a href="docs/os-support.md">OS support list</a>.
<br>
<br>
<br>

# [2] How to install?

<br>

If you have four servers,kubernetes master software is installed on the three servers (192.168.1.11, 192.168.1.12, 192.168.1.13), and kubernetes node software is installed on the four servers (192.168.1.11, 192.168.1.12, 192.168.1.13, 192.168.1.14). The operating system of the server is pure CentOS Linux or RHEL(Red Hat Enterprise Linux). It's like this:
<table>
<tr><td><b>IP Address</b></td><td><b>Role</b></td><td><b>OS Version</b></td><td><b>Root Password</b></td></tr>
<tr><td>192.168.1.11</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.12</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.13</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.14</td><td>k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
</table>
You expect the architecture after installation to be as follows:

![kube-install-arch](docs/images/kube-install-arch-1.jpg)

<br>

## 2.1 Download kube-install package file

<br>

You can download the `kube-install-*.tgz` package from https://github.com/cloudnativer/kube-install/releases. <br>

For example, we have downloaded the `kube-install-v0.7.0-beta2.tgz` package.<br>


```
# cd /root/
# curl -O https://github.com/cloudnativer/kube-install/releases/download/v0.7.0-beta2/kube-install-v0.7.0-beta2.tgz
# tar -zxvf kube-install-v0.7.0-beta2.tgz
# cd /root/kube-install/
```

<br>

## 2.2 Initialize system environment

<br>
Please operate in the root user environment. Perform the system environment initialization operation on the k8s-master selected above: <br>

```
# cd /root/kube-install/
# ./kube-install -opt init -ostype "centos7"
```

Notice: Please make sure that the `-ostype` flag you entered is correct, only support `rhel7`, `rhel8`, `centos7`, `centos8`, `ubuntu20`, `suse15` these types of "ostype".<br>
In addition, if you need to specify the directory path to the Kubernetes cluster installation, you can set it using the `-softdir` parameter.

<br>

## 2.3 open the SSH password free channel

<br>
Before using the web platform for installation, please open the SSH password free channel from localhost to the target host.

You can open the SSH password free channel by manually, or through the `kube-install -exec sshcontrol` command.<br>

```
# cd /root/kube-install/
# ./kube-install -exec sshcontrol -sship "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -sshpass "cloudnativer"
```

Or click the `Open SSH Channel of Host` button in the web platform to SSH through. Here is the process of SSH connection, <a href="webssh0.7.md">click here to view more details</a> !<br>

<br>

## 2.3 Install kubernetes cluster

<br>
Please operate in the root user environment. Execute on the k8s-master selected above:<br>

```
# cd /root/kube-install/
# ./kube-install -opt install -master "192.168.1.11,192.168.1.12,192.168.1.13" -node "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -sshpwd "cloudnativer" -ostype "centos7"
```

Notice: Please make sure that the `-ostype` flag you entered is correct, only support `rhel7`, `rhel8`, `centos7`, `centos8`, `ubuntu20`, `suse15` these types of "ostype".<br>
In addition, if you need to specify the directory path to the Kubernetes cluster installation, you can set it using the `-softdir` parameter.

<br>


## 2.4 Login kubernetes dashboard UI

<br>
Execute the following command on the k8s-master you selected to view the kube-dashboard console URL and key:<br>

```
# cat /opt/kube-install/loginkey.txt
```


![loginkey](docs/images/loginkey2.jpg)

Login to the kube-dashboard console UI using the URL and key in the `/opt/kube-install/loginkey.txt` document.Here are the relevant screenshots:

![kube-dashboard](docs/images/kube-dashboard3.jpg)


![kube-dashboard](docs/images/kube-dashboard4.jpg)

<br>

## 2.5 Use the web platform to install

<br>
You can also install the Kubernetes cluster through the Kube-Install web platform. 

Before using the web platform for installation, please open the SSH password free channel from localhost to the target host.You can use the `kube-install -exec sshcontrol` command to SSH through, or click the `Open SSH Channel of Host` button in the upper right corner to SSH through. Here is the process of SSH connection, <a href="webssh0.7.md">click here to view more details</a> !<br>

![kube-dashboard](docs/images/webinstall001.jpg)

Click the `Install Kubernetes` button in the upper right corner to start the installation operation.

![kube-dashboard](docs/images/webinstall002.jpg)

For the installation process using the web platform, <a href="docs/webinstall0.7.md">click here to view more details</a> !<br>

<br>
<br>
<br>

# [3] Operation and maintenance

<br>

After kube-install is installed, you can directly execute `kube-install` command in any directory of any k8s-master server.Kube-install can not only quickly install the highly available kubernetes cluster, but also add k8s-node, delete k8s-node, delete k8s-master and rebuild k8s-master.<br>

Suppose you expect to install two servers (192.168.1.15 and 192.168.1.16) as k8s-nodes and join the kubernetets cluster in Chapter [2].

<table>
<tr><td><b>IP Address</b></td><td><b>Role</b></td><td><b>OS Version</b></td><td><b>Root Password</b></td></tr>
<tr><td>192.168.1.11</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.12</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.13</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.14</td><td>k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td><b>192.168.1.15</b></td><td>k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td><b>192.168.1.16</b></td><td>k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
</table>

Select any k8s-master server, and execute the following command on it:<br>

```
# kube-install -opt addnode -node "192.168.1.15,192.168.1.16" -sshpwd "cloudnativer" -ostype "centos7"
```

Notice: Please make sure that the `-ostype` flag you entered is correct, only support `rhel7`, `rhel8`, `centos7`, `centos8`, `ubuntu20`, `suse15` these types of "ostype".<br>
In addition, if you need to specify the directory path to add Kubernetes node, you can set it using the `-softdir` parameter.

<br>

The architecture after installation is shown in the following figure:

![kube-install-arch](docs/images/kube-install-arch-2.jpg)

<br>
You can also install the Kubernetes cluster through the Kube-Install web platform. For the installation process using the web platform, <a href="docs/webinstall0.7.md">click here to view more details</a> ! <br>

![kube-dashboard](docs/images/webinstall002.jpg)

<br>

Notice: you can <a href="docs/operation0.6.md">click here to view more operation documents</a> about add k8s-node, delete k8s-node, delete k8s-master, rebuild k8s-master, and uninstall cluster.

<br>
<br>


# [4] Parameter introduction

<br>

The parameters about kube-install can be viewed using the `kube-install help` command. You can also <a href="docs/parameters0.7.md">see more detailed parameter introduction here</a>.<br>

<br>
<br>


# [5] How to build it?

<br>

The build can be completed automatically by executing the `make` command. You can also <a href="docs/build.md">see more detailed build instructions here</a>.<br>

<br>
<br>


# [6] How to Contribute

Fork it <br>
Create your feature branch (git checkout -b my-new-feature) <br>
Commit your changes (git commit -am 'Add some feature') <br>
Push to the branch (git push origin my-new-feature) <br>
Create new Pull Request <br>
<br>
<br>



