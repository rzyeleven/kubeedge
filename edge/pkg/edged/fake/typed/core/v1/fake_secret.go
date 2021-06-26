package v1

import (
	"context"
	
	"github.com/kubeedge/kubeedge/edge/pkg/metamanager/client"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fakecorev1 "k8s.io/client-go/kubernetes/typed/core/v1/fake"
)

// FakePersistentVolumeClaims implements PersistentVolumeClaimInterface
type FakeSecrets struct {
	fakecorev1.FakeSecrets
	ns         string
	MetaClient client.CoreInterface
}

// Get takes name of the persistentVolumeClaim, and returns the corresponding persistentVolumeClaim object
func (c *FakeSecrets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *corev1.Secret, err error) {
	return c.MetaClient.Secrets(c.ns).Get(name)
}
