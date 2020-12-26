package kilib

import (
    "os"
)


func InstallGenfile(currentdir string) {
    genfile_file, err := os.Create(currentdir+"/sys/0.base/genfile/tasks/main.yml") 
    CheckErr(err)
    defer genfile_file.Close() 
    genfile_file.WriteString("- name: 1.Create {{software_home}} directory\n  file:\n    path: \"{{software_home}}\"\n    state: directory\n- name: 2.Distributing deployment files to kubernetes master, please wait...\n  copy:\n    src: \""+currentdir+"/{{item}}\"\n    dest: \"{{software_home}}/\"\n  with_items:\n    - sys\n    - docs\n    - pkg\n    - workflow\n    - yaml\n    - kube-install\n- copy:\n    src: \""+currentdir+"/kube-install\"\n    dest: \"/usr/local/bin/kube-install\"\n    mode: 0755\n- name: 3.Configure permissions for executables\n  file: path={{software_home}}/{{ item }} mode=755 owner=root group=root\n  with_items:\n    - workflow/sshkey-init.sh\n    - workflow/sshops-init.sh\n    - workflow/getmasterconfig.sh\n- name: 4.Install distributed control software\n  shell: \"{{software_home}}/workflow/sshops-init.sh {{software_home}} {{software_home}}\"\n  ignore_errors: yes\n")

}





