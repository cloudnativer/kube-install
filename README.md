<br>
<font size="5">**One key tool for rapid deployment of highly available kubernetes clusters**</font><br>
<br>
<br>
<font size="4">**Kube-install and kubernetes version correspondence:**</font><br>
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
<br>
# yum -y install ansible git <br>

2,Download Kube install package<br>


Select a k8s-master and execute:<br>
# cd /root/<br>
# git clone https://github.com/cloudnativer/kube-install.git <br>
# cd /root/kube-install <br>

Download the kube-install-pkg-*.*.tgz package from this link https://github.com/cloudnative/kube-install/releases <br>
# tar -zxvf kube-install-pkg-*.*.tgz <br>


3,Deploy k8s cluster<br>
Execute on the k8s master selected above:<br>
# ansible-playbook -i inventory k8scluster-install.yml<br>


4,Add k8s-node to k8s cluster<br>
Execute on the k8s master selected above:<br>
ansible-playbook -i inventory k8scluster-addnode.yml <br>

<br>
<br>
<br>

