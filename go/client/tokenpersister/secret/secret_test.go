package secret_test

import (
	"context"
	"testing"

	"github.com/metal-stack/api/go/client/tokenpersister/secret"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/stretchr/testify/require"
)

func TestNewPersisterFunc_validation(t *testing.T) {
	clientset := fake.NewSimpleClientset()

	_, err := secret.NewPersistTokenFunc("", "secret", "key", clientset)
	require.Error(t, err)

	_, err = secret.NewPersistTokenFunc("ns", "", "key", clientset)
	require.Error(t, err)

	_, err = secret.NewPersistTokenFunc("ns", "secret", "", clientset)
	require.Error(t, err)

	_, err = secret.NewPersistTokenFunc("ns", "secret", "key", nil)
	require.Error(t, err)
}

func TestPersister_persist_withExistingSecret(t *testing.T) {
	existing := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-secret",
			Namespace: "my-namespace",
		},
		Data: map[string][]byte{
			"other": []byte("value"),
		},
	}
	clientset := fake.NewSimpleClientset(existing)

	persister, err := secret.NewPersistTokenFunc("my-namespace", "my-secret", "token", clientset)
	require.NoError(t, err)

	err = persister("new-token")
	require.NoError(t, err)

	s, err := clientset.CoreV1().Secrets("my-namespace").Get(context.Background(), "my-secret", metav1.GetOptions{})
	require.NoError(t, err)
	require.Equal(t, "new-token", string(s.Data["token"]))
	require.Equal(t, "value", string(s.Data["other"]))
}

func TestPersister_persist_withNilData(t *testing.T) {
	existing := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-secret",
			Namespace: "my-namespace",
		},
		Data: nil,
	}
	clientset := fake.NewSimpleClientset(existing)

	persister, err := secret.NewPersistTokenFunc("my-namespace", "my-secret", "token", clientset)
	require.NoError(t, err)

	err = persister("new-token")
	require.NoError(t, err)

	s, err := clientset.CoreV1().Secrets("my-namespace").Get(context.Background(), "my-secret", metav1.GetOptions{})
	require.NoError(t, err)
	require.Equal(t, "new-token", string(s.Data["token"]))
}

func TestPersister_persist_secretNotFound(t *testing.T) {
	clientset := fake.NewSimpleClientset()

	persister, err := secret.NewPersistTokenFunc("my-namespace", "missing", "token", clientset)
	require.NoError(t, err)

	err = persister("new-token")
	require.Error(t, err)
}
