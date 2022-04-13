package kilib

import (
//    "fmt"
    "os"
)



// Adaptively update the genfile process file. 
func InstallGenfile(osType string, mode string, currentDir string, subProcessDir string, logName string) {
    var copyPkgStr string
    if osType != "" {
        copyPkgStr = "- copy:\n    src: \""+currentDir+"/pkg/"+osType+"\"\n    dest: \"/tmp/.kubeinstalltemp/data"+subProcessDir+"/pkg/\"\n"
    }
    genfile_file, err := os.Create(currentDir+"/data/output"+subProcessDir+"/sys/0x0000000000base/genfile/tasks/main.yml") 
    CheckErr(err,currentDir,logName,mode)
    defer genfile_file.Close() 
    genfile_file.WriteString("- name: 0.Distributing deployment files to target host, please wait...\n  file:\n    path: \"{{software_home}}\"\n    state: directory\n- file:\n    path: \"/tmp/.kubeinstalltemp/data"+subProcessDir+"/pkg/\"\n    state: directory\n"+copyPkgStr+"- copy:\n    src: \""+currentDir+"/sys\"\n    dest: \"/tmp/.kubeinstalltemp/data"+subProcessDir+"/\"\n")

}

// Adaptively update the IpvsYaml process file.
func InstallIpvsYaml(mode string, currentDir string, masterArray []string, kubeApiPort string, subProcessDir string, logName string) {
    var ipvsinitYaml string
    master_num := len(masterArray)
    for i := 0; i < master_num; i++ {
        ipvsinitYaml = ipvsinitYaml+"  - ip: "+masterArray[i]+" \n"
    }
    ipvsinitYamlFile, err := os.Create(currentDir+"/data/output"+subProcessDir+"/sys/0x000certificate/copycfssl/templates/ipvsinit_ep.yaml.j2") 
    CheckErr(err,currentDir,logName,mode)
    defer ipvsinitYamlFile.Close()
    ipvsinitYamlFile.WriteString("apiVersion: v1\nkind: Endpoints\nmetadata:\n  name: ipvsinit-lb\n  namespace: kube-system\n  labels:\n    k8sapp: ipvsinit-lb\nsubsets:\n- addresses:\n"+ipvsinitYaml+"  ports:\n  - name: k8s-api\n    port: "+kubeApiPort+"\n    protocol: TCP\n")
}

// Adaptively update the preinstall shell.
func InstallPreShell(osType string, mode string, currentDir string, subProcessDir string, logName string) {
    var preInstallShell string
    if osType == "ubuntu20" {
        preInstallShell = "dpkg --force-depends -i /tmp/.kubeinstalltemp/data{{sub_process_dir}}/pkg/ubuntu20/basepkg-ubuntu20/*.deb ; \n\n exit 0 \n\n "
    } else {
        preInstallShell = "setenforce 0 >/dev/null 2>&1 ; \n\n rpm -U /tmp/.kubeinstalltemp/data{{sub_process_dir}}/pkg/"+osType+"/basepkg-"+osType+"/*.rpm --force --nodeps ; \n\n exit 0 \n\n "
    }
    InstallPreShellFile, err := os.Create(currentDir+"/data/output"+subProcessDir+"/sys/0x0000000000base/all/templates/preinstall.sh.j2")
    CheckErr(err,currentDir,logName,mode)
    defer InstallPreShellFile.Close()
    InstallPreShellFile.WriteString("#!/bin/bash \n \n " + preInstallShell)
}



