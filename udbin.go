package main

import (
    "os"
)


func installGenfile(currentdir string) {
    genfile_file, err := os.Create(currentdir+"/bin/0.base/genfile/tasks/main.yml") //新建inventory配置文件
    checkErr(err)
    defer genfile_file.Close() //main函数结束前， 关闭文件
    genfile_file.WriteString("- name: 1.创建{{software_home}}目录\n  file:\n    path: \"{{software_home}}\"\n    state: directory\n")
    genfile_file.WriteString("- name: 2.正在将部署文件分发到k8s-master\n  copy:\n    src: \""+currentdir+"/{{item}}\"\n    dest: \"{{software_home}}/\"\n  with_items:\n    - bin\n    - docs\n    - pkg\n    - workflow\n    - yaml\n    - kube-install\n- copy:\n    src: \""+currentdir+"/kube-install\"\n    dest: \"/usr/local/bin/kube-install\"\n    mode: 0755\n")
    genfile_file.WriteString("- name: 3.配置可执行文件的权限\n  file: path={{software_home}}/{{ item }} mode=755 owner=root group=root\n  with_items:\n    - workflow/sshkey-init.sh\n    - workflow/sshops-init.sh\n    - workflow/getmasterconfig.sh\n")

}





