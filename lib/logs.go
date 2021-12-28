package kilib

import (
    "fmt"
)


// Displays the log information of the kube-install.
func ShowLogs(opt string, label string, currentDir string) {
    _, _, _, subProcessDir, _ := ParameterConvert("", "", "", "", label, "")
    if opt != "" {
        logStr, _ := ReadFile(currentDir + "/data/logs" + subProcessDir + "/logs/" + opt + ".log")
        fmt.Println(logStr)
    } else {
        fmt.Println("\"-exec\" and \"-label\" parameters cannot be empty, please check! \n")
        return
    }
}


