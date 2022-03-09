package kilib

import (
//    "fmt"
    "os"
    "time"
    "syscall"
    "sort"
    "math"
    "strings"
    "strconv"
    "github.com/shopspring/decimal"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/disk"
)


type Pair struct {
    Key string
    Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func SortMapByValue(m map[string]int) PairList {
    p := make(PairList, len(m))
    i := 0
    for k, v := range m {
        p[i] = Pair{k, v}
        i++
    }
    sort.Sort(p)
    return p
}

// Schedule task schedule in seconds.
func SecondPeriodSchedule() {
    return
}

// Schedule task schedule in minutes.
func MinutePeriodSchedule(currentDir string, kissh string, logName string, mode string) {
    for {
        var k8sNum,stuOk,stuInstall,stuUninstall,stuNotok,stuUnknow,stuAll int
        var k8sNumMap map[string]int
        k8sNumMap = make(map[string]int)
        var labelArray = make([]string,5)
        var nodeNumArray = make([]string,5)
        var labelStr,nodeNumStr,scheduleList string
        curTimeOriginal := time.Now()
        curTime := curTimeOriginal.Format("2006-01-02 15:04")
        curTimeRelative, _ := time.ParseDuration("-1h")
        relativeTime,_ := time.Parse("2006-01-02 15:04", curTimeOriginal.Add(curTimeRelative).Format("2006-01-02 15:04"))
        k8sArrayLocal,_ := GetAllDir(currentDir+"/data/output",currentDir,logName,mode)
        k8sNum = len(k8sArrayLocal)
        for i := 0; i < len(k8sArrayLocal); i++ {
            _,_,_,subProcessDir,_ := ParameterConvert(mode, "", "", "", k8sArrayLocal[i], "")
	    masterArray,_ := GetAllDir(currentDir+"/data/output"+subProcessDir+"/masters",currentDir,logName,mode)
	    nodeArray,_ := GetAllDir(currentDir+"/data/output"+subProcessDir+"/nodes",currentDir,logName,mode)
            sshPort,_ := ReadFile(currentDir+"/data/output"+subProcessDir+"/sshport.txt")
            scheduler,_ := ReadFile(currentDir+"/data/output"+subProcessDir+"/scheduler.txt")
            instime,_ := ReadFile(currentDir+"/data/output"+subProcessDir+"/installtime.txt")
            stu,err_stu := ReadFile(currentDir+"/data/output"+subProcessDir+"/status.txt")
            statusFileInfo, _ := os.Stat(currentDir+"/data/output"+subProcessDir+"/status.txt")
            statusFileStat := statusFileInfo.Sys().(*syscall.Stat_t)
            modfiyTime,_ := time.Parse("2006-01-02 15:04", time.Unix(int64(statusFileStat.Mtim.Sec), 0).Format("2006-01-02 15:04"))
            currentTime,_ := time.Parse("2006-01-02 15:04", curTime)
            installTime,_ := time.Parse("2006-01-02 15:04", instime)
            layoutName := "install"
            if scheduler == "on" {
                // Generate scheduled installation schedule list.
                scheduleListTime := strings.Replace(strings.Replace(strings.Replace(strings.Replace(instime, "-", ",", 1), "-", "-1,", 1), " ", ",", -1), ":", ",", -1)
                scheduleList = scheduleList + "{title:'("+k8sArrayLocal[i]+")',start: new Date("+scheduleListTime+"),allDay: false}"
                if i < len(k8sArrayLocal)-1 { scheduleList = scheduleList + "," }
                // Trigger takes effect and start scheduled installation!
                if installTime.Before(currentTime) {
                    if len(masterArray) == 1{
                        layoutName = "onemasterinstall"
                    }
                    go InstallScheduler(k8sArrayLocal[i], masterArray, nodeArray, kissh, currentDir, "install", layoutName, subProcessDir, sshPort, logName, mode)
                }
            } else {
                //Refresh progress bar time 
                progressBar,err_pb := ReadFile(currentDir+"/data/output"+subProcessDir+"/progressbar.txt")
                if progressBar != "" && err_pb == nil {
                    progressBarInt,_ := strconv.Atoi(progressBar)
                    progressBarInt = progressBarInt + 1
                    DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/progressbar.txt", strconv.Itoa(progressBarInt), currentDir, logName, mode)
                }
            }
            if err_stu == nil {
                switch {
                    case stu == "ok" :
                        err_health := DetectK8sHealth(k8sArrayLocal[i], currentDir, logName, mode)
                        if err_health == nil {
                            stuOk = stuOk + 1
                            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/k8shealth.txt", "healthy", currentDir, logName, mode)
                        } else {
                            stuUnknow = stuUnknow + 1
                            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/k8shealth.txt", "unhealthy", currentDir, logName, mode)
                        }
                    case stu == "installing":
                        stuInstall = stuInstall + 1
                        if modfiyTime.Before(relativeTime) {
                            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "unknow", currentDir, logName, mode)
                            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/progressbar.txt", "", currentDir, logName, mode)
                        }
                        // Generate installed schedule list.
                        scheduleListTime := strings.Replace(strings.Replace(strings.Replace(strings.Replace(instime, "-", ",", 1), "-", "-1,", 1), " ", ",", -1), ":", ",", -1)
                        scheduleList = scheduleList + "{title:'("+k8sArrayLocal[i]+")',start: new Date("+scheduleListTime+"),allDay: false}"
                        if i < len(k8sArrayLocal)-1 { scheduleList = scheduleList + "," }
                    case stu == "restarting" :
                        if modfiyTime.Before(relativeTime) {
                            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "unknow", currentDir, logName, mode)
                            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/progressbar.txt", "", currentDir, logName, mode)
                        } else {
                            err_health := DetectK8sHealth(k8sArrayLocal[i], currentDir, logName, mode)
                            if err_health == nil {
                                stuOk = stuOk + 1
                                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "ok", currentDir, logName, mode)
                                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/k8shealth.txt", "healthy", currentDir, logName, mode)
                            } else {
                                stuInstall = stuInstall + 1
                                DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/k8shealth.txt", "unhealthy", currentDir, logName, mode)
                            }
                        }
                    case stu == "uninstalling":
                        stuUninstall = stuUninstall + 1
                        if modfiyTime.Before(relativeTime) {
                            DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/status.txt", "unknow", currentDir, logName, mode)
                        }
                    case stu == "notok":
                        stuNotok = stuNotok + 1
                    case stu == "" :
                        k8sNum = k8sNum - 1
                    default :
                        stuUnknow = stuUnknow + 1
                        DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/k8shealth.txt", "unhealthy", currentDir, logName, mode)
                }
            } else {
                k8sNum = k8sNum - 1
            }
            nodeMap := GetClusterNode(k8sArrayLocal[i], currentDir, logName, mode)
            k8sNum := len(nodeMap)
            k8sNumMap [ k8sArrayLocal[i] ] = k8sNum  
        }
        DatabaseUpdate(currentDir+"/data/statistics/schedulelist.txt", scheduleList, currentDir, logName, mode)
        stuAll = stuOk + stuInstall + stuUninstall + stuNotok + stuUnknow
        if stuAll == 0 {
            stuUnknow = 1
            stuAll = 1
        }
        stuOkResult,_ := decimal.NewFromFloat(float64(stuOk)/float64(stuAll)*100 + 0/5).Round(2).Float64()
        stuInstallResult,_ := decimal.NewFromFloat(float64(stuInstall)/float64(stuAll)*100 + 0/5).Round(2).Float64()
        stuUninstallResult,_ := decimal.NewFromFloat(float64(stuUninstall)/float64(stuAll)*100 + 0/5).Round(2).Float64()
        stuNotokResult,_ := decimal.NewFromFloat(float64(stuNotok)/float64(stuAll)*100 + 0/5).Round(2).Float64()
        stuUnknowResult,_ := decimal.NewFromFloat(float64(stuUnknow)/float64(stuAll)*100 + 0/5).Round(2).Float64()
        k8smap := SortMapByValue(k8sNumMap)
        i := 0
        for k8s := range k8smap {
            labelArray[i] = k8smap[k8s].Key
            nodeNumArray[i] = strconv.Itoa(k8smap[k8s].Value)
            i++
            if i > 4 {
                break
            }
        }
        for j:=4; j>=0; j-- {
            labelStr = labelStr+labelArray[j]
            nodeNumStr = nodeNumStr+nodeNumArray[j]
            if j<1 {
                break
            }
            labelStr = labelStr+","
            nodeNumStr = nodeNumStr+","
        }
        cpuInfo,_ := cpu.Percent(time.Duration(time.Second), false)
        cpuInfoResult := int(math.Floor(cpuInfo[0] + 0/5))
        memInfo,_ := mem.VirtualMemory()
        memInfoResult := int(math.Floor(memInfo.UsedPercent + 0/5))
        diskInfo,_ := disk.Usage("/")
        diskInfoResult := int(math.Floor(diskInfo.UsedPercent + 0/5))
        DatabaseUpdate(currentDir+"/data/statistics/k8snum.txt", strconv.Itoa(k8sNum), currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/statistics/cpuinfo.txt", strconv.Itoa(cpuInfoResult), currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/statistics/meminfo.txt", strconv.Itoa(memInfoResult), currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/statistics/diskinfo.txt", strconv.Itoa(diskInfoResult), currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/statistics/stuok.txt", strconv.FormatFloat(stuOkResult,'E', -1, 64), currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/statistics/stuinstall.txt", strconv.FormatFloat(stuInstallResult,'E', -1, 64), currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/statistics/stuuninstall.txt", strconv.FormatFloat(stuUninstallResult,'E', -1, 64), currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/statistics/stunotok.txt", strconv.FormatFloat(stuNotokResult,'E', -1, 64), currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/statistics/stuunkonw.txt", strconv.FormatFloat(stuUnknowResult,'E', -1, 64), currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/statistics/labellist.txt", labelStr, currentDir, logName, mode)
        DatabaseUpdate(currentDir+"/data/statistics/nodenumlist.txt", nodeNumStr, currentDir, logName, mode)
        time.Sleep(time.Duration(60)*time.Second)
    }
}

// Schedule task schedule in hours.
func HourPeriodSchedule() {
    return
}

// Schedule task plan in days.
func DayPeriodSchedule() {
    return
}



