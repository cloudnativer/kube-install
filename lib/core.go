package kilib

import (
    "fmt"
    "os"
    "time"
)


// Rebuild the core operation part of the master.
func RebuildMasterCore(mode string, masterArray []string, currentDir string, kissh string, subProcessDir string, currentUser string, label string, softDir string, logName string) {
    opt := "rebuildmaster"
    logStr := LogStr(mode)
    CreateDir(currentDir+"/data/output"+subProcessDir, currentDir, logName, mode)
    os.OpenFile(currentDir+"/data/logs"+subProcessDir+"/logs/rebuildmaster.log", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
    ShellExecute("echo \"*************************************************************************************\n\n[Info] "+time.Now().String()+" Rebuilding kubernetes master, please wait ...\n\n    Kubernetes cluster label: "+label+"\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/rebuildmaster.log")
    ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
    ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" Rebuilding kubernetes master of "+label+" cluster ...</div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    sshPort := GetSshPort(label, currentDir, mode)
    if !RebuildmasterConfig(mode, masterArray, currentDir, subProcessDir, logName, sshPort) {
        ShellExecute("echo [Error] "+time.Now().String()+" \"The parameters you entered are incorrect, please check! \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/rebuildmaster.log")
        return
    }
    kubeApiPort := GetkubeApiPort(label, currentDir, mode)
    _, err := os.Stat(currentDir+"/data/output"+subProcessDir+"/sys/0x000certificate/copycfssl/templates/ipvsinit_ep.yaml")
    if err != nil || os.IsNotExist(err) {
        ShellExecute("cp -rf "+currentDir+"/sys "+currentDir+"/data/output"+subProcessDir+"/")
        InstallIpvsYaml(mode, currentDir, masterArray, kubeApiPort, subProcessDir, logName)
    }
    InstallGenfile("", mode, currentDir, subProcessDir, logName)
    RebuildmasterYML("",currentDir+"/data/output"+subProcessDir, currentDir, currentUser, logName)
    for i := 0; i < len(masterArray); i++ {
        CreateDir(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i], currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/ttystatus.txt", currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "rebuilding", currentDir, logName, mode)
    }
    err_rebuildmaster := ExecuteOpt(kissh, currentDir, opt, opt, subProcessDir, "")
    if err_rebuildmaster != nil {
        for i := 0; i < len(masterArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "notok", currentDir, logName, mode)
        }
        ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes master rebuild failed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/rebuildmaster.log")
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Failed to rebuild master of "+label+" cluster! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    } else {
        notokMaster := ""
        for i := 0; i < len(masterArray); i++ {
            chkMasterStr := ShellOutput("curl -k -s --head 'https://"+masterArray[i]+":"+kubeApiPort+"' | awk '/HTTP/ {print $2}' ")
            fmt.Println(chkMasterStr)
            if chkMasterStr == "401\n" {
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "ok", currentDir, logName, mode)
            } else {
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "notok", currentDir, logName, mode)
                notokMaster = notokMaster+" "+masterArray[i]
            }
        }
        if notokMaster == "" {
            ShellExecute("echo [Info] "+time.Now().String()+" \"Kubernetes master rebuild operation completed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/rebuildmaster.log")
            ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
            ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" The master of "+label+" cluster has been rebuilt successfully! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
        } else {
            ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes master rebuild failed! Please repair it manually or solve the etcd problem manually and try again. \n\n    List of failed masters: "+notokMaster+" \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/rebuildmaster.log")
            ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
            ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Failed to rebuild master of "+label+" cluster! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
        }
    }
}

