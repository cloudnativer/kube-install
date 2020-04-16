package main

import (
    "fmt"
    "net"
    "os"
    "os/exec"
    "bufio"
    "io"
    "io/ioutil"
    "path/filepath"
    "log"
    "time"
    "flag"
    "strings"
    "strconv"
)

func main() {

    var opt string
    var master string
    var node string
    var mvip string
    var addnode string
    var delnode string
    var rebuildmaster string
    var delmaster string
    var sshpwd string

    flag.StringVar(&opt,"opt","","Available options: init | install | addnode | delnode | rebuildmaster | delmaster")
    flag.StringVar(&master,"master","","The IP address of k8s master server filled in for the first installation.")
    flag.StringVar(&node,"node","","The IP address of k8s node server filled in for the first installation.")
    flag.StringVar(&mvip,"mvip","","K8s master cluster virtual IP address filled in for the first installation.")
    flag.StringVar(&addnode,"addnode","","IP address of k8s node server to be added.")
    flag.StringVar(&delnode,"delnode","","IP address of k8s node server to be deleted.")
    flag.StringVar(&rebuildmaster,"rebuildmaster","","IP address of k8s master server to be rebuilt.")
    flag.StringVar(&delmaster,"delmaster","","IP address of k8s master server to be deleted.")
    flag.StringVar(&sshpwd,"sshpwd","","SSH login root password of each server.")
    flag.Parse()

    master_array := strings.Split(master, ",")
    master_str := strings.Replace(master, "," , " " , -1)
    node_array := strings.Split(node, ",")
    node_str := strings.Replace(node, "," , " " , -1)
    addnode_array := strings.Split(addnode, ",")
    addnode_str := strings.Replace(addnode, "," , " " , -1)
    delnode_array := strings.Split(delnode, ",")
    delnode_str := strings.Replace(delnode, "," , " " , -1)
    rebuildmaster_array := strings.Split(rebuildmaster, ",")
    rebuildmaster_str := strings.Replace(rebuildmaster, "," , " " , -1)
    delmaster_array := strings.Split(delmaster, ",")
    delmaster_str := strings.Replace(delmaster, "," , " " , -1)

//设置各种path
    softdir := "/opt/kube-install"
    path, err := os.Executable()
    checkErr(err)
    currentdir := filepath.Dir(path)
    if currentdir == "/usr/local/bin" {
	currentdir = softdir
    }


//执行部署指令
    /**
      使用举例如下：
        初始化
        ./kube-install -opt init
        安装k8s集群
        ./kube-install -opt install -master "192.168.122.11,192.168.122.12,192.168.122.13" -node "192.168.122.11,192.168.122.12,192.168.122.13,192.168.122.14" -mvip "192.168.122.100" -sshpwd "cloudnativer"
        添加k8s-node节点
        ./kube-install -opt addnode -addnode "192.168.122.15,192.168.122.16" -sshpwd "cloudnativer"
        删除k8s-node节点
        ./kube-install -opt delnode -delnode "192.168.122.13,192.168.122.15" -sshpwd "cloudnativer"
        删除k8s-master节点
        ./kube-install -opt delmaster -delmaster "192.168.122.13" -sshpwd "cloudnativer"
        重建k8s-master节点
        ./kube-install -opt rebuildmaster -rebuildmaster "192.168.122.13" -sshpwd "cloudnativer"
    **/


    switch {

      //执行init指令
      case opt == "init" :
        fmt.Println("正在进行初始化，请稍后……") 
        //显示初始化进度条
        time.Sleep(1 * time.Second)
        for i := 1; i <= 100; i = i + 1 {
          fmt.Fprintf(os.Stdout, "%d%% [%s]\r",i,progressBar(i,"#") + progressBar(100-i," "))
          time.Sleep(time.Second * 1)
        } 
        shellExecute(currentdir+"/workflow/sshops-init.sh \""+softdir+"\" \""+currentdir+"\"")
        fmt.Println("初始化完毕！") 

      //执行install指令
      case opt == "install" :
        fmt.Println("正在部署kubernetes集群，请稍后……") 
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
        fmt.Println("kubernetes集群部署完毕！")

      //执行addnode指令
      case opt == "addnode" :
        fmt.Println("正在添加k8s-node节点，请稍后……") 
        checkParam(opt,addnode)
        checkParam(opt,sshpwd)
        shellExecute(softdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \""+addnode_str+"\" \""+softdir+"\" \""+softdir+"\" \"addnode\"")
        _, err_addnode := copyFile(softdir+"/workflow/general.inventory", softdir+"/workflow/addnode.inventory")
        checkErr(err_addnode)
        addnodeConfig(addnode_array, softdir)
        addnodeYML(softdir)
        shellExecute("ansible-playbook -i "+softdir+"/workflow/addnode.inventory "+softdir+"/workflow/k8scluster-addnode.yml")
        fmt.Println("k8s-node节点添加完毕！")

      //执行delnode指令
      case opt == "delnode" :
        fmt.Println("正在删除k8s-node节点，请稍后……") 
        checkParam(opt,delnode)
        checkParam(opt,sshpwd)
        shellExecute(softdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \""+delnode_str+"\" \""+softdir+"\" \""+softdir+"\" \"delnode\"")
        _, err_delnode := copyFile(softdir+"/workflow/general.inventory", softdir+"/workflow/delnode.inventory")
        checkErr(err_delnode)
        delnodeConfig(delnode_array, softdir)
        delnodeYML(softdir)
        shellExecute("ansible-playbook -i "+softdir+"/workflow/delnode.inventory "+softdir+"/workflow/k8scluster-delnode.yml")
        for i := 0; i < len(delnode_array); i++ {
            shellExecute("kubectl delete node "+delnode_array[i])
        } 
        fmt.Println("k8s-node节点删除完毕！")

      //执行rebuildmaster指令
      case opt == "rebuildmaster" :
        fmt.Println("正在重建k8s-master节点，请稍后……")
        checkParam(opt,rebuildmaster)
        checkParam(opt,sshpwd)
        shellExecute(softdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \""+rebuildmaster_str+"\" \""+softdir+"\" \""+softdir+"\" \"rebuildmaster\"")
        _, err_rebuildmaster := copyFile(softdir+"/workflow/general.inventory", softdir+"/workflow/rebuildmaster.inventory")
        checkErr(err_rebuildmaster)
        rebuildmasterConfig(rebuildmaster_array, softdir)
        rebuildmasterYML(softdir)
        shellExecute("ansible-playbook -i "+softdir+"/workflow/rebuildmaster.inventory "+softdir+"/workflow/k8scluster-rebuildmaster.yml")
        fmt.Println("k8s-master节点重建完毕！")

      //执行delmaster指令
      case opt == "delmaster" :
        fmt.Println("正在删除k8s-master节点，请稍后……")
        checkParam(opt,delmaster)
        checkParam(opt,sshpwd)
        shellExecute(softdir+"/workflow/sshkey-init.sh \""+sshpwd+"\" \""+delmaster_str+"\" \""+softdir+"\" \""+softdir+"\" \"delmaster\"")
        _, err_delmaster := copyFile(softdir+"/workflow/general.inventory", softdir+"/workflow/delmaster.inventory")
        checkErr(err_delmaster)
        delmasterConfig(delmaster_array, softdir)
        delmasterYML(softdir)
        shellExecute("ansible-playbook -i "+softdir+"/workflow/delmaster.inventory "+softdir+"/workflow/k8scluster-delmaster.yml")
        fmt.Println("k8s-master节点删除完毕！")


      default:
        panic("您输入的-opt参数有误，请检查！")

    }

}




