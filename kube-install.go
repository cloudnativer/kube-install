package main

import (
    "fmt"
    "time"
    "os"
    "path/filepath"
    "flag"
    "strings"
//    "strconv"
    "kube-install/lib"
)



func main() {

    var exec string
    var master string
    var node string
    var k8sver string
    var ostype string
    var softdir string
    var label string
    var sship string
    var sshpass string
    var listen string

    initFlag := flag.Bool("init",false,"Initialize the local system environment.")
    iFlag := flag.Bool("i",false,"Initialize the local system environment.")
    daemonFlag := flag.Bool("daemon",false,"Run as a daemon service. (Enable this switch to use the web console for management)")
    dFlag := flag.Bool("d",false,"Run as a daemon service. (Enable this switch to use the web console for management)")
    showk8sFlag := flag.Bool("showk8s",false,"Display all deployed kubernetes cluster information.")
    sFlag := flag.Bool("s",false,"Display all deployed kubernetes cluster information.")
    versionFlag := flag.Bool("version",false,"Display software version information of kube-install.")
    vFlag := flag.Bool("v",false,"Display software version information of kube-install.")
    helpFlag := flag.Bool("help",false,"Display usage help information of kube-install.")
    hFlag := flag.Bool("h",false,"Display usage help information of kube-install.")
    flag.StringVar(&exec,"exec","","Deploy and uninstall kubernetes cluster.(Use with \"init | sshcontrol | install | addnode | delnode | delmaster | rebuildmaster | uninstall\")")
    flag.StringVar(&exec,"e","","Deploy and uninstall kubernetes cluster.(Use with \"init | sshcontrol | install | addnode | delnode | delmaster | rebuildmaster | uninstall\")")
    flag.StringVar(&listen,"listen","","Set the IP and port on which the daemon service listens. (Default is \"0.0.0.0:9080\")")
    flag.StringVar(&master,"master","","The IP address of k8s master server filled in for the first installation")
    flag.StringVar(&node,"node","","The IP address of k8s node server filled in for the first installation")
    flag.StringVar(&ostype,"ostype","","Specifies the distribution operating system type: centos7 | centos8 | rhel7 | rhel8 | ubuntu20 | suse15")
    flag.StringVar(&k8sver,"k8sver","","Specifies the version of k8s software installed.(Default is \"kubernetes 1.22\")")
    flag.StringVar(&softdir,"softdir","","Specifies the installation directory of kubernetes cluster.(Default is \"/exec/kube-install\")")
    flag.StringVar(&label,"label",".default","In the case of deploying and operating multiple kubernetes clusters, it is necessary to specify a label to uniquely identify a kubernetes cluster.")
    flag.StringVar(&sship,"sship","","The IP address of the target host through which the SSH channel is opened.(Use with \"sshcontrol\")")
    flag.StringVar(&sshpass,"sshpass","","The SSH password of the target host through which the SSH channel is opened.(Use with \"sshcontrol\")")
    flag.Parse()

    // Get the current execution path and user of the program, and set log file name.
    logName := time.Now().Format("date_20060102_150405")
    path, err := os.Executable()
    kilib.CheckErr(err, "", logName, "")
    currentDir := filepath.Dir(path)
    currentUser := strings.Replace(kilib.ShellOutput("echo $USER"), "\n", "", -1)
    kissh := currentDir+"/proc/kissh/bin/ansible-playbook"

    // Set the version number and release date of Kube-Install.
    const (
        Version string = "v0.7.0-beta2"
        ReleaseDate string = "8/19/2021"
        CompatibleK8S string = "1.18, 1.19, 1.20, 1.21, and 1.22"
        CompatibleOS string = "CentOS linux 7, CentOS linux 8, RHEL 7, RHEL 8, Ubuntu 20, and SUSE 15"
    )

    switch {

        //Initialize local host environment
        case *initFlag , *iFlag :
            kilib.CheckOS(CompatibleOS, ostype, currentDir, logName, "")
            fmt.Println("\nInitialization in progress, please wait...\n")
            kilib.CreateDir(currentDir+"/data/config", currentDir, logName, "")
            kilib.CreateFile(currentDir+"/data/config/language.txt", currentDir, logName, "")
            kilib.CreateDir(currentDir+"/data/output", currentDir, logName, "")
            kilib.CreateDir(currentDir+"/data/logs", currentDir, logName, "")
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
            kilib.DatabaseUpdate(currentDir+"/data/msg/msg.txt", " \n \n \n \n \n \n \n", currentDir, logName, "")
            _,_,_,_,ostype := kilib.ParameterConvert("", "", "", "", "", ostype)
            err_ops := kilib.SshOpsInit(currentDir, ostype, "")
            if err_ops != nil {
                panic("[Error] "+time.Now().String()+" Initialization failed, the basic dependency package is missing!\n")
            } else {
                fmt.Println("Notice: If you are prompted to enter the password below, please enter the root password again! ")
                var ipArray = []string{"127.0.0.1"}
                err_host := kilib.SshKey(ipArray, "", currentDir)
                if err_host != nil {
                    panic("[Error] "+time.Now().String()+" Initialization failed ! There is a problem with the local SSH key. \n        (Please try again with root user)\n")
                } else {
                    fmt.Println("\n\nInitialization completed!\n")
                }
            }

        // Run as a daemon process.
        case *daemonFlag , *dFlag :
            fmt.Println("Notice: If you are prompted to enter the password below, please enter the root password again! \n")
            var ipArray = []string{"127.0.0.1"}
            err_host := kilib.SshKey(ipArray, "", currentDir)
            if err_host != nil {
                panic("[Error] "+time.Now().String()+" Kube-Install start failed ! \nThere is a problem with the local SSH key. \n")
            }
            if listen != "" {
                kilib.DaemonRun(Version,ReleaseDate,CompatibleK8S,CompatibleOS,listen,currentDir,currentUser,kissh,logName,"DAEMON")
            } else {
                kilib.DaemonRun(Version,ReleaseDate,CompatibleK8S,CompatibleOS,"0.0.0.0:9080",currentDir,currentUser,kissh,logName,"DAEMON")
            }

        // View all kubernetes cluster information
        case *showk8sFlag , *sFlag :
            labelArray,err := kilib.GetAllDir(currentDir+"/data/output",currentDir,logName,"")
            kilib.CheckErr(err,currentDir,logName,"")
            fmt.Println("---------------------------┬-----------┬---------------------┬-----------------------------------┬-------------------------------\n    Label Of K8S Cluster   |  Version  | Installation Status |      Resource File Directory      |    K8S-Master Information     \n---------------------------┼-----------┼---------------------┼-----------------------------------┼-------------------------------")
            for _, i := range labelArray {
                label1 := "                          | "
                k8sver1 := "          | "
                status1 := "                    | "
                softDir1 := "                                  |"
                label2 := string(i)
                k8sver2 := "k8s "+kilib.GetClusterK8sVer(label2,currentDir,"")
                status2,_ := kilib.GetClusterStatus(label2,currentDir,logName,"")
                softDir2 := kilib.GetClusterSoftdir(label2,currentDir,"")
                if len(softDir2) < len(softDir1) { softDir2 = softDir2 + softDir1[len(softDir2):len(softDir1)] }
                master2 := kilib.GetClusterMaster(label2,currentDir,logName,"")
                fmt.Printf(" " + label2 + label1[len(label2):len(label1)] + k8sver2 + k8sver1[len(k8sver2):len(k8sver1)] + status2 + status1[len(status2):len(status1)] + softDir2+" ")
                fmt.Println(master2)
            }
            fmt.Println("---------------------------┴-----------┴---------------------┴-----------------------------------┴-------------------------------")

        // View software version details.
        case *versionFlag , *vFlag :
            kilib.ShowVersion(Version,ReleaseDate,CompatibleK8S,CompatibleOS)

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
                   err := kilib.SshKey(ipArray, sshpass, currentDir)
                   if err != nil {
                       fmt.Println("[Error] "+time.Now().String()+" Failed to open the SSH channel. Please use \"root\" user to manually open the SSH channel from the local host to the target host, or try to open the SSH channel again after executing the following command on the target host:\n  ----------------------------------------------------------------- \n   sudo sed -i \"/PermitRootLogin/d\" /etc/ssh/sshd_config \n   sudo sh -c \"echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config\" \n   sudo sed -i \"/StrictHostKeyChecking/s/^#//; /StrictHostKeyChecking/s/ask/no/\" /etc/ssh/ssh_config \n   sudo systemctl restart sshd \n  ----------------------------------------------------------------- \n\n\nFailed to open SSH tunnel!\n")
                   } else {
                       fmt.Println("[Info] "+time.Now().String()+" Successfully open the SSH channel from local host to the target host ("+sship+")！\n\n\nThe SSH tunnel is opened!\n")
                   }

              //Execute install command
               case exec == "install" :
                   kilib.CheckParam(exec,"master",master)
                   kilib.CheckParam(exec,"node",node)
                   masterArray,nodeArray,softdir,subProcessDir,ostypeResult := kilib.ParameterConvert("", master, node, softdir, label, ostype)
                   kilib.DatabaseInit(currentDir,subProcessDir,logName,"")
                   kilib.InstallCore("",master,masterArray,node,nodeArray,softdir,currentDir,kissh,subProcessDir,currentUser,label,ostypeResult,ostype,k8sver,logName,Version,CompatibleK8S,CompatibleOS,"","newinstall")

               //Execute addnode command
               case exec == "addnode" :
                   kilib.CheckParam(exec,"node",node)
                   _,nodeArray,_,subProcessDir,ostypeResult := kilib.ParameterConvert("", "", node, softdir, label, ostype)
                   kilib.AddNodeCore("",node,nodeArray,currentDir,kissh,subProcessDir,currentUser,label,softdir,ostypeResult,logName,CompatibleOS)

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
                   kilib.UninstallCore("",master,masterArray,node,nodeArray,softdir,currentDir,kissh,subProcessDir,currentUser,label,ostype,logName,CompatibleOS) 
           }

        default:
           fmt.Println("Notice: the command parameters you entered are incorrect. Please refer to the following help document and re-enter after checking! \n")
           kilib.ShowHelp()

    }

}

