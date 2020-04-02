<center><b>kube-install description:</b></center><br>
This is a one click rapid deployment tool for highly available kubernetes clusters.
<br>

![avatar](docs/images/kube-install-logo.jpg)

<br>
<br>
<br><br>
<b> Kube-install and kubernetes version correspondence: </b><br>
<table>
<tr><td>Kube-install version</td><td>Corresponding installation</td><td>kubernetes version</td>
<tr><td>v0.1.*</td><td>-----></td><td>v1.14.*</td>
<tr><td>v0.2.*</td><td>-----></td><td>v1.17.*</td>
</table>
<br>
<br>
<br>
<b>Installation instructions:</b><br>
<br>
<br>
1.Install dependent tools<br>
<table>
<tr><td>yum -y install ansible git</td>
</table>
<br>
2.Set SSH password free login<br>
Select a k8s-master and execute:<br>
<table>
<tr><td>
cat <<EOF> hostname.txt
192.168.122.11 22 123456789
192.168.122.12 22 123456789
192.168.122.13 22 123456789
192.168.122.14 22 123456789
192.168.122.15 22 123456789
EOF
cat hostname.txt | while read ip port pawd;do sshpass -p $pawd ssh-copy-id -p $port root@$ip;done
sed -i '/StrictHostKeyChecking/s/^#//; /StrictHostKeyChecking/s/ask/no/' /etc/ssh/ssh_config
systemctl restart sshd 
</td>
</table>
<br>
3.Download Kube install package<br>
Execute on the k8s-master selected above:<br>
<table>
<tr><td>
cd /opt/<br>
git clone https://github.com/cloudnativer/kube-install.git <br>
cd /opt/kube-install <br>
 </td>
</table>
Download the kube-install-pkg-*.*.tgz package from this link https://github.com/cloudnative/kube-install/releases <br>
<table>
<tr><td>
 cd /opt/kube-install <br>
 tar -zxvf kube-install-pkg-*.*.tgz<br></td>
</table>
<br>
4.Deploy k8s cluster<br>
Execute on the k8s-master selected above:<br>
<table>
<tr><td>
 cd /opt/kube-install <br>
 ansible-playbook -i inventory k8scluster-install.yml <br></td>
</table>
<br>
5.Add k8s-node to k8s cluster<br>
Execute on the k8s-master selected above:<br>
<table>
<tr><td>
 cd /opt/kube-install <br>
 ansible-playbook -i inventory k8scluster-addnode.yml <br></td>
</table>
<br>
<br>
<br>

