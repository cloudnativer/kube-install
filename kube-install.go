package main


import (
    "fmt"
    "os"
    "path/filepath"
    "time"
    "flag"
    "strings"
    "kube-install/lib"
)


func main() {

    var opt string
    var master string
    var node string
    var sshpwd string
    var ostype string
    var softdir string
    var currentdir string

    flag.StringVar(&opt,"opt","","Available options: init | install | addnode | delnode | rebuildmaster | delmaster | uninstall")
    flag.StringVar(&master,"master","","The IP address of k8s master server filled in for the first installation")
    flag.StringVar(&node,"node","","The IP address of k8s node server filled in for the first installation")
    flag.StringVar(&sshpwd,"sshpwd","","The root password used to SSH login to each server")
    flag.StringVar(&ostype,"ostype","","Specifies the distribution operating system type: centos7 | centos8 | rhel7 | rhel8 | suse15")
    flag.StringVar(&softdir,"softdir","/opt/kube-install","Specify the installation path of kubernetes cluster.")
    flag.Parse()

    master_array := strings.Split(master, ",")
    master_str := strings.Replace(master, "," , " " , -1)
    node_array := strings.Split(node, ",")
    node_str := strings.Replace(node, "," , " " , -1)
    path, err := os.Executable()
    kilib.CheckErr(err)
    currentdir = filepath.Dir(path)
    if currentdir == "/usr/local/bin" {
	currentdir = softdir
    }



    //Execute opt
    switch {

      //Execute init command
      case opt == "init" :
          fmt.Println("\nInitialization in progress, please wait...\n")
          ostype = kilib.CheckOS(ostype)
          time.Sleep(1 * time.Second)
          for i := 1; i <= 100; i = i + 1 {
              fmt.Fprintf(os.Stdout, "%d%% [%s]\r",i,kilib.ProgressBar(i,"#") + kilib.ProgressBar(100-i," "))
              time.Sleep(time.Second * 1)
          }
          kilib.SshOpsInit(softdir, currentdir, ostype) 
          fmt.Println("\n\nInitialization completed!\n") 

      //Execute install command
      case opt == "install" :
          fmt.Println("\nDeploying kubernetes cluster, please wait...\n") 
          kilib.CheckParam(opt,"\"-master\"",master)
          kilib.CheckParam(opt,"\"-node\"",node)
          kilib.CheckParam(opt,"\"-sshpwd\"",sshpwd)
          ostype = kilib.CheckOS(ostype)
          kilib.SshKeyInit(sshpwd, master_str+" "+node_str, softdir, currentdir, opt)
          kilib.GeneralConfig(master_array, node_array, currentdir, softdir, ostype)
          _, err_cpfile := kilib.CopyFile(currentdir+"/config/general.inventory", currentdir+"/config/install.inventory")
          kilib.CheckErr(err_cpfile)
          kilib.InstallConfig(master_array, node_array, currentdir, softdir)
          kilib.InstallGenFile(currentdir)
          kilib.InstallIpvsYaml(currentdir, master_array)
          if len(master_array) == 1{
              kilib.OnemasterinstallYML(currentdir, ostype)
              kilib.Operation(opt, currentdir, "onemasterinstall")
          }else{
              kilib.InstallYML(currentdir, ostype)
              kilib.Operation(opt, currentdir, opt)
          }
          fmt.Println("Please check the file "+softdir+"/loginkey.txt to get the login key. \n")
          fmt.Println("=============================================================================\nKubernetes cluster installation completed! \n=============================================================================\n")

      //Execute addnode command
      case opt == "addnode" :
          fmt.Println("\nAdding k8s-node, please wait...\n") 
          kilib.CheckParam(opt,"\"-node\"",node)
          kilib.CheckParam(opt,"\"-sshpwd\"",sshpwd)
          ostype = kilib.CheckOS(ostype)
          kilib.SshKeyInit(sshpwd, node_str, softdir, currentdir, opt)
          _, err_addnode := kilib.CopyFile(currentdir+"/config/general.inventory", currentdir+"/config/addnode.inventory")
          kilib.CheckErr(err_addnode)
          kilib.AddnodeConfig(node_array, currentdir)
          kilib.AddnodeYML(currentdir, ostype)
          kilib.Operation(opt, currentdir, opt)
          fmt.Println("=============================================================================\nK8s-node has been added to the kubernetes cluster! \n=============================================================================\n")

      //Execute delnode command
      case opt == "delnode" :
          fmt.Println("\nDeleting k8s-node, please wait...\n") 
          kilib.CheckParam(opt,"\"-node\"",node)
          kilib.CheckParam(opt,"\"-sshpwd\"",sshpwd)
          kilib.SshKeyInit(sshpwd, node_str, softdir, currentdir, opt)
          _, err_cpfile := kilib.CopyFile(currentdir+"/config/general.inventory", currentdir+"/config/delnode.inventory")
          kilib.CheckErr(err_cpfile)
          kilib.DelnodeConfig(node_array, currentdir)
          kilib.DelnodeYML(currentdir)
          delnodeiplist := "{"+node+"}"
          if len(node_array) == 1 { delnodeiplist = node }
          kilib.ShellExecute("kubectl delete node "+delnodeiplist )
          kilib.Operation(opt, currentdir, opt)
          fmt.Println("=============================================================================\nK8s-node has been removed from the kubernetes cluster! \n=============================================================================\n")

      //Execute rebuildmaster command
      case opt == "rebuildmaster" :
          fmt.Println("\nRebuilding k8s-master, please wait...\n")
          kilib.CheckParam(opt,"\"-master\"",master)
          kilib.CheckParam(opt,"\"-sshpwd\"",sshpwd)
          ostype = kilib.CheckOS(ostype)
          kilib.SshKeyInit(sshpwd, master_str, softdir, currentdir, opt)
          _, err_cpfile := kilib.CopyFile(currentdir+"/config/general.inventory", currentdir+"/config/rebuildmaster.inventory")
          kilib.CheckErr(err_cpfile)
          kilib.RebuildmasterConfig(master_array, currentdir)
          kilib.InstallGenFile(softdir)
          kilib.RebuildmasterYML(currentdir)
          kilib.Operation(opt, currentdir, opt)
          fmt.Println("=============================================================================\nK8s-master in the kubernetes cluster has been rebuilt! \n=============================================================================\n")

      //Execute delmaster command
      case opt == "delmaster" :
          fmt.Println("\nDeleting k8s-master, please wait...\n")
          kilib.CheckParam(opt,"\"-master\"",master)
          kilib.CheckParam(opt,"\"-sshpwd\"",sshpwd)
          kilib.SshKeyInit(sshpwd, master_str, softdir, currentdir, opt)
          _, err_cpfile := kilib.CopyFile(currentdir+"/config/general.inventory", currentdir+"/config/delmaster.inventory")
          kilib.CheckErr(err_cpfile)
          kilib.DelmasterConfig(master_array, currentdir)
          kilib.DelmasterYML(currentdir)
          kilib.Operation(opt, currentdir, opt)
          fmt.Println("=============================================================================\nK8s-master has been removed from the kubernetes cluster! \n=============================================================================\n")

      //Execute uninstall command
      case opt == "uninstall" :
          fmt.Println("\nUninstalling kubernetes cluster, please wait...\n\n")
          kilib.CheckParam(opt,"\"-master\"",master)
          kilib.CheckParam(opt,"\"-node\"",node)
          kilib.CheckParam(opt,"\"-sshpwd\"",sshpwd)
          kilib.SshKeyInit(sshpwd, master_str+" "+node_str, softdir, currentdir, opt)
          //Create tmp kube-install config dir
          os.RemoveAll("/tmp/.kube-install/config/")
          err_mkdir := os.MkdirAll("/tmp/.kube-install/config/", 0666)
          kilib.CheckErr(err_mkdir)
          //Create tmp kube-install config files
          kilib.GeneralConfig(master_array, node_array, currentdir, softdir, ostype)
          _, err_cpfile := kilib.CopyFile(currentdir+"/config/general.inventory", "/tmp/.kube-install/config/uninstall.inventory")
          kilib.CheckErr(err_cpfile)
          kilib.UninstallConfig(node_array, master_array, "/tmp/.kube-install")
          kilib.UninstallYML(currentdir)
          //Uninstall kubernetes cluster process now
          delnodeiplist := "{"+node+"}"
          if len(node_array) == 1 { delnodeiplist = node }
          fmt.Println("K8s-node list: "+delnodeiplist+" \n")
          kilib.ShellExecute("kubectl delete node "+delnodeiplist+">/dev/null 2>&1")
          fmt.Println("k8s-node delete operation execution completed!\n\nUninstall k8s-master and k8s-node software, please wait...")
          kilib.Operation(opt, "/tmp/.kube-install", opt)
          fmt.Println("K8s-master and k8s-node software uninstall operation execution completed!\n")
          os.RemoveAll("/tmp/.kube-install/")
          fmt.Println("=============================================================================\nKubernetes cluster has been installed! \n=============================================================================\n")

      //Default output help information
      default:
          kilib.ShowHelp()

    }

}



