package kilib

import (
    "net"
)


func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}

func CheckIP(ipv4 string) {
    address := net.ParseIP(ipv4)  
    if address == nil {
         panic("The format of IP address you entered is wrong, please check! \n--------------------------------------------------------\n")
    }
}

func CheckOS(osType string) (string) {
    switch {
      case osType == "centos7" :
          return "rhel7"
      case osType == "rhel7" :
          return "rhel7"
      case osType == "centos8" :
          return "rhel8"
      case osType == "rhel8" :
          return "rhel8"
      case osType == "suse15" :
          return "suse15"
      default:
          panic("Please make sure that the \"-ostype\" parameter you entered is correct! Only support rhel7, rhel8, centos7, centos8, suse15 these types of \"ostype\": \n--------------------------------------------------------\n    rhel7   --> Red Hat Enterprise Linux 7 \n    rhel8   --> Red Hat Enterprise Linux 8 \n    centos7 --> CentOS Linux 7 \n    centos8 --> CentOS Linux 8 \n    suse15  --> OpenSUSE Linux 15 \n\n ")
    }
}

func CheckParam(option string, paramname string, param string) {
    if param == "" {
         panic("When you execute the "+option+" operation, you must enter the "+paramname+" parameter, please check! \n--------------------------------------------------------\n")
    }
}


