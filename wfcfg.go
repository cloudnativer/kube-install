package main

import (
    "os"
    "bufio"
    "strconv"
)



func generalConfig(master_array []string, node_array []string, mvip string, currentdir string, softdir string) {
    //生成通用配置
    inventory_file, err := os.Create(currentdir+"/workflow/general.inventory") 
    checkErr(err)
    defer inventory_file.Close() 
    inventory_file.WriteString("###--------------------------------------k8s通用配置---------------------------------###\n")
    inventory_file.WriteString("\n[master1]\n")
    inventory_file.WriteString(master_array[0]+" ip="+master_array[0]+"\n")
    inventory_file.WriteString("\n[k8s:vars]\n"+"k8s_install_home=\""+softdir+"/k8s\"\n")
    inventory_file.WriteString("software_home=\""+softdir+"\"\n")
    inventory_file.WriteString("\n### k8s-master配置 ###\n")
    var master_iplist,etcd_initial,etcd_endpoints,nginx_upstream string
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
    inventory_file.WriteString("master_iplist=\""+master_iplist+"\"\n")
    inventory_file.WriteString("etcd_initial=\""+etcd_initial+"\"\n")
    inventory_file.WriteString("etcd_endpoints=\""+etcd_endpoints+"\"\n")
    inventory_file.WriteString("nginx_upstream=\""+nginx_upstream+"\"\n")
    if master_num == 1{
      inventory_file.WriteString("master_vip=\""+master_array[0]+"\"\n")
      inventory_file.WriteString("master_vport=\"6443\"\n")
    }else{
      inventory_file.WriteString("master_vip=\""+mvip+"\"\n")
      inventory_file.WriteString("master_vport=\"8443\"\n")
    }
    if node_num > 1{
      inventory_file.WriteString("\n### dashboard配置 ###\n")
      inventory_file.WriteString("dashboard_ip=\""+node_array[0]+"\"\n")
      inventory_file.WriteString("\n### registry配置 ###\n")
      inventory_file.WriteString("registry_ip=\""+node_array[1]+"\"\n")
    }else{
      inventory_file.WriteString("\n### dashboard配置 ###\n")
      inventory_file.WriteString("dashboard_ip=\""+node_array[0]+"\"\n")
      inventory_file.WriteString("\n### registry配置 ###\n")
      inventory_file.WriteString("registry_ip=\""+node_array[0]+"\"\n")
    }
    inventory_file.WriteString("\n### traefik配置 ###\n")
    inventory_file.WriteString("traefik_admin_port=\"80\"\n")
    inventory_file.WriteString("traefik_data_port=\"8080\"\n")
    inventory_file.WriteString("\n### k8s-network配置 ###\n")
    inventory_file.WriteString("service_cidr=\"10.254.0.0/16\"\n")
    inventory_file.WriteString("service_svc_ip=\"10.254.0.1\"\n")
    inventory_file.WriteString("service_dns_svc_ip=\"10.254.0.2\"\n")
    inventory_file.WriteString("pod_cidr=\"172.30.0.0/16\"\n\n\n")

}

func installConfig(master_array []string, node_array []string, currentdir string, softdir string) {
    //生成master配置
    inventory_file, err := os.OpenFile(currentdir+"/workflow/install.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    checkErr(err)
    defer inventory_file.Close() 
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###----------------------------------k8s-master主机列表------------------------------###\n")
    write.WriteString("\n[master]\n")
    master_num := len(master_array)
    for i := 0; i < master_num; i++ {
      checkIP(master_array[i])
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
    //生成node配置
    write.WriteString("\n\n\n###-----------------------------------k8s-node主机列表-------------------------------###\n")
    write.WriteString("\n[node]\n")
    for i := 0; i < len(node_array); i++ {
      checkIP(node_array[i])
      write.WriteString(node_array[i]+" ip="+node_array[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"master\n"+"etcd\n"+"node\n"+"nginx\n\n\n")
    write.Flush()

}

func addnodeConfig(addnode_array []string, softdir string) {
    //生成addnode配置
    inventory_file, err := os.OpenFile(softdir+"/workflow/addnode.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    checkErr(err)
    defer inventory_file.Close() 
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------新增的k8s-node主机列表------------------------------###\n")
    write.WriteString("\n[addnode]\n")
    for i := 0; i < len(addnode_array); i++ {
      checkIP(addnode_array[i])
      write.WriteString(addnode_array[i]+" ip="+addnode_array[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"addnode\n\n\n")
    write.Flush()

}

func delnodeConfig(delnode_array []string, softdir string) {
    //生成delnode配置
    inventory_file, err := os.OpenFile(softdir+"/workflow/delnode.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    checkErr(err)
    defer inventory_file.Close() 
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------被删除的k8s-node主机列表------------------------------###\n")
    write.WriteString("\n[delnode]\n")
    for i := 0; i < len(delnode_array); i++ {
      checkIP(delnode_array[i])
      write.WriteString(delnode_array[i]+" ip="+delnode_array[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"delnode\n\n\n")
    write.Flush()

}

func rebuildmasterConfig(rebuildmaster_array []string, softdir string) {
    //生成rebuildmaster配置
    _, err := os.Stat(softdir+"/workflow/install.inventory")
    if err != nil {
        panic(softdir+"/workflow/install.inventory文件已被您误删除，请手工恢复该文件or联系管理员！")
    }
    if os.IsNotExist(err) {
        panic(softdir+"/workflow/install.inventory文件已被您误删除，请手工恢复该文件or联系管理员！")
    }
    inventory_file, err := os.OpenFile(softdir+"/workflow/rebuildmaster.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    checkErr(err)
    defer inventory_file.Close() 
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###----------------------------------要重建的k8s-master主机列表------------------------------###\n")
    write.WriteString("\n[master]\n")
    rebuildmaster_num := len(rebuildmaster_array)
    for i := 0; i < rebuildmaster_num; i++ {
      checkIP(rebuildmaster_array[i])
      write.WriteString(rebuildmaster_array[i]+" ip="+rebuildmaster_array[i]+"\n")
    }
    write.WriteString("\n[etcd]\n")
    for i := 0; i < rebuildmaster_num; i++ {
      etcdname := shellOutput(softdir+"/workflow/getmasterconfig.sh "+softdir+" etcdname "+rebuildmaster_array[i])
      write.WriteString(rebuildmaster_array[i]+" ip="+rebuildmaster_array[i]+" etcdname="+etcdname+"\n")
    }
    write.WriteString("\n[nginx]\n")
    for i := 0; i < rebuildmaster_num; i++ {
      priority := shellOutput(softdir+"/workflow/getmasterconfig.sh "+softdir+" priority "+rebuildmaster_array[i])
      write.WriteString(rebuildmaster_array[i]+" ip="+rebuildmaster_array[i]+" priority="+priority+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"master\n"+"etcd\n"+"nginx\n\n\n")
    write.Flush()

}

func delmasterConfig(delmaster_array []string, softdir string) {
    //生成delmaster配置
    inventory_file, err := os.OpenFile(softdir+"/workflow/delmaster.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    checkErr(err)
    defer inventory_file.Close() 
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------被删除的k8s-master主机列表------------------------------###\n")
    write.WriteString("\n[delmaster]\n")
    for i := 0; i < len(delmaster_array); i++ {
      checkIP(delmaster_array[i])
      write.WriteString(delmaster_array[i]+" ip="+delmaster_array[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"delmaster\n\n\n")
    write.Flush()

}


