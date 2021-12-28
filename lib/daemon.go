package kilib

import (
	//"fmt"
        "time"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

        "github.com/gin-contrib/sessions"
        "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username     string `form:"username" binding:"required"`
	Password     string `form:"password" binding:"required"`
}

type passwordForm struct {
        Password         string `form:"password" binding:"required"`
        Newpassword1     string `form:"newpassword1" binding:"required"`
        Newpassword2     string `form:"newpassword2" binding:"required"`
}

type Version struct {
	Version     string `json:"version"`
	ReleaseDate string `json:"releaseDate"`
}

type ClusterList struct {
	Label          string   `json:"label"`
	K8sver         string   `json:"k8sver"`
	Softdir        string   `json:"softdir"`
	Ostype         string   `json:"ostype"`
	K8t            string   `json:"k8t"`
	Master         []string `json:"master"`
	Status         string   `json:"status"`
	Progressbar    string   `json:"progressbar"`
	Scheduler      string   `json:"scheduler"`
	Ost            string   `json:"ost"`
	Sdr            string   `json:"sdr"`
	K8sdashboardip string   `json:"k8sdashboardip"`
	Instime        string   `json:"instime"`
	K8v            string   `form:"k8v"`
	Lang           string   `form:"lang"`
}

type ClusterAddForm struct {
	Master             string `form:"master" binding:"required"`
	Node               string `form:"node" binding:"required"`
	Ostype             string `form:"ostype" binding:"required"`
	K8sver             string `form:"k8sver" binding:"required"`
	Label              string `form:"label" binding:"required"`
	Softdir            string `form:"softdir"`
	Installtime        string `form:"installtime"`
	Way                string `form:"way"`
        Upgradekernel      string `form:"upgradekernel"`
        K8sdashboard      string `form:"k8sdashboard"`
        Cniplugin          string `form:"cniplugin"`
}

type ClusterDelForm struct {
	Master  string `form:"master" binding:"required"`
	Node    string `form:"node" binding:"required"`
	Label   string `form:"label" binding:"required"`
	K8sver  string `form:"k8sver"`
	Softdir string `form:"softdir"`
	Ostype  string `form:"ostype"`
}

type MasterList struct {
	Label        string `json:"label"`
	K8sver       string `json:"k8sver"`
	Master       string `json:"master"`
	Masterstatus string `json:"masterstatus"`
	TTYstatus    string `json:"ttystatus"`
	Softdir      string `json:"softdir"`
	Ostype       string `form:"ostype"`
	Lang         string `json:"lang"`
}

type MasterrebuildForm struct {
	Master  []string `form:"master" binding:"required"`
	Label   string   `form:"label" binding:"required"`
	K8sver  string   `form:"k8sver" binding:"required"`
	Softdir string   `form:"softdir" binding:"required"`
	Ostype  string   `form:"ostype" binding:"required"`
}

type MasterDelForm struct {
	Master  []string `form:"master" binding:"required"`
	Label   string   `form:"label" binding:"required"`
	K8sver  string   `form:"k8sver" binding:"required"`
	Softdir string   `form:"softdir" binding:"required"`
	Ostype  string   `form:"ostype" binding:"required"`
}

type K8sList struct {
	K8s     string `json:"k8s"`
	Nodenum string `json:"Nodenum"`
}

type NodeList struct {
	Label      string `json:"label"`
	K8sver     string `json:"k8sver"`
	Ostype     string `json:"ostype"`
	Node       string `json:"node"`
	Nodestatus string `json:"nodestatus"`
	TTYstatus  string `json:"ttystatus"`
	Softdir    string `json:"softdir"`
	Lang       string `json:"lang"`
}

type NodeAddForm struct {
	Node               string `form:"addnode" binding:"required"`
	Label              string `form:"label" binding:"required"`
	K8sver             string `form:"k8sver" binding:"required"`
	Softdir            string `form:"softdir" binding:"required"`
	Ostype             string `form:"ostype" binding:"required"`
        Upgradekernel  string `form:"upgradekernel"`
}

type NodeDelForm struct {
	Node    []string `form:"node" binding:"required"`
	Label   string   `form:"label" binding:"required"`
	K8sver  string   `form:"k8sver" binding:"required"`
	Softdir string   `form:"softdir" binding:"required"`
	Ostype  string   `form:"ostype" binding:"required"`
}

type SelectList struct {
	Label       string `json:"label"`
	Labelnow    string `json:"labelnow"`
	Optnow      string `json:"optnow"`
	K8sver      string `json:"k8sver"`
	Softdir     string `json:"softdir"`
	Ostype      string `json:"ostype"`
	Status      string `json:"status"`
}

type ToolSwitchForm struct {
	Sshtool      string `form:"sshtool" binding:"required"`
	Installtool  string `form:"installtool" binding:"required"`
	Calendartool string `form:"calendartool" binding:"required"`
        Usertool     string `form:"usertool" binding:"required"`
}

type SshKeyForm struct {
	Sship   string `form:"sship" binding:"required"`
	Sshpass string `form:"sshpass" binding:"required"`
}

