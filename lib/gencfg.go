package kilib

import (
    "os"
    "bufio"
    "strconv"
    "strings"
)



// Set the node IP address for addons scheduling.
func CreateAddonsNode(nodeArray []string) (string,string,string){
    var addonsIp1,addonsIp2,addonsIp3 string
    node_num := len(nodeArray)
    switch {
        case node_num == 1 :
            addonsIp1,addonsIp2,addonsIp3 = nodeArray[0],nodeArray[0],nodeArray[0]
        case node_num == 2 :
            addonsIp1,addonsIp2,addonsIp3 = nodeArray[0],nodeArray[0],nodeArray[1]
        case node_num >= 3 :
            addonsIp1,addonsIp2,addonsIp3 = nodeArray[0],nodeArray[1],nodeArray[2]
    }
    return addonsIp1,addonsIp2,addonsIp3 
}

// Generate basic general configuration information.
func GeneralConfig(mode string, masterArray []string, nodeArray []string, currentDir string, softDir string, subProcessDir string, ostype string, k8sVer string, logName string) {
    inventory_file, err := os.Create(currentDir+"/data/output"+subProcessDir+"/general.inventory")
    CheckErr(err,currentDir,logName,mode)
    defer inventory_file.Close()

    inventory_file.WriteString("###--------------------------------------General Configuration---------------------------------###\n")
    inventory_file.WriteString("\n[kissh]\n127.0.0.1 ip=127.0.0.1\n\n[master1]\n"+masterArray[0]+" ip="+masterArray[0]+"\n\n[k8s:vars]\n"+"k8s_install_home=\""+softDir+"/k8s\"\n"+"k8s_addons_home=\""+currentDir+"/data/output"+subProcessDir+"/addons\"\nsoftware_home=\""+softDir+"\"\nkipath=\""+currentDir+"\"\nsub_process_dir=\""+subProcessDir+"\"\n")
    if mode == "DAEMON" {
        inventory_file.WriteString("rebootime=\"1\"\nrebootxt=\"The operating system will automatically restart to take effect on the cluster configuration.\"")
    } else {
        inventory_file.WriteString("rebootime=\"5\"\nrebootxt=\"The operating system will automatically restart in 10 seconds to take effect on the cluster configuration.\"")
    }
    inventory_file.WriteString("\n### Kubernetes Master Configuration ###\n")
    var master_iplist,etcd_initial,etcd_endpoints,ingress_upstream string
    var ipvsinit_shell string = "ipvsadm -A -t 10.254.0.3:6443 -s rr "
    var master_vip string = "10.254.0.3"
    master_num := len(masterArray)
    if master_num == 1 {
        master_vip = masterArray[0]
    }
    for i := 0; i < master_num; i++ {
        if i > 0{
            master_iplist = master_iplist + ","
            etcd_initial = etcd_initial + ","
            etcd_endpoints = etcd_endpoints + ","
        }
        master_iplist = master_iplist+"\\\""+masterArray[i]+"\\\""
        etcd_initial = etcd_initial+masterArray[i]+"=https://"+masterArray[i]+":2380"
        etcd_endpoints = etcd_endpoints+"https://"+masterArray[i]+":2379"
        ipvsinit_shell = ipvsinit_shell+" && ipvsadm -a -t 10.254.0.3:6443 -r "+masterArray[i]+":6443 -m"
    }
    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/etcdendpoints.txt", etcd_endpoints, currentDir, logName, mode)
    inventory_file.WriteString("ostype=\""+ostype+"\"\nk8sver=\""+k8sVer+"\"\nmaster_iplist=\""+master_iplist+"\"\netcd_initial=\""+etcd_initial+"\"\netcd_endpoints=\""+etcd_endpoints+"\"\ningress_upstream=\""+ingress_upstream+"\"\nipvsinit_shell = \""+ipvsinit_shell+"\"\nmaster1_ip = \""+masterArray[0]+"\"\nmaster_vip = \""+master_vip+"\"\nmaster_vport=\"6443\"\n")
    if k8sVer == "1.18" {
        inventory_file.WriteString("k8sdashboardversion=\"v1.10.1\"\n")
    } else {
        inventory_file.WriteString("k8sdashboardversion=\"v2.2.0\"\n")
    }
    // Addons IP Configuration
    addonsIp1,addonsIp2,addonsIp3 := CreateAddonsNode(nodeArray)
    inventory_file.WriteString("\n### addons_ip配置 ###\naddons_ip1=\""+addonsIp1+"\"\naddons_ip2=\""+addonsIp2+"\"\naddons_ip3=\""+addonsIp3+"\"\n")
    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/addons/addonsip/registryip.txt", addonsIp1, currentDir, logName, mode)
    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/addons/addonsip/k8sdashboardip.txt", addonsIp3, currentDir, logName, mode)
    inventory_file.WriteString("\n### traefik configuration ###\ntraefik_admin_port=\"80\"\ntraefik_data_port=\"8080\"\n")
    inventory_file.WriteString("\n### k8s network配置 ###\nservice_cidr=\"10.254.0.0/16\"\nservice_svc_ip=\"10.254.0.1\"\nservice_dns_svc_ip=\"10.254.0.2\"\npod_cidr=\"10.244.0.0/16\"\n\n\n")
}

