package velero

import (
	"context"
	"fmt"
	"log"

	velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	velerov1client "github.com/vmware-tanzu/velero/pkg/generated/clientset/versioned/typed/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupsGetter has a method to return a BackupInterface.
// A group's client should implement this interface.
type BackupsGetter interface {
	Backups(namespace string) BackupExpansionInterface
}

type BackupExpansionInterface interface {
	IsBackupDone(name string) (bool, error)
	IsBackupCompletedSuccessfully(name string) (bool, error)
	velerov1client.BackupInterface
}

type backupexpansion struct {
	velerov1client.BackupInterface
}

func newBackupExpansion(backupInterface velerov1client.BackupInterface) *backupexpansion {
	return &backupexpansion{
		backupInterface,
	}
}

func (b *backupexpansion) IsBackupCompletedSuccessfully(name string) (bool, error) {
	backup, err := b.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}
	if backup.Status.Phase == velerov1.BackupPhaseCompleted {
		return true, nil
	}
	return false, fmt.Errorf("backup phase is: %s; expected: %s\nvalidation errors: %v",
		backup.Status.Phase, velerov1.BackupPhaseCompleted, backup.Status.ValidationErrors)
}

func (b *backupexpansion) IsBackupDone(name string) (bool, error) {
	backup, err := b.Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return false, err
	}
	if len(backup.Status.Phase) > 0 {
		log.Printf(fmt.Sprintf("backup phase: %s\n", backup.Status.Phase))
	}
	if backup.Status.Phase != "" && backup.Status.Phase != velerov1.BackupPhaseNew && backup.Status.Phase != velerov1.BackupPhaseInProgress {
		return true, nil
	}
	return false, nil
}
