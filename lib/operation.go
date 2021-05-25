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
    if err != nil {
        fmt.Println("\nWarning: There may be some problems in initialization, some hosts' SSH service or network is unreachable! \nThis may cause the installation process to fail, Please check! \n")
    }
}

func Operation(opt string, currentDir string, layoutYml string) {
    err := ShellExecute("ansible-playbook -i "+currentDir+"/config/"+opt+".inventory "+currentDir+"/config/k8scluster-"+layoutYml+".yml")
    CheckErr(err)
    fmt.Println(opt+" operation executed successfully! \n")
}

