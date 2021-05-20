package kilib

import (
    "fmt"
    "os"
    "os/exec"
    "bufio"
    "io"
    "io/ioutil"
    "log"
    "strings"
)



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

func ShellOutput(strCommand string) (string, error) {
    cmd := exec.Command("/bin/bash", "-c", strCommand) 
    stdout, _ := cmd.StdoutPipe()
    if err := cmd.Start(); err != nil{
        return "",err
    }
    out_bytes, _ := ioutil.ReadAll(stdout)
    stdout.Close()
    if err := cmd.Wait(); err != nil {
        return "",err
    }
    return string(out_bytes),nil
}

func ProgressBar(n int, char string) (s string) {
    for i:=1;i<=n;i++{
        s+=char
    }
    return
}