// Run the web server program and schedule program.
func DaemonRun(Version string, ReleaseDate string, CompatibleK8S string, CompatibleOS string, listenIPandPort string, currentDir string, currentUser string, kissh string, logName string, mode string) {

	// Create kube-install daemon log file
	CreateDir(currentDir+"/data/logs/kubeinstalld/", currentDir, logName, mode)
	f, _ := os.Create(currentDir + "/data/logs/kubeinstalld/" + logName + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Create http router
	router := gin.Default()
	router.LoadHTMLGlob(currentDir + "/static/html/*")
	router.StaticFS("/static", http.Dir(currentDir+"/static/staticfs"))
        store := cookie.NewStore([]byte("ju65yjhi72jjsjdc721gsfc62"))
        store.Options(sessions.Options{
            MaxAge: int(30 * time.Minute), //30min
            Path:   "/",
        })
        router.Use(sessions.Sessions("0x000000cnrki", store))

	// Background regular inspection statistics of various states
	go MinutePeriodSchedule(currentDir, kissh, logName, mode)

	/*********************************************************************
	   Kube-Install Web Backend
	**********************************************************************/

	// home page
	router.GET("/", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get index information
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		var k8slist []K8sList
		label := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		osType := c.DefaultQuery("ostype", "")
		msg, _ := ReadFile(currentDir + "/data/msg/msg.txt")
		k8sNum, _ := ReadFile(currentDir + "/data/statistics/k8snum.txt")
		cpuInfoResult, _ := ReadFile(currentDir + "/data/statistics/cpuinfo.txt")
		memInfoResult, _ := ReadFile(currentDir + "/data/statistics/meminfo.txt")
		diskInfoResult, _ := ReadFile(currentDir + "/data/statistics/diskinfo.txt")
		stuOk, _ := ReadFile(currentDir + "/data/statistics/stuok.txt")
		stuOkResult, _ := strconv.ParseFloat(stuOk, 64)
		stuInstall, _ := ReadFile(currentDir + "/data/statistics/stuinstall.txt")
		stuInstallResult, _ := strconv.ParseFloat(stuInstall, 64)
		stuUninstall, _ := ReadFile(currentDir + "/data/statistics/stuuninstall.txt")
		stuUninstallResult, _ := strconv.ParseFloat(stuUninstall, 64)
		stuNotok, _ := ReadFile(currentDir + "/data/statistics/stunotok.txt")
		stuNotokResult, _ := strconv.ParseFloat(stuNotok, 64)
		stuUnknow, _ := ReadFile(currentDir + "/data/statistics/stuunkonw.txt")
		stuUnknowResult, _ := strconv.ParseFloat(stuUnknow, 64)
		labelStr, _ := ReadFile(currentDir + "/data/statistics/labellist.txt")
		nodeNumStr, _ := ReadFile(currentDir + "/data/statistics/nodenumlist.txt")
		labelArray := strings.Split(labelStr, ",")
		nodeNumArray := strings.Split(nodeNumStr, ",")
		for i := 0; i < len(labelArray); i++ {
			k8slist = append(k8slist, K8sList{labelArray[i], nodeNumArray[i]})
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         label,
			"K8sver":        k8sVer,
			"Softdir":       softDir,
			"Ostype":        osType,
			"K8snum":        k8sNum,
			"Syscpu":        cpuInfoResult,
			"Sysmem":        memInfoResult,
			"Sysdisk":       diskInfoResult,
			"Status1":       stuOkResult,
			"Status2":       stuInstallResult,
			"Status3":       stuUninstallResult,
			"Status4":       stuNotokResult,
			"Status5":       stuUnknowResult,
			"K8slist":       k8slist,
			"Msg":           template.HTML(msg),
			"Version":       Version,
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

        // login page
        router.GET("/login", func(c *gin.Context) {
                session := sessions.Default(c)
                session.Get("admin")
                langFromWeb := c.Query("lang")
                Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
                c.HTML(http.StatusOK, "login.tmpl", gin.H{"Lang": Lang})
        })

	router.GET("/clusterlist", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get cluster list information
		var clusterlist []ClusterList
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		label := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		osType := c.DefaultQuery("ostype", "")
		labelArray, err := GetAllDir(currentDir+"/data/output", currentDir, logName, mode)
		CheckErr(err, currentDir, logName, mode)
		for _, i := range labelArray {
			k8t := string(i)
			stu, sch := GetClusterStatus(k8t, currentDir, logName, mode)
			sdr := GetClusterSoftdir(k8t, currentDir, mode)
			k8v := GetClusterK8sVer(k8t, currentDir, mode)
			ost := GetClusterOstype(k8t, currentDir, mode)
			_, k8sDashboardIp, _ := GetClusterAddons(k8t, currentDir, mode)
			_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", k8t, "")
			mst, err_mst := GetAllDir(currentDir+"/data/output"+subProcessDir+"/masters", currentDir, logName, mode)
			CheckErr(err_mst, currentDir, logName, mode)
			instime, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/installtime.txt")
			progressBar, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/progressbar.txt")
			clusterlist = append(clusterlist, ClusterList{label, k8sVer, softDir, osType, k8t, mst, stu, progressBar, sch, ost, sdr, k8sDashboardIp, instime, k8v, Lang})
		}
		c.HTML(http.StatusOK, "cluster.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         label,
			"K8sver":        k8sVer,
			"Softdir":       softDir,
			"Ostype":        osType,
			"Clusterlist":   clusterlist,
			"Sshuser":       currentUser,
			"Version":       Version,
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

	// cluster info page
	router.GET("/clusterinfo", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get cluster information
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		label := c.DefaultQuery("label", "")
		stu, _ := GetClusterStatus(label, currentDir, logName, mode)
		_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
		mst, err := GetAllDir(currentDir+"/data/output"+subProcessDir+"/masters", currentDir, logName, mode)
		CheckErr(err, currentDir, logName, mode)
		k8sVer, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/k8sver.txt")
                cniPlugin, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/cniplugin.txt")
		osType, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/ostype.txt")
		softDir, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/softdir.txt")
		etcdEndpoints, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/etcdendpoints.txt")
		nd := ListNode(label, currentDir, logName, mode)
		registryIp, k8sDashboardIp, k8sdashboardtoken := GetClusterAddons(label, currentDir, mode)
		kubeCfg := GetClusterKubecfg(label, currentDir, mode)
		if err != nil {
			kubeCfg = ""
		} else {
			kubeCfg = strings.Replace(strings.Replace(kubeCfg, "\n", "<br>\n", -1), " ", "&nbsp;", -1)
		}
		var registryUsage, k8sDashboardUsage string
		if Lang == "cn" {
			registryUsage = "&nbsp;&nbsp;&nbsp;&nbsp; [root@localhost ~]# docker pull " + registryIp + ":5000/镜像名称:镜像Tag<br>&nbsp;&nbsp;&nbsp;&nbsp; [root@localhost ~]# docker push " + registryIp + ":5000/镜像名称:镜像Tag<br>&nbsp;&nbsp;&nbsp;&nbsp; [root@localhost ~]# ctr -n=k8s.io images pull " + registryIp + ":5000/镜像名称:镜像Tag<br>&nbsp;&nbsp;&nbsp;&nbsp; [root@localhost ~]# ctr -n=k8s.io images push " + registryIp + ":5000/镜像名称:镜像Tag\n"
			k8sDashboardUsage = "使用浏览器访问，登录令牌如下: <br> " + k8sdashboardtoken
		} else {
			registryUsage = "&nbsp;&nbsp;&nbsp;&nbsp; [root@localhost ~]# docker pull " + registryIp + ":5000/Your_image_name:Image_Tag<br>&nbsp;&nbsp;&nbsp;&nbsp; [root@localhost ~]# docker push " + registryIp + ":5000/Your_image_name:Image_Tag<br>&nbsp;&nbsp;&nbsp;&nbsp; [root@localhost ~]# ctr -n=k8s.io images pull " + registryIp + ":5000/Your_image_name:Image_Tag<br>&nbsp;&nbsp;&nbsp;&nbsp; [root@localhost ~]# ctr -n=k8s.io images push " + registryIp + ":5000/Your_image_name:Image_Tag\n"
			k8sDashboardUsage = "Use a browser to access. The login token is as follows: <br> " + k8sdashboardtoken
		}
		c.HTML(http.StatusOK, "clusterinfo.tmpl", gin.H{
			"Lang":              Lang,
			"Label":             label,
			"K8sver":            k8sVer,
                        "Cniplugin":         cniPlugin,
			"Softdir":           softDir,
			"Ostype":            osType,
			"Status":            stu,
			"Master":            mst,
			"Node":              nd,
			"Etcdendpoints":     etcdEndpoints,
			"Registryip":        registryIp,
			"Registryusage":     template.HTML(registryUsage),
			"K8sdashboardip":    k8sDashboardIp,
			"K8sdashboardusage": template.HTML(k8sDashboardUsage),
			"Kubeconfig":        template.HTML(kubeCfg),
			"Sshuser":           currentUser,
			"Tools":             "no",
			"Opt":               "",
			"Version":           Version,
			"Releasedate":       ReleaseDate,
			"Compatiblek8s":     CompatibleK8S,
			"Compatibleos":      CompatibleOS,
		})
	})

	// page of add cluster
	router.GET("/clusteradd", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get cluster information
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		label := ""
		mst := ""
		osType := ""
		k8sVer := ""
                cniPlugin := ""
		softDir := "/opt/kube-install"
		nd := ""
		way := c.DefaultQuery("way", "newinstall")
		tools := c.Query("tools")
		if way == "reinstall" {
			label = c.DefaultQuery("label", "")
			mst = c.DefaultQuery("master", "")
			mst = strings.Replace(mst[1:len(mst)-1], " ", ",", -1)
			osType = c.DefaultQuery("ostype", "")
			k8sVer = c.DefaultQuery("k8sver", "")
                        cniPlugin = GetClusterCNI(label, currentDir, mode)
			softDir = c.DefaultQuery("softdir", "/opt/kube-install")
			nodemap := GetClusterNode(label, currentDir, logName, mode)
			for node := range nodemap {
				nd = nd + "," + node
			}
			if nd != "" {
				nd = nd[1:len(nd)]
			}
		}
		c.HTML(http.StatusOK, "clusteradd.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         label,
			"Master":        mst,
			"Ostype":        osType,
			"K8sver":        k8sVer,
                        "Cniplugin":     cniPlugin,
			"Node":          nd,
			"Softdir":       softDir,
			"Sshuser":       currentUser,
			"Way":           way,
			"Tools":         tools,
			"Opt":           "install",
			"Version":       Version,
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

	// page of delete cluster
	router.GET("/clusterdel", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get cluster information
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		label := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		osType := c.DefaultQuery("ostype", "")
		mst := c.DefaultQuery("master", "")
		tools := c.Query("tools")
		nd := ""
		nodemap := GetClusterNode(label, currentDir, logName, mode)
		for node := range nodemap {
			nd = nd + "," + node
		}
		if nd != "" {
			nd = nd[1:len(nd)]
		}
		c.HTML(http.StatusOK, "clusterdel.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         label,
			"K8sver":        k8sVer,
			"Master":        strings.Replace(mst[1:len(mst)-1], " ", ",", -1),
			"Node":          nd,
			"Softdir":       softDir,
			"Ostype":        osType,
			"Sshuser":       currentUser,
			"Tools":         tools,
			"Opt":           "uninstall",
			"Version":       Version,
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

	// page of delete schedule
	router.GET("/deleteschedule", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get schedule information
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		label := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
                mst := c.DefaultQuery("master", "")
                mst = strings.Replace(mst[1:len(mst)-1], " ", ",", -1)
                nd := ""
                nodemap := GetClusterNode(label, currentDir, logName, mode)
                for node := range nodemap {
                    nd = nd + "," + node
                }
                if nd != "" {
                    nd = nd[1:len(nd)]
                }
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		osType := c.DefaultQuery("ostype", "")
		scheduler := c.Query("scheduler")
		instime := c.Query("instime")
		tools := c.Query("tools")
		c.HTML(http.StatusOK, "deleteschedule.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         label,
			"K8sver":        k8sVer,
                        "Master":        mst,
                        "Node":          nd,
			"Softdir":       softDir,
			"Ostype":        osType,
			"Sshuser":       currentUser,
			"Scheduler":     scheduler,
			"Instime":       instime,
			"Tools":         tools,
			"Opt":           "uninstall",
			"Version":       Version,
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

	// master page
	router.GET("/masteradmin", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get master list information
		var masterlist []MasterList
		var selectlist []SelectList
		var osType string
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		labelNow := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		stu := c.DefaultQuery("status", "unknow")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		if labelNow != "" {
			_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", labelNow, "")
			osType, _ = ReadFile(currentDir + "/data/output" + subProcessDir + "/ostype.txt")
			mastermap := GetClusterMaster(labelNow, currentDir, logName, mode)
			for master := range mastermap {
				ttyStu, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/masters/" + master + "/ttystatus.txt")
				masterlist = append(masterlist, MasterList{labelNow, k8sVer, master, mastermap[master], ttyStu, softDir, osType, Lang})
			}
		} else {
			osType = ""
		}
		labelArray, err := GetAllDir(currentDir+"/data/output", currentDir, logName, mode)
		CheckErr(err, currentDir, logName, mode)
		var k8v string
		for _, i := range labelArray {
			label := string(i)
			_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
			k8v, _ = ReadFile(currentDir + "/data/output" + subProcessDir + "/k8sver.txt")
			stus, _ := GetClusterStatus(label, currentDir, logName, mode)
			selectlist = append(selectlist, SelectList{label, labelNow, "", k8v, softDir, osType, stus})
		}
		c.HTML(http.StatusOK, "master.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         labelNow,
			"K8sver":        k8sVer,
			"Status":        stu,
			"Softdir":       softDir,
			"Ostype":        osType,
			"Selectlist":    selectlist,
			"Masterlist":    masterlist,
			"Version":       Version,
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

	// page of master info
	router.GET("/masterinfo", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get master information
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		label := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		masterIp := c.DefaultQuery("masterip", "127.0.0.1")
		masterStatus := c.DefaultQuery("masterstatus", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		osType := c.DefaultQuery("ostype", "")
		c.HTML(http.StatusOK, "masterinfo.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         label,
			"K8sver":        k8sVer,
			"Masterip":      masterIp,
			"Masterstatus":  masterStatus,
			"Ostype":        osType,
			"Softdir":       softDir,
			"Version":       Version,
			"Tools":         "no",
			"Opt":           "",
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

	// node page
	router.GET("/nodeadmin", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get node list information
		var nodelist []NodeList
		var selectlist []SelectList
		var status, osType string
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		labelNow := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		if labelNow != "" {
			_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", labelNow, "")
			status, _ = ReadFile(currentDir + "/data/output" + subProcessDir + "/status.txt")
			osType, _ = ReadFile(currentDir + "/data/output" + subProcessDir + "/ostype.txt")
			nodemap := GetClusterNode(labelNow, currentDir, logName, mode)
			for node := range nodemap {
				ttyStu, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/nodes/" + node + "/ttystatus.txt")
				nodelist = append(nodelist, NodeList{labelNow, k8sVer, osType, node, nodemap[node], ttyStu, softDir, Lang})
			}
		} else {
			status = "unknow"
			osType = ""
		}
		labelArray, err := GetAllDir(currentDir+"/data/output", currentDir, logName, mode)
		CheckErr(err, currentDir, logName, mode)
		var k8v string
		for _, i := range labelArray {
			label := string(i)
			_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
			k8v, _ = ReadFile(currentDir + "/data/output" + subProcessDir + "/k8sver.txt")
			stus, _ := GetClusterStatus(label, currentDir, logName, mode)
			selectlist = append(selectlist, SelectList{label, labelNow, "", k8v, softDir, osType, stus})
		}
		c.HTML(http.StatusOK, "node.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         labelNow,
			"K8sver":        k8sVer,
			"Status":        status,
			"Softdir":       softDir,
			"Ostype":        osType,
			"Selectlist":    selectlist,
			"Nodelist":      nodelist,
			"Sshuser":       currentUser,
			"Version":       Version,
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

	// page of node info
	router.GET("/nodeinfo", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get node information
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		label := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		osType := c.DefaultQuery("ostype", "")
		nodeIp := c.DefaultQuery("nodeip", "127.0.0.1")
		nodeStatus := c.DefaultQuery("nodestatus", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		runStatus, nodeLabels, runcVer, osVer, kernelVer, cpu, memory, createTime := GetNodeInfo(label, nodeIp, currentDir, logName, mode)
		c.HTML(http.StatusOK, "nodeinfo.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         label,
			"K8sver":        k8sVer,
			"Nodeip":        nodeIp,
			"Nodestatus":    nodeStatus,
			"Runstatus":     runStatus,
			"Nodelabels":    nodeLabels,
			"Ostype":        osType,
			"Osver":         osVer,
			"Softdir":       softDir,
			"Runc":          runcVer,
			"Kernelver":     kernelVer,
			"Cpu":           cpu,
			"Memory":        memory,
			"Createtime":    createTime,
			"Version":       Version,
			"Tools":         "no",
			"Opt":           "",
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

	// page of add node
	router.GET("/nodeadd", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get node information
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		label := c.DefaultQuery("label", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		osType := c.DefaultQuery("ostype", "")
		c.HTML(http.StatusOK, "nodeadd.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         label,
			"Softdir":       softDir,
			"Sshuser":       currentUser,
			"Ostype":        osType,
			"Version":       Version,
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

	// tools page
	router.GET("/tools", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get tools list information
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		toolSwitch := GetToolSwitch(currentDir, logName, mode)
		sshTool := toolSwitch[0]
		installTool := toolSwitch[1]
		calendarTool := toolSwitch[2]
                userTool := toolSwitch[3]
		label := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		osType := c.DefaultQuery("ostype", "")
		c.HTML(http.StatusOK, "tools.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         label,
			"Sshtool":       sshTool,
			"Installtool":   installTool,
			"Calendartool":  calendarTool,
                        "Usertool":      userTool,
			"K8sver":        k8sVer,
			"Softdir":       softDir,
			"Ostype":        osType,
			"Sshuser":       currentUser,
			"Version":       Version,
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

	// page of calendar scheduler
	router.GET("/calendarscheduler", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get calendar information
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		label := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		osType := c.DefaultQuery("ostype", "")
		scheduleList, _ := ReadFile(currentDir + "/data/statistics/schedulelist.txt")
		c.HTML(http.StatusOK, "calendarscheduler.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         label,
			"K8sver":        k8sVer,
			"Softdir":       softDir,
			"Ostype":        osType,
			"Schedulelist":  template.JS(scheduleList),
			"Sshuser":       currentUser,
			"Version":       Version,
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

        // page of user info
        router.GET("/userinfo", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get user information
                langFromWeb := c.Query("lang")
                Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
                label := c.DefaultQuery("label", "")
                k8sVer := c.DefaultQuery("k8sver", "")
                softDir := c.DefaultQuery("softdir", "/opt/kube-install")
                osType := c.DefaultQuery("ostype", "")
                backUrl := c.DefaultQuery("backurl", "tools")
                c.HTML(http.StatusOK, "userinfo.tmpl", gin.H{
                        "Lang":          Lang,
                        "Label":         label,
                        "K8sver":        k8sVer,
                        "Softdir":       softDir,
                        "Ostype":        osType,
                        "Backurl":        backUrl,
                        "Sshuser":       currentUser,
                        "Version":       Version,
                        "Releasedate":   ReleaseDate,
                        "Compatiblek8s": CompatibleK8S,
                        "Compatibleos":  CompatibleOS,
                })
        })

	// logs page
	router.GET("/logs", func(c *gin.Context) {
                //get and check session
                session := sessions.Default(c)
                checkSession(session, c)
                //get logs information
		var selectlist []SelectList
		var clog string
		var err error
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		labelNow := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		osType := c.DefaultQuery("ostype", "")
		opt := c.DefaultQuery("opt", "")
		labelArray, err := GetAllDir(currentDir+"/data/output", currentDir, logName, mode)
		CheckErr(err, currentDir, logName, mode)
		for _, i := range labelArray {
			label := string(i)
			stu, _ := GetClusterStatus(label, currentDir, logName, mode)
			selectlist = append(selectlist, SelectList{label, labelNow, opt, "", "", "", stu})
		}
		if opt == "systemlog" {
			clog, err = ReadFile(currentDir + "/data/logs/kubeinstalld/" + logName + ".log")
			if err != nil {
				clog = "No log information..."
			}
			clog = strings.Replace(strings.Replace(clog, "\n", "<br>\n", -1), " ", "&nbsp;", -1)
		} else if opt == "" {
			if Lang == "cn" {
				clog = "请点击右上角的 [选择操作日志类型] 按钮选择日志类型..."
			} else {
				clog = "Please click the [Select Log Type] button in the upper right corner to select the log type..."
			}
		} else {
			if labelNow == "" {
				if Lang == "cn" {
					clog = "请点击右上角的 [切换kubernetes集群] 按钮，选择一个集群来查看日志..."
				} else {
					clog = "Please click the [Switch K8S Cluster] button in the upper right corner to select a cluster to view logs..."
				}
			} else {
				_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", labelNow, "")
				clog, err = ReadFile(currentDir + "/data/logs" + subProcessDir + "/logs/" + opt + ".log")
				if err != nil {
					clog = "No log information..."
				}
				clog = strings.Replace(strings.Replace(clog, "\n", "<br>\n", -1), " ", "&nbsp;", -1)
			}
		}
		c.HTML(http.StatusOK, "logs.tmpl", gin.H{
			"Lang":          Lang,
			"Label":         labelNow,
			"K8sver":        k8sVer,
			"Softdir":       softDir,
			"Ostype":        osType,
			"Opt":           opt,
			"Selectlist":    selectlist,
			"Clog":          template.HTML(clog),
			"Version":       Version,
			"Releasedate":   ReleaseDate,
			"Compatiblek8s": CompatibleK8S,
			"Compatibleos":  CompatibleOS,
		})
	})

	/*********************************************************************
	   Kube-Install Operation Process
	**********************************************************************/

        // operation of login
        router.POST("/loginset", func(c *gin.Context) {
                var form LoginForm
                var username,password,errorInfo string
                langFromWeb := c.Query("lang")
                Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
                if c.ShouldBind(&form) == nil {
                        username = form.Username
                        password = form.Password
                } else {
                        username = c.DefaultQuery("username", "")
                        password = c.DefaultQuery("password", "")
                }
                // in this case proper binding will be automatically selected
                if c.ShouldBind(&form) == nil {
                        pwdByte, err := ReadFile(currentDir+"/data/.key/admin")
                        if err != nil {
                            errorInfo = "System error! Please use the 'kube-install -init' command to reinitialize your system."
                            c.HTML(http.StatusOK, "loginset.tmpl", gin.H{"Errorinfo": errorInfo, "Lang": Lang})
                            return
                        }
                        key, _ := DePwdCode(pwdByte)
                        if username == "admin" && password == string(key) {
                                    session := sessions.Default(c)
                                    session.Set(username, password)
                                    session.Save()
                                    c.Redirect(http.StatusMovedPermanently, "/")
                        } else {
                                if Lang == "cn" {
                                    errorInfo = "你输入的用户名或密码错误，"
                                } else {
                                    errorInfo = "The username or password you entered is incorrect."
                                }
                                c.HTML(http.StatusOK, "loginset.tmpl", gin.H{"Errorinfo": errorInfo, "Lang": Lang})
                                return
                        }
                } else {
                        if Lang == "cn" {
                            errorInfo = "用户名或密码不能为空，"
                        } else {
                            errorInfo = "User name or password cannot be empty."
                        }
                        c.HTML(http.StatusOK, "loginset.tmpl", gin.H{"Errorinfo": errorInfo, "Lang": Lang})
                        return
                }
        })

        // operation of logout
        router.GET("/logoutset", func(c *gin.Context) {
            var username string
            username = c.DefaultQuery("username", "admin")
            session := sessions.Default(c)
            session.Delete(username)
            session.Clear()
            c.Redirect(http.StatusMovedPermanently, "/login")
        })

	// operation of install
	router.POST("/install", func(c *gin.Context) {
		var form ClusterAddForm
		var master, node, osType, k8sVer, softDir, label, installTime, way, upgradeKernel, k8sDashboard, cniPlugin string
		if c.ShouldBind(&form) == nil {
			master = form.Master
			node = form.Node
			osType = form.Ostype
			k8sVer = form.K8sver
			softDir = form.Softdir
			label = form.Label
			installTime = form.Installtime
			way = form.Way
                        upgradeKernel  = form.Upgradekernel
                        k8sDashboard = form.K8sdashboard
                        cniPlugin = form.Cniplugin
		} else {
			master = c.Query("master")
			node = c.Query("node")
			osType = c.DefaultQuery("ostype", "")
			k8sVer = c.DefaultQuery("k8ever", "")
			softDir = c.DefaultQuery("softdir", "/opt/kube-install")
			label = c.DefaultQuery("label", "")
			installTime = c.DefaultQuery("installtime", "")
			way = c.DefaultQuery("way", "newinstall")
                        upgradeKernel  = c.DefaultQuery("upgradekernel", "")
                        k8sDashboard  = c.DefaultQuery("k8sdashboard", "")
                        cniPlugin = c.DefaultQuery("cniplugin", "")
		}
                langFromWeb := c.Query("lang")
                Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
                tools := c.Query("tools")
                if !CheckLabel(label) {
                        c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "install", "Result": "failure", "Info": "Installation failed! The \"label\" parameter length must be less than 32 strings, please check!", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
                        return
                }
		if osType == "unknow" {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "install", "Result": "failure", "Info": "Installation failed! Please set the operating system type correctly!", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
			return
		}
		if way == "newinstall" {
			labelArray, err := GetAllDir(currentDir+"/data/output", currentDir, logName, mode)
			CheckErr(err, currentDir, logName, mode)
			if StrInArray(label, labelArray) {
				c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "install", "Result": "failure", "Info": "Installation failed! Label cannot be repeated!", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
				return
			}
		}
		if master != "" && node != "" {
			masterArray, nodeArray, softDir, subProcessDir, osTypeResult := ParameterConvert(mode, master, node, softDir, label, osType)
			for i := 0; i < len(masterArray); i++ {
				if !CheckIP(masterArray[i]) {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "install", "Result": "failure", "Info": "Installation failed! The Master IP address format you entered is incorrect, please check!", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					return
				}
			}
			for j := 0; j < len(nodeArray); j++ {
				if !CheckIP(nodeArray[j]) {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "install", "Result": "failure", "Info": "Installation failed! The Node IP address format you entered is incorrect, please check!", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					return
				}
			}
			DatabaseInit(currentDir, subProcessDir, logName, mode)
			stu, sch := GetClusterStatus(label, currentDir, logName, mode)
			var installInfo string
			if stu != "installing" && stu != "restarting" && stu != "uninstalling" && sch != "on" {
				go InstallCore(mode, master, masterArray, node, nodeArray, softDir, currentDir, kissh, subProcessDir, currentUser, label, osTypeResult, osType, k8sVer, logName, Version, CompatibleK8S, CompatibleOS, installTime, way, upgradeKernel, k8sDashboard, cniPlugin)
				if installTime == "" {
					if Lang == "cn" {
						installInfo = "Kubernetes集群正在后台安装中 ... "
					} else {
						installInfo = "Kubernetes cluster installation started in the background ... "
					}
				} else {
					if Lang == "cn" {
						installInfo = "计划任务已经生成，系统将会在" + installTime + "进行Kubernetest集群的安装部署操作！"
					} else {
						installInfo = "The planning task has been generated. The system will install the kubernetest cluster at " + installTime + " !"
					}
				}
				c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "install", "Result": "success", "Info": installInfo, "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
			} else {
				c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "install", "Result": "failure", "Info": "Installation failed! There are scheduled tasks in the background, or someone else is installing or uninstalling the current cluster!", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
			}
		} else {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "install", "Result": "failure", "Info": "Installation failed! The parameter you entered is wrong!", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
		}
	})

	// operation of rebuild master
	router.POST("/rebuildmaster", func(c *gin.Context) {
		var form MasterrebuildForm
		var masterArray []string
		var softDir, label, k8sVer, osType string
		if c.ShouldBind(&form) == nil {
			masterArray = form.Master
			softDir = form.Softdir
			label = form.Label
			k8sVer = form.K8sver
			osType = form.Ostype
		} else {
			masterArray = strings.Split(c.Query("master"), ",")
			softDir = c.DefaultQuery("softdir", "/opt/kube-install")
			label = c.DefaultQuery("label", "")
			k8sVer = c.DefaultQuery("k8sver", "")
			osType = c.DefaultQuery("ostype", "")
		}
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		if osType == "unknow" {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "rebuildmaster", "Result": "failure", "Info": "Rebuild failed! Please operate correctly System type!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
			return
		}
		_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", softDir, label, osType)
		masterArraylen := len(masterArray)
		if masterArraylen < 1 || masterArray[0] == "" {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "rebuildmaster", "Result": "failure", "Info": "Rebuild failed! The IP address of kubernetes master cannot be empty!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
			return
		} else {
			for i := 0; i < masterArraylen; i++ {
				if !CheckIP(masterArray[i]) {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "rebuildmaster", "Result": "failure", "Info": "Rebuild failed! The Master IP address format you entered is incorrect, please check!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					return
				}
				stu, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/masters/" + masterArray[i] + "/status.txt")
				if stu == "adding" || stu == "rebuilding" || stu == "deleting" {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "rebuildmaster", "Result": "failure", "Info": "Rebuild failed! K8s master is being deleted or repaired by others!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					return
				}
			}
		}
		go RebuildMasterCore(mode, masterArray, currentDir, kissh, subProcessDir, currentUser, label, softDir, logName)
		if Lang == "cn" {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "rebuildmaster", "Result": "success", "Info": "Kubernetes master正在后台修复中 ...", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
		} else {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "rebuildmaster", "Result": "success", "Info": "Started repairing kubernetes master in the background ...", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
		}
	})

	// operation of delete master
	router.POST("/delmaster", func(c *gin.Context) {
		var form MasterDelForm
		var masterArray []string
		var softDir, label, k8sVer, osType string
		if c.ShouldBind(&form) == nil {
			masterArray = form.Master
			softDir = form.Softdir
			label = form.Label
			k8sVer = form.K8sver
			osType = form.Ostype
		} else {
			masterArray = strings.Split(c.Query("master"), ",")
			softDir = c.DefaultQuery("softdir", "/opt/kube-install")
			label = c.DefaultQuery("label", "")
			k8sVer = c.DefaultQuery("k8sver", "")
			osType = c.DefaultQuery("ostype", "")
		}
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", softDir, label, "")
		masterArraylen := len(masterArray)
		if masterArraylen < 1 || masterArray[0] == "" {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "delmaster", "Result": "failure", "Info": "Delete failed! The IP address of kubernetes master cannot be empty!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
			return
		} else {
			for i := 0; i < masterArraylen; i++ {
				if !CheckIP(masterArray[i]) {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "rebuildmaster", "Result": "failure", "Info": "Delete failed! The Master IP address format you entered is incorrect, please check!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					return
				}
				stu, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/masters/" + masterArray[i] + "/status.txt")
				if stu == "adding" || stu == "rebuilding" || stu == "deleting" {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "rebuildmaster", "Result": "failure", "Info": "Delete failed! K8s master is being deleted or repaired by others!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					return
				}
			}
		}
		go DeleteMasterCore(mode, masterArray, currentDir, kissh, subProcessDir, currentUser, label, softDir, logName)
		if Lang == "cn" {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "delmaster", "Result": "success", "Info": "Kubernetes master 正在后台销毁中 ...", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
		} else {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "delmaster", "Result": "success", "Info": "Started deleting kubernetes master in the background ...", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
		}
	})

	// operation of add node
	router.POST("/addnode", func(c *gin.Context) {
		var form NodeAddForm
		var node, softDir, label, k8sVer, osType, upgradeKernel  string
		if c.ShouldBind(&form) == nil {
			node = form.Node
			softDir = form.Softdir
			label = form.Label
			k8sVer = form.K8sver
			osType = form.Ostype
                        upgradeKernel  = form.Upgradekernel
		} else {
			node = c.Query("node")
			softDir = c.DefaultQuery("softdir", "/opt/kube-install")
			label = c.DefaultQuery("label", "")
			k8sVer = c.DefaultQuery("k8sver", "")
			osType = c.DefaultQuery("ostype", "")
                        upgradeKernel  = c.DefaultQuery("upgradekernel", "")
		}
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		if osType == "unknow" {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "addnode", "Result": "failure", "Info": "Failed to add! Please select the operating system type of kubernetes node correctly!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
			return
		}
		_, nodeArray, _, subProcessDir, osTypeResult := ParameterConvert(mode, "", node, softDir, label, osType)
		nodeArraylen := len(nodeArray)
		nd := ListNode(label, currentDir, logName, mode)
		for j := 0; j < len(nd); j++ {
			if strings.Contains(node, nd[j]) {
				c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "addnode", "Result": "failure", "Info": "Failed to add! kubernetes node already exists, cannot add repeatedly!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
				return
			}
		}
		if nodeArraylen < 1 || nodeArray[0] == "" {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "addnode", "Result": "failure", "Info": "Failed to add! The IP address of kubernetes node cannot be empty!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
			return
		} else {
			for i := 0; i < nodeArraylen; i++ {
				if !CheckIP(nodeArray[i]) {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "addnode", "Result": "failure", "Info": "Failed to add! The Node IP address format you entered is incorrect, please check!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					return
				}
				stu, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/nodes/" + nodeArray[i] + "/status.txt")
				if stu == "adding" || stu == "deleting" {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "addnode", "Result": "failure", "Info": "Failed to add! kubernetes node is being deleted or added by others!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					return
				}
			}
		}
		go AddNodeCore(mode, node, nodeArray, currentDir, kissh, subProcessDir, currentUser, label, softDir, osTypeResult, k8sVer, logName, CompatibleOS, upgradeKernel )
		if Lang == "cn" {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "addnode", "Result": "success", "Info": "Kubernetes node正在后台添加中 ...", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
		} else {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "addnode", "Result": "success", "Info": "Started adding kubernetes node in the background ...", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
		}
	})

	// operation of delete node
	router.POST("/delnode", func(c *gin.Context) {
		var form NodeDelForm
		var nodeArray []string
		var softDir, label, k8sVer, osType string
		if c.ShouldBind(&form) == nil {
			nodeArray = form.Node
			softDir = form.Softdir
			label = form.Label
			k8sVer = form.K8sver
			osType = form.Ostype
		} else {
			nodeArray = strings.Split(c.Query("node"), ",")
			softDir = c.DefaultQuery("softdir", "/opt/kube-install")
			label = c.DefaultQuery("label", "")
			k8sVer = c.DefaultQuery("k8sver", "")
			osType = c.DefaultQuery("ostype", "")
		}
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", softDir, label, "")
		nodeArraylen := len(nodeArray)
		if nodeArraylen < 1 || nodeArray[0] == "" {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "addnode", "Result": "failure", "Info": "Delete failed! The IP address of kubernetes node cannot be empty!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
			return
		} else {
			for i := 0; i < nodeArraylen; i++ {
				if !CheckIP(nodeArray[i]) {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "addnode", "Result": "failure", "Info": "Delete failed! The Node IP address format you entered is incorrect, please check!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					return
				}
				stu, _ := ReadFile(currentDir + "/data/output" + subProcessDir + "/nodes/" + nodeArray[i] + "/status.txt")
				if stu == "adding" || stu == "deleting" {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "addnode", "Result": "failure", "Info": "Delete failed! K8s node is being deleted or added by others!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					return
				}
			}
		}
		go DeleteNodeCore(mode, nodeArray, currentDir, kissh, subProcessDir, currentUser, label, softDir, logName)
		if Lang == "cn" {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "delnode", "Result": "success", "Info": "Kubernetes node正在后台销毁中 ...", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
		} else {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "delnode", "Result": "success", "Info": "Started deleting kubernetes node in the background ...", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
		}
	})

	// operation of uninstall
	router.POST("/uninstall", func(c *gin.Context) {
		var form ClusterDelForm
		var master, node, softDir, label, k8sVer, osType string
		if c.ShouldBind(&form) == nil {
			master = form.Master
			node = form.Node
			softDir = form.Softdir
			label = form.Label
			k8sVer = form.K8sver
			osType = form.Ostype
		} else {
			master = c.Query("master")
			node = c.Query("node")
			softDir = c.DefaultQuery("softdir", "/opt/kube-install")
			label = c.DefaultQuery("label", "")
			k8sVer = c.DefaultQuery("k8sver", "")
			osType = c.DefaultQuery("ostype", "")
		}
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		if master != "" && node != "" && osType != "unknow" {
			masterArray, nodeArray, softDir, subProcessDir, osTypeResult := ParameterConvert(mode, master, node, softDir, label, osType)
			for i := 0; i < len(masterArray); i++ {
				if !CheckIP(masterArray[i]) {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "uninstall", "Result": "failure", "Info": "Uninstall failed! The Master IP address format you entered is incorrect, please check!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					return
				}
			}
			for j := 0; j < len(nodeArray); j++ {
				if !CheckIP(nodeArray[j]) {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "uninstall", "Result": "failure", "Info": "Uninstall failed! The Node IP address format you entered is incorrect, please check!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					return
				}
			}
			stu, sch := GetClusterStatus(label, currentDir, logName, mode)
			if stu != "installing" && stu != "restarting" && stu != "uninstalling" && sch != "on" {
				go UninstallCore(mode, master, masterArray, node, nodeArray, softDir, currentDir, kissh, subProcessDir, currentUser, label, osTypeResult, logName, CompatibleOS)
				if Lang == "cn" {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "uninstall", "Result": "success", "Info": "Kubernetes集群正在后台卸载中 ...", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
				} else {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "uninstall", "Result": "success", "Info": "Started to uninstall kubernetes cluster in the background ...", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
				}
			} else {
				c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "uninstall", "Result": "failure", "Info": "Uninstall failed! There are scheduled tasks in the background, or someone else is installing or uninstalling the current cluster!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
			}
		} else {
			c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "uninstall", "Result": "failure", "Info": "Uninstall failed! The parameter you input is wrong, please check!", "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
		}
	})

	// operation of delete install schedule
	router.POST("/deleteinstallschedule", func(c *gin.Context) {
		tools := c.Query("tools")
		label := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		osType := c.DefaultQuery("ostype", "")
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", softDir, label, "")
		err_installtime := DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/installtime.txt", "", currentDir, logName, mode)
		err_scheduler := DatabaseUpdate(currentDir+"/data/output"+subProcessDir+"/scheduler.txt", "", currentDir, logName, mode)
		var scheduleResult string
		var scheduleInfo string
		if err_installtime == nil && err_scheduler == nil {
			scheduleResult = "success"
			if Lang == "cn" {
				scheduleInfo = "定时安装计划任务已删除！"
			} else {
				scheduleInfo = "Success: scheduled installation task deleted successfully!"
			}
		} else {
			scheduleResult = "failure"
			if Lang == "cn" {
				scheduleInfo = "定时安装计划任务删除失败！"
			} else {
				scheduleInfo = "Failed: scheduled installation task deletion failed!"
			}
		}
		c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "deleteschedule", "Result": scheduleResult, "Info": scheduleInfo, "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
	})

        // operation of change password
        router.POST("/changepw", func(c *gin.Context) {
                var form passwordForm
                var oldPassword,newPassword1,newPassword2,changeResult,changeInfo string
                label := c.DefaultQuery("label", "")
                k8sVer := c.DefaultQuery("k8sver", "")
                softDir := c.DefaultQuery("softdir", "/opt/kube-install")
                osType := c.DefaultQuery("ostype", "")
                langFromWeb := c.Query("lang")
                Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
                if c.ShouldBind(&form) == nil {
                        pwdByte, err := ReadFile(currentDir+"/data/.key/admin")
                        if err != nil {
                            changeResult = "failure"
                            changeInfo = "System error! Please use the 'kube-install -init' command to reinitialize your system."
                            c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "updateuserinfo", "Result": changeResult, "Info": changeInfo, "Softdir": softDir, "Ostype": osType, "Tools": "yes", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
                            return
                        }
                        currentPassword, _ := DePwdCode(pwdByte)
                        oldPassword = form.Password
                        newPassword1 = form.Newpassword1
                        newPassword2 = form.Newpassword2
                        if oldPassword == string(currentPassword) && newPassword1 == newPassword2 {
                                key, _ := EnPwdCode([]byte(newPassword1))
                                err := DatabaseUpdate(currentDir+"/data/.key/admin", string(key), currentDir, logName, mode)
                                if err != nil {
                                        changeResult = "failure"
                                        changeInfo = "System error! Please use the 'kube-install -init' command to reinitialize your system."
                                } else {
                                        changeResult = "success"
                                        if Lang == "cn" {
                                                changeInfo = "用户信息更新成功！"
                                        } else {
                                                changeInfo = "User information updated successfully!"
                                        }
                                }
                        } else {
                                changeResult = "failure"
                                if Lang == "cn" {
                                        changeInfo = "用户信息更新失败，密码不匹配！"
                                } else {
                                        changeInfo = "User information updated failed. Passwords do not match!"
                                }
                        }
                } else {
                        changeResult = "failure"
                        if Lang == "cn" {
                                changeInfo = "用户信息更新失败，密码不能为空！"
                        } else {
                                changeInfo = "User information updated failed. Password cannot be empty!"
                        }
                }
                c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "updateuserinfo", "Result": changeResult, "Info": changeInfo, "Softdir": softDir, "Ostype": osType, "Tools": "yes", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
                return
	})

        // Set tools switch
        router.POST("/toolswitch", func(c *gin.Context) {
                var form ToolSwitchForm
                var sshTool, installTool, calendarTool, userTool string
                tools := c.Query("tools")
                label := c.DefaultQuery("label", "")
                k8sVer := c.DefaultQuery("k8sver", "")
                softDir := c.DefaultQuery("softdir", "/opt/kube-install")
                osType := c.DefaultQuery("ostype", "")
                langFromWeb := c.Query("lang")
                Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
                if c.ShouldBind(&form) == nil {
                        sshTool = form.Sshtool
                        installTool = form.Installtool
                        calendarTool = form.Calendartool
                        userTool = form.Usertool
                        err := SetToolSwitch(sshTool, installTool, calendarTool, userTool, currentDir, logName, mode)
                        if err != nil {
                                if Lang == "cn" {
                                        c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "toolswitch", "Result": "failure", "Info": "工具面板设置操作失败！", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
                                } else {
                                        c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "toolswitch", "Result": "failure", "Info": "Failed to set the tool panel!", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
                                }
                        } else {
                                if Lang == "cn" {
                                        c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "toolswitch", "Result": "success", "Info": "工具面板设置操作成功！", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
                                } else {
                                        c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "toolswitch", "Result": "success", "Info": "Tool panel setting operation succeeded!", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
                                }
                        }
                } else {
                        if Lang == "cn" {
                                c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "toolswitch", "Result": "failure", "Info": "设置失败！请确保你设置了正确的选项。", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
                        } else {
                                c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "toolswitch", "Result": "failure", "Info": "Setting failed! Please make sure you set the correct options.", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
                        }
                }
        })

	// Open the SSH key channel
	router.POST("/sshkey", func(c *gin.Context) {
		var form SshKeyForm
		var sshIp, sshPass string
		tools := c.Query("tools")
		label := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		osType := c.DefaultQuery("ostype", "")
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		if c.ShouldBind(&form) == nil {
			sshIp = form.Sship
			sshPass = form.Sshpass
			ipArray := strings.Split(sshIp, ",")
			for i := 0; i < len(ipArray); i++ {
				if !CheckIP(ipArray[i]) {
					if Lang == "cn" {
						c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "sshkey", "Result": "failure", "Info": "目标主机SSH打通失败！你输入的IP地址格式有误，请检查！", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					} else {
						c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "sshkey", "Result": "failure", "Info": "Failed to connect SSH channel, the IP address format you entered is incorrect, please check!", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
					}
					return
				}
			}
			err := SshKey(ipArray, sshPass, currentDir)
			if err != nil {
				if Lang == "cn" {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "sshkey", "Result": "failure", "Info": template.HTML("目标主机SSH打通失败！请使用\"root\"用户手动打通从kube-install到目标主机的SSH通道，或者在目标主机上执行以下命令后再次尝试打通：<br> <div class='cli_font' style='text-shadow: 0 0px; text-align:left; font-family: Droid Sans; font-size: 13px;'><br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[root@localhost ~]# &nbsp; sudo sed -i \"/PermitRootLogin/d\" /etc/ssh/sshd_config <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[root@localhost ~]# &nbsp; sudo sh -c \"echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config\" <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[root@localhost ~]# &nbsp; sudo sed -i \"/StrictHostKeyChecking/s/^#//; /StrictHostKeyChecking/s/ask/no/\" /etc/ssh/ssh_config <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[root@localhost ~]# &nbsp; sudo systemctl restart sshd <br><br> </div>"), "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
				} else {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "sshkey", "Result": "failure", "Info": template.HTML("Failed to connect SSH channel! Please use \"root\" user to manually open the SSH channel from the local host to the target host, or try to open the SSH channel again after executing the following command on the target host:<br> <div class='cli_font' style='text-shadow: 0 0px; text-align:left; font-family: Droid Sans; font-size: 13px;'><br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[root@localhost ~]# &nbsp; sudo sed -i \"/PermitRootLogin/d\" /etc/ssh/sshd_config <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[root@localhost ~]# &nbsp; sudo sh -c \"echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config\" <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[root@localhost ~]# &nbsp; sudo sed -i \"/StrictHostKeyChecking/s/^#//; /StrictHostKeyChecking/s/ask/no/\" /etc/ssh/ssh_config <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[root@localhost ~]# &nbsp; sudo systemctl restart sshd <br><br> </div>"), "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
				}
			} else {
				if Lang == "cn" {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "sshkey", "Result": "success", "Info": "已成功打通到目标主机(" + sshIp + ")的SSH通道！", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
				} else {
					c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "sshkey", "Result": "success", "Info": "Successfully open the SSH channel from Kube-Install to the target host (" + sshIp + ")!", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
				}
			}
		} else {
			if Lang == "cn" {
				c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "sshkey", "Result": "failure", "Info": "目标主机SSH打通失败！请使用“root”用户手动打通从Kube-Install到目标主机的SSH通道！", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
			} else {
				c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": "sshkey", "Result": "failure", "Info": "Failed to open SSH channel, please open SSH channel to target host manually!", "Softdir": softDir, "Ostype": osType, "Tools": tools, "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
			}
		}
	})

	// page of remote TTY
	router.GET("/tty", func(c *gin.Context) {
		var ttyResult, ttyInfo string
		ttyIP := c.DefaultQuery("ttyip", "")
		ttyOption := c.DefaultQuery("ttyoption", "")
		opt := c.DefaultQuery("opt", "")
		softDir := c.DefaultQuery("softdir", "/opt/kube-install")
		label := c.DefaultQuery("label", "")
		k8sVer := c.DefaultQuery("k8sver", "")
		osType := c.DefaultQuery("ostype", "")
		langFromWeb := c.Query("lang")
		Lang := ChangeLang(langFromWeb, currentDir, logName, mode)
		err := SwitchTTY(label, currentDir, ttyIP, ttyOption, logName, mode)
		if err == nil {
			ttyResult = "success"
			if ttyOption == "enable" {
				if Lang == "cn" {
					ttyInfo = "Web终端已成功开启！ (你可以使用浏览器访问 http://" + ttyIP + ":7681 )"
				} else {
					ttyInfo = "The Web terminal has been enabled!  (You can use your browser to access http://" + ttyIP + ":7681 )"
				}
			} else {
				if Lang == "cn" {
					ttyInfo = "Web终端已成功禁用！"
				} else {
					ttyInfo = "The Web terminal has been disabled !"
				}
			}
		} else {
			ttyResult = "failure"
			if Lang == "cn" {
				ttyInfo = "操作失败！请检查" + ttyIP + "的系统网络和端口占用情况！"
			} else {
				ttyInfo = "Operation failed! Please check the network and port occupation of " + ttyIP + " !"
			}
		}
		c.HTML(http.StatusOK, "optresult.tmpl", gin.H{"Label": label, "K8sver": k8sVer, "Opt": opt, "Result": ttyResult, "Info": ttyInfo, "Softdir": softDir, "Ostype": osType, "Tools": "no", "Lang": Lang, "Version": Version, "Releasedate": ReleaseDate, "Compatiblek8s": CompatibleK8S, "Compatibleos": CompatibleOS})
		return
	})

	// Run kube-install Daemon
	router.Run(listenIPandPort)

}
