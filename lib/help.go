package kilib

import (
    "fmt"
)



func ShowHelp(){
    fmt.Println("kube-install version 0.7.0-beta (Creation Date: 6/17/2021)\n=================================================================\n\nUsage of kube-install: -opt [OPTION] COMMAND [ARGS]...\n\nOPTIONS:\n  init             Initialize the system environment.\n  install          Install kubernetes cluster.\n  delnode          Remove the node from the kubernetes cluster.\n  addnode          Add node to the kubernetes cluster.\n  delmaster        Remove the master from the kubernetes cluster.\n  rebuildmaster    Rebuild the damaged kubernetes master.\n  uninstall        Uninstall kubernetes cluster.\n  help             Display help information.\n\nCOMMANDS:\n  master           The IP address of kubernetes master server.\n  node             The IP address of kubernetes node server.\n  sshpwd           The root password used to SSH login to each server.\n  ostype           Specifies the distribution OS type: centos7|centos8|rhel7|rhel8|suse15.\n  softdir          Specifies the installation path of kubernetes cluster.(Default value is \"/opt/kube-install\")\n\n--------------------------------------------------------------\n\nEXAMPLES:\n  Initialize the system environment:\n    kube-install -opt init -ostype \"rhel7\" \n  Install kubernetes cluster:\n    kube-install -opt install -master \"192.168.1.11,192.168.1.12,192.168.1.13\" -node \"192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14\" -sshpwd \"cloudnativer\" -ostype \"rhel7\" \n  Add node to the kubernetes cluster:\n    kube-install -opt addnode -node \"192.168.1.15,192.168.1.16\" -sshpwd \"cloudnativer\" -ostype \"rhel7\" \n  Remove the node from the kubernetes cluster:\n    kube-install -opt delnode -node \"192.168.1.13,192.168.1.15\" -sshpwd \"cloudnativer\"\n  Remove the master from the kubernetes cluster:\n    kube-install -opt delmaster -master \"192.168.1.13\" -sshpwd \"cloudnativer\"\n  Rebuild the damaged kubernetes master:\n    kube-install -opt rebuildmaster -master \"192.168.1.13\" -sshpwd \"cloudnativer\" -ostype \"rhel7\" \n  Uninstall kubernetes cluster:\n    kube-install -opt uninstall -master \"192.168.1.11,192.168.1.12,192.168.1.13\" -node \"192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14\" -sshpwd \"cloudnativer\"\n  Display help information:\n    kube-install -opt help\n    kube-install help\n\n=================================================================\n")
}