func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func checkIP(ipv4 string) {
    address := net.ParseIP(ipv4)  
    if address == nil {
         panic("您输入的IP地址格式有误，请检查！")
    }
}

func checkParam(option string, param string) {
    if param == "" {
         panic("执行"+option+"操作时，必须输入"+param+"参数，请检查！")
    }
}

func progressBar(n int,char string) (s string) {
    for i:=1;i<=n;i++{
        s+=char
    }
    return
}

func copyFile(srcFileName string, dstFileName string) (written int64, err error) {
    //用于拷贝文件的函数，接收两个文件路径 srcFileName dstFileName
    srcFile, err := os.Open(srcFileName)
    if err != nil {
            fmt.Printf("open file err = %v\n", err)
            return
    }
    defer srcFile.Close()

    //通过srcFile，获取到Reader
    reader := bufio.NewReader(srcFile)

    //打开dstFileName
    dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666)
    if err != nil {
            fmt.Printf("open file err = %v\n", err)
            return
    }

    writer := bufio.NewWriter(dstFile)
    defer func() {
            writer.Flush() //把缓冲区的内容写入到文件
            dstFile.Close()
    }()

    return io.Copy(writer, reader)

}

func installGenfile(currentdir string) {
    genfile_file, err := os.Create(currentdir+"/bin/0.base/genfile/tasks/main.yml") //新建inventory配置文件
    checkErr(err)
    defer genfile_file.Close() //main函数结束前， 关闭文件
    genfile_file.WriteString("- name: 1.创建{{software_home}}目录\n  file:\n    path: \"{{software_home}}\"\n    state: directory\n")
    genfile_file.WriteString("- name: 2.正在将部署文件分发到k8s-master\n  copy:\n    src: \""+currentdir+"/{{item}}\"\n    dest: \"{{software_home}}/\"\n  with_items:\n    - bin\n    - docs\n    - pkg\n    - workflow\n    - yaml\n    - kube-install\n- copy:\n    src: \""+currentdir+"/kube-install\"\n    dest: \"/usr/local/bin/kube-install\"\n    mode: 0755\n")
    genfile_file.WriteString("- name: 3.配置可执行文件的权限\n  file: path={{software_home}}/{{ item }} mode=755 owner=root group=root\n  with_items:\n    - workflow/sshkey-init.sh\n    - workflow/sshops-init.sh\n    - workflow/getmasterconfig.sh\n")

}

