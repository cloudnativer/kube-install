package kilib

import (
    "fmt"
    "net"
    "os"
    "os/exec"
    "bufio"
    "io"
    "io/ioutil"
    "log"
    "strings"
)


func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}

func CheckIP(ipv4 string) {
    address := net.ParseIP(ipv4)  
    if address == nil {
         panic("The format of IP address you entered is wrong, please check!")
    }
}

func CheckParam(option string, param string) {
    if param == "" {
         panic("When performing the "+option+" operation, you must enter the "+param+" parameter, please check!")
    }
}

func ProgressBar(n int,char string) (s string) {
    for i:=1;i<=n;i++{
        s+=char
    }
    return
}

func CopyFile(srcFileName string, dstFileName string) (written int64, err error) {
    //Functions for copying files
    srcFile, err := os.Open(srcFileName)
    if err != nil {
            fmt.Printf("open file err = %v\n", err)
            return
    }
    defer srcFile.Close()
    reader := bufio.NewReader(srcFile)

    dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
    if err != nil {
            fmt.Printf("open file err = %v\n", err)
            return
    }
    writer := bufio.NewWriter(dstFile)
    defer func() {
            writer.Flush()
            dstFile.Close()
    }()

    return io.Copy(writer, reader)

}


 
func ShellAsynclog(reader io.ReadCloser) error {
    cache := ""
    buf := make([]byte, 2048)
    for {
        num, err := reader.Read(buf)
        if err != nil && err!=io.EOF{
            return err
        }
        if num > 0 {
            b := buf[:num]
            s := strings.Split(string(b), "\n")
            line := strings.Join(s[:len(s)-1], "\n")
            fmt.Printf("%s%s\n", cache, line)
            cache = s[len(s)-1]
        }
    }
    return nil
}
 
func ShellExecute(shellfile string) error {
    cmd := exec.Command("sh", "-c", shellfile)
    stdout, _ := cmd.StdoutPipe()
    stderr, _ := cmd.StderrPipe()
    if err := cmd.Start(); err != nil {
        log.Printf("Error starting command: %s......", err.Error())
        return err
    }
    go ShellAsynclog(stdout)
    go ShellAsynclog(stderr)
    if err := cmd.Wait(); err != nil {
        log.Printf("Error waiting for command execution: %s......", err.Error())
        return err
    }
    return nil
}

func ShellOutput(strCommand string)(string){
    cmd := exec.Command("/bin/bash", "-c", strCommand) 
    stdout, _ := cmd.StdoutPipe()
    if err := cmd.Start(); err != nil{
        fmt.Println("Execute failed when Start:" + err.Error())
        return ""
    }
    out_bytes, _ := ioutil.ReadAll(stdout)
    stdout.Close()
    if err := cmd.Wait(); err != nil {
        fmt.Println("Execute failed when Wait:" + err.Error())
        return ""
    }
    return string(out_bytes)
}

func ShowHelp(){
    fmt.Println("Version 0.2.0 (Creation Date: 04/21/2020)")
    fmt.Println("Usage of kube-install: -opt [OPTIONS] COMMAND [ARGS]...\n")
    fmt.Println("Options:")
    fmt.Println("  init             Initialize the system environment.")
    fmt.Println("  install          Install k8s cluster.")
    fmt.Println("  delnode          Remove the k8s-node from the cluster.")
    fmt.Println("  addnode          Add k8s-node to the cluster.")
    fmt.Println("  delmaster        Remove the k8s-master from the cluster.")
    fmt.Println("  rebuildmaster    Rebuild the damaged k8s-master.")
    fmt.Println("  help             Display help information.\n")
    fmt.Println("Commands:")
    fmt.Println("  master           The IP address of k8s-master server.")
    fmt.Println("  mvip             K8s-master cluster virtual IP address.")
    fmt.Println("  node             The IP address of k8s-node server.")
    fmt.Println("  sshpwd           SSH login root password of each server.\n")
    fmt.Println("For example:")
    fmt.Println("  Initialize the system environment:")
    fmt.Println("    kube-install -opt init")
    fmt.Println("  Install k8s cluster:")
    fmt.Println("    kube-install -opt install -master \"192.168.1.11,192.168.1.12,192.168.1.13\" -node \"192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14\" -mvip \"192.168.1.88\" -sshpwd \"cloudnativer\"")
    fmt.Println("  Add k8s-node to the cluster:")
    fmt.Println("    kube-install -opt addnode -node \"192.168.1.15,192.168.1.16\" -sshpwd \"cloudnativer\"")
    fmt.Println("  Remove the k8s-node from the cluster:")
    fmt.Println("    kube-install -opt delnode -node \"192.168.1.13,192.168.1.15\" -sshpwd \"cloudnativer\"")
    fmt.Println("  Remove the k8s-master from the cluster:")
    fmt.Println("    kube-install -opt delmaster -master \"192.168.1.13\" -sshpwd \"cloudnativer\"")
    fmt.Println("  Rebuild the damaged k8s-master:")
    fmt.Println("    kube-install -opt rebuildmaster -master \"192.168.1.13\" -sshpwd \"cloudnativer\"")
    fmt.Println("  Display help information:")
    fmt.Println("    kube-install -opt help\n")
}