// Delete the core operation part of the master.
func DeleteMasterCore(mode string, masterArray []string, currentDir string, kissh string, subProcessDir string, currentUser string, label string, softDir string, logName string) {
    opt := "delmaster"
    logStr := LogStr(mode)
    CreateDir(currentDir+"/data/output"+subProcessDir, currentDir, logName, mode)
    os.OpenFile(currentDir+"/data/logs"+subProcessDir+"/logs/delmaster.log", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
    ShellExecute("echo \"*************************************************************************************\n\n[Info] "+time.Now().String()+" Deleting kubernetes master, please wait ...\n\n    Kubernetes cluster label: "+label+"\n\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/delmaster.log")
    ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
    ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" Deleting kubernetes master of "+label+" cluster ... </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    sshPort := GetSshPort(label, currentDir, mode)
    if !DelmasterConfig(mode, masterArray, currentDir, subProcessDir, logName, sshPort) {
        ShellExecute("echo [Error] "+time.Now().String()+" \"The parameters you entered are incorrect, please check! \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/delmaster.log")
        return
    }
    DelmasterYML("",currentDir+"/data/output"+subProcessDir, currentDir, currentUser, logName, false)
    for i := 0; i < len(masterArray); i++ {
        CreateDir(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i], currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "deleting", currentDir, logName, mode)
    }
    err_delmaster := ExecuteOpt(kissh, currentDir, opt, opt, subProcessDir, "")
    if err_delmaster != nil {
        for i := 0; i < len(masterArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "notok", currentDir, logName, mode)
        }
        ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes master delete failed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/delmaster.log")
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Failed to delete master of "+label+" cluster! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    } else {
        for i := 0; i < len(masterArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "notinstall", currentDir, logName, mode)
        }
        ShellExecute("echo [Info] "+time.Now().String()+" \"Kubernetes master delete operation completed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/delmaster.log")
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" The master of "+label+" cluster has been deleted successfully! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    }
}

// Add the core operation part of the node.
func AddNodeCore(mode string, node string, nodeArray []string, currentDir string, kissh string, subProcessDir string, currentUser string, label string, softDir string, osTypeResult string, k8sVer string, logName string, CompatibleOS string, upgradeKernel string) {
    opt := "addnode"
    logStr := LogStr(mode)
    upgradeKernelStr := "\n    Automatically Upgrade OS Kernel: Not Support"
    if osTypeResult == "centos7" || osTypeResult == "rhel7" {
        if upgradeKernel == "yes"{
            upgradeKernelStr = "\n    Automatically Upgrade OS Kernel: YES"
        } else {
            upgradeKernelStr = "\n    Automatically Upgrade OS Kernel: NO"
        }
    }
    CheckOS(CompatibleOS, osTypeResult, currentDir, logName, mode)
    CreateDir(currentDir+"/data/output"+subProcessDir, currentDir, logName, mode)
    sshPort := GetSshPort(label, currentDir, mode)
    os.OpenFile(currentDir+"/data/logs"+subProcessDir+"/logs/addnode.log", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
    ShellExecute("echo \"*************************************************************************************\n\n[Info] "+time.Now().String()+" Adding kubernetes node, please wait ... \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/addnode.log")
    ShellExecute("echo \"    Kubernetes Cluster Label: "+label+"\n    Kubernetes Node: "+node+"\n    SSH Operation Port: "+sshPort+"\n    Operating System Type: "+osTypeResult+ upgradeKernelStr +"\n    System User for Operation: "+currentUser+"\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/addnode.log")
    ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
    ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" Adding kubernetes node to "+label+" cluster ... </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    if !AddnodeConfig(mode, nodeArray, currentDir, subProcessDir, logName, sshPort) {
        for i := 0; i < len(nodeArray); i++ {
            os.RemoveAll(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i])
        }
        ShellExecute("echo [Error] "+time.Now().String()+" \"The parameters you entered are incorrect, please check! \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/addnode.log")
        return
    }
    AddnodeYML("",currentDir+"/data/output"+subProcessDir,currentDir,currentUser,logName,upgradeKernel,osTypeResult,k8sVer)
    InstallGenfile(osTypeResult, mode, currentDir, subProcessDir, logName)
    InstallPreShell(osTypeResult, mode, currentDir, subProcessDir, logName)
    for i := 0; i < len(nodeArray); i++ {
        CreateDir(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i], currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i]+"/status.txt", currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i]+"/ttystatus.txt", currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i]+"/status.txt", "adding", currentDir, logName, mode)
    }
    err_addnode := ExecuteOpt(kissh, currentDir, opt, opt, subProcessDir, "")
    if err_addnode != nil {
        for i := 0; i < len(nodeArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i]+"/status.txt", "notok", currentDir, logName, mode)
        }
        ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes node add failed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/addnode.log")
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Failed to add node of "+label+" cluster! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    } else {
        for i := 1; i <= 3; i++ {
            runStatus,_,_,_,_,_,_,_ := GetNodeInfo(label, nodeArray[0], currentDir, logName, mode)
            if runStatus == "True" {
                for i := 0; i < len(nodeArray); i++ {
                    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i]+"/status.txt", "ok", currentDir, logName, mode)
                }
                ShellExecute("echo [Info] "+time.Now().String()+" \"The node of "+label+" cluster has been added successfully! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/addnode.log")
                ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
                ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" The node of "+label+" cluster has been added successfully! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
                return
            } 
            time.Sleep(time.Duration(i*60)*time.Second)
        }
        for i := 0; i < len(nodeArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i]+"/status.txt", "notok", currentDir, logName, mode)
        }
        ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes node add failed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/addnode.log")
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Failed to add node of "+label+" cluster! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    }
}

