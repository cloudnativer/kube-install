<br>
<font size="5">**One key tool for rapid deployment of highly available kubernetes clusters**</font><br>
<br>
<br>
<font size="4">**kube-install与kubernetes版本对应关系：**</font><br>
<table>
<tr><td>kube-install版本</td><td>对应安装</td><td>kubernetes版本</td>
<tr><td>v0.1.*</td><td>对应安装</td><td>v1.14.*</td>
<tr><td>v0.2.*</td><td>对应安装</td><td>v1.17.*</td>
</table>


<br>
<br>
<br>


<font size="4">**安装说明：**</font><br>
1，ansible的安装
选择一个k8s-master，然后执行：
yum -y install ansible

2，下载kube-install包
在上一步选定的那个k8s-master上执行：
cd /root/
wget xxxx
tar -zxvf kube-install.tgz


cd /root/kube-install
wget xxxx
tar -zxvf kube-install-pkg-0.1.tgz
mv kube-install-pkg pkg

3，部署k8s集群
在上面选定的那个k8s-master上执行：
ansible-playbook -i inventory k8scluster-install.yml


4，往k8s集群追加k8s-node节点
在上面选定的那个k8s-master上执行：
ansible-playbook -i inventory k8scluster-addnode.yml 

<br>
<br>
<br>

