package kilib

import (
    "os"
)


func InstallYML(softdir string, ostype string) {
    osCompatibilityLayout := ""
    if (ostype == "centos7") || (ostype == "rhel7") {
        osCompatibilityLayout = "- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/kernel\n"
    }
    install_file, err := os.Create(softdir+"/config/k8scluster-install.yml")
    CheckErr(err)
    defer install_file.Close()
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/genfile\n"+osCompatibilityLayout+"- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/all\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000action/delnode\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/docker\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x000certificate/copycfssl\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x000certificate/createssl\n- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000storage\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000network/flanneld\n- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/kubectl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/apiserver\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/api-rbac\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/controller-manager\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/scheduler\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000node/kubelet\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000node/approve-csr\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000node/kube-proxy\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000addons\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/admintoken\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000action/pushimages\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000finish/install\n")
}

func OnemasterinstallYML(softdir string, ostype string) {
    osCompatibilityLayout := ""
    if (ostype == "centos7") || (ostype == "rhel7") {
        osCompatibilityLayout = "- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/kernel\n"
    }
    onemasterinstall_file, err := os.Create(softdir+"/config/k8scluster-onemasterinstall.yml") 
    CheckErr(err)
    defer onemasterinstall_file.Close() 
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/genfile\n"+osCompatibilityLayout+"- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/all\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000action/delnode\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/docker\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x000certificate/copycfssl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x000certificate/createssl\n- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000storage\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000network/flanneld\n- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/kubectl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/apiserver\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/api-rbac\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/controller-manager\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/scheduler\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000node/kubelet\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000node/approve-csr\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000node/kube-proxy\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000addons\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/admintoken\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000action/pushimages\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000finish/install\n")
}

func AddnodeYML(softdir string, ostype string) {
    osCompatibilityLayout := ""
    if (ostype == "centos7") || (ostype == "rhel7") {
        osCompatibilityLayout = "- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/kernel\n"
    }
    addnode_file, err := os.Create(softdir+"/config/k8scluster-addnode.yml") 
    CheckErr(err)
    defer addnode_file.Close() 
    addnode_file.WriteString(osCompatibilityLayout+"- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/all\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000action/delnode\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/docker\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000network/flanneld\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000node/kubelet\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000node/kube-proxy\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000node/approve-csr\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000finish/addnode\n")
}

func DelnodeYML(softdir string) {
    delnode_file, err := os.Create(softdir+"/config/k8scluster-delnode.yml") 
    CheckErr(err)
    defer delnode_file.Close() 
    delnode_file.WriteString("- remote_user: root\n  hosts: delnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000action/delnode\n")
}

func RebuildmasterYML(softdir string) {
    rebuildmaster_file, err := os.Create(softdir+"/config/k8scluster-rebuildmaster.yml") 
    CheckErr(err)
    defer rebuildmaster_file.Close() 
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/genfile\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000000base/all\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x000certificate/copycfssl\n- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x0000000storage\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/kubectl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/apiserver\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/controller-manager\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/scheduler\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000master/admintoken\n")
}

func DelmasterYML(softdir string) {
    delmaster_file, err := os.Create(softdir+"/config/k8scluster-delmaster.yml") 
    CheckErr(err)
    defer delmaster_file.Close() 
    delmaster_file.WriteString("- remote_user: root\n  hosts: delmaster\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000action/delmaster\n")
}

func UninstallYML(softdir string) {
    uninstall_file, err := os.Create("/tmp/config/k8scluster-uninstall.yml")
    CheckErr(err)
    defer uninstall_file.Close()
    uninstall_file.WriteString("\n- remote_user: root\n  hosts: delnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000action/delnode\n- remote_user: root\n  hosts: delmaster\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000action/delmaster\n- remote_user: root\n  hosts: delmaster,delnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/sys/0x00000000finish/uninstall\n")
}