func installYML(softdir string) {
    install_file, err := os.Create(softdir+"/workflow/k8scluster-install.yml") //新建inventory配置文件
    checkErr(err)
    defer install_file.Close() //main函数结束前， 关闭文件
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/genfile\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/kernel\n")
    install_file.WriteString("- remote_user: root\n  hosts: master,node,nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/all\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delnode\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/docker\n")
    install_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/copycfssl\n")
    install_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/createssl\n")
    install_file.WriteString("- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/2.etcd\n")
    install_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/etcd_network\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/flanneld\n")
    install_file.WriteString("- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/nginx\n")
    install_file.WriteString("- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/kubectl\n")
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/apiserver\n")
    install_file.WriteString("- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/keepalived\n")
    install_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/api-rbac\n")
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/controller-manager\n")
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/scheduler\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kubelet\n")
    install_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/approve-csr\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kube-proxy\n")
    install_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/7.kube-addons\n")
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/admintoken\n")
    install_file.WriteString("- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/alter\n")
    install_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/9.finish/install\n")
}

func onemasterinstallYML(softdir string) {
    onemasterinstall_file, err := os.Create(softdir+"/workflow/k8scluster-onemasterinstall.yml") //新建inventory配置文件
    checkErr(err)
    defer onemasterinstall_file.Close() //main函数结束前， 关闭文件
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/kernel\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/all\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delnode/stopnode\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/docker\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/copycfssl\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/createssl\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/2.etcd\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/etcd_network\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/flanneld\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/kubectl\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/apiserver\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/api-rbac\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/controller-manager\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/scheduler\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kubelet\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/approve-csr\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kube-proxy\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/7.kube-addons\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/admintoken\n")
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/9.finish/install\n")
}

func addnodeYML(softdir string) {
    addnode_file, err := os.Create(softdir+"/workflow/k8scluster-addnode.yml") //新建inventory配置文件
    checkErr(err)
    defer addnode_file.Close() //main函数结束前， 关闭文件
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/kernel\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/all\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delnode\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/docker\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/flanneld\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kubelet\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kube-proxy\n")
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/9.finish/addnode\n")
}

func delnodeYML(softdir string) {
    delnode_file, err := os.Create(softdir+"/workflow/k8scluster-delnode.yml") //新建inventory配置文件
    checkErr(err)
    defer delnode_file.Close() //main函数结束前， 关闭文件
    delnode_file.WriteString("- remote_user: root\n  hosts: delnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delnode\n")
}

func rebuildmasterYML(softdir string) {
    rebuildmaster_file, err := os.Create(softdir+"/workflow/k8scluster-rebuildmaster.yml") //新建inventory配置文件
    checkErr(err)
    defer rebuildmaster_file.Close() //main函数结束前， 关闭文件
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/genfile\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/all\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/copycfssl\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/2.etcd\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/nginx\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/kubectl\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/apiserver\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/keepalived\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/controller-manager\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/scheduler\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/admintoken\n")
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/alter\n")
}

func delmasterYML(softdir string) {
    delmaster_file, err := os.Create(softdir+"/workflow/k8scluster-delmaster.yml") //新建inventory配置文件
    checkErr(err)
    defer delmaster_file.Close() //main函数结束前， 关闭文件
    delmaster_file.WriteString("- remote_user: root\n  hosts: delmaster\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delmaster\n")
}

