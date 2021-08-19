package kilib

import (
    "fmt"
    "os"
    "os/exec"
    "io"
    "io/ioutil"
    "log"
    "strings"
)



func CopyFile(src string, dst string) (int64, error) {
        sourceFileStat, err := os.Stat(src)
        if err != nil {
                return 0, err
        }

        if !sourceFileStat.Mode().IsRegular() {
                return 0, fmt.Errorf("%s is not a regular file", src)
        }

        source, err := os.Open(src)
        if err != nil {
                return 0, err
        }
        defer source.Close()

        destination, err := os.Create(dst)
        if err != nil {
                return 0, err
        }
        defer destination.Close()
        nBytes, err := io.Copy(destination, source)
        return nBytes, err
}

func CreateDir(dir string, currentDir string, logName string, mode string) {
     _,err := os.Stat(dir)
     if err != nil {
         err_mk := os.MkdirAll(dir , 0755)
         CheckErr(err_mk, currentDir, logName, mode)
     }
}

func CreateFile(filePth string, currentDir string, logName string, mode string) {
     _,err := os.Stat(filePth)
     if err != nil {
         if os.IsNotExist(err) {
             _,err_ct := os.Create(filePth)
             CheckErr(err_ct, currentDir, logName, mode)
         }
     }
}

func GetAllDir(pathname string, currentDir string, logName string, mode string) ([]string, error) {
    var s []string
    rd, err := ioutil.ReadDir(pathname)
    if err != nil {
        CheckErr(err,currentDir,logName,mode)
    }
    for _, fi := range rd {
        if fi.IsDir(){
            s = append(s, fi.Name())
        }
    }
    return s,err
}

func ChangeLang(langFromWeb string, currentDir string, logName string, mode string) string {
    var Lang string
    if langFromWeb != "" {
        Lang = langFromWeb
        DatabaseUpdate(currentDir+"/data/config/language.txt", Lang, currentDir, logName, mode)
    } else {
        langFromData,_ := ReadFile(currentDir+"/data/config/language.txt")
        if langFromData != "" {
            Lang = langFromData
        } else {
            Lang = "en"
        }
    }
    return Lang
}

func ProgressBar(n int, char string) (s string) {
    for i:=1;i<=n;i++{
        s+=char
    }
    return
}

func ReadFile(filePth string) (string, error) {
    file, err := os.Open(filePth)
    if err != nil{
        return "",err
    }
    defer file.Close()
    fileContent, err := ioutil.ReadAll(file)
    if err == nil{
        return string(fileContent),nil
    }else{
        return "",err
    }
}

func ReadFileAsArray(filePath string) ([]string, error) {
    result := []string{}
    b, err := ioutil.ReadFile(filePath)
    if err != nil {
        return result, err
    }
    s := string(b)
    for _, lineStr := range strings.Split(s, "\n") {
        lineStr = strings.TrimSpace(lineStr)
        if lineStr == "" {
            continue
        }
        result = append(result, lineStr)
    }
    return result, nil
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
    cmd := exec.Command("sh", "-c", strCommand) 
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

func StrInArray(value string, list []string) bool {
    for _, v := range list {
        if v == value {
            return true
        }
    }
    return false
}

func LogStr(mode string) string {
    var logStr string
    if mode == "DAEMON" {
        logStr = " >> "
    } else {
        logStr = " | tee -a "
    }
    return logStr
}


