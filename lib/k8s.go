package kilib

import (
	"context"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//kubernetes Clientset
func k8sClientSet(currentDir string, subProcessDir string, logName string, mode string) *kubernetes.Clientset {
	var kubeconfig string
	kubeconfig = filepath.Join(currentDir + "/data/output" + subProcessDir + "/cert/ssl/kube-install.kubeconfig")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	CheckErr(err, currentDir, logName, mode)
	clientset, err := kubernetes.NewForConfig(config)
	CheckErr(err, currentDir, logName, mode)
	return clientset
}

//Detect kubernetes cluster health
func DetectK8sHealth(label string, currentDir string, logName string, mode string) error {
	var err error
	_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
	kubecfg, err_kcfg := ReadFile(currentDir + "/data/output" + subProcessDir + "/cert/ssl/kube-install.kubeconfig")
	if err_kcfg == nil {
		if kubecfg == "" {
			err = fmt.Errorf("kubeconfig did not return any value!")
		} else {
			_, err = k8sClientSet(currentDir, subProcessDir, logName, mode).CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
		}
	} else {
		err = err_kcfg
	}
	return err
}

//Query node list of the kubernetes cluster
func ListNode(label string, currentDir string, logName string, mode string) []string {
	_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
	kubecfg, err := ReadFile(currentDir + "/data/output" + subProcessDir + "/cert/ssl/kube-install.kubeconfig")
	if kubecfg == "" || err != nil {
		return nil
	} else {
		var nodeArray []string
		todo := context.Background()
		list, err := k8sClientSet(currentDir, subProcessDir, logName, mode).CoreV1().Nodes().List(todo, metav1.ListOptions{})
		CheckErr(err, currentDir, logName, mode)
		for _, d := range list.Items {
			nodeArray = append(nodeArray, d.Name)
		}
		return nodeArray
	}
}

//Query node info of the kubernetes cluster
func GetNodeInfo(label string, nodeIP string, currentDir string, logName string, mode string) (string, map[string]string, string, string, string, string, string, string) {
	_, _, _, subProcessDir, _ := ParameterConvert(mode, "", "", "", label, "")
	kubecfg, err := ReadFile(currentDir + "/data/output" + subProcessDir + "/cert/ssl/kube-install.kubeconfig")
	if kubecfg == "" || err != nil {
		return "", nil, "", "", "", "", "", ""
	} else {
		todo := context.Background()
		info, err := k8sClientSet(currentDir, subProcessDir, logName, mode).CoreV1().Nodes().Get(todo, nodeIP, metav1.GetOptions{})
		CheckErr(err, currentDir, logName, mode)
		runStatus := "Unknow"
		infoConditionsLen := len(info.Status.Conditions)
		if infoConditionsLen > 0 {
			runStatus = fmt.Sprintf("%s", info.Status.Conditions[infoConditionsLen-1].Status)
		}
		nodeLabels := info.ObjectMeta.Labels
		runcVer := info.Status.NodeInfo.ContainerRuntimeVersion
		kernelVer := info.Status.NodeInfo.KernelVersion
		cpu := info.Status.Capacity.Cpu().String()
		osVer := info.Status.NodeInfo.OSImage
		memory := info.Status.Allocatable.Memory().String()
		createTime := fmt.Sprintf("%s", info.CreationTimestamp)
		return runStatus, nodeLabels, runcVer, osVer, kernelVer, cpu, memory, createTime
	}
}
