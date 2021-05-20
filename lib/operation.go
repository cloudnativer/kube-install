package kilib

import (
    "fmt"
)


func SshOpsInit(softDir string, currentDir string, osType string) {
    err := ShellExecute(currentDir+"/proc/sshops-init.sh \""+softDir+"\" \""+currentDir+"\" \""+osType+"\"")
    CheckErr(err)
}

func SshKeyInit(sshPwd string, ip string, softDir string, currentDir string, opt string) {
    err := ShellExecute(currentDir+"/proc/sshkey-init.sh \""+sshPwd+"\" \"127.0.0.1 "+ip+"\" \""+softDir+"\" \""+currentDir+"\" \"+opt+\"")
    CheckErr(err)
}

func Operation(opt string, currentDir string) {
    err := ShellExecute("ansible-playbook -i "+currentDir+"/config/"+opt+".inventory "+currentDir+"/config/k8scluster-"+opt+".yml")
    CheckErr(err)
    fmt.Println(opt+" operation executed successfully! \n")
}

