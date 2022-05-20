package velero

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	rest "k8s.io/client-go/rest"
)

// PodsGetter has a method to return a PodInterface.
// A group's client should implement this interface.
type PodsGetter interface {
	Pods(namespace string) PodInterface
}

// PodInterface has methods to work with Pod resources.
type PodInterface interface {
	List(opts metav1.ListOptions) (*v1.PodList, error)
	GetLogsByLevel(level string) []string
	GetErrorLogs() []string
}

// pods implements PodInterface
type pods struct {
	client rest.Interface
	ns     string
}

// newPods returns a Pods
func newPods(c *corev1client.CoreV1Client, namespace string) *pods {
	return &pods{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// List takes label and field selectors, and returns the list of Pods that match those selectors.
func (c *pods) List(opts metav1.ListOptions) (result *v1.PodList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PodList{}
	opts.LabelSelector = "app.kubernetes.io/name=velero"
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

func (p *pods) GetLogsByLevel(level string) []string {
	containerLogs, err := p.GetLogsToString()
	if err != nil {
		log.Printf("cannot get velero container logs")
		return nil
	}
	containerLogsArray := strings.Split(containerLogs, "\n")
	var arr = []string{}
	for i, line := range containerLogsArray {
		if strings.Contains(line, "level="+level) {
			arr = append(arr, fmt.Sprintf("velero container %s line#%d: "+line+"\n", level, i))
		}
	}
	return arr
}

func (p *pods) GetErrorLogs() []string {
	return p.GetLogsByLevel("error")
}

// Returns logs from velero container on velero pod
func (p *pods) GetLogsToString() (string, error) {
	podList, err := p.List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	podOpts := v1.PodLogOptions{Container: "velero"}
	var logs string
	for _, podInfo := range (*podList).Items {
		req := p.client.Get().Namespace(p.ns).Name(podInfo.Name).Resource("pods").SubResource("log").VersionedParams(&podOpts, scheme.ParameterCodec)
		podLogs, err := req.Stream(context.Background())
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

func (p *pods) PodRunning() (bool, error) {
	podList, err := p.List(metav1.ListOptions{})
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