// Delete the core operation part of the node.
func DeleteNodeCore(mode string, nodeArray []string, currentDir string, kissh string, subProcessDir string, currentUser string, label string, softDir string, logName string) {
    opt := "delnode"
    logStr := LogStr(mode)
    CreateDir(currentDir+"/data/output"+subProcessDir, currentDir, logName, mode)
    os.OpenFile(currentDir+"/data/logs"+subProcessDir+"/logs/delnode.log", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
    ShellExecute("echo \"*************************************************************************************\n\n[Info] "+time.Now().String()+" Deleting kubernetes node, please wait ... \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/delnode.log")
    ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
    ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" Deleting kubernetes node of "+label+" cluster ... </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    sshPort := GetSshPort(label, currentDir, mode)
    if !DelnodeConfig(mode, nodeArray, currentDir, subProcessDir, logName, sshPort) {
        ShellExecute("echo [Error] "+time.Now().String()+" \"The parameters you entered are incorrect, please check! \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/delnode.log")
        for i := 0; i < len(nodeArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i]+"/status.txt", "unknow", currentDir, logName, mode)
        }
        return
    }
    DelnodeYML("", currentDir+"/data/output"+subProcessDir, currentDir, currentUser, logName, false)
    for i := 0; i < len(nodeArray); i++ {
        // update install status of kubernetes node
        CreateDir(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i], currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i]+"/status.txt", currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i]+"/status.txt", "deleting", currentDir, logName, mode)
        // delete node for the kubernetes cluster
        if !DeleteNode(label, nodeArray[i], currentDir, logName, mode) {
            ShellExecute("echo [Warning] "+time.Now().String()+" \"There is an exception in deleting node("+nodeArray[i]+") from kubernetes cluster, or node("+nodeArray[i]+") information does not exist in kubernetes cluster. \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/delnode.log")
        }
    }
    //ExecuteDeleteNode(nodeArray, currentDir, subProcessDir, opt, mode)
    ShellExecute("echo [Info] "+time.Now().String()+" \"The system is scheduling pod to other healthy nodes in the cluster. Please wait... \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/delnode.log")
    time.Sleep(time.Duration(30)*time.Second)
    err_delnode := ExecuteOpt(kissh, currentDir, opt, opt, subProcessDir, "")
    if err_delnode != nil {
        for i := 0; i < len(nodeArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i]+"/status.txt", "notok", currentDir, logName, mode)
        }
        ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes node delete failed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/delnode.log")
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Failed to delete node of "+label+" cluster! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    } else {
        for i := 0; i < len(nodeArray); i++ {
            os.RemoveAll(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[i])
        }
        ShellExecute("echo [Info] "+time.Now().String()+" \"Kubernetes node delete operation completed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/delnode.log")
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" The node of "+label+" cluster has been deleted successfully! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    }
}

