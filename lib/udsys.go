package kilib

import (
//    "fmt"
    "os"
)


// Adaptively update the genfile process file. 
func InstallGenfile(mode string, currentDir string, subProcessDir string, logName string) {
    genfile_file, err := os.Create(currentDir+"/data/output"+subProcessDir+"/sys/0x0000000000base/genfile/tasks/main.yml") 
    CheckErr(err,currentDir,logName,mode)
    defer genfile_file.Close() 
    genfile_file.WriteString("- name: 1.Create {{software_home}} directory\n  file:\n    path: \"{{software_home}}\"\n    state: directory\n- file:\n    path: \"/tmp/.kubeinstalltemp/data"+subProcessDir+"/pkg/\"\n    state: directory\n- name: 2.Distributing deployment files to target host, please wait...\n  copy:\n    src: \""+currentDir+"/pkg/{{ostype}}\"\n    dest: \"/tmp/.kubeinstalltemp/data"+subProcessDir+"/pkg/\"\n- copy:\n    src: \""+currentDir+"/sys\"\n    dest: \"/tmp/.kubeinstalltemp/data"+subProcessDir+"/\"\n")

}

// Adaptively update the IpvsYaml process file.
func InstallIpvsYaml(mode string, currentDir string, masterArray []string, subProcessDir string, logName string) {
    var ipvsinitYaml string
    master_num := len(masterArray)
    for i := 0; i < master_num; i++ {
        ipvsinitYaml = ipvsinitYaml+"  - ip: "+masterArray[i]+" \n"
    }
    ipvsinitYamlFile, err := os.Create(currentDir+"/data/output"+subProcessDir+"/sys/0x000certificate/copycfssl/templates/ipvsinit_ep.yaml") 
    CheckErr(err,currentDir,logName,mode)
    defer ipvsinitYamlFile.Close()
    ipvsinitYamlFile.WriteString("apiVersion: v1\nkind: Endpoints\nmetadata:\n  name: ipvsinit-lb\n  namespace: kube-system\n  labels:\n    k8sapp: ipvsinit-lb\nsubsets:\n- addresses:\n"+ipvsinitYaml+"  ports:\n  - name: k8s-api\n    port: 6443\n    protocol: TCP\n")
}



