package kilib

import (
    "os"
    "bufio"
)


// Automatically generate Systemd service startup file.
func CreateSystemdService(mode string, currentDir string, logName string) bool {
    os.RemoveAll(currentDir+"/kube-install.service")
    os.RemoveAll("/etc/systemd/system/kube-install.service")
    CreateFile(currentDir+"/kube-install.service", currentDir, logName, mode)
    CheckFileExist(currentDir+"/", "kube-install.service", currentDir, logName, mode)

    inventory_file, err := os.OpenFile(currentDir+"/kube-install.service",os.O_WRONLY | os.O_APPEND, 0666)
    CheckErr(err,currentDir,logName,mode)
    defer inventory_file.Close()
    write := bufio.NewWriter(inventory_file)
    write.WriteString("[Unit] \nDescription=kube-install One click fast installation of highly available kubernetes cluster. \nDocumentation=https://cloudnativer.github.io \nAfter=sshd.service \nRequires=sshd.service \n\n[Service] \nEnvironment=\"USER=root\" \nExecStart="+currentDir+"/kube-install -daemon \nUser=root \nPrivateTmp=true \nLimitNOFILE=65536 \nTimeoutStartSec=5 \nRestartSec=10 \nRestart=always \n\n[Install] \nWantedBy=multi-user.target \n\n")
    write.Flush()

    _,err_cp := CopyFile(currentDir+"/kube-install.service", "/etc/systemd/system/kube-install.service")
    if err_cp != nil {
        CheckErr(err_cp,currentDir,logName,mode)
    } else {
        CheckFileExist("/etc/systemd/system/", "kube-install.service", currentDir, logName, mode)
    }

    return true
}


