
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


