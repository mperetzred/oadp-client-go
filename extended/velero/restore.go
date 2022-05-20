package velero

import (
	"context"
	"fmt"
	"log"

	velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	velerov1client "github.com/vmware-tanzu/velero/pkg/generated/clientset/versioned/typed/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RestoresGetter has a method to return a RestoreExpansionInterface.
// A group's client should implement this interface.
type RestoresGetter interface {
	Restores(namespace string) RestoreExpansionInterface
}

type RestoreExpansionInterface interface {
	IsRestoreDone(name string) (bool, error)
	IsRestoreCompletedSuccessfully(name string) (bool, error)
	velerov1client.RestoreInterface
}

type restoreexpansion struct {
	velerov1client.RestoreInterface
}

func newRestoreExpansion(restoreInterface velerov1client.RestoreInterface) *restoreexpansion {
	return &restoreexpansion{
		restoreInterface,
	}
}

func (r *restoreexpansion) IsRestoreCompletedSuccessfully(name string) (bool, error) {
	restore, err := r.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}
	if restore.Status.Phase == velerov1.RestorePhaseCompleted {
		return true, nil
	}
	return false, fmt.Errorf("backup phase is: %s; expected: %s\nvalidation errors: %v",
		restore.Status.Phase, velerov1.BackupPhaseCompleted, restore.Status.ValidationErrors)
}

func (r *restoreexpansion) IsRestoreDone(name string) (bool, error) {
	restore, err := r.Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return false, err
	}
	if len(restore.Status.Phase) > 0 {
		log.Printf(fmt.Sprintf("backup phase: %s\n", restore.Status.Phase))
	}
	if restore.Status.Phase != "" && restore.Status.Phase != velerov1.RestorePhaseNew &&
		restore.Status.Phase != velerov1.RestorePhaseInProgress {
		return true, nil
	}
	return false, nil
}
