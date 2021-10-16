package kilib

import (
    "fmt"
)


// Displays the detailed version information of the kube-install.
func ShowVersion(Version string, ReleaseDate string, CompatibleK8S string, CompatibleOS string){
    fmt.Println("[Version]\n    Version: "+Version+"\n    Release Date: "+ReleaseDate+"\n\n[Compatibility] \n    Kubernetes: "+CompatibleK8S+"\n    OS: "+CompatibleOS+"\n    Hardware: x86 | amd64 \n")
}



