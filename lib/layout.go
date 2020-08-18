package kilib

import (
    "os"
)


func InstallYML(softdir string) {
    install_file, err := os.Create(softdir+"/workflow/k8scluster-install.yml")
    CheckErr(err)
    defer install_file.Close()
    install_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/genfile\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/kernel\n- remote_user: root\n  hosts: master,node,nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/all\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delnode\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/docker\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/copycfssl\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/createssl\n- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/2.etcd\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/etcd_network\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/flanneld\n- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/nginx\n- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/kubectl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/apiserver\n- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/keepalived\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/api-rbac\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/controller-manager\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/scheduler\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kubelet\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/approve-csr\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kube-proxy\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/7.kube-addons\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/admintoken\n- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/alter\n- remote_user: "+ currentuser +"\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/pushimages\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/9.finish/install\n")
}

func OnemasterinstallYML(softdir string) {
    onemasterinstall_file, err := os.Create(softdir+"/workflow/k8scluster-onemasterinstall.yml") 
    CheckErr(err)
    defer onemasterinstall_file.Close() 
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/genfile\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/kernel\n- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/all\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delnode\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/docker\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/copycfssl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/createssl\n- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/2.etcd\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/etcd_network\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/flanneld\n- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/kubectl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/apiserver\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/api-rbac\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/controller-manager\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/scheduler\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kubelet\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/approve-csr\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kube-proxy\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/7.kube-addons\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/admintoken\n- remote_user: "+ currentuser +"\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/pushimages\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/9.finish/install\n")
}

func AddnodeYML(softdir string) {
    addnode_file, err := os.Create(softdir+"/workflow/k8scluster-addnode.yml") 
    CheckErr(err)
    defer addnode_file.Close() 
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/kernel\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/all\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delnode\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/docker\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/3.network/flanneld\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kubelet\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/kube-proxy\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/6.kube-node/approve-csr\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/9.finish/addnode\n")
}

func DelnodeYML(softdir string) {
    delnode_file, err := os.Create(softdir+"/workflow/k8scluster-delnode.yml") 
    CheckErr(err)
    defer delnode_file.Close() 
    delnode_file.WriteString("- remote_user: root\n  hosts: delnode\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delnode\n")
}

func RebuildmasterYML(softdir string) {
    rebuildmaster_file, err := os.Create(softdir+"/workflow/k8scluster-rebuildmaster.yml") 
    CheckErr(err)
    defer rebuildmaster_file.Close() 
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/genfile\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/0.base/all\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/1.cfssl/copycfssl\n- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/2.etcd\n- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/nginx\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/kubectl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/apiserver\n- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/keepalived\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/controller-manager\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/scheduler\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/5.kube-master/admintoken\n- remote_user: root\n  hosts: nginx\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/4.kube-nginx/alter\n")
}

func DelmasterYML(softdir string) {
    delmaster_file, err := os.Create(softdir+"/workflow/k8scluster-delmaster.yml") 
    CheckErr(err)
    defer delmaster_file.Close() 
    delmaster_file.WriteString("- remote_user: root\n  hosts: delmaster\n  gather_facts: no\n  roles:\n    - "+softdir+"/bin/8.action/delmaster\n")
}

