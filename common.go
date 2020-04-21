
package main

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


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func checkIP(ipv4 string) {
    address := net.ParseIP(ipv4)  
    if address == nil {
         panic("您输入的IP地址格式有误，请检查！")
    }
}

func checkParam(option string, param string) {
    if param == "" {
         panic("执行"+option+"操作时，必须输入"+param+"参数，请检查！")
    }
}

func progressBar(n int,char string) (s string) {
    for i:=1;i<=n;i++{
        s+=char
    }
    return
}

func copyFile(srcFileName string, dstFileName string) (written int64, err error) {
    //用于拷贝文件的函数，接收两个文件路径 srcFileName dstFileName
    srcFile, err := os.Open(srcFileName)
    if err != nil {
            fmt.Printf("open file err = %v\n", err)
            return
    }
    defer srcFile.Close()

    //通过srcFile，获取到Reader
    reader := bufio.NewReader(srcFile)

    //打开dstFileName
    dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
    if err != nil {
            fmt.Printf("open file err = %v\n", err)
            return
    }

    writer := bufio.NewWriter(dstFile)
    defer func() {
            writer.Flush() //把缓冲区的内容写入到文件
            dstFile.Close()
    }()

    return io.Copy(writer, reader)

}


 
func shellAsynclog(reader io.ReadCloser) error {
    cache := "" //缓存不足一行的日志信息
    buf := make([]byte, 2048)
    for {
        num, err := reader.Read(buf)
        if err != nil && err!=io.EOF{
            return err
        }
        if num > 0 {
            b := buf[:num]
            s := strings.Split(string(b), "\n")
            line := strings.Join(s[:len(s)-1], "\n") //取出整行的日志
            fmt.Printf("%s%s\n", cache, line)
            cache = s[len(s)-1]
        }
    }
    return nil
}
 
func shellExecute(shellfile string) error {
    cmd := exec.Command("sh", "-c", shellfile)
    stdout, _ := cmd.StdoutPipe()
    stderr, _ := cmd.StderrPipe()
    if err := cmd.Start(); err != nil {
        log.Printf("Error starting command: %s......", err.Error())
        return err
    }
    go shellAsynclog(stdout)
    go shellAsynclog(stderr)
    if err := cmd.Wait(); err != nil {
        log.Printf("Error waiting for command execution: %s......", err.Error())
        return err
    }
    return nil
}

func shellOutput(strCommand string)(string){
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

func showHelp(){
    fmt.Println("Version 0.1.1")
    fmt.Println("Usage of kube-install: -opt [OPTIONS] COMMAND [ARGS]...\n")
    fmt.Println("Options: \n")
    fmt.Println("  init             Initialize the system environment.")
    fmt.Println("  install          Install k8s cluster.")
    fmt.Println("  delnode          Remove the k8s-node from the cluster.")
    fmt.Println("  addnode          Add k8s-node to the cluster.")
    fmt.Println("  delmaster        Remove the k8s-master from the cluster.")
    fmt.Println("  rebuildmaster    Rebuild the damaged k8s-master.")
    fmt.Println("  help             Display help information.\n")
    fmt.Println("Commands:\n")
    fmt.Println("  master           The IP address of k8s-master server.")
    fmt.Println("  mvip             K8s-master cluster virtual IP address.")
    fmt.Println("  node             The IP address of k8s-node server.")
    fmt.Println("  sshpwd           SSH login root password of each server.\n\n")
    fmt.Println("For example：\n")
    fmt.Println("  Initialize the system environment:")
    fmt.Println("    kube-install -opt init")
    fmt.Println("  Install k8s cluster:")
    fmt.Println("    kube-install -opt install -master \"192.168.122.11,192.168.122.12,192.168.122.13\" -node \"192.168.122.11,192.168.122.12,192.168.122.13,192.168.122.14\" -mvip \"192.168.122.100\" -sshpwd \"cloudnativer\"")
    fmt.Println("  Add k8s-node to the cluster:")
    fmt.Println("    kube-install -opt addnode -node \"192.168.122.15,192.168.122.16\" -sshpwd \"cloudnativer\"")
    fmt.Println("  Remove the k8s-node from the cluster:")
    fmt.Println("    kube-install -opt delnode -node \"192.168.122.13,192.168.122.15\" -sshpwd \"cloudnativer\"")
    fmt.Println("  Remove the k8s-master from the cluster:")
    fmt.Println("    kube-install -opt delmaster -master \"192.168.122.13\" -sshpwd \"cloudnativer\"")
    fmt.Println("  Rebuild the damaged k8s-master:")
    fmt.Println("    kube-install -opt rebuildmaster -master \"192.168.122.13\" -sshpwd \"cloudnativer\"")
    fmt.Println("  Display help information:")
    fmt.Println("    kube-install -opt help\n\n")
}


