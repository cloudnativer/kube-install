
# Add Node, Delete Node, Rebuild Master, and Uninstall

<br>

If you have four servers, kubernetes master software is installed on the three servers (192.168.1.11, 192.168.1.12, 192.168.1.13), and kubernetes node software is installed on the three servers (192.168.1.11, 192.168.1.12, 192.168.1.13, 192.168.1.14). <br>
<table>
<tr><td><b>IP Address</b></td><td><b>Role</b></td><td><b>OS Version</b></td><td><b>Root Password</b></td></tr>
<tr><td>192.168.1.11</td><td>k8s-master,k8s-node,kube-install</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.12</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.13</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.14</td><td>k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
</table>

The current kubernetes cluster architecture is as follows:

![kube-install-arch](images/kube-install-arch-1.jpg)

<br>

Next, we will carry out daily operation and maintenance operations such as add-k8s node, delete k8s-node, delete k8s-master and rebuild k8s-master.

<br>

## Add k8s-node to k8s cluster

<br>
You will install two servers (192.168.1.15 and 192.168.1.16) as k8s-node and join the kubernetets cluster.
<table>
<tr><td><b>IP Address</b></td><td><b>Role</b></td><td><b>OS Version</b></td><td><b>Root Password</b></td></tr>
<tr><td>192.168.1.11</td><td>k8s-master,k8s-node,kube-install</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.12</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.13</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.14</td><td>k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td><b>192.168.1.15</b></td><td>k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
<tr><td><b>192.168.1.16</b></td><td>k8s-node</td><td>CentOS Linux release 7 or Red Hat Enterprise Linux(RHEL) 7</td><td>cloudnativer</td></tr>
</table>

Execute the following command on kube-install host:<br>

```
# kube-install -exec addnode -node "192.168.1.15,192.168.1.16" -k8sver "1.22" -ostype "centos7" -label "192168001011"
```

Notice: Please make sure that the `-ostype` flag you entered is correct, only support `rhel7`, `rhel8`, `centos7`, `centos8`, `ubuntu20`, `suse15` these types of "ostype".<br>
If you need to specify the directory path to add Kubernetes node, you can set it using the `-softdir` parameter.<br>

![kube-dashboard](images/webnodeadd001.jpg)

In addition, you can also add Kubernetes node using the Kube-Install web platform. For the installation process using the web platform, <a href="docs/webinstall0.7.md">click here to view more details</a> ! <br>

<br>

The architecture after installation is shown in the following figure:

![kube-install-arch](images/kube-install-arch-2.jpg)

<br>

## Delete k8s-node from k8s cluster

<br>
You will delete the two k8s-nodes (192.168.1.15 and 192.168.1.16) from the kubernetets cluster.
Execute the following command on kube-install host:<br>

```
# kube-install -exec delnode -node "192.168.1.13,192.168.1.15" -label "192168001011"
```

Notice: If you specify the `-softdir` parameter value during the installation or addnode operation, please specify the same `-softdir` parameter value during the delnode operation.<br>

![kube-dashboard](images/webnodeadd001.jpg)

In addition, you can also delete Kubernetes node using the Kube-Install web platform. For the installation process using the web platform, <a href="docs/webinstall0.7.md">click here to view more details</a> ! <br>

<br>

The architecture after installation is shown in the following figure:

![kube-install-arch](images/kube-install-arch-3.jpg)

<br>

## Delete k8s-master from k8s cluster

<br>
You will Delete the k8s-master (192.168.1.13) from the kubernetets cluster.
Execute the following command on kube-install host:<br>

```
# kube-install -exec delmaster -master "192.168.1.13" -label "192168001011"
```

Notice: If you specify the `-softdir` parameter value during the installation operation, please specify the same `-softdir` parameter value during the delmaster operation.<br>
In addition, you can also delete Kubernetes master using the Kube-Install web platform. For the installation process using the web platform, <a href="docs/webinstall0.7.md">click here to view more details</a> ! <br>

The architecture after installation is shown in the following figure:

![kube-install-arch](images/kube-install-arch-4.jpg)

<br>

## Rebuild k8s-master to k8s cluster

<br>
You will rebuild the damaged k8s-master (192.168.1.13) in the kubernetets cluster.
Execute the following command on kube-install host:<br>

```
# kube-install -exec rebuildmaster -rebuildmaster "192.168.1.13" -k8sver "1.22" -ostype "centos7" -label "192168001011"
```

Notice: Please make sure that the `-ostype` flag you entered is correct, only support `rhel7`, `rhel8`, `centos7`, `centos8`, `ubuntu20`, `suse15` these types of "ostype".<br>
If you specify the `-softdir` parameter value during the installation operation, please specify the same `-softdir` parameter value during the rebuildmaster operation.<br>
In addition, you can also rebuild Kubernetes master using the Kube-Install web platform. For the installation process using the web platform, <a href="docs/webinstall0.7.md">click here to view more details</a> ! <br>

<br>

The architecture after installation is shown in the following figure:

![kube-install-arch](images/kube-install-arch-5.jpg)

<br>


## Uninstall kubernetes cluster

<br>
You will uninstall kubernetets cluster.
Execute the following command on kube-install host:<br>

```
# kube-install -exec uninstall -master "192.168.1.11,192.168.1.12,192.168.1.13" -node "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -label "192168001011"
```

Notice: If you specify the `-softdir` parameter value during the installation operation, please specify the same `-softdir` parameter value during the uninstall operation.<br>

![kube-dashboard](images/webinstall002.jpg)

In addition, you can also uninstall the Kubernetes cluster using the Kube-Install web platform. For the installation process using the web platform, <a href="docs/webinstall0.7.md">click here to view more details</a> ! 

<br>
<br>
<br>