// Install the core operation part of the cluster.
func InstallCore(mode string, master string, masterArray []string, node string, nodeArray []string, softDir string, currentDir string, kissh string, subProcessDir string, currentUser string, label string, osTypeResult string, osType string, k8sVer string, logName string, Version string, CompatibleK8S string, CompatibleOS string, installTime string, way string, upgradeKernel string, kubeDashboard string, kubeApiPort string, cniPlugin string, sshPort string) {
    opt := "install"
    layoutName := "install"
    logStr := LogStr(mode)
    upgradeKernelStr := "\n    Automatically Upgrade OS Kernel: Not Support"
    if osTypeResult == "centos7" || osTypeResult == "rhel7" {
        if upgradeKernel == "yes"{
            upgradeKernelStr = "\n    Automatically Upgrade OS Kernel: YES"
        } else {
            upgradeKernelStr = "\n    Automatically Upgrade OS Kernel: NO"
        }
    } 
    CheckOS(CompatibleOS, osTypeResult, currentDir, logName, mode)
    CheckK8sVersion(Version, CompatibleK8S, k8sVer, currentDir, logName, mode)
    CheckCNI(cniPlugin, currentDir, logName, mode)
    CreateDir(currentDir+"/data/output"+subProcessDir, currentDir, logName, mode)
    CreateDir(currentDir+"/data/logs"+subProcessDir, currentDir, logName, mode)
    if way == "newinstall" {
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/softdir.txt", softDir, currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/ostype.txt", osType, currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/k8sver.txt", k8sVer, currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/k8sapiport.txt", kubeApiPort, currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/cniplugin.txt", cniPlugin, currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "unknow", currentDir, logName, mode)
    }
    os.OpenFile(currentDir+"/data/logs"+subProcessDir+"/logs/install.log", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
    GeneralConfig(mode, masterArray, nodeArray, currentDir, softDir, subProcessDir, osTypeResult, k8sVer, kubeDashboard, kubeApiPort, logName, sshPort)
    if !InstallConfig(mode, masterArray, nodeArray, currentDir, subProcessDir, logName, sshPort) {
        ShellExecute("echo [Error] "+time.Now().String()+" \"The parameters you entered are incorrect, please check! \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
        if way == "newinstall" { os.RemoveAll(currentDir+"/data/output"+subProcessDir) }
        return
    }
    ShellExecute("cp -rf "+currentDir+"/sys "+currentDir+"/data/output"+subProcessDir+"/")
    InstallGenfile(osTypeResult, mode, currentDir, subProcessDir, logName)
    InstallIpvsYaml(mode, currentDir, masterArray, kubeApiPort, subProcessDir, logName)
    var err_install error
    if len(masterArray) == 1{
        OnemasterInstallYML(mode,currentDir+"/data/output"+subProcessDir,currentDir,currentUser,logName,upgradeKernel,osTypeResult,k8sVer,cniPlugin)
        layoutName = "onemasterinstall"
    }else{
        InstallYML(mode,currentDir+"/data/output"+subProcessDir, currentDir, currentUser, logName,upgradeKernel,osTypeResult,k8sVer,cniPlugin)
    }
    CreateFile(currentDir+"/data/output"+subProcessDir+"/sshport.txt", currentDir, logName, mode)
    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/sshport.txt", sshPort, currentDir, logName, mode)
    for i := 0; i < len(masterArray); i++ {
        CreateDir(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i], currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/ttystatus.txt", currentDir, logName, mode)
    }
    for j := 0; j < len(nodeArray); j++ {
        CreateDir(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j], currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j]+"/status.txt", currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j]+"/ttystatus.txt", currentDir, logName, mode)
    }
    if installTime != "" {
        ShellExecute("echo \"*************************************************************************************\n\n[Info] "+time.Now().String()+" Start scheduled installation task, please wait ...  \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
        ShellExecute("echo \"    Kubernetes Cluster Label: "+label+"\n    Kubernetes Version: Kubernetes v"+k8sVer+"\n    Kubernetes Master: "+master+"\n    Kubernetes Node: "+node+"\n    SSH Operation Port: "+sshPort+"\n    CNI Plug-in Type: "+cniPlugin+"\n    Operating System Type: "+osType+ upgradeKernelStr +"\n    System User for Installation: "+currentUser+"\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/softdirtemp.txt", softDir, currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/ostypetemp.txt", osType, currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/k8svertemp.txt", k8sVer, currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/cniplugintemp.txt", cniPlugin, currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/installtime.txt", installTime, currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/scheduler.txt", "on", currentDir, logName, mode)
        return
    } else {
        ShellExecute("echo \"*************************************************************************************\n\n[Info] "+time.Now().String()+" Installing kubernetes cluster, please wait ... \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
        ShellExecute("echo \"    Kubernetes Cluster Label: "+label+"\n    Kubernetes Version: Kubernetes v"+k8sVer+"\n    Kubernetes Master: "+master+"\n    Kubernetes Node: "+node+"\n    SSH Operation Port: "+sshPort+"\n    CNI Plug-in Type: "+cniPlugin+"\n    Operating System Type: "+osType+ upgradeKernelStr +"\n    System User for Installation: "+currentUser+"\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
        sch,_ := ReadFile(currentDir+"/data/output"+subProcessDir+"/scheduler.txt")
        if sch == "on" {
            ShellExecute("echo [Error] "+time.Now().String()+" \"Installation conflict! Background scheduled tasks exist and installation is in progress.\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
            if way == "newinstall" { os.RemoveAll(currentDir+"/data/output"+subProcessDir) }
        } else {
            if way == "reinstall" {
                _,_,_,subProcessDir,_ := ParameterConvert(mode, "", "", "", label, "")
                err_cert := os.RemoveAll(currentDir+"/data/output"+subProcessDir+"/cert")
                if err_cert != nil {
                    ShellExecute("echo [Error] "+time.Now().String()+" \"Failed to install! There are residual files that cannot be cleaned up. Please delete the directory ("+currentDir+"/data/output"+subProcessDir+"cert) manually and try to install again!\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
                    return
                }
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/softdir.txt", softDir, currentDir, logName, mode)
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/ostype.txt", osType, currentDir, logName, mode)
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/k8sver.txt", k8sVer, currentDir, logName, mode)
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/k8sapiport.txt", kubeApiPort, currentDir, logName, mode)
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/cniplugin.txt", cniPlugin, currentDir, logName, mode)
            }
            for i := 0; i < len(masterArray); i++ {
                CreateDir(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i], currentDir, logName, mode)
                CreateFile(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", currentDir, logName, mode)
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "adding", currentDir, logName, mode)
            }
            for j := 0; j < len(nodeArray); j++ {
                CreateDir(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j], currentDir, logName, mode)
                CreateFile(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j]+"/status.txt", currentDir, logName, mode)
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j]+"/status.txt", "adding", currentDir, logName, mode)
            }
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "installing", currentDir, logName, mode)
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/progressbar.txt", "1", currentDir, logName, mode)
            ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
            ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" Installing kubernetes cluster of "+label+" cluster ...</div></div>\" >> "+currentDir+"/data/msg/msg.txt")
            err_install = ExecuteOpt(kissh, currentDir, opt, layoutName, subProcessDir, mode)
        }
    }
    if err_install != nil {
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "notok", currentDir, logName, mode)
        for i := 0; i < len(masterArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "unknow", currentDir, logName, mode)
        }
        ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes install failed! There is an error in the process! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Kubernetes cluster ("+label+") install failed! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    } else {
        ShellExecute("echo [Info] "+time.Now().String()+" \"Cleaning and detection after installation are in progress. Please wait ... \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
        for i := 1; i <= 3; i++ {
            if len(ListNode(label,currentDir,logName,mode)) >= len(nodeArray) {
                if mode == "DAEMON" {
                    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "restarting", currentDir, logName, mode)
                } else {
                    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "ok", currentDir, logName, mode)
                    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/progressbar.txt", "", currentDir, logName, mode)
                }
                for i := 0; i < len(masterArray); i++ {
                    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "ok", currentDir, logName, mode)
                }
                for j := 0; j < len(nodeArray); j++ {
                    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j]+"/status.txt", "ok", currentDir, logName, mode)
                }
                ShellExecute("echo [Info] "+time.Now().String()+" \"Kubernetes cluster install completed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
                ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
                ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" Kubernetes cluster ("+label+") has been installed successfully! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
                return
            }
            time.Sleep(time.Duration(i*90)*time.Second)
        }
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "notok", currentDir, logName, mode)
        for i := 0; i < len(masterArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "unknow", currentDir, logName, mode)
        }
        ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes cluster install failed! "+label+" cluster status is unhealthy! Please check the network or other factors between the kubernetes master and the node. \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Kubernetes cluster ("+label+") install failed! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    }
}

