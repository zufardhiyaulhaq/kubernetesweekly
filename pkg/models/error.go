package models

type KubeweeklyNameNotFoundError struct{}

func (k *KubeweeklyNameNotFoundError) Error() string {
	return "Kubeweekly name not found"
}
