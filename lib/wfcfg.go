package kilib

import (
    "os"
    "bufio"
    "strconv"
)



func GeneralConfig(master_array []string, node_array []string, mvip string, currentdir string, softdir string) {
    //Generate generic configuration
    inventory_file, err := os.Create(currentdir+"/workflow/general.inventory") 
    CheckErr(err)
    defer inventory_file.Close() 
    inventory_file.WriteString("###--------------------------------------General configuration---------------------------------###\n")
    inventory_file.WriteString("\n[master1]\n127.0.0.1 ip=127.0.0.1\n\n[k8s:vars]\n"+"k8s_install_home=\""+softdir+"/k8s\"\nsoftware_home=\""+softdir+"\"\n")
    inventory_file.WriteString("\n### k8s-master configuration ###\n")
    var master_iplist,etcd_initial,etcd_endpoints,nginx_upstream,ingress_upstream string
    master_num := len(master_array)
    node_num := len(node_array)
    for i := 0; i < master_num; i++ {
      if i > 0{
        master_iplist = master_iplist + ","
        etcd_initial = etcd_initial + ","
        etcd_endpoints = etcd_endpoints + ","
      }
      master_iplist = master_iplist+"\\\""+master_array[i]+"\\\""
      etcd_initial = etcd_initial+"kube"+strconv.Itoa(i)+"=https://"+master_array[i]+":2380"
      etcd_endpoints = etcd_endpoints+"https://"+master_array[i]+":2379"
      nginx_upstream = nginx_upstream+"server "+master_array[i]+":6443 max_fails=3 fail_timeout=30s;"
    }
    for i := 0; i < node_num; i++ {
      ingress_upstream = ingress_upstream+"server "+node_array[i]+":80 max_fails=3 fail_timeout=30s;"
    }
    inventory_file.WriteString("master_iplist=\""+master_iplist+"\"\netcd_initial=\""+etcd_initial+"\"\netcd_endpoints=\""+etcd_endpoints+"\"\nnginx_upstream=\""+nginx_upstream+"\"\ningress_upstream=\""+ingress_upstream+"\"\n")
    if master_num == 1{
      inventory_file.WriteString("master_vip=\""+master_array[0]+"\"\nmaster_vport=\"6443\"\n")
    }else{
      inventory_file.WriteString("master_vip=\""+mvip+"\"\nmaster_vport=\"8443\"\n")
    }
    //Setting the scheduling IP for addons
    switch {
      case node_num == 1 :
        inventory_file.WriteString("\n### addons_ip configuration ###\naddons_ip1=\""+node_array[0]+"\"\naddons_ip2=\""+node_array[0]+"\"\naddons_ip3=\""+node_array[0]+"\"\n")
      case node_num == 2 :
        inventory_file.WriteString("\n### addons_ip configuration ###\naddons_ip1=\""+node_array[0]+"\"\naddons_ip2=\""+node_array[1]+"\"\naddons_ip3=\""+node_array[1]+"\"\n")
      case node_num >= 3 :
        inventory_file.WriteString("\n### addons_ip configuration ###\naddons_ip1=\""+node_array[0]+"\"\naddons_ip2=\""+node_array[1]+"\"\naddons_ip3=\""+node_array[2]+"\"\n")
      default:
        panic("You must install at least one k8s-node to ensure that the cluster is running properly!")
    }
    inventory_file.WriteString("\n### traefik configuration ###\ntraefik_admin_port=\"80\"\ntraefik_data_port=\"8080\"\n")
    inventory_file.WriteString("\n### k8s-network configuration ###\nservice_cidr=\"10.254.0.0/16\"\nservice_svc_ip=\"10.254.0.1\"\nservice_dns_svc_ip=\"10.254.0.2\"\npod_cidr=\"172.30.0.0/16\"\n\n\n")

}