// Uninstall the core operation part of the cluster.
func UninstallCore(mode string, master string, masterArray []string, node string, nodeArray []string, softDir string, currentDir string, kissh string, subProcessDir string, currentUser string, label string, osTypeResult string, logName string, CompatibleOS string, sshPort string) {
    opt := "uninstall"
    logStr := LogStr(mode)
    //CheckOS(CompatibleOS, osTypeResult, currentDir, logName, mode)
    CreateDir(currentDir+"/data/output"+subProcessDir, currentDir, logName, mode)
    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "uninstalling", currentDir, logName, mode)
    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/progressbar.txt", "1", currentDir, logName, mode)
    os.OpenFile(currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
    ShellExecute("echo \"*************************************************************************************\n\n[Info] "+time.Now().String()+" Uninstalling kubernetes cluster, please wait ... \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
    ShellExecute("echo \"    Kubernetes Cluster Label: "+label+"\n    Kubernetes Master: "+master+"\n    Kubernetes Node: "+node+"\n    SSH Operation Port: "+sshPort+"\n    System User for Uninstallation: "+currentUser+"\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
    ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
    ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" Uninstalling kubernetes cluster of "+label+" cluster ... </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    GeneralConfig(mode, masterArray, nodeArray, currentDir, softDir, subProcessDir, osTypeResult, "", "", "", logName, sshPort)
    if !InstallConfig(mode,masterArray, nodeArray, currentDir, subProcessDir, logName, sshPort) {
        ShellExecute("echo [Error] "+time.Now().String()+"The parameters you entered are incorrect, please check! \" \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
        return
    }
    if !DelmasterConfig(mode,masterArray, currentDir, subProcessDir, logName, sshPort) {
        ShellExecute("echo [Error] "+time.Now().String()+" \"The parameters you entered are incorrect, please check! \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
        return
    }
    if !DelnodeConfig(mode,nodeArray, currentDir, subProcessDir, logName, sshPort) {
        ShellExecute("echo [Error] "+time.Now().String()+" \"The parameters you entered are incorrect, please check! \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
        return
    }
    DelmasterYML("",currentDir+"/data/output"+subProcessDir, currentDir, currentUser, logName, true)
    DelnodeYML("",currentDir+"/data/output"+subProcessDir, currentDir, currentUser, logName, true)
    for i := 0; i < len(masterArray); i++ {
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "deleting", currentDir, logName, mode)
    }
    for j := 0; j < len(nodeArray); j++ {
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j]+"/status.txt", "deleting", currentDir, logName, mode)
    }
    ShellExecute("echo [Info] "+time.Now().String()+" \"Loading operation configuration ... \""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
    err_delnode := ExecuteOpt(kissh, currentDir, opt, "delnode", subProcessDir, "")
    if err_delnode != nil {
        ShellExecute("echo [Error] "+time.Now().String()+" \"Failed to delete node of "+label+" cluster! \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
        ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes cluster ("+label+") uninstall failed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "notok", currentDir, logName, mode)
        for i := 0; i < len(nodeArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+masterArray[i]+"/status.txt", "unknow", currentDir, logName, mode)
        }
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Kubernetes cluster ("+label+") uninstall failed! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    } else {
        ShellExecute("echo [Info] "+time.Now().String()+" \"All kubernetes node of "+label+" cluster has been uninstalled successfully! \n\nPlease wait 15 seconds before uninstalling all kubernetes masters ... \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
    }
    time.Sleep(time.Duration(15)*time.Second)
    err_delmaster := ExecuteOpt(kissh, currentDir, opt, "delmaster", subProcessDir, "")
    if err_delmaster != nil {
        ShellExecute("echo [Error] "+time.Now().String()+" \"Failed to delete master of "+label+" cluster! \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
        ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes cluster ("+label+") uninstall failed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "notok", currentDir, logName, mode)
        for i := 0; i < len(masterArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "unknow", currentDir, logName, mode)
        }
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Kubernetes uninstall failed! There is an error in the process! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    } else {
        ShellExecute("echo [Info] "+time.Now().String()+" \"Cleaning and detection after uninstallation are in progress. Please wait ... \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
        for i := 1; i <= 3; i++ {
            err_health := DetectK8sHealth(label, currentDir, logName, mode)
            if err_health != nil {
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "notinstall", currentDir, logName, mode)
                ShellExecute("echo [Info] "+time.Now().String()+" \"Cleaning up temporary cache files ... \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
                err_rm := os.RemoveAll("/tmp/.kubeinstalltemp"+subProcessDir)
                err_gc := os.RemoveAll(currentDir+"/data/output"+subProcessDir)
                if err_rm != nil || err_gc != nil {
                   ShellExecute("echo [Warning] "+time.Now().String()+" \"There are residual files that cannot be cleaned up. Please delete the following directories manually: \n    "+currentDir+"/data/output"+subProcessDir+"\n    /tmp/.kubeinstalltemp\n    ……\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log") 
                }
                ShellExecute("echo [Info] "+time.Now().String()+" \"Kubernetes cluster uninstall completed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
                ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
                ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" Kubernetes cluster ("+label+") has been uninstalled successfully! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
                return
            }
            time.Sleep(time.Duration(i*25)*time.Second)
        }
        ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes cluster ("+label+") uninstall failed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "notok", currentDir, logName, mode)
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Kubernetes cluster ("+label+") uninstall failed! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    }
}

// The core operation part of schedule installation.
func InstallScheduler(label string, masterArray []string, nodeArray []string, kissh string, currentDir string, opt string, layoutName string, subProcessDir string, sshPort string, logName string, mode string) {
    logStr := LogStr(mode)
    _,err_sdr := CopyFile(currentDir+"/data/output"+subProcessDir+"/softdirtemp.txt", currentDir+"/data/output"+subProcessDir+"/softdir.txt")
    CheckErr(err_sdr,currentDir,logName,mode)
    _,err_ost := CopyFile(currentDir+"/data/output"+subProcessDir+"/ostypetemp.txt", currentDir+"/data/output"+subProcessDir+"/ostype.txt")
    CheckErr(err_ost,currentDir,logName,mode)
    _,err_k8v := CopyFile(currentDir+"/data/output"+subProcessDir+"/k8svertemp.txt", currentDir+"/data/output"+subProcessDir+"/k8sver.txt")
    CheckErr(err_k8v,currentDir,logName,mode)
    _,err_cni := CopyFile(currentDir+"/data/output"+subProcessDir+"/cniplugintemp.txt", currentDir+"/data/output"+subProcessDir+"/cniplugin.txt")
    CheckErr(err_cni,currentDir,logName,mode)
    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "installing", currentDir, logName, mode)
    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/progressbar.txt", "1", currentDir, logName, mode)
    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/installtime.txt", "", currentDir, logName, mode)
    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/scheduler.txt", "off", currentDir, logName, mode)
    ShellExecute("echo \"\n*************************************************************************************\n\n[Info] "+time.Now().String()+" Now start the installation process: \n\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
    ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
    ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" Installing kubernetes cluster of "+label+" cluster ...</div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    CreateFile(currentDir+"/data/output"+subProcessDir+"/sshport.txt", currentDir, logName, mode)
    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/sshport.txt", sshPort, currentDir, logName, mode)
    for i := 0; i < len(masterArray); i++ {
        CreateDir(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i], currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/ttystatus.txt", currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "adding", currentDir, logName, mode)
    }
    for j := 0; j < len(nodeArray); j++ {
        CreateDir(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j], currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j]+"/status.txt", currentDir, logName, mode)
        CreateFile(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j]+"/ttystatus.txt", currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j]+"/status.txt", "adding", currentDir, logName, mode)
    }
    err_install := ExecuteOpt(kissh, currentDir, opt, layoutName, subProcessDir, mode)
    if err_install != nil {
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "notok", currentDir, logName, mode)
        for i := 0; i < len(masterArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "unknow", currentDir, logName, mode)
        }
        ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes install failed! There is an error in the process! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Kubernetes cluster ("+label+") install failed! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    } else {
        ShellExecute("echo [Info] "+time.Now().String()+" \"Cleaning and detection after installation are in progress. Please wait ... \n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/uninstall.log")
        for i := 1; i <= 3; i++ {
            if len(ListNode(label,currentDir,logName,mode)) >= len(nodeArray) {
                if mode == "DAEMON" {
                    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "restarting", currentDir, logName, mode)
                } else {
                    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "ok", currentDir, logName, mode)
                    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/progressbar.txt", "", currentDir, logName, mode)
                }
                for i := 0; i < len(masterArray); i++ {
                    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "ok", currentDir, logName, mode)
                }
                for j := 0; j < len(nodeArray); j++ {
                    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+nodeArray[j]+"/status.txt", "ok", currentDir, logName, mode)
                }
                ShellExecute("echo [Info] "+time.Now().String()+" \"Kubernetes cluster install completed! \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
                ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
                ShellExecute("echo \"<div class='g_12'><div class='info iDialog'>[Info] "+time.Now().String()+" Kubernetes cluster ("+label+") has been installed successfully! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
                return
            }
            time.Sleep(time.Duration(i*90)*time.Second)
        }
        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "notok", currentDir, logName, mode)
        for i := 0; i < len(masterArray); i++ {
            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+masterArray[i]+"/status.txt", "unknow", currentDir, logName, mode)
        }
        ShellExecute("echo [Error] "+time.Now().String()+" \"Kubernetes cluster install failed! "+label+" cluster status is unhealthy! Please check the network or other factors between the kubernetes master and the node. \n\n*************************************************************************************\n\""+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/install.log")
        ShellExecute("sed -i '1d' "+currentDir+"/data/msg/msg.txt")
        ShellExecute("echo \"<div class='g_12'><div class='error iDialog'>[Error] "+time.Now().String()+" Kubernetes cluster ("+label+") install failed! </div></div>\" >> "+currentDir+"/data/msg/msg.txt")
    }
}

