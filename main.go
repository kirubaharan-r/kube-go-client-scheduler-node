package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	//"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var cout int = 0
var pod1 Pod

type Pod struct {
	name      string
	namespace string
	ip        string
	status    string
}

func main() {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//config, err := rest.InClusterConfig()

	if err != nil {
		fmt.Println(os.Getenv("GOOS"))
		panic(err.Error())
	}

	ctx := context.Background()
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	poes, err := clientset.CoreV1().Pods(os.Args[3]).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(len(poes.Items))
	//fmt.Println("Pod_Name  Status   life_status")

	for _, tes := range poes.Items {
		var a string = string(tes.Status.Phase)
		//var b int = int(len(poes.Items))
		cmd := exec.Command("kubectl", "label", "nodes", "--all", "kiru=be")
		cmd.Run() //kubectl label nodes --all hello=bye
		c, _ := strconv.Atoi(os.Args[1])
		exr, _ := exec.LookPath("cmd")
		for _, dem := range poes.Items {

			pod1.name = string(dem.Name)
			pod1.namespace = string(dem.Namespace)
			pod1.ip = string(dem.Name)
			pod1.status = string(dem.Status.Phase)

			fmt.Println(pod1, dem.)

			if string(dem.Status.Phase) == "Running" {
				cout++
			}

		}
		fmt.Print(cout)
		//exr, _ := exec.LookPath("/usr/local/bin/")  //for linux

		if cout <= c {
			fmt.Println(tes.Name, "-", a, "-", "schdeuled pod")
			schdeuler := exec.Command(exr, "/K", "kubectl taint node -l "+os.Args[2]+" "+os.Args[2]+":NoSchedule-")
			//schdeuler := exec.Command(exr, "kubectl", "taint", "node", "-l", os.Args[1], os.Args[1]+":NoSchedule-") //for linux
			schdeuler.Run()
			//fmt.Println(schdeuler)

		} else {
			fmt.Println(tes.Name, "-", a, "-", "unschdeuled pod")
			schdeuler := exec.Command(exr, "/K", "kubectl taint node -l "+os.Args[2]+" "+os.Args[2]+":NoSchedule")
			//schdeuler := exec.Command(exr, "kubectl", "taint", "node", "-l", os.Args[1], os.Args[1]+":NoSchedule") //for linux
			schdeuler.Run()
			//fmt.Println(schdeuler)
		}
		break

	}

}

