package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"

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
	//corev1 "k8s.io/client-go/applyconfigurations/core/v1"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//swagger impemetations
)

var name, id, label, value, key, remove string
var pod1 Pod
var test1 test
var pars1 pars
var kubeconfig *string

type Pod struct {
	name      string
	namespace string
	ip        string
	status    string
}
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}
type test struct {
	name string
}

type pars struct {
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type NodeSpecApplyConfiguration struct {
}

func main() {
	r := gin.Default()

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	r.GET("/kiruba", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/user/:id/:name/:label/", act)
	r.POST("/label/node/:key/:value/:remove", lab)
	r.POST("/nodes/scheduleble/:names", nodes)

	r.Run(":8084")
	r.Use(act)

}

func act(c *gin.Context) {

	var cout int
	// var noo Pod
	name = c.Param("name")
	id = c.Param("id")
	label = c.Param("label")

	//c.String(http.StatusOK, id)
	//c.String(http.StatusOK, name)
	//c.String(http.StatusOK, label)

	//c.JSON(http.StatusOK, gin.H{"name": noo.name, "namespace": noo.namespace, "ip": noo.ip, "status": noo.status})

	c.String(http.StatusBadRequest, http.MethodGet, http.MethodTrace, "please use correct path for using the scheduler")
	fmt.Println(name)

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
	poes, err := clientset.CoreV1().Pods(string("kube-system")).List(ctx, metav1.ListOptions{})
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
		c, err := strconv.Atoi(id)
		//c, err := strconv.ParseInt(id)
		fmt.Println(c)
		fmt.Println(err)

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

		//fmt.Print(cout)
		//exr, _ := exec.LookPath("/usr/local/bin/")  //for linux
		if runtime.GOOS == "windows" {
			exr, _ := exec.LookPath("cmd")
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
			//break
		} else {
			exr, _ := exec.LookPath("/usr/local/bin/")
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
			//break
		}

	}
	z := len(poes.Items)
	c.JSON(http.StatusOK, gin.H{"name": pod1.name, "namespace": pod1.namespace, "ip": pod1.ip, "status": pod1.status})

	c.String(http.StatusOK, "  The update value of allocation size is  "+id+" .And the current Running size is ", z)
}

func lab(c *gin.Context) {
	key = c.Param("key")
	value = c.Param("value")
	remove = c.Param("remove")
	if runtime.GOOS == "windows" {
		exr, _ := exec.LookPath("cmd")
		if remove == "-" {
			//exr, _ := exec.LookPath("/usr/local/bin/")
			fmt.Print(key)
			fmt.Print(value)

			schdeuler := exec.Command(exr, "/K", "kubectl label node --all "+key+"="+value)
			schdeuler.Run()
		} else {

			//exr, _ := exec.LookPath("/usr/local/bin/")
			fmt.Print(key)
			fmt.Print(value)

			schdeuler := exec.Command(exr, "/K", "kubectl label node --all "+key+"-")
			schdeuler.Run()
		}
	} else {
		exr, _ := exec.LookPath("/usr/local/bin/")
		if remove == "-" {

			schdeuler := exec.Command(exr, "/K", "kubectl label node --all "+key+"-")
			schdeuler.Run()
		} else {

			//exr, _ := exec.LookPath("cmd")

			schdeuler := exec.Command(exr, "/K", "kubectl label node --all "+key+"="+value)
			schdeuler.Run()
		}
	}
}

func nodes(c *gin.Context) {

	c.JSON(test1.name)
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	ctx := context.TODO()
	clientset, _ := kubernetes.NewForConfig(config)

	test, _ := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{}) //TimeoutSeconds: func() *int64 { i := int64(3); return &i }()
	//name := "docker-desktop"
	//make, _ := clientset.CoreV1().Nodes().Apply(ctx, *corev1.NodeApplyConfiguration{Spec.Unschedulable: true}, metav1.ApplyOptions{Force: true})
	//node, _ := clientset.CoreV1().Nodes().Update(ctx, P, metav1.UpdateOptions{FieldManager: "node_field_mgr"})

	//panic(err)
	fmt.Println(test.Items)
	//c.JSON(test.name)
	//fmt.Println(make, node)
}
