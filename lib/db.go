package kilib

import (
	//   "fmt"
	//    "os"
	"io/ioutil"
	//    "strings"
	//    "time"
)

// Initialize the key and value of creating database (temporarily replaced by file) to prepare for later transformation into database.
func DatabaseInit(currentDir string, subProcessDir string, logName string, mode string) {
	CreateDir(currentDir+"/data/logs"+subProcessDir+"/logs", currentDir, logName, mode)
	CreateDir(currentDir+"/data/output"+subProcessDir+"/masters", currentDir, logName, mode)
	CreateDir(currentDir+"/data/output"+subProcessDir+"/nodes", currentDir, logName, mode)
	CreateDir(currentDir+"/data/output"+subProcessDir+"/cert/ssl", currentDir, logName, mode)
	CreateDir(currentDir+"/data/output"+subProcessDir+"/addons/addonsip", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/addons/addonsip/registryip.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/addons/addonsip/k8sdashboardip.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/status.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/softdir.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/softdirtemp.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/ostype.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/ostypetemp.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/k8sver.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/k8svertemp.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/k8shealth.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/etcdendpoints.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/installtime.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/scheduler.txt", currentDir, logName, mode)
	CreateFile(currentDir+"/data/output"+subProcessDir+"/progressbar.txt", currentDir, logName, mode)
}

// Update the contents of the database.
func DatabaseUpdate(key string, value string, currentDir string, logName string, mode string) error {
	err := ioutil.WriteFile(key, []byte(value), 0666)
	return err
}

// Query the master information in the cluster.
func GetClusterMaster(label string, currentDir string, logName string, mode string) map[string]string {
	var masterStatusMap map[string]string
	masterStatusMap = make(map[string]string)
	_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
	mstArrayLocal, err := GetAllDir(currentDir+"/data/output"+subProcessDir+"/masters", currentDir, logName, mode)
	CheckErr(err, currentDir, logName, mode)
	health, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/k8shealth.txt")
	for i := 0; i < len(mstArrayLocal); i++ {
		masterstu, err := ReadFile(currentDir + "/data/output" + subProcessDir + "/masters/" + mstArrayLocal[i] + "/status.txt")
		CheckErr(err, currentDir, logName, mode)
		switch {
		case masterstu == "adding", masterstu == "rebuilding", masterstu == "notok", masterstu == "notinstall":
			masterStatusMap[mstArrayLocal[i]] = masterstu
		case masterstu == "ok":
			if health == "healthy" {
				masterStatusMap[mstArrayLocal[i]] = masterstu
			} else {
				masterStatusMap[mstArrayLocal[i]] = "unknow"
			}
		case masterstu == "":
			masterStatusMap[mstArrayLocal[i]] = "notinstall"
		case masterstu == "deleting":
			masterStatusMap[mstArrayLocal[i]] = "deleting"
		default:
			masterStatusMap[mstArrayLocal[i]] = "unknow"
		}
	}
	return masterStatusMap
}

// Query the node information in the cluster.
func GetClusterNode(label string, currentDir string, logName string, mode string) map[string]string {
	var ndArray []string
	var nodeStatusMap map[string]string
	nodeStatusMap = make(map[string]string)
	_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
	health, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/k8shealth.txt")
	ndArrayLocal, err := GetAllDir(currentDir+"/data/output"+subProcessDir+"/nodes", currentDir, logName, mode)
	if health == "healthy" {
		ndArrayK8s := ListNode(label, currentDir, logName, mode)
		if ndArrayK8s == nil {
			ndArray = ndArrayLocal
		} else {
			ndArray = ndArrayK8s
			CheckErr(err, currentDir, logName, mode)
			for i := 0; i < len(ndArrayLocal); i++ {
				if !StrInArray(ndArrayLocal[i], ndArrayK8s) {
					ndArray = append(ndArray, ndArrayLocal[i])
				}
			}
		}
		for j := 0; j < len(ndArray); j++ {
			nodestu, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/nodes/" + ndArray[j] + "/status.txt")
			if ndArrayK8s != nil && StrInArray(ndArray[j], ndArrayK8s) {
				switch {
				case nodestu == "notok", nodestu == "ok", nodestu == "notinstall", nodestu == "":
					nodeStatusMap[ndArray[j]] = "ok"
				case nodestu == "adding", nodestu == "deleting":
					nodeStatusMap[ndArray[j]] = nodestu
				default:
					nodeStatusMap[ndArray[j]] = "unknow"
				}
			} else {
				switch {
				case nodestu == "ok", nodestu == "unknow", nodestu == "":
					nodeStatusMap[ndArray[j]] = "unknow"
				case nodestu == "notok", nodestu == "notinstall", nodestu == "adding", nodestu == "deleting":
					nodeStatusMap[ndArray[j]] = nodestu
				default:
					nodeStatusMap[ndArray[j]] = "unknow"
				}
			}
		}
	} else {
		for i := 0; i < len(ndArrayLocal); i++ {
			nodestu, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/nodes/" + ndArrayLocal[i] + "/status.txt")
			switch {
			case nodestu == "notinstall", nodestu == "":
				nodeStatusMap[ndArrayLocal[i]] = "notinstall"
			case nodestu == "notok", nodestu == "notinstall", nodestu == "adding", nodestu == "deleting":
				nodeStatusMap[ndArrayLocal[i]] = nodestu
			default:
				nodeStatusMap[ndArrayLocal[i]] = "unknow"
			}
		}
	}
	return nodeStatusMap
}

// Query the kubeconfig in the cluster.
func GetClusterKubecfg(label string, currentDir string, mode string) string {
	_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
	kubecfg, err := ReadFile(currentDir + "/data/output" + subProcessDir + "/cert/ssl/kube-install.kubeconfig")
	if err != nil || kubecfg == "" {
		return ""
	} else {
		return kubecfg
	}
}

// Query the addons information in the cluster.
func GetClusterAddons(label string, currentDir string, mode string) (string, string, string) {
	_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
	registryip, err_registryip := ReadFile(currentDir + "/data/output" + subProcessDir + "/addons/addonsip/registryip.txt")
	k8sdashboardip, err_k8sdashboardip := ReadFile(currentDir + "/data/output" + subProcessDir + "/addons/addonsip/k8sdashboardip.txt")
	k8sdashboardtoken, err_k8sdashboardtoken := ReadFile(currentDir + "/data/output" + subProcessDir + "/cert/dashboard_login_token.txt")
	if err_registryip == nil && err_k8sdashboardip == nil && err_k8sdashboardtoken == nil {
		return registryip, k8sdashboardip, k8sdashboardtoken
	} else {
		return "", "", ""
	}
}

// Query the directory information of cluster installation.
func GetClusterSoftdir(label string, currentDir string, mode string) string {
	_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
	kdr, err := ReadFile(currentDir + "/data/output" + subProcessDir + "/softdir.txt")
	if err != nil {
		return ""
	} else {
		return kdr
	}
}

// Query the status information in the cluster.
func GetClusterStatus(label string, currentDir string, logName string, mode string) (string, string) {
	var sch string
	var err_sch error
	_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
	sch, err_sch = ReadFile(currentDir + "/data/output" + subProcessDir + "/scheduler.txt")
	if err_sch != nil {
		sch = "off"
	}
	stu, err := ReadFile(currentDir + "/data/output" + subProcessDir + "/status.txt")
	if err == nil {
		switch {
		case stu == "installing", stu == "restarting", stu == "uninstalling", stu == "notok", stu == "notinstall":
			return stu, sch
		case stu == "ok":
			health, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/k8shealth.txt")
			if health == "healthy" {
				return stu, sch
			} else {
				return "unknow", sch
			}
		case stu == "":
			return "notinstall", sch
		default:
			return "unknow", sch
		}
	} else {
		return "notinstall", sch
	}
}

// Query the OS type of cluster installation.
func GetClusterOstype(label string, currentDir string, mode string) string {
	_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
	ostype, err := ReadFile(currentDir + "/data/output" + subProcessDir + "/ostype.txt")
	if err != nil {
		return ""
	} else {
		return ostype
	}
}

// Query the version of kubernetes cluster installation.
func GetClusterK8sVer(label string, currentDir string, mode string) string {
	_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
	k8sver, err := ReadFile(currentDir + "/data/output" + subProcessDir + "/k8sver.txt")
	if err != nil {
		return ""
	} else {
		return k8sver
	}
}

