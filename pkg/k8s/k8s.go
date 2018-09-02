package k8s

import (
	"fmt"
	"os"

	corev1 "k8s.io/api/core/v1"
	extensionv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//client pool of k8s cluster context
var clientpool map[string]*Client

type Client struct {
	clientset *kubernetes.Clientset
	context   string
}

func NewClient(context string) (*Client, error) {
	client := new(Client)
	client.context = context

	kubeconfig, err := clientcmd.LoadFromFile(os.Getenv("HOME") + "/.kube/config")
	if err != nil {
		panic(err)
	}
	if kubeconfig == nil || kubeconfig.Contexts == nil {
		panic("something's wrong with kube config file")
	}

	if _, ok := kubeconfig.Contexts[context]; !ok {
		return nil, fmt.Errorf("do not have the context:[%s] in kube config file", context)
	}

	kubeconfig.CurrentContext = context

	config, err := clientcmd.NewDefaultClientConfig(*kubeconfig, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	client.clientset = clientset
	return client, nil
}

func GetClient(context string) (*Client, error) {
	if client, ok := clientpool[context]; ok {
		return client, nil
	}
	return nil, fmt.Errorf("did not find the client for context: %s", context)
}

func CreateRabbitMQService(context string) error {
	svc := &corev1.Service{}
	svc.Name = "rabbitmq-svc"
	svc.Namespace = "default"
	svc.Spec.Selector = map[string]string{
		"app": "rabbitmq",
	}
	svc.Spec.Ports = make([]corev1.ServicePort, 1)
	svc.Spec.Ports[0].Port = int32(5672)
	svc.Spec.Ports[0].TargetPort = intstr.IntOrString{IntVal: int32(5672)}

	client, err := GetClient(context)
	if err != nil {
		return err
	}
	_, err = client.clientset.CoreV1().Services("default").Create(svc)

	return err
}

func CreateRabbitMQReplicaset(context string) error {
	rs := &extensionv1beta1.ReplicaSet{}
	rs.Labels["app"] = "rabbitmq"
	if rs.Spec.Template.Labels == nil {
		rs.Spec.Template.Labels = make(map[string]string)
	}
	rs.Spec.Template.Labels["app"] = "rabbitmq"

	client, err := GetClient(context)
	if err != nil {
		return err
	}
	_, err = client.clientset.ExtensionsV1beta1().ReplicaSets("default").Create(rs)

	return err
}

func CreateCIJob(repo, dependances, build, context string) error {
	return nil
}