// Generate install configuration information.
func InstallConfig(mode string, masterArray []string, nodeArray []string, currentDir string, subProcessDir string, logName string) bool {
    _,err_cp := CopyFile(currentDir+"/data/output"+subProcessDir+"/general.inventory", currentDir+"/data/output"+subProcessDir+"/install.inventory")
    if err_cp != nil {
        CheckErr(err_cp,currentDir,logName,mode)
    } else {
        CheckFileExist(currentDir+"/data/output"+subProcessDir+"/", "install.inventory", currentDir, logName, mode)
    }

    // Kubernetes Master IP List
    inventory_file, err := os.OpenFile(currentDir+"/data/output"+subProcessDir+"/install.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    CheckErr(err,currentDir,logName,mode)
    defer inventory_file.Close() 

    write := bufio.NewWriter(inventory_file)
    write.WriteString("###----------------------------------Kubernetes Master IP List------------------------------###\n")
    write.WriteString("\n[master]\n")
    master_num := len(masterArray)
    for i := 0; i < master_num; i++ {
      if !CheckIP(masterArray[i]) {
          return false
      }
      write.WriteString(masterArray[i]+" ip="+masterArray[i]+"\n")
    }
    write.WriteString("\n[etcd]\n")
    for i := 0; i < master_num; i++ {
      write.WriteString(masterArray[i]+" ip="+masterArray[i]+" etcdname=kube"+strconv.Itoa(i)+"\n")
    }

    // Kubernetes Node IP List
    write.WriteString("\n\n\n###-----------------------------------Kubernetes Node IP List-------------------------------###\n")
    write.WriteString("\n[node]\n")
    for i := 0; i < len(nodeArray); i++ {
      if !CheckIP(nodeArray[i]) {
          return false
      }
      write.WriteString(nodeArray[i]+" ip="+nodeArray[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"kissh\n"+"master1\n"+"master\n"+"etcd\n"+"node\n\n\n")
    write.Flush()

    return true
}

// Generate configuration information of add node.
func AddnodeConfig(mode string, addNodeArray []string, currentDir string, subProcessDir string, logName string) bool {
    _,err_cp := CopyFile(currentDir+"/data/output"+subProcessDir+"/general.inventory", currentDir+"/data/output"+subProcessDir+"/addnode.inventory")
    if err_cp != nil {
        CheckErr(err_cp,currentDir,logName,mode)
    } else {
        CheckFileExist(currentDir+"/data/output"+subProcessDir+"/", "addnode.inventory", currentDir, logName, mode)
    }

    //Add Node IP List
    inventory_file, err := os.OpenFile(currentDir+"/data/output"+subProcessDir+"/addnode.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    CheckErr(err,currentDir,logName,mode)
    defer inventory_file.Close() 

    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------Add Node IP List------------------------------###\n")
    write.WriteString("\n[addnode]\n")
    for i := 0; i < len(addNodeArray); i++ {
      if !CheckIP(addNodeArray[i]) {
          return false
      }
      write.WriteString(addNodeArray[i]+" ip="+addNodeArray[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"kissh\n"+"master1\n"+"addnode\n\n\n")
    write.Flush()

    return true
}

// Generate configuration information of delete node.
func DelnodeConfig(mode string, delNodeArray []string, currentDir string, subProcessDir string, logName string) bool {
    _,err_cp := CopyFile(currentDir+"/data/output"+subProcessDir+"/general.inventory", currentDir+"/data/output"+subProcessDir+"/delnode.inventory")
    if err_cp != nil {
        CheckErr(err_cp,currentDir,logName,mode)
    } else {
        CheckFileExist(currentDir+"/data/output"+subProcessDir+"/", "delnode.inventory", currentDir, logName, mode)
    }

    //Delete Node IP List
    inventory_file, err := os.OpenFile(currentDir+"/data/output"+subProcessDir+"/delnode.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    CheckErr(err,currentDir,logName,mode)
    defer inventory_file.Close() 

    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------Delete Node IP List------------------------------###\n")
    write.WriteString("\n[delnode]\n")
    for i := 0; i < len(delNodeArray); i++ {
      if !CheckIP(delNodeArray[i]) {
          return false
      }
      write.WriteString(delNodeArray[i]+" ip="+delNodeArray[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"kissh\n"+"master1\n"+"delnode\n\n\n")
    write.Flush()

    return true
}

// Generate configuration information of rebuild master.
func RebuildmasterConfig(mode string, rebuildMasterArray []string, currentDir string, subProcessDir string, logName string) bool {
    _,err_cp := CopyFile(currentDir+"/data/output"+subProcessDir+"/general.inventory", currentDir+"/data/output"+subProcessDir+"/rebuildmaster.inventory")
    if err_cp != nil {
        CheckErr(err_cp,currentDir,logName,mode)
    } else {
        CheckFileExist(currentDir+"/data/output"+subProcessDir+"/", "rebuildmaster.inventory", currentDir, logName, mode)
    }

    //Rebuild Master IP List
    inventory_file, err := os.OpenFile(currentDir+"/data/output"+subProcessDir+"/rebuildmaster.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    CheckErr(err,currentDir,logName,mode)
    defer inventory_file.Close() 

    write := bufio.NewWriter(inventory_file)
    write.WriteString("###----------------------------------Rebuild Master IP List------------------------------###\n")
    write.WriteString("\n[master]\n")
    rebuildmaster_num := len(rebuildMasterArray)
    for i := 0; i < rebuildmaster_num; i++ {
      if !CheckIP(rebuildMasterArray[i]) {
          return false
      }
      write.WriteString(rebuildMasterArray[i]+" ip="+rebuildMasterArray[i]+"\n")
    }
    write.WriteString("\n[etcd]\n")
    for i := 0; i < rebuildmaster_num; i++ {
      etcdname := ShellOutput(currentDir+"/data/output"+subProcessDir+"/sys/0x0000000000base/prestart/getmasterconfig.sh "+currentDir+"/data/output"+subProcessDir+"' etcdname "+rebuildMasterArray[i]+" 4")
      if !strings.Contains(etcdname, "etcdname") {
        etcdname = ShellOutput(currentDir+"/data/output"+subProcessDir+"/sys/0x0000000000base/prestart/getmasterconfig.sh "+currentDir+"/data/output"+subProcessDir+"' etcdname "+rebuildMasterArray[i]+" 3")
      }
      write.WriteString(rebuildMasterArray[i]+" ip="+rebuildMasterArray[i]+" "+etcdname+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"kissh\n"+"master1\n"+"master\n"+"etcd\n\n\n")
    write.Flush()

    return true
}

// Generate configuration information of delete master.
func DelmasterConfig(mode string, delMasterArray []string, currentDir string, subProcessDir string, logName string) bool {
    _,err_cp := CopyFile(currentDir+"/data/output"+subProcessDir+"/general.inventory", currentDir+"/data/output"+subProcessDir+"/delmaster.inventory")
    if err_cp != nil {
        CheckErr(err_cp,currentDir,logName,mode)
    } else {
        CheckFileExist(currentDir+"/data/output"+subProcessDir+"/", "delmaster.inventory", currentDir, logName, mode)
    }

    //Delete Master IP List  
    inventory_file, err := os.OpenFile(currentDir+"/data/output"+subProcessDir+"/delmaster.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    CheckErr(err,currentDir,logName,mode)
    defer inventory_file.Close() 

    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------Delete Master IP List  ------------------------------###\n")
    write.WriteString("\n[delmaster]\n")
    for i := 0; i < len(delMasterArray); i++ {
      if !CheckIP(delMasterArray[i]) {
          return false
      }
      write.WriteString(delMasterArray[i]+" ip="+delMasterArray[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"kissh\n"+"master1\n"+"delmaster\n\n\n")
    write.Flush()

    return true
}
 

