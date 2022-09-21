package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"k8s.io/client-go/rest"
	// "k8s.io/client-go/tools/clientcmd"
	// "k8s.io/client-go/util/homedir"
)

func main() {
	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()

	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	config, err := rest.InClusterConfig()

	if err != nil {
		//fmt.Println(os.Getenv("GOOS"))
		panic(err.Error())
	}

	ctx := context.Background()
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	poes, err := clientset.CoreV1().Pods(os.Args[2]).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(len(poes.Items))
	//fmt.Println("Pod_Name  Status   life_status")

	for _, tes := range poes.Items {
		var a string = string(tes.Status.Phase)
		var b int = int(len(poes.Items))
		cmd := exec.Command("kubectl", "label", "nodes", "--all", os.Args[1])
		cmd.Run() //kubectl label nodes --all hello=bye
		c, _ := strconv.Atoi(os.Args[3])
		exr, _ := exec.LookPath("/usr/local/bin/")

		if b <= c {
			fmt.Println("schdeuleing pod")
			schdeuler := exec.Command(exr, "kubectl", "taint", "node", "-l", os.Args[1], os.Args[1]+":NoSchedule-")
			stdout, _ := schdeuler.StdoutPipe()
			schdeuler.Run()
			fmt.Println(schdeuler, stdout)

		} else {
			fmt.Println("unschdeuleing pod")
			schdeuler := exec.Command(exr, "kubectl", "taint", "node", "-l", os.Args[1], os.Args[1]+":NoSchedule")
			stdout, _ := schdeuler.StdoutPipe()
			schdeuler.Run()
			fmt.Println(schdeuler, stdout)
		}
		break

	}

}
