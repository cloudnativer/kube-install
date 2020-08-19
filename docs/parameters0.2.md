<center><font size=5>Parameter introduction of kube-install</font></center><br>
<br>
<b>Introduction:</b><br>
<br>
The parameters about kube-install can be viewed using the "kube-install help" or "kube-install -opt help" command. <br>
<table width=100%>
<tr><td bgcolor=#000000><font color=#C0FF3E># kube-install -opt help</font></td></tr>
<tr><td bgcolor=#000000></td></tr>
<tr><td bgcolor=#000000><font color=#C0FF3E>
<pre>
Usage of kube-install: -opt [OPTIONS] COMMAND [ARGS]...<br>
<br>
Options: <br>
<br>
  init             Initialize the system environment.<br>
  install          Install k8s cluster.<br>
  addnode          Add k8s-node to the cluster.<br>
  delnode          Remove the k8s-node from the cluster.<br>
  rebuildmaster    Rebuild the damaged k8s-master.<br>
  delmaster        Remove the k8s-master from the cluster.<br>
  help             Display help information.<br>
<br>
Commands:<br>
<br>
  master           The IP address of k8s-master server.<br>
  mvip             K8s-master cluster virtual IP address.<br>
  node             The IP address of k8s-node server.<br>
  sshpwd           SSH login root password of each server.<br>
 </pre>
</font></td></tr>
<tr><td bgcolor=#000000></td></tr>
<tr><td bgcolor=#000000><font color=#C0FF3E>
<pre>
For example：<br>
<br>
  Initialize the system environment:<br>
    kube-install -opt init<br>
  Install k8s cluster:<br>
    kube-install -opt install -master "192.168.1.11,192.168.1.12,192.168.1.13" -node "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -mvip "192.168.1.88" -sshpwd "cloudnativer"<br>
  Add k8s-node to the cluster:<br>
    kube-install -opt addnode -node "192.168.1.15,192.168.1.16" -sshpwd "cloudnativer"<br>
  Remove the k8s-node from the cluster:<br>
    kube-install -opt delnode -node "192.168.1.13,192.168.1.15" -sshpwd "cloudnativer"<br>
  Remove the k8s-master from the cluster:<br>
    kube-install -opt delmaster -master "192.168.1.13" -sshpwd "cloudnativer"<br>
  Rebuild the damaged k8s-master:<br>
    kube-install -opt rebuildmaster -master "192.168.1.13" -sshpwd "cloudnativer"<br>
  Display help information:<br>
    kube-install -opt help<br>
  </pre>
</font></td></tr>
<tr><td bgcolor=#000000></td></tr>
</table>
<br>
<br>
<br>