func InstallConfig(master_array []string, node_array []string, currentdir string, softdir string) {
    //Generate master configuration
    inventory_file, err := os.OpenFile(currentdir+"/workflow/install.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    CheckErr(err)
    defer inventory_file.Close() 
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###----------------------------------k8s-master host list------------------------------###\n")
    write.WriteString("\n[master]\n")
    master_num := len(master_array)
    for i := 0; i < master_num; i++ {
      CheckIP(master_array[i])
      write.WriteString(master_array[i]+" ip="+master_array[i]+"\n")
    }
    write.WriteString("\n[etcd]\n")
    for i := 0; i < master_num; i++ {
      write.WriteString(master_array[i]+" ip="+master_array[i]+" etcdname=kube"+strconv.Itoa(i)+"\n")
    }
    write.WriteString("\n[nginx]\n")
    j := 120
    role := "MASTER"
    for i := 0; i < master_num; i++ {
      if i > 0 {
        role = "BACKUP"
      }
      write.WriteString(master_array[i]+" ip="+master_array[i]+" priority="+strconv.Itoa(j)+" role="+role+"\n")
      j = j - 10
    }
    //Generate node configuration
    write.WriteString("\n\n\n###-----------------------------------k8s-node host list-------------------------------###\n")
    write.WriteString("\n[node]\n")
    for i := 0; i < len(node_array); i++ {
      CheckIP(node_array[i])
      write.WriteString(node_array[i]+" ip="+node_array[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"master\n"+"etcd\n"+"node\n"+"nginx\n\n\n")
    write.Flush()

}

func AddnodeConfig(addnode_array []string, softdir string) {
    //Generate addnode configuration
    inventory_file, err := os.OpenFile(softdir+"/workflow/addnode.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    CheckErr(err)
    defer inventory_file.Close() 
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------New k8s-node host list------------------------------###\n")
    write.WriteString("\n[addnode]\n")
    for i := 0; i < len(addnode_array); i++ {
      CheckIP(addnode_array[i])
      write.WriteString(addnode_array[i]+" ip="+addnode_array[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"addnode\n\n\n")
    write.Flush()

}

func DelnodeConfig(delnode_array []string, softdir string) {
    //Generate delnode configuration
    inventory_file, err := os.OpenFile(softdir+"/workflow/delnode.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    CheckErr(err)
    defer inventory_file.Close() 
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------To delete k8s-node host list------------------------------###\n")
    write.WriteString("\n[delnode]\n")
    for i := 0; i < len(delnode_array); i++ {
      CheckIP(delnode_array[i])
      write.WriteString(delnode_array[i]+" ip="+delnode_array[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"delnode\n\n\n")
    write.Flush()

}

func RebuildmasterConfig(rebuildmaster_array []string, softdir string) {
    //Generate rebuildmaster configuration
    _, err := os.Stat(softdir+"/workflow/install.inventory")
    if err != nil {
        panic(softdir+"/workflow/install.inventory file has been deleted by mistake. Please restore the file manually or contact the administrator!")
    }
    if os.IsNotExist(err) {
        panic(softdir+"/workflow/install.inventory file has been deleted by mistake. Please restore the file manually or contact the administrator!")
    }
    inventory_file, err := os.OpenFile(softdir+"/workflow/rebuildmaster.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    CheckErr(err)
    defer inventory_file.Close() 
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###----------------------------------To rebuild k8s-master host list------------------------------###\n")
    write.WriteString("\n[master]\n")
    rebuildmaster_num := len(rebuildmaster_array)
    for i := 0; i < rebuildmaster_num; i++ {
      CheckIP(rebuildmaster_array[i])
      write.WriteString(rebuildmaster_array[i]+" ip="+rebuildmaster_array[i]+"\n")
    }
    write.WriteString("\n[etcd]\n")
    for i := 0; i < rebuildmaster_num; i++ {
      etcdname := ShellOutput(softdir+"/workflow/getmasterconfig.sh "+softdir+" etcdname "+rebuildmaster_array[i])
      write.WriteString(rebuildmaster_array[i]+" ip="+rebuildmaster_array[i]+" etcdname="+etcdname+"\n")
    }
    write.WriteString("\n[nginx]\n")
    for i := 0; i < rebuildmaster_num; i++ {
      priority := ShellOutput(softdir+"/workflow/getmasterconfig.sh "+softdir+" priority "+rebuildmaster_array[i])
      write.WriteString(rebuildmaster_array[i]+" ip="+rebuildmaster_array[i]+" priority="+priority+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"master\n"+"etcd\n"+"nginx\n\n\n")
    write.Flush()

}

func DelmasterConfig(delmaster_array []string, softdir string) {
    //Generate delmaster configuration
    inventory_file, err := os.OpenFile(softdir+"/workflow/delmaster.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    CheckErr(err)
    defer inventory_file.Close() 
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------To delete k8s-master host list------------------------------###\n")
    write.WriteString("\n[delmaster]\n")
    for i := 0; i < len(delmaster_array); i++ {
      CheckIP(delmaster_array[i])
      write.WriteString(delmaster_array[i]+" ip="+delmaster_array[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"delmaster\n\n\n")
    write.Flush()

}

func UninstallConfig(delnode_array []string, delmaster_array []string, softdir string) {
    //Generate uninstall configuration
    inventory_file, err := os.OpenFile(softdir+"/workflow/uninstall.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    CheckErr(err)
    defer inventory_file.Close()
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------To delete k8s-node host list------------------------------###\n")
    write.WriteString("\n[delnode]\n")
    for i := 0; i < len(delnode_array); i++ {
      CheckIP(delnode_array[i])
      write.WriteString(delnode_array[i]+" ip="+delnode_array[i]+"\n")
    }
    write.WriteString("###---------------------------------To delete k8s-master host list------------------------------###\n")
    write.WriteString("\n[delmaster]\n")
    for i := 0; i < len(delmaster_array); i++ {
      CheckIP(delmaster_array[i])
      write.WriteString(delmaster_array[i]+" ip="+delmaster_array[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"delmaster\n"+"delnode\n\n\n")
    write.Flush()

}


