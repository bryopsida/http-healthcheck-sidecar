package kubernetes

import (
	"context"
	"strings"

	"github.com/bryopsida/http-healthcheck-sidecar/config"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func returnOverrideState() bool {
	return config.GetStatusOverrideState()
}

func fetchPodState(podName string) *v1.Pod {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pod, err := clientset.CoreV1().Pods("").Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	return pod
}

func IsPodHealthy(podName string) bool {
	if config.IsStatusOverriden() {
		return config.GetStatusOverrideState()
	} else {
		podState := fetchPodState(podName)
		for _, containerStatus := range podState.Status.ContainerStatuses {
			if !containerStatus.Ready {
				return false
			}
		}
		return true
	}
}

func IsPodContainerHealthy(podName string, containerName string) bool {
	if config.IsStatusOverriden() {
		return config.GetStatusOverrideState()
	} else {
		podState := fetchPodState(podName)
		for _, containerStatus := range podState.Status.ContainerStatuses {
			if strings.ToLower(containerStatus.Name) == strings.ToLower(containerName) {
				return containerStatus.Ready
			}
		}
		return false
	}
}
