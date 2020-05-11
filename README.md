This is a one click rapid deployment tool for highly available kubernetes clusters.
<br>

![avatar](docs/images/kube-install-logo.jpg)

<br>
<b>[1] Corresponding relation: </b><br>
<br>
Kube-install and kubernetes version correspondence:
<table>
<tr><td>kube-install Version</td><td>Corresponding Relation</td><td>Kubernetes Version</td>
<tr><td>v0.1.* <br>v0.2.* </td><td>  -----> </td><td>v1.14.* </td></tr>
<tr><td>v0.3.* </td><td>  -----> </td><td>v1.17.* </td></tr>
</table>
<br>
<br>
<br>
<b>[2] How to install?</b><br>
<br>
2.1 Download kube-install file<br>
Select a k8s-master and execute:<br>
<table>
<tr><td bgcolor=#000000>
<font color=#C0FF3E># cd /root/</font><br>
<font color=#C0FF3E># git clone https://github.com/cloudnativer/kube-install.git </font><br>
<font color=#C0FF3E># cd /root/kube-install/</font><br>
 </td></tr>
</table>
<br>
2.2 Download the kube-install-pkg-0.1.tgz package from this link https://github.com/cloudnativer/kube-install/releases <br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E>
 # cd /root/kube-install/<br>
 # tar -zxvf kube-install-pkg-0.1.tgz
 </font></td></tr>
</table>
<br>
2.3 Initialization<br>
Perform pre installation initialization<br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E># cd /root/kube-install/<br> # ./kube-install -opt init</font></td></tr>
</table>
<br>
2.4 Install k8s cluster<br>
If your server environment is as follows:<br>
<table>
<tr><td>IP Address</td><td>Role</td><td>OS Version</td><td>Root Password</td></tr>
<tr><td>192.168.1.11</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.12</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.13</td><td>k8s-master,k8s-node</td><td>CentOS Linux release 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.14</td><td>k8s-node</td><td>CentOS Linux release 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.15</td><td>k8s-node</td><td>CentOS Linux release 7</td><td>cloudnativer</td></tr>
<tr><td>192.168.1.16</td><td>k8s-node</td><td>CentOS Linux release 7</td><td>cloudnativer</td></tr>
</table>
Well,Execute on the k8s-master selected above:<br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E># cd /root/kube-install/<br> # ./kube-install -opt install -master "192.168.1.11,192.168.1.12,192.168.1.13" -node "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -mvip "192.168.1.88" -sshpwd "cloudnativer"</font></td></tr>
</table>
Note: in the above command, the "-mvip" parameter is the k8s cluster virtual IP address.<br>
<br>
<br>
<br>
<b>[3] Operation and maintenance:</b><br>
<br>
3.1 Add k8s-node to k8s cluster<br>
Select any k8s-master server, and execute the following command on it:<br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E># kube-install -opt addnode -node "192.168.1.15,192.168.1.16" -sshpwd "cloudnativer"</font></td></tr>
</table>
<br>
3.2 Delete k8s-node from k8s cluster<br>
Select any k8s-master server, and execute the following command on it:<br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E># kube-install -opt delnode -node "192.168.1.13,192.168.1.15" -sshpwd "cloudnativer"</font></td></tr>
</table>
<br>
3.3 Delete k8s-master from k8s cluster<br>
Select any k8s-master server, and execute the following command on it:<br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E># kube-install -opt delmaster -master "192.168.1.13" -sshpwd "cloudnativer"</font></td></tr>
</table>
<br>
3.4 Rebuild k8s-master to k8s cluster<br>
Select any k8s-master server, and execute the following command on it:<br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E># kube-install -opt rebuildmaster -master "192.168.1.13" -sshpwd "cloudnativer"</font></td></tr>
</table>
<br>
<br>
<br>
<b>[4] Parameter introduction:</b><br>
<br>
The parameters about kube-install can be viewed using the "kube-install help" or "kube-install -opt help" command. <a href="docs/parameters0.2.md">You can also see more detailed parameter introduction here.</a><br>
<br>
<br>
<br>

