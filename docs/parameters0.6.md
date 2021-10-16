<center><font size=5>Parameter introduction of kube-install</font></center><br>
<br>
<b>Introduction:</b><br>
<br>
The parameters about kube-install can be viewed using the `kube-install help` or `kube-install -opt help` command. <br>
<table width=100%>
<tr><td>
 
 ```
  # kube-install help
 ```
  
</td></tr>
<tr><td></td></tr>
<tr><td>

```
Usage of kube-install: -opt [OPTIONS] COMMAND [ARGS]...

Options: 
  init             Initialize the system environment.
  install          Install kubernetes cluster.
  addnode          Add k8s-node to the cluster.
  delnode          Remove the k8s-node from the cluster.
  rebuildmaster    Rebuild the damaged k8s-master.
  delmaster        Remove the k8s-master from the cluster.
  uninstall        Uninstall kubernetes cluster.
  help             Display help information.

Commands:
  master           The IP address of k8s-master server.
  node             The IP address of k8s-node server.
  sshpwd           SSH login root password of each server.
  ostype           Specifies the distribution OS type: centos7|centos8|rhel7|rhel8|suse15.
```

</td></tr>
<tr><td></td></tr>
<tr><td>

```
For example：
  Initialize the system environment:
    kube-install -opt init
  Install k8s cluster:
    kube-install -opt install -master "192.168.1.11,192.168.1.12,192.168.1.13" -node "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -sshpwd "cloudnativer" -ostype "centos7"
  Add k8s-node to the cluster:
    kube-install -opt addnode -node "192.168.1.15,192.168.1.16" -sshpwd "cloudnativer" -ostype "centos7"
  Remove the k8s-node from the cluster:
    kube-install -opt delnode -node "192.168.1.13,192.168.1.15" -sshpwd "cloudnativer"
  Remove the k8s-master from the cluster:
    kube-install -opt delmaster -master "192.168.1.13" -sshpwd "cloudnativer"
  Rebuild the damaged k8s-master:
    kube-install -opt rebuildmaster -master "192.168.1.13" -sshpwd "cloudnativer" -ostype "centos7"
  Uninstall k8s cluster:
    kube-install -opt uninstall -master "192.168.1.11,192.168.1.12,192.168.1.13" -node "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -sshpwd "cloudnativer"
  Display help information:
    kube-install -opt help
```

</td></tr>
<tr><td></td></tr>
</table>
<br>
<br>
<br>


