package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
    "flag"
    "strings"
)

func main() {

    var opt string
    var master string
    var node string
    var mvip string
    var sshpwd string

    flag.StringVar(&opt,"opt","","Available options: init | install | addnode | delnode | rebuildmaster | delmaster")
    flag.StringVar(&master,"master","","The IP address of k8s master server filled in for the first installation.")
    flag.StringVar(&node,"node","","The IP address of k8s node server filled in for the first installation.")
    flag.StringVar(&mvip,"mvip","","K8s master cluster virtual IP address filled in for the first installation.")
    flag.StringVar(&sshpwd,"sshpwd","","SSH login root password of each server.")
    flag.Parse()

    master_array := strings.Split(master, ",")
    master_str := strings.Replace(master, "," , " " , -1)
    node_array := strings.Split(node, ",")
    node_str := strings.Replace(node, "," , " " , -1)

    softdir := "/opt/kube-install"
    path, err := os.Executable()
    checkErr(err)
    currentdir := filepath.Dir(path)
    if currentdir == "/usr/local/bin" {
	currentdir = softdir
    }


    switch {

      //Execute init command
      case opt == "init" :
        fmt.Println("Initialization in progress, please wait……") 
        time.Sleep(1 * time.Second)
        for i := 1; i <= 100; i = i + 1 {
          fmt.Fprintf(os.Stdout, "%d%% [%s]\r",i,progressBar(i,"#") + progressBar(100-i," "))
          time.Sleep(time.Second * 1)
        } 
        shellExecute(currentdir+"/workflow/sshops-init.sh \""+softdir+"\" \""+currentdir+"\"")
        fmt.Println("Initialization completed!") 

      //Execute install command
      case opt == "install" :
        fmt.Println("Deploying kubernetes cluster, please wait……") 
        checkParam(opt,master)
        checkParam(opt,node)
        checkParam(opt,mvip)
        checkParam(opt,sshpwd)
        shellExecute(currentdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \""+master_str+" "+node_str+"\" \""+softdir+"\" \""+currentdir+"\" \"install\"")
        generalConfig(master_array, node_array, mvip, currentdir, softdir)
        _, err_install := copyFile(currentdir+"/workflow/general.inventory", currentdir+"/workflow/install.inventory")
        checkErr(err_install)
        installConfig(master_array, node_array, currentdir, softdir)
        installGenfile(currentdir)
        if len(master_array) == 1{
          onemasterinstallYML(currentdir)
          shellExecute("ansible-playbook -i "+currentdir+"/workflow/install.inventory "+currentdir+"/workflow/k8scluster-onemasterinstall.yml")
        }else{
          installYML(currentdir)
          shellExecute("ansible-playbook -i "+currentdir+"/workflow/install.inventory "+currentdir+"/workflow/k8scluster-install.yml")
        }
        fmt.Println("Kubernetes cluster deployment completed!")

      //Execute addnode command
      case opt == "addnode" :
        fmt.Println("Adding k8s-node, please wait……") 
        checkParam(opt,node)
        checkParam(opt,sshpwd)
        shellExecute(softdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \""+node_str+"\" \""+softdir+"\" \""+softdir+"\" \"addnode\"")
        _, err_addnode := copyFile(softdir+"/workflow/general.inventory", softdir+"/workflow/addnode.inventory")
        checkErr(err_addnode)
        addnodeConfig(node_array, softdir)
        addnodeYML(softdir)
        shellExecute("ansible-playbook -i "+softdir+"/workflow/addnode.inventory "+softdir+"/workflow/k8scluster-addnode.yml")
        fmt.Println("K8s-node added completed!")

      //Execute delnode command
      case opt == "delnode" :
        fmt.Println("Deleting k8s-node, please wait……") 
        checkParam(opt,node)
        checkParam(opt,sshpwd)
        shellExecute(softdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \""+node_str+"\" \""+softdir+"\" \""+softdir+"\" \"delnode\"")
        _, err_delnode := copyFile(softdir+"/workflow/general.inventory", softdir+"/workflow/delnode.inventory")
        checkErr(err_delnode)
        delnodeConfig(node_array, softdir)
        delnodeYML(softdir)
        shellExecute("ansible-playbook -i "+softdir+"/workflow/delnode.inventory "+softdir+"/workflow/k8scluster-delnode.yml")
        for i := 0; i < len(node_array); i++ {
            shellExecute("kubectl delete node "+node_array[i])
        } 
        fmt.Println("K8s-node deleted completed!")

      //Execute rebuildmaster command
      case opt == "rebuildmaster" :
        fmt.Println("Rebuilding k8s-master, please wait……")
        checkParam(opt,master)
        checkParam(opt,sshpwd)
        shellExecute(softdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \""+master_str+"\" \""+softdir+"\" \""+softdir+"\" \"rebuildmaster\"")
        _, err_rebuildmaster := copyFile(softdir+"/workflow/general.inventory", softdir+"/workflow/rebuildmaster.inventory")
        checkErr(err_rebuildmaster)
        rebuildmasterConfig(master_array, softdir)
        installGenfile(softdir)
        rebuildmasterYML(softdir)
        shellExecute("ansible-playbook -i "+softdir+"/workflow/rebuildmaster.inventory "+softdir+"/workflow/k8scluster-rebuildmaster.yml")
        fmt.Println("K8s-master ebuilt completed!")

      //Execute delmaster command
      case opt == "delmaster" :
        fmt.Println("Deleting k8s-master, please wait……")
        checkParam(opt,master)
        checkParam(opt,sshpwd)
        shellExecute(softdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \""+master_str+"\" \""+softdir+"\" \""+softdir+"\" \"delmaster\"")
        _, err_delmaster := copyFile(softdir+"/workflow/general.inventory", softdir+"/workflow/delmaster.inventory")
        checkErr(err_delmaster)
        delmasterConfig(master_array, softdir)
        delmasterYML(softdir)
        shellExecute("ansible-playbook -i "+softdir+"/workflow/delmaster.inventory "+softdir+"/workflow/k8scluster-delmaster.yml")
        fmt.Println("K8s-master deleted completed!")

      //Default output help information
      default:
        showHelp()

    }

}