// Perform the operation of deleting node.
func ExecuteDeleteNode(nodeArray []string, currentDir string, subProcessDir string, opt string, mode string){
    var delNodeList string
    logStr := LogStr(mode)
    nodeArrayLen := len(nodeArray)
    if nodeArrayLen == 1 {
        delNodeList = nodeArray[0]
    } else {
        delNodeList = "{"
        for i := 0; i < nodeArrayLen; i++ {
            delNodeList = delNodeList + nodeArray[i]
            if i == nodeArrayLen-1 {
                delNodeList = delNodeList + "}"
            } else {
                delNodeList = delNodeList + ","
            }
        }
    }
    ShellExecute(currentDir+"/pkg/proc/.bin/kubectl --kubeconfig "+currentDir+"/data/output"+subProcessDir+"/cert/ssl/kube-install.kubeconfig delete node "+delNodeList+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/"+opt+".log")
}

// Execute the operation commands of the core part.
func ExecuteOpt(kiCommand string, currentDir string, opt string, layoutName string, subProcessDir string, mode string) error {
    logStr := LogStr(mode)
    inventoryName := opt
    if opt == "uninstall" {
        inventoryName = layoutName
    }
    err := ShellExecute(kiCommand+" -i "+currentDir+"/data/output"+subProcessDir+"/"+inventoryName+".inventory "+currentDir+"/data/output"+subProcessDir+"/k8scluster-"+layoutName+".yml"+logStr+currentDir+"/data/logs"+subProcessDir+"/logs/"+opt+".log" )
    return err
}

// Perform remote TTY disable or enable operation
func SwitchTTY(label string, currentDir string, ttyIP string, ttyOption string, sshPort string, logName string, mode string) error {
    var err error
    _,_,_,subProcessDir,_ := ParameterConvert(mode, "", "", "", label, "")
    mstArrayLocal, err := GetAllDir(currentDir+"/data/output"+subProcessDir+"/masters", currentDir, logName, mode)
    ndArrayLocal, err := GetAllDir(currentDir+"/data/output"+subProcessDir+"/nodes", currentDir, logName, mode)
    if ttyOption == "enable" {
        err = ShellExecute(currentDir+"/pkg/proc/kissh/bin/ansible all -i "+ttyIP+":"+sshPort+", -m command -a 'systemctl restart kube-install-ttyd.service'")
        if err == nil {
            if StrInArray(ttyIP, mstArrayLocal) {
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+ttyIP+"/ttystatus.txt", "enable", currentDir, logName, mode)
            }
            if StrInArray(ttyIP, ndArrayLocal) {
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+ttyIP+"/ttystatus.txt", "enable", currentDir, logName, mode)
            }
            ShellExecute(currentDir+"/pkg/proc/kissh/bin/ansible all -i "+ttyIP+":"+sshPort+", -m command -a 'systemctl enable kube-install-ttyd.service'")
        }
    } else {
        err = ShellExecute(currentDir+"/pkg/proc/kissh/bin/ansible all -i "+ttyIP+":"+sshPort+", -m command -a 'systemctl stop kube-install-ttyd.service'")
        if err == nil {
            if StrInArray(ttyIP, mstArrayLocal) {
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/masters/"+ttyIP+"/ttystatus.txt", "disable", currentDir, logName, mode)
            }
            if StrInArray(ttyIP, ndArrayLocal) {
                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/nodes/"+ttyIP+"/ttystatus.txt", "disable", currentDir, logName, mode)
            }
            ShellExecute(currentDir+"/pkg/proc/kissh/bin/ansible all -i "+ttyIP+":"+sshPort+", -m command -a 'systemctl disable kube-install-ttyd.service'")
        }
    }
    return err
}


