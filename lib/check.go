package kilib

import (
//    "fmt"
    "net"
    "time"
    "os"
)



func CheckErr(err error, currentDir string, logName string, mode string) {
    logStr := LogStr(mode)
    if err != nil {
        if mode == "DAEMON"{
            ShellExecute("echo [Error] "+time.Now().String()+" An error occurred: "+err.Error()+logStr+currentDir+"/data/logs/kubeinstalld/"+logName+".log")
            return
        } else {
            panic(err)
        }
    }
}

func CheckIP(ipv4 string) bool {
    address := net.ParseIP(ipv4)  
    if address == nil {
        return false
    } else {
        return true
    }
}

func CheckOS(CompatibleOS string, osType string, currentDir string, logName string, mode string) {
    logStr := LogStr(mode)
    if osType == "unknow" || osType == "" {
        if mode == "DAEMON" {
            ShellExecute("echo [Info] "+time.Now().String()+" \"The \"ostype\" parameter you entered is incorrect, please check! \n\""+logStr+currentDir+"/data/logs/kubeinstalld/"+logName+".log")
            return
        } else {
            panic("Please make sure that the \"-ostype\" parameter you entered is correct! Only supports "+CompatibleOS+" versions of \"ostype\": \n--------------------------------------------------------\n    rhel7   --> Red Hat Enterprise Linux 7 \n    rhel8   --> Red Hat Enterprise Linux 8 \n    centos7 --> CentOS Linux 7 \n    centos8 --> CentOS Linux 8 \n    ubuntu20 --> Ubuntu Server 20 \n    suse15  --> OpenSUSE 15 \n\n ")
        }
    } else {
        return
    }
}

func CheckK8sVersion(Version string, CompatibleK8S string, k8sVer string, currentDir string, logName string, mode string) {
    logStr := LogStr(mode)
    if k8sVer == "1.17" || k8sVer == "1.18" || k8sVer == "1.19" || k8sVer == "1.20" || k8sVer == "1.21" || k8sVer == "1.22" || k8sVer == "1.22" {
        return
    } else {
        if mode == "DAEMON" {
            ShellExecute("echo [Info] "+time.Now().String()+" \"The \"k8sver\" parameter you entered is incorrect, please check! \n\""+logStr+currentDir+"/data/logs/kubeinstalld/"+logName+".log")
            return
        } else {
            panic("Please make sure that the \"-k8sver\" parameter you entered is correct! \n--------------------------------------------------------------------------\nKube-Install "+Version+" only supports "+CompatibleK8S+" versions of kubernetes. \n\nNotice: If you want to install the old version(1.14, 1.15, 1.16) of kubernetes, you can use the historical release of kube-install.\n")
        }
    }
}

func CheckPort(port int) bool {
    if ( port <= 1 ) || ( port >= 65535 ) {
        return false
    } else {
        return true
    }
}

func CheckParam(option string, paramName string, param string) {
    if param == "" {
         panic("When performing "+option+" operation, you must enter "+paramName+" parameter, please check!")
    }
}

func CheckFileExist(path string,fileName string, currentDir string, logName string, mode string) {
    logStr := LogStr(mode)
    _, err := os.Stat(path+fileName)
    if err != nil || os.IsNotExist(err) {
        if mode == "DAEMON"{
            ShellExecute("echo [Error] "+time.Now().String()+" An error occurred: "+fileName+"File generation failed!"+logStr+"/data/logs/kubeinstalld/"+logName+".log")
            return
        } else {
            panic(fileName+"File generation failed!")
        }
    }
}


