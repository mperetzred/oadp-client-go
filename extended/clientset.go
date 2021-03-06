package extended

import (
	"fmt"
	"net/http"
	"sync"

	veleroexpansion "github.com/mperetzred/oadp-client-go/extended/velero"
	oadpv1alpha1 "github.com/mperetzred/oadp-client-go/generated/oadp/clientset/versioned/typed/v1alpha1"

	//operators "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var once sync.Once
var oadpclient *Clientset

type Interface interface {
	VeleroClient() veleroexpansion.VeleroV1Interface
	oadpv1alpha1.OadpV1alpha1Interface
}

type Clientset struct {
	veleroclient *veleroexpansion.VeleroV1Client
	*oadpv1alpha1.OadpV1alpha1Client
}

func (c *Clientset) VeleroClient() veleroexpansion.VeleroV1Interface {
	return c.veleroclient
}

func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	veleroclient, err := veleroexpansion.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	oadpclient, err := oadpv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	return &Clientset{
		veleroclient,
		oadpclient,
	}, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

func GetOadpClient() (Clientset, error) {
	var err error
	once.Do(func() {
		oadpclient, err = NewForConfig(config.GetConfigOrDie())
	})
	return *oadpclient, err
}
