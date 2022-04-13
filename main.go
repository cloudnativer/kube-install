package main

import (
    "fmt"
    "time"
    "os"
    "path/filepath"
    "flag"
    "strings"
    "kube-install/lib"
)



func main() {

    var exec,master,node,k8sver,ostype,softdir,label,sship,sshport,sshpass,listen,upgradekernel,k8sdashboard,k8sapiport,cniplugin string

    initFlag := flag.Bool("init",false,"Initialize the local system environment.")
    iFlag := flag.Bool("i",false,"Initialize the local system environment.")
    daemonFlag := flag.Bool("daemon",false,"Run as a daemon service. (enable this switch to use the web console for management)")
    dFlag := flag.Bool("d",false,"Run as a daemon service. (enable this switch to use the web console for management)")
    logsFlag := flag.Bool("logs",false,"View Installation, uninstall, master and node management log information.")
    lFlag := flag.Bool("l",false,"View Installation, uninstall, master and node management log information.")
    showk8sFlag := flag.Bool("showk8s",false,"Display all deployed kubernetes cluster information.")
    sFlag := flag.Bool("s",false,"Display all deployed kubernetes cluster information.")
    versionFlag := flag.Bool("version",false,"Display software version information of kube-install.")
    vFlag := flag.Bool("v",false,"Display software version information of kube-install.")
    helpFlag := flag.Bool("help",false,"Display usage help information of kube-install.")
    hFlag := flag.Bool("h",false,"Display usage help information of kube-install.")
    flag.StringVar(&exec,"exec","","Deploy and uninstall kubernetes cluster.(use with \"init | sshcontrol | install | addnode | delnode | delmaster | rebuildmaster | uninstall\")")
    flag.StringVar(&exec,"e","","Deploy and uninstall kubernetes cluster.(use with \"init | sshcontrol | install | addnode | delnode | delmaster | rebuildmaster | uninstall\")")
    flag.StringVar(&listen,"listen","","Set the IP and port on which the daemon service listens. (default is \"0.0.0.0:9080\")")
    flag.StringVar(&master,"master","","The IP address of k8s master server filled in for the first installation")
    flag.StringVar(&node,"node","","The IP address of k8s node server filled in for the first installation")
    flag.StringVar(&ostype,"ostype","","Specifies the distribution operating system type: \"centos7 | centos8 | rhel7 | rhel8 | ubuntu20 | suse15\".")
    flag.StringVar(&k8sver,"k8sver","","Specifies the version of k8s software installed.(default is \"kubernetes 1.23\")")
    flag.StringVar(&upgradekernel,"upgradekernel","no","Because the lower versions of CentOS 7 and redhat 7 may lack kernel modules, only the kernel automatic upgrade of CentOS 7 and rhel7 operating systems is supported here, and other operating systems do not need to be upgraded.")
    flag.StringVar(&k8sdashboard,"k8sdashboard","yes","Automatically deploy kube-dashboard to kubernetes cluster.")
    flag.StringVar(&k8sapiport,"k8sapiport","6443","The TCP Port of the k8s kube-apiserver.")
    flag.StringVar(&cniplugin,"cniplugin","flannel","Specifies the CNI plug-in type: \"flannel | calico | kuberouter | weave | cilium\".")
    flag.StringVar(&softdir,"softdir","","Specifies the installation directory of kubernetes cluster.(default is \"/opt/kube-install\")")
    flag.StringVar(&label,"label",".default","In the case of deploying and operating multiple kubernetes clusters, it is necessary to specify a label to uniquely identify a kubernetes cluster. (length must be less than 32 strings)")
    flag.StringVar(&sship,"sship","","The IP address of the target host through which the SSH channel is opened.(use with \"sshcontrol\")")
    flag.StringVar(&sshport,"sshport","22","The TCP Port of the target host through which the SSH channel is opened.")
    flag.StringVar(&sshpass,"sshpass","","The SSH password of the target host through which the SSH channel is opened.(use with \"sshcontrol\")")
    flag.Parse()

    // Get the current execution path and user of the program, and set log file name.
    logName := time.Now().Format("date_20060102_150405")
    path, err := os.Executable()
    kilib.CheckErr(err, "", logName, "")
    currentDir := filepath.Dir(path)
    currentUser := strings.Replace(kilib.ShellOutput("echo $USER"), "\n", "", -1)
    kissh := currentDir+"/pkg/proc/kissh/bin/ansible-playbook"

    // Set the version number and release date of Kube-Install.
    const (
        Version string = "v0.8.0-beta2"
        ReleaseDate string = "4/13/2022"
        CompatibleK8S string = "1.18, 1.19, 1.20, 1.21, 1.22, 1.23, and 1.24"
        CompatibleOS string = "CentOS linux 7, CentOS linux 8, RHEL 7, RHEL 8, Ubuntu 20, and SUSE 15"
    )

    switch {

        //Initialize local host environment
        case *initFlag , *iFlag :
            kilib.CheckOS(CompatibleOS, ostype, currentDir, logName, "")
            fmt.Println("\nInitialization in progress, please wait...\n")
            kilib.CreateDir(currentDir+"/data/config", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/config/language.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/config/tools.txt", currentDir, logName, "")
            kilib.CreateDir(currentDir+"/data/output", currentDir, logName, "")
            kilib.CreateDir(currentDir+"/data/logs", currentDir, logName, "")
            kilib.CreateDir(currentDir+"/data/.key", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/.key/admin", currentDir, logName, "")
            kilib.CreateDir(currentDir+"/data/msg", currentDir, logName, "")
            kilib.CreateDir(currentDir+"/data/statistics", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/msg/msg.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/statistics/k8snum.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/statistics/cpuinfo.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/statistics/meminfo.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/statistics/diskinfo.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/statistics/stuok.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/statistics/stuinstall.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/statistics/stuuninstall.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/statistics/stunotok.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/statistics/stuunkonw.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/statistics/labellist.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/statistics/nodenumlist.txt", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/statistics/schedulelist.txt", currentDir, logName, "")
            key, _ := kilib.EnPwdCode([]byte("CloudNativeR"))
            kilib.DatabaseUpdate(currentDir+"/data/.key/admin", string(key), currentDir, logName, "")
            kilib.DatabaseUpdate(currentDir+"/data/msg/msg.txt", " \n \n \n \n \n \n \n", currentDir, logName, "")
            _,_,_,_,ostype := kilib.ParameterConvert("", "", "", "", "", ostype)
            err_ops := kilib.SshOpsInit(currentDir, ostype, "")
            if err_ops != nil {
                fmt.Println("[Error] "+time.Now().String()+" Initialization failed, the basic dependency package is missing!\n")
                return
            } else {
                fmt.Println("Notice: If you are prompted to enter the password below, please enter the root password again! ")
                var ipArray = []string{"127.0.0.1"}
                err_host := kilib.SshKey(ipArray, sshport, "", currentDir)
                if err_host != nil {
                    fmt.Println("\n[Error] "+time.Now().String()+" Initialization failed ! There is a problem with the local SSH key. \n\nRecommendations:\n    If the SSH port of the host is not \"22\", use the \"-sshport\" to specify the correct port.\n    (Please try again with root user)\n\nInitialization failed!\n")
                    return
                } else {
                    kilib.CreateSystemdService("", currentDir, logName)
                    fmt.Println("\n\nInitialization completed!\n")
                }
            }

        // Run as a daemon process.
        case *daemonFlag , *dFlag :
            fmt.Println("Notice: If you are prompted to enter the password below, please enter the root password again! \n")
            if listen != "" {
                kilib.DaemonRun(Version,ReleaseDate,CompatibleK8S,CompatibleOS,listen,currentDir,currentUser,kissh,logName,"DAEMON")
            } else {
                kilib.DaemonRun(Version,ReleaseDate,CompatibleK8S,CompatibleOS,"0.0.0.0:9080",currentDir,currentUser,kissh,logName,"DAEMON")
            }

        // View all kubernetes cluster information
        case *showk8sFlag , *sFlag :
            labelArray,err := kilib.GetAllDir(currentDir+"/data/output",currentDir,logName,"")
            kilib.CheckErr(err,currentDir,logName,"")
            fmt.Println("---------------------------------┬-----------┬----------------┬-------------------------------------┬---------------------------------\n    Kubernetes Cluster Label     |  Version  | Install Status |     Install Resource Directory      |  Kubernetes Master Information  \n---------------------------------┼-----------┼----------------┼-------------------------------------┼---------------------------------")
            for _, i := range labelArray {
                var label2,softDir2 string
                labelInit := "                                | "
                k8sverInit := "          | "
                statusInit := "               | "
                softDirInit := "                                    |"
                label := string(i)
                if len(label) < len(labelInit) { label2 = label + labelInit[len(label):len(labelInit)] } else { label2 = label + " | " }
                k8sver := "k8s "+kilib.GetClusterK8sVer(label,currentDir,"")
                status,_ := kilib.GetClusterStatus(label,currentDir,logName,"")
                softDir := kilib.GetClusterSoftdir(label,currentDir,"")
                if len(softDir) < len(softDirInit) { softDir2 = softDir + softDirInit[len(softDir):len(softDirInit)] } else { softDir2 = softDir + " |" }
                master := kilib.GetClusterMaster(label,currentDir,logName,"")
                fmt.Printf(" " + label2 + k8sver + k8sverInit[len(k8sver):len(k8sverInit)] + status + statusInit[len(status):len(statusInit)] + softDir2+" ")
                fmt.Println(master)
            }
            fmt.Println("---------------------------------┴-----------┴----------------┴-------------------------------------┴---------------------------------")

        // View log information
        case *logsFlag , *lFlag :
            kilib.ShowLogs(exec, label, currentDir)

        // View software version details.
        case *versionFlag , *vFlag :
            kilib.ShowVersion(Version, ReleaseDate, CompatibleK8S, CompatibleOS)

        // Help information of Kube-Install
        case *helpFlag , *hFlag :
            kilib.ShowHelp()

        case exec != "" :
            switch {
               //Execute sshcontrol command
               case exec == "sshcontrol" :
                   kilib.CheckParam(exec,"sship",sship)
                   kilib.CheckParam(exec,"sshpass",sshpass)
                   fmt.Println("\nOpening SSH tunnel, please wait...\n")
                   ipArray := strings.Split(sship, ",")
                   err := kilib.SshKey(ipArray, sshport, sshpass, currentDir)
                   if err != nil {
                       fmt.Println("[Error] "+time.Now().String()+" Failed to open the SSH channel. Please use \"root\" user to manually open the SSH channel from the local host to the target host, or try to open the SSH channel again after executing the following command on the target host:\n\n  ----------------------------------------------------------------- \n   sudo sed -i \"/PermitRootLogin/d\" /etc/ssh/sshd_config \n   sudo sh -c \"echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config\" \n   sudo sed -i \"/StrictHostKeyChecking/s/^#//; /StrictHostKeyChecking/s/ask/no/\" /etc/ssh/ssh_config \n   sudo systemctl restart sshd \n  ----------------------------------------------------------------- \n\n   (If the SSH port of the host is not \"22\", use the \"-sshport\" to specify the correct port.)\n\nFailed to open SSH tunnel!\n")
                   } else {
                       fmt.Println("[Info] "+time.Now().String()+" Successfully open the SSH channel from local host to the target host ("+sship+":"+sshport+")！\n\n\nThe SSH tunnel is opened!\n")
                   }

              //Execute install command
               case exec == "install" :
                   kilib.CheckParam(exec,"master",master)
                   kilib.CheckParam(exec,"node",node)
                   if !kilib.CheckLabel(label) {
                       fmt.Println("\nThe \"-ostype\" parameter length must be less than 32 strings, please check! \n\n ")
                       return
                   }
                   masterArray,nodeArray,softdir,subProcessDir,ostypeResult := kilib.ParameterConvert("", master, node, softdir, label, ostype)
                   kilib.DatabaseInit(currentDir,subProcessDir,logName,"")
                   kilib.InstallCore("",master,masterArray,node,nodeArray,softdir,currentDir,kissh,subProcessDir,currentUser,label,ostypeResult,ostype,k8sver,logName,Version,CompatibleK8S,CompatibleOS,"","newinstall",upgradekernel,k8sdashboard,k8sapiport,cniplugin,sshport)

               //Execute addnode command
               case exec == "addnode" :
                   kilib.CheckParam(exec,"node",node)
                   _,nodeArray,_,subProcessDir,ostypeResult := kilib.ParameterConvert("", "", node, softdir, label, ostype)
                   kilib.AddNodeCore("",node,nodeArray,currentDir,kissh,subProcessDir,currentUser,label,softdir,ostypeResult,k8sver,logName,CompatibleOS,upgradekernel)

               //Execute delnode command
               case exec == "delnode" :
                   kilib.CheckParam(exec,"node",node)
                   _,nodeArray,_,subProcessDir,_ := kilib.ParameterConvert("", "", node, softdir, label, ostype)
                   kilib.DeleteNodeCore("",nodeArray,currentDir,kissh,subProcessDir,currentUser,label,softdir,logName)

               //Execute rebuildmaster command
               case exec == "rebuildmaster" :
                   kilib.CheckParam(exec,"master",master)
                   masterArray,_,_,subProcessDir,_ := kilib.ParameterConvert("", master, "", softdir, label, "")
                   kilib.RebuildMasterCore("",masterArray,currentDir,kissh,subProcessDir,currentUser,label,softdir,logName)

               //Execute delmaster command
               case exec == "delmaster" :
                   kilib.CheckParam(exec,"master",master)
                   masterArray,_,_,subProcessDir,_ := kilib.ParameterConvert("", master, "", softdir, label, ostype)
                   kilib.DeleteMasterCore("",masterArray,currentDir,kissh,subProcessDir,currentUser,label,softdir,logName)

               //Execute uninstall command
               case exec == "uninstall" :
                   kilib.CheckParam(exec,"master",master)
                   kilib.CheckParam(exec,"node",node)
                   masterArray,nodeArray,softdir,subProcessDir,ostype := kilib.ParameterConvert("", master, node, softdir, label, ostype)
                   kilib.UninstallCore("",master,masterArray,node,nodeArray,softdir,currentDir,kissh,subProcessDir,currentUser,label,ostype,logName,CompatibleOS,sshport) 
           }

        default:
           fmt.Println("Notice: the command parameters you entered are incorrect. Please refer to the following help document and re-enter after checking! \n")
           kilib.ShowHelp()

    }

}