func generalConfig(master_array []string, node_array []string, mvip string, currentdir string, softdir string) {
    //生成通用配置
    inventory_file, err := os.Create(currentdir+"/workflow/general.inventory") //新建inventory配置文件
    checkErr(err)
    defer inventory_file.Close() //main函数结束前， 关闭文件

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
    defer inventory_file.Close() //main函数结束前， 关闭文件

    //写入文件时，使用带缓存的 *Writer
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
    for i := 0; i < master_num; i++ {
      write.WriteString(master_array[i]+" ip="+master_array[i]+" priority="+strconv.Itoa(j)+"\n")
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

    //Flush将缓存的文件真正写入到文件中
    write.Flush()

}

func addnodeConfig(addnode_array []string, softdir string) {
    //生成addnode配置
    inventory_file, err := os.OpenFile(softdir+"/workflow/addnode.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    checkErr(err)
    defer inventory_file.Close() //main函数结束前， 关闭文件

    //写入文件时，使用带缓存的 *Writer
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------新增的k8s-node主机列表------------------------------###\n")
    write.WriteString("\n[addnode]\n")
    for i := 0; i < len(addnode_array); i++ {
      checkIP(addnode_array[i])
      write.WriteString(addnode_array[i]+" ip="+addnode_array[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"addnode\n\n\n")
    //Flush将缓存的文件真正写入到文件中
    write.Flush()

}

func delnodeConfig(delnode_array []string, softdir string) {
    //生成delnode配置
    inventory_file, err := os.OpenFile(softdir+"/workflow/delnode.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    checkErr(err)
    defer inventory_file.Close() //main函数结束前， 关闭文件

    //写入文件时，使用带缓存的 *Writer
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------被删除的k8s-node主机列表------------------------------###\n")
    write.WriteString("\n[delnode]\n")
    for i := 0; i < len(delnode_array); i++ {
      checkIP(delnode_array[i])
      write.WriteString(delnode_array[i]+" ip="+delnode_array[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"delnode\n\n\n")
    //Flush将缓存的文件真正写入到文件中
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
    defer inventory_file.Close() //main函数结束前， 关闭文件

    //写入文件时，使用带缓存的 *Writer
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

    //Flush将缓存的文件真正写入到文件中
    write.Flush()

}

func delmasterConfig(delmaster_array []string, softdir string) {
    //生成delmaster配置
    inventory_file, err := os.OpenFile(softdir+"/workflow/delmaster.inventory",os.O_WRONLY | os.O_APPEND, 0666)
    checkErr(err)
    defer inventory_file.Close() //main函数结束前， 关闭文件

    //写入文件时，使用带缓存的 *Writer
    write := bufio.NewWriter(inventory_file)
    write.WriteString("###---------------------------------被删除的k8s-master主机列表------------------------------###\n")
    write.WriteString("\n[delmaster]\n")
    for i := 0; i < len(delmaster_array); i++ {
      checkIP(delmaster_array[i])
      write.WriteString(delmaster_array[i]+" ip="+delmaster_array[i]+"\n")
    }
    write.WriteString("\n[k8s:children]\n"+"master1\n"+"delmaster\n\n\n")
    //Flush将缓存的文件真正写入到文件中
    write.Flush()

}
 
func shellAsynclog(reader io.ReadCloser) error {
    cache := "" //缓存不足一行的日志信息
    buf := make([]byte, 2048)
    for {
        num, err := reader.Read(buf)
        if err != nil && err!=io.EOF{
            return err
        }
        if num > 0 {
            b := buf[:num]
            s := strings.Split(string(b), "\n")
            line := strings.Join(s[:len(s)-1], "\n") //取出整行的日志
            fmt.Printf("%s%s\n", cache, line)
            cache = s[len(s)-1]
        }
    }
    return nil
}
 
func shellExecute(shellfile string) error {
    cmd := exec.Command("sh", "-c", shellfile)
    stdout, _ := cmd.StdoutPipe()
    stderr, _ := cmd.StderrPipe()
    if err := cmd.Start(); err != nil {
        log.Printf("Error starting command: %s......", err.Error())
        return err
    }
    go shellAsynclog(stdout)
    go shellAsynclog(stderr)
    if err := cmd.Wait(); err != nil {
        log.Printf("Error waiting for command execution: %s......", err.Error())
        return err
    }
    return nil
}

func shellOutput(strCommand string)(string){
    cmd := exec.Command("/bin/bash", "-c", strCommand) 
    stdout, _ := cmd.StdoutPipe()
    if err := cmd.Start(); err != nil{
        fmt.Println("Execute failed when Start:" + err.Error())
        return ""
    }
    out_bytes, _ := ioutil.ReadAll(stdout)
    stdout.Close()
    if err := cmd.Wait(); err != nil {
        fmt.Println("Execute failed when Wait:" + err.Error())
        return ""
    }
    return string(out_bytes)
}









