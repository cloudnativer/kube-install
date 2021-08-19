package kilib

import (
//    "fmt"
)


func SshKey(ipArray []string, sshPass string, currentDir string) error {
    var err error
    idrsa,err_idrsa := ReadFile("/root/.ssh/id_rsa")
    if err_idrsa != nil || idrsa == "" {
        ShellExecute("ssh-keygen -t rsa -P \"\" -f /root/.ssh/id_rsa >/dev/null 2>&1")
    }
    if len(ipArray) == 1 && sshPass == "" {
        err = ShellExecute("ssh-copy-id -p 22 root@"+ipArray[0]+" >/dev/null 2>&1")
    } else {
        for _, ip := range ipArray {
            err = ShellExecute("sshpass -p \""+sshPass+"\" ssh-copy-id -p 22 root@"+ip+" >/dev/null 2>&1")
            if err != nil {
                break
            }
        }
    }
    return err
}

func SshOpsInit(currentDir string, osType string, mode string) error {
    err := ShellExecute(currentDir+"/proc/sshops-init.sh \""+currentDir+"\" \""+osType+"\"")
    return err
}



