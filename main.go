package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"context"
	"flag"
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

var name string
var id string
var label string
var pod1 Pod
var pars1 pars
var kubeconfig *string

type Pod struct {
	name      string
	namespace string
	ip        string
	status    string
}

type pars struct {
}

func main() {
	r := gin.Default()

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/user/:id/:name/:label/", act)

	r.Run(":8084")

}
func act(c *gin.Context) {
	var cout int
	name = c.Param("name")
	id = c.Param("id")
	label = c.Param("label")
	c.String(http.StatusOK, id)
	c.String(http.StatusOK, name)
	c.String(http.StatusOK, label)
	fmt.Print(name)

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
	poes, err := clientset.CoreV1().Pods(string(label)).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(len(poes.Items))
	//fmt.Println("Pod_Name  Status   life_status")

	for _, tes := range poes.Items {
		var _ string = string(tes.Status.Phase)
		//var b int = int(len(poes.Items))
		cmd := exec.Command("kubectl", "label", "nodes", "--all", label)
		cmd.Run() //kubectl label nodes --all hello=bye
		c, _ := strconv.Atoi(id)
		exr, _ := exec.LookPath("cmd")
		for _, dem := range poes.Items {

			pod1.name = string(dem.Name)
			pod1.namespace = string(dem.Namespace)
			pod1.ip = string(dem.Status.PodIP)
			pod1.status = string(dem.Status.Phase)

			fmt.Println(pod1.name, pod1.namespace, pod1.ip, pod1.status)

			if string(dem.Status.Phase) == "Running" {
				cout++
			}

		}
		fmt.Print(cout)
		//exr, _ := exec.LookPath("/usr/local/bin/")  //for linux

		if cout <= c {
			fmt.Println("\n ------------")
			fmt.Println("\n making schdeule pods")
			fmt.Println("\n ------------")
			schdeuler := exec.Command(exr, "/K", "kubectl taint node -l "+label+" "+label+":NoSchedule-")
			//schdeuler := exec.Command(exr, "kubectl", "taint", "node", "-l", os.Args[1], os.Args[1]+":NoSchedule-") //for linux
			schdeuler.Run()
			//fmt.Println(schdeuler)

		} else {
			fmt.Println("\n ------------")
			fmt.Println("\n making unschdeule pods")
			fmt.Println("\n ------------")
			schdeuler := exec.Command(exr, "/K", "kubectl taint node -l "+label+" "+label+":NoSchedule")
			//schdeuler := exec.Command(exr, "kubectl", "taint", "node", "-l", os.Args[1], os.Args[1]+":NoSchedule") //for linux
			schdeuler.Run()
			//fmt.Println(schdeuler)
		}
		break

	}
}
