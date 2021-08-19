package kilib

import (
    "fmt"
)



func ShowHelp(){
    fmt.Println(`Usage of kube-install: 
kube-install [OPTION] { [COMMAND] OBJECT [ARGS]... } 

OPTIONS:
  d, daemon        Run as a daemon service. (Enable this switch to use the web console for management)
  e, exec          Deploy and uninstall kubernetes cluster.(Use with "init | sshcontrol | install | addnode | delnode | delmaster | rebuildmaster | uninstall")
  h, help          Display usage help information of kube-install.
  i, init          Initialize the local system environment.
  s, showk8s       Display all deployed kubernetes cluster information.
  v, version       Display software version information of kube-install.

COMMAND:
  addnode          Add node to the kubernetes cluster.
  delmaster        Remove the master from the kubernetes cluster.
  delnode          Remove the node from the kubernetes cluster.
  install          Install kubernetes cluster.
  rebuildmaster    Rebuild the damaged kubernetes master.
  sshcontrol       Open the SSH channel from the local to the target host.(You can also get through manually)
  uninstall        Uninstall kubernetes cluster.

OBJECT:
  label            In the case of deploying and operating multiple kubernetes clusters, it is necessary to specify a label to uniquely identify a kubernetes cluster.
  listen           Set the IP and port on which the daemon service listens. (Default is "0.0.0.0:9080")
  master           The IP address of kubernetes master host.
  node             The IP address of kubernetes node host.
  ostype           Specifies the distribution OS type: "centos7 | centos8 | rhel7 | rhel8 | ubuntu20 | suse15".
  softdir          Specifies the installation directory of kubernetes cluster.(Default is "/opt/kube-install")
  sship            The IP address of the target host through which the SSH channel is opened.(Use with "sshcontrol")
  sshpass          The root password of the target host through which the SSH channel is opened.(Use with "sshcontrol")

--------------------------------------------------------------

EXAMPLES:
  Initialize the system environment:
    kube-install -init -ostype "rhel7" 
  Open the SSH channel from the local to the target host (You can also get through manually):
    kube-install -exec sshcontrol -sship "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -sshpass "cloudnativer"
  Install kubernetes cluster:
    kube-install -exec install -master "192.168.1.11,192.168.1.12,192.168.1.13" -node "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -k8sver "1.22" -ostype "rhel7" -label "192168001011"
  Add node to the kubernetes cluster:
    kube-install -exec addnode -node "192.168.1.15,192.168.1.16" -k8sver "1.22" -ostype "rhel7" -label "192168001011"
  Remove the node from the kubernetes cluster:
    kube-install -exec delnode -node "192.168.1.13,192.168.1.15" -label "192168001011"
  Remove the master from the kubernetes cluster:
    kube-install -exec delmaster -master "192.168.1.13" -label "192168001011"
  rebuild the damaged kubernetes master:
    kube-install -exec rebuildmaster -master "192.168.1.13" -k8sver "1.22" -ostype "rhel7" -label "192168001011"
  Uninstall kubernetes cluster:
    kube-install -exec uninstall -master "192.168.1.11,192.168.1.12,192.168.1.13" -node "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -label "192168001011"
  Enable this switch to use the web console for management:
    kube-install -daemon -listen "0.0.0.0:8888"
  Display software version information
    kube-install -version

=================================================================
    `)
}



