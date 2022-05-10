package versioned

import (
	snapshotclientset "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned"
	snapshotV1 "github.com/mperetzred/oadp-client-go/extended/velero/snapshot/v1"
	snapshotV1beta1 "github.com/mperetzred/oadp-client-go/extended/velero/snapshot/v1beta1"
	"k8s.io/client-go/rest"
)

type Interface interface {
	SnapshotV1beta1() snapshotV1beta1.SnapshotV1beta1Interface
	SnapshotV1() snapshotV1.SnapshotV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	snapshotV1beta1 *snapshotV1beta1.SnapshotV1beta1Client
	snapshotV1      *snapshotV1.SnapshotV1Client
}

// SnapshotV1beta1 retrieves the SnapshotV1beta1Client
func (c *Clientset) SnapshotV1beta1() snapshotV1beta1.SnapshotV1beta1Interface {
	return c.snapshotV1beta1
}

// SnapshotV1 retrieves the SnapshotV1Client
func (c *Clientset) SnapshotV1() snapshotV1.SnapshotV1Interface {
	return c.snapshotV1
}

func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	snapshotClientset, err := snapshotclientset.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return &Clientset{
		&snapshotV1beta1.SnapshotV1beta1Client{
			snapshotClientset.SnapshotV1beta1(),
		},
		&snapshotV1.SnapshotV1Client{
			snapshotClientset.SnapshotV1(),
		},
	}, nil
}
