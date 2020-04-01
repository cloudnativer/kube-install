<br>
 # One key tool for rapid deployment of highly available kubernetes clusters <br>
<br>
<br>
 # Kube-install and kubernetes version correspondence: <br>
<table>
<tr><td>Kube-install version</td><td>Corresponding installation</td><td>kubernetes version</td>
<tr><td>v0.1.*</td><td>-----></td><td>v1.14.*</td>
<tr><td>v0.2.*</td><td>-----></td><td>v1.17.*</td>
</table>
<br>
<br>
<br>
<font size="4">**Installation instructions:**</font><br>
1,Install dependent tools<br>
<table>
<tr><td>yum -y install ansible git</td>
</table>
<br>
2,Download Kube install package<br>
Select a k8s-master and execute:<br>
<table>
<tr><td>
cd /root/<br>
git clone https://github.com/cloudnativer/kube-install.git <br>
cd /root/kube-install <br>
 </td>
</table>
Download the kube-install-pkg-*.*.tgz package from this link https://github.com/cloudnative/kube-install/releases <br>
<table>
<tr><td>tar -zxvf kube-install-pkg-*.*.tgz</td>
</table>
<br>
3,Deploy k8s cluster<br>
Execute on the k8s master selected above:<br>
<table>
<tr><td>ansible-playbook -i inventory k8scluster-install.yml</td>
</table>
<br>
4,Add k8s-node to k8s cluster<br>
Execute on the k8s master selected above:<br>
<table>
<tr><td>ansible-playbook -i inventory k8scluster-addnode.yml</td>
</table>
<br>
<br>
<br>

