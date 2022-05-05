package client

import (
	snapshotclientset "github.com/kubernetes-csi/external-snapshotter/client/v6/clientset/versioned"
	oadpv1alpha1 "github.com/mperetzred/oadp-client-go/generated/oadp/clientset/versioned/typed/v1alpha1"
	velerov1 "github.com/mperetzred/oadp-client-go/generated/velero/clientset/versioned"
	appsv1 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	configv1 "github.com/openshift/client-go/config/clientset/versioned"
	secv1 "github.com/openshift/client-go/security/clientset/versioned/typed/security/v1"
	operators "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

type OadpClient interface {
	VeleroClient() velerov1.Interface
	SnapshotClient() snapshotclientset.Interface
	OcpConfigV1Client() configv1.Interface
	Config() *rest.Config
	SecClient() secv1.SecurityV1Interface
	OperatorClient() operators.Interface
	kubernetes.Interface
}

type oadp struct {
	kubeconfig      string
	config          *rest.Config
	configv1Client  *configv1.Clientset
	secClient       *secv1.SecurityV1Client
	appsClient      *appsv1.AppsV1Client
	veleroClient    *velerov1.Clientset
	snapshotClient  *snapshotclientset.Clientset
	operatorsClient *operators.Clientset
	*kubernetes.Clientset
	*oadpv1alpha1.OadpV1alpha1Client
}

func (o oadp) Config() *rest.Config {
	return o.config
}

func (o oadp) SnapshotClient() snapshotclientset.Interface {
	return o.snapshotClient
}

func (o oadp) VeleroClient() velerov1.Interface {
	return o.veleroClient
}

func (o oadp) OcpConfigV1Client() configv1.Interface {
	return o.configv1Client
}

func (o oadp) SecClient() secv1.SecurityV1Interface {
	return o.secClient
}

func (o oadp) OperatorClient() operators.Interface {
	return o.operatorsClient
}

func GetOadpSubresourceClient(master string, kubeconfig string) (OadpClient, error) {
	restConfig := config.GetConfigOrDie()
	coreClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	configClient, err := configv1.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	secClient, err := secv1.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	appsClient, err := appsv1.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	veleroClient, err := velerov1.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	snapshotClient, err := snapshotclientset.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	oadpClient, err := oadpv1alpha1.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	operatorsClient, err := operators.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	return &oadp{
		kubeconfig,
		restConfig,
		configClient,
		secClient,
		appsClient,
		veleroClient,
		snapshotClient,
		operatorsClient,
		coreClient,
		oadpClient,
	}, nil
}
