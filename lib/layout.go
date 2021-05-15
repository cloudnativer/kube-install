package kilib

import (
    "os"
)


func InstallYML(softdir string, ostype string) {
    osCompatibilityLayout := ""
    if (ostype == "centos7") || (ostype == "rhel7") {
        osCompatibilityLayout = "- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/kernel\n"
    }
    install_file, err := os.Create(softdir+"/workflow/k8scluster-install.yml")
    CheckErr(err)
    defer install_file.Close()
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/genfile\n"+osCompatibilityLayout+"- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/all\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0008.action/delnode\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/docker\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0002.cfssl/copycfssl\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0002.cfssl/createssl\n- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0003.etcd\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0004.network/etcd_network\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0004.network/flanneld\n- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/kubectl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/apiserver\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/api-rbac\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/controller-manager\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/scheduler\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0006.kube-node/kubelet\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0006.kube-node/approve-csr\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0006.kube-node/kube-proxy\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0007.kube-addons\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/admintoken\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0008.action/pushimages\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0009.finish/install\n")
}

func OnemasterinstallYML(softdir string, ostype string) {
    osCompatibilityLayout := ""
    if (ostype == "centos7") || (ostype == "rhel7") {
        osCompatibilityLayout = "- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/kernel\n"
    }
    onemasterinstall_file, err := os.Create(softdir+"/workflow/k8scluster-onemasterinstall.yml") 
    CheckErr(err)
    defer onemasterinstall_file.Close() 
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/genfile\n"+osCompatibilityLayout+"- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/all\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0008.action/delnode\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/docker\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0002.cfssl/copycfssl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0002.cfssl/createssl\n- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0003.etcd\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0004.network/etcd_network\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0004.network/flanneld\n- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/kubectl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/apiserver\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/api-rbac\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/controller-manager\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/scheduler\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0006.kube-node/kubelet\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0006.kube-node/approve-csr\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0006.kube-node/kube-proxy\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0007.kube-addons\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/admintoken\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0008.action/pushimages\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0009.finish/install\n")
}

func AddnodeYML(softdir string, ostype string) {
    osCompatibilityLayout := ""
    if (ostype == "centos7") || (ostype == "rhel7") {
        osCompatibilityLayout = "- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/kernel\n"
    }
    addnode_file, err := os.Create(softdir+"/workflow/k8scluster-addnode.yml") 
    CheckErr(err)
    defer addnode_file.Close() 
    addnode_file.WriteString(osCompatibilityLayout+"- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/all\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0008.action/delnode\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/docker\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0004.network/flanneld\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0006.kube-node/kubelet\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0006.kube-node/kube-proxy\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0006.kube-node/approve-csr\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0009.finish/addnode\n")
}

func DelnodeYML(softdir string) {
    delnode_file, err := os.Create(softdir+"/workflow/k8scluster-delnode.yml") 
    CheckErr(err)
    defer delnode_file.Close() 
    delnode_file.WriteString("- remote_user: root\n  hosts: delnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0008.action/delnode\n")
}

func RebuildmasterYML(softdir string) {
    rebuildmaster_file, err := os.Create(softdir+"/workflow/k8scluster-rebuildmaster.yml") 
    CheckErr(err)
    defer rebuildmaster_file.Close() 
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/genfile\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0001.base/all\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0002.cfssl/copycfssl\n- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0003.etcd\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/kubectl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/apiserver\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/controller-manager\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/scheduler\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0005.kube-master/admintoken\n")
}

func DelmasterYML(softdir string) {
    delmaster_file, err := os.Create(softdir+"/workflow/k8scluster-delmaster.yml") 
    CheckErr(err)
    defer delmaster_file.Close() 
    delmaster_file.WriteString("- remote_user: root\n  hosts: delmaster\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0008.action/delmaster\n")
}

func UninstallYML(softdir string) {
    uninstall_file, err := os.Create("/tmp/workflow/k8scluster-uninstall.yml")
    CheckErr(err)
    defer uninstall_file.Close()
    uninstall_file.WriteString("\n- remote_user: root\n  hosts: delnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0008.action/delnode\n- remote_user: root\n  hosts: delmaster\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0008.action/delmaster\n- remote_user: root\n  hosts: delmaster,delnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0009.finish/uninstall\n")
}



