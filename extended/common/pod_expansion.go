package common

import (
	"bytes"
	"context"
	"io"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

type PodsExpansionInterface interface {
	ListByLabelSelector() (*corev1.PodList, error)
	GetLogsToString() (string, error)
	PodRunning() (bool, error)
}

type PodsExpansion struct {
	Opts          *corev1.PodLogOptions
	LabelSelector string
	corev1client.PodInterface
}

func (p *PodsExpansion) ListByLabelSelector() (*corev1.PodList, error) {
	// get pods in test namespace with labelSelector
	podList, err := p.List(context.TODO(), metav1.ListOptions{LabelSelector: p.LabelSelector})
	if err != nil {
		return nil, err
	}
	return podList, nil
}

// Returns logs from velero container on velero pod
func (p *PodsExpansion) GetLogsToString() (string, error) {
	podList, err := p.ListByLabelSelector()
	if err != nil {
		return "", err
	}

	var logs string
	for _, podInfo := range (*podList).Items {
		req := p.GetLogs(podInfo.Name, p.Opts)
		podLogs, err := req.Stream(context.TODO())
		if err != nil {
			return "", err
		}
		defer podLogs.Close()
		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, podLogs)
		if err != nil {
			return "", err
		}
		logs = buf.String()
	}
	return logs, nil
}

func (p *PodsExpansion) PodRunning() (bool, error) {
	podList, err := p.ListByLabelSelector()
	if err != nil || len(podList.Items) == 0 {
		return false, err
	}
	for _, podInfo := range (*podList).Items {
		if podInfo.Status.Phase != corev1.PodRunning {
			log.Printf("pod: %s is not yet running with status: %v", podInfo.Name, podInfo.Status)
			return false, nil
		}
	}

	return true, nil

}
