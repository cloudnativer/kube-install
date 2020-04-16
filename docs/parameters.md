<center><font size=5>Parameter introduction of kube-install</font></center><br>
<br>
<b>Introduction:</b><br>
<br>
The parameters about kube-install can be viewed using the "kube-install --help" command. <br>
<table width=100%>
<tr><td colspan="3" bgcolor=#000000><font color=#C0FF3E># kube-install --help</font></td></tr>
<tr><td colspan="3" bgcolor=#000000></td></tr>
<tr><td bgcolor=#000000><font color=#C0FF3E>-addnode</font></td><td bgcolor=#000000><font color=#C0FF3E> string </font></td><td bgcolor=#000000><font color=#C0FF3E>IP address of k8s node server to be added.</font></td></tr>
<tr><td bgcolor=#000000><font color=#C0FF3E>-delmaster</font></td><td bgcolor=#000000><font color=#C0FF3E> string </font></td><td bgcolor=#000000><font color=#C0FF3E>IP address of k8s master server to be deleted.</font></td></tr>
<tr><td bgcolor=#000000><font color=#C0FF3E>-delnode</font></td><td bgcolor=#000000><font color=#C0FF3E> string </font></td><td bgcolor=#000000><font color=#C0FF3E>IP address of k8s node server to be deleted.</font></td></tr>
<tr><td bgcolor=#000000><font color=#C0FF3E>-master</font></td><td bgcolor=#000000><font color=#C0FF3E> string </font></td><td bgcolor=#000000><font color=#C0FF3E>The IP address of k8s master server filled in for the first installation</font></td></tr>
<tr><td bgcolor=#000000><font color=#C0FF3E>-mvip</font></td><td bgcolor=#000000><font color=#C0FF3E> string </font></td><td bgcolor=#000000><font color=#C0FF3E>K8s master cluster virtual IP address filled in for the first installation.</font></td></tr>
<tr><td bgcolor=#000000><font color=#C0FF3E>-node</font></td><td bgcolor=#000000><font color=#C0FF3E> string </font></td><td bgcolor=#000000><font color=#C0FF3E>The IP address of k8s node server filled in for the first installation.</font></td></tr>
<tr><td bgcolor=#000000><font color=#C0FF3E>-opt</font></td><td bgcolor=#000000><font color=#C0FF3E> string </font></td><td bgcolor=#000000><font color=#C0FF3E>Available options：init | install | addnode | delnode | rebuildmaster | delmaster</font></td></tr>
<tr><td bgcolor=#000000><font color=#C0FF3E>-rebuildmaster</font></td><td bgcolor=#000000><font color=#C0FF3E> string </font></td><td bgcolor=#000000><font color=#C0FF3E>IP address of k8s master server to be rebuilt.</font></td></tr>
<tr><td bgcolor=#000000><font color=#C0FF3E>-sshpwd</font></td><td bgcolor=#000000><font color=#C0FF3E> string </font></td><td bgcolor=#000000><font color=#C0FF3E>SSH login root password of each server.</font></td></tr>
<tr><td bgcolor=#000000><font color=#C0FF3E>-k8sdir</font></td><td bgcolor=#000000><font color=#C0FF3E> string </font></td><td bgcolor=#000000><font color=#C0FF3E>Target path of k8s cluster software installation.</font></
td></tr>
<tr><td colspan="3" bgcolor=#000000></td></tr>
</table>
<br>
<br>
<b>For Example:</b><br>
<br>
Perform pre installation initialization<br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E># ./kube-install -opt init</font></td></tr>
</table>
Install k8s cluster<br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E># ./kube-install -opt install -master "192.168.122.11,192.168.122.12,192.168.122.13" -node "192.168.122.11,192.168.122.12,192.168.122.13,192.168.122.14" -mvip "192.168.122.100" -sshpwd "cloudnativer"</font></td></tr>
</table>
Add k8s-node to k8s cluster<br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E># kube-install -opt addnode -addnode "192.168.122.15,192.168.122.16" -sshpwd "cloudnativer"</font></td></tr>
</table>
Delete k8s-node from k8s cluster<br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E># kube-install -opt delnode -delnode "192.168.122.13,192.168.122.15" -sshpwd "cloudnativer"</font></td></tr>
</table>
Delete k8s-master from k8s cluster<br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E># kube-install -opt delnode -delnode "192.168.122.13,192.168.122.15" -sshpwd "cloudnativer"</font></td></tr>
</table>
Rebuild k8s-master to k8s cluster<br>
<table>
<tr><td bgcolor=#000000><font color=#C0FF3E># kube-install -opt rebuildmaster -rebuildmaster "192.168.122.13" -sshpwd "cloudnativer"</font></td></tr>
</table>
<br>
<br>
<br>





