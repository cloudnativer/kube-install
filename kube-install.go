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

    flag.StringVar(&opt,"opt","","Available options: init | install | addnode | delnode | rebuildmaster | delmaster | uninstall")
    flag.StringVar(&master,"master","","The IP address of k8s master server filled in for the first installation.")
    flag.StringVar(&node,"node","","The IP address of k8s node server filled in for the first installation.")
    flag.StringVar(&sshpwd,"sshpwd","","The root password used to SSH login to each server.")
    flag.Parse()

    master_array := strings.Split(master, ",")
    master_str := strings.Replace(master, "," , " " , -1)
    node_array := strings.Split(node, ",")
    node_str := strings.Replace(node, "," , " " , -1)

    softdir := "/opt/kube-install"
    path, err := os.Executable()
    kilib.CheckErr(err)
    currentdir := filepath.Dir(path)
    if currentdir == "/usr/local/bin" {
	currentdir = softdir
    }


    switch {

      //Execute init command
      case opt == "init" :
          fmt.Println("\nInitialization in progress, please wait...\n") 
          time.Sleep(1 * time.Second)
          for i := 1; i <= 100; i = i + 1 {
              fmt.Fprintf(os.Stdout, "%d%% [%s]\r",i,kilib.ProgressBar(i,"#") + kilib.ProgressBar(100-i," "))
              time.Sleep(time.Second * 1)
        } 
        kilib.ShellExecute(currentdir+"/workflow/sshops-init.sh \""+softdir+"\" \""+currentdir+"\"")
        fmt.Println("\n\nInitialization completed!\n") 

      //Execute install command
      case opt == "install" :
          fmt.Println("\nDeploying kubernetes cluster, please wait...\n") 
          kilib.CheckParam(opt,"master",master)
          kilib.CheckParam(opt,"node",node)
          kilib.CheckParam(opt,"sshpwd",sshpwd)
          kilib.ShellExecute(currentdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \"127.0.0.1 "+master_str+" "+node_str+"\" \""+softdir+"\" \""+currentdir+"\" \"install\"")
          kilib.GeneralConfig(master_array, node_array, currentdir, softdir)
          _, err_install := kilib.CopyFile(currentdir+"/workflow/general.inventory", currentdir+"/workflow/install.inventory")
          kilib.CheckErr(err_install)
          kilib.InstallConfig(master_array, node_array, currentdir, softdir)
          kilib.InstallGenfile(currentdir)
          if len(master_array) == 1{
              kilib.OnemasterinstallYML(currentdir)
              kilib.ShellExecute("ansible-playbook -i "+currentdir+"/workflow/install.inventory "+currentdir+"/workflow/k8scluster-onemasterinstall.yml")
          }else{
              kilib.InstallYML(currentdir)
              kilib.ShellExecute("ansible-playbook -i "+currentdir+"/workflow/install.inventory "+currentdir+"/workflow/k8scluster-install.yml")
          }
          fmt.Println("Kubernetes cluster deployment operation execution completed!")

      //Execute addnode command
      case opt == "addnode" :
          fmt.Println("\nAdding k8s-node, please wait...\n") 
          kilib.CheckParam(opt,"node",node)
          kilib.CheckParam(opt,"sshpwd",sshpwd)
          kilib.ShellExecute(softdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \"127.0.0.1 "+node_str+"\" \""+softdir+"\" \""+softdir+"\" \"addnode\"")
          _, err_addnode := kilib.CopyFile(softdir+"/workflow/general.inventory", softdir+"/workflow/addnode.inventory")
          kilib.CheckErr(err_addnode)
          kilib.AddnodeConfig(node_array, softdir)
          kilib.AddnodeYML(softdir)
          kilib.ShellExecute("ansible-playbook -i "+softdir+"/workflow/addnode.inventory "+softdir+"/workflow/k8scluster-addnode.yml")
          fmt.Println("K8s-node add operation execution completed!")

      //Execute delnode command
      case opt == "delnode" :
          fmt.Println("\nDeleting k8s-node, please wait...\n") 
          kilib.CheckParam(opt,"node",node)
          kilib.CheckParam(opt,"sshpwd",sshpwd)
          kilib.ShellExecute(softdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \"127.0.0.1 "+node_str+"\" \""+softdir+"\" \""+softdir+"\" \"delnode\"")
          _, err_delnode := kilib.CopyFile(softdir+"/workflow/general.inventory", softdir+"/workflow/delnode.inventory")
          kilib.CheckErr(err_delnode)
          kilib.DelnodeConfig(node_array, softdir)
          kilib.DelnodeYML(softdir)
          delnodeiplist := "{"+node+"}"
          if len(node_array) == 1 { delnodeiplist = node }
          kilib.ShellExecute("kubectl delete node "+delnodeiplist )
          kilib.ShellExecute("ansible-playbook -i "+softdir+"/workflow/delnode.inventory "+softdir+"/workflow/k8scluster-delnode.yml")
          fmt.Println("K8s-node delete operation execution completed!")

      //Execute rebuildmaster command
      case opt == "rebuildmaster" :
          fmt.Println("\nRebuilding k8s-master, please wait...\n")
          kilib.CheckParam(opt,"master",master)
          kilib.CheckParam(opt,"sshpwd",sshpwd)
          kilib.ShellExecute(softdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \"127.0.0.1 "+master_str+"\" \""+softdir+"\" \""+softdir+"\" \"rebuildmaster\"")
          _, err_rebuildmaster := kilib.CopyFile(softdir+"/workflow/general.inventory", softdir+"/workflow/rebuildmaster.inventory")
          kilib.CheckErr(err_rebuildmaster)
          kilib.RebuildmasterConfig(master_array, softdir)
          kilib.InstallGenfile(softdir)
          kilib.RebuildmasterYML(softdir)
          kilib.ShellExecute("ansible-playbook -i "+softdir+"/workflow/rebuildmaster.inventory "+softdir+"/workflow/k8scluster-rebuildmaster.yml")
          fmt.Println("K8s-master rebuilt operation execution completed!")

      //Execute delmaster command
      case opt == "delmaster" :
          fmt.Println("\nDeleting k8s-master, please wait...\n")
          kilib.CheckParam(opt,"master",master)
          kilib.CheckParam(opt,"sshpwd",sshpwd)
          kilib.ShellExecute(softdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \"127.0.0.1 "+master_str+"\" \""+softdir+"\" \""+softdir+"\" \"delmaster\"")
          _, err_delmaster := kilib.CopyFile(softdir+"/workflow/general.inventory", softdir+"/workflow/delmaster.inventory")
          kilib.CheckErr(err_delmaster)
          kilib.DelmasterConfig(master_array, softdir)
          kilib.DelmasterYML(softdir)
          kilib.ShellExecute("ansible-playbook -i "+softdir+"/workflow/delmaster.inventory "+softdir+"/workflow/k8scluster-delmaster.yml")
          fmt.Println("K8s-master delete operation execution completed!")

      //Execute uninstall command
      case opt == "uninstall" :
          fmt.Println("\nUninstalling kubernetes cluster, please wait...\n\n")
          kilib.CheckParam(opt,"master",master)
          kilib.CheckParam(opt,"node",node)
          kilib.CheckParam(opt,"sshpwd",sshpwd)
          kilib.ShellExecute(currentdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \""+master_str+" "+node_str+"\" \""+softdir+"\" \""+currentdir+"\" \"install\"")
          //Create tmp workflow dir
          err_rmdir:= os.RemoveAll("/tmp/workflow/")
          kilib.CheckErr(err_rmdir)
          err_mkdir := os.Mkdir("/tmp/workflow", 0666)
          kilib.CheckErr(err_mkdir)
          //Create tmp config files
          kilib.GeneralConfig(master_array, node_array, currentdir, softdir)
          _, err_uninstall := kilib.CopyFile(currentdir+"/workflow/general.inventory", "/tmp/workflow/uninstall.inventory")
          kilib.CheckErr(err_uninstall)
          kilib.UninstallConfig(node_array, master_array, "/tmp")
          kilib.UninstallYML(currentdir)
          //Uninstall kubernetes cluster process now
          delnodeiplist := "{"+node+"}"
          if len(node_array) == 1 { delnodeiplist = node }
          fmt.Println("K8s-node list: "+delnodeiplist+" \n")
          kilib.ShellExecute("kubectl delete node "+delnodeiplist+">/dev/null 2>&1")
          fmt.Println("k8s-node delete operation execution completed!\n\nUninstall k8s-master and k8s-node software, please wait...")
          kilib.ShellExecute("ansible-playbook -i /tmp/workflow/uninstall.inventory /tmp/workflow/k8scluster-uninstall.yml")
          fmt.Println("K8s-master and k8s-node software uninstall operation execution completed!\n\n**********************************************************************************\n\n")
          os.RemoveAll("/tmp/workflow/")
          fmt.Println("Kubernetes cluster uninstall operation execution completed!\n\n")

      //Default output help information
      default:
          kilib.ShowHelp()

    }

}



