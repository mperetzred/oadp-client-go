package velero

import (
	"fmt"
	"log"
	"strings"

	"github.com/mperetzred/oadp-client-go/extended/common"
	corev1 "k8s.io/api/core/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

type PodsGetter interface {
	Pods(namespace string) VeleroPodsInterface
}

type VeleroPodsInterface interface {
	common.PodsExpansionInterface
}

type veleropods struct {
	*common.PodsExpansion
}

func newDefaultVeleroPods(pods corev1client.PodInterface) *veleropods {
	podLogOpts := corev1.PodLogOptions{
		Container: "velero",
	}

	return &veleropods{
		&common.PodsExpansion{
			&podLogOpts,
			"app.kubernetes.io/name=velero",
			pods,
		},
	}
}

func (vp *veleropods) GetLogsByLevel(level string) []string {
	containerLogs, err := vp.GetLogsToString()
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

func (vp *veleropods) GetErrorLogs(level string) []string {
	return vp.GetLogsByLevel("error")
}
