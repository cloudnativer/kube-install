package kilib

import (
    "strings"
)


// Serialize or convert the input parameters.
func ParameterConvert(mode string, master string, node string, softDir string, label string, osType string) ([]string, []string, string, string, string) {

    //Convert OS type
    switch {
      case osType == "centos7" :
          osType = "rhel7"
      case osType == "rhel7" :
          osType = "rhel7"
      case osType == "centos8" :
          osType = "rhel8"
      case osType == "rhel8" :
          osType = "rhel8"
      case osType == "ubuntu20" :
          osType = "ubuntu20"
      case osType == "suse15" :
          osType = "suse15"
      default:
          osType = "unknow"
    }

    //Convert masterArray and nodeArray
    masterArray := strings.Split(master, ",")
    nodeArray := strings.Split(node, ",")

    //Set the directory where kube-install is installed on the target host
    if softDir == "" {
        softDir = "/opt/kube-install"
    }

    //In the multi k8s cluster deployment scenario, use label to generate the subprocess path
    subProcessDir := "/"+label

    return masterArray,nodeArray,softDir,subProcessDir,osType

}





