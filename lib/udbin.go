package kilib

import (
    "os"
)


func InstallGenFile(currentdir string) {
    genfile_file, err := os.Create(currentdir+"/sys/0x0000000000base/genfile/tasks/main.yml") 
    CheckErr(err)
    defer genfile_file.Close() 
    genfile_file.WriteString("- name: 1.Create {{software_home}} directory\n  file:\n    path: \"{{software_home}}\"\n    state: directory\n- name: 2.Distributing deployment files to kubernetes master, please wait...\n  copy:\n    src: \""+currentdir+"/{{item}}\"\n    dest: \"{{software_home}}/\"\n  with_items:\n    - sys\n    - docs\n    - pkg\n    - proc\n    - config\n    - yaml\n    - kube-install\n- copy:\n    src: \""+currentdir+"/kube-install\"\n    dest: \"/usr/local/bin/kube-install\"\n    mode: 0755\n- name: 3.Configure permissions for executables\n  file: path={{software_home}}/{{ item }} mode=755 owner=root group=root\n  with_items:\n    - proc/sshkey-init.sh\n    - proc/sshops-init.sh\n    - proc/getmasterconfig.sh\n- name: 4.Install distributed control software\n  shell: \"{{software_home}}/proc/sshops-init.sh {{software_home}} {{software_home}}\"\n  ignore_errors: yes\n")
}


func InstallIpvsYaml(currentdir string, master_array []string) {
    var ipvsinit_yaml string
    master_num := len(master_array)
    for i := 0; i < master_num; i++ {
        ipvsinit_yaml = ipvsinit_yaml+"  - ip: "+master_array[i]+" \n"
    }
    ipvsinityaml_file, err := os.Create(currentdir+"/sys/0x000certificate/copycfssl/templates/ipvsinit_ep.yaml") 
    CheckErr(err)
    defer ipvsinityaml_file.Close()
    ipvsinityaml_file.WriteString("apiVersion: v1\nkind: Endpoints\nmetadata:\n  name: ipvsinit-lb\n  namespace: kube-system\n  labels:\n    k8sapp: ipvsinit-lb\nsubsets:\n- addresses:\n"+ipvsinit_yaml+"  ports:\n  - name: k8s-api\n    port: 6443\n    protocol: TCP\n")
}




