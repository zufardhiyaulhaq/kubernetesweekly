package models

import (
	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
	"gopkg.in/yaml.v2"
)

type Weekly struct {
	APIVersion string                       `yaml:"apiVersion"`
	Kind       string                       `yaml:"kind"`
	Metadata   WeeklyMetadata               `yaml:"metadata"`
	Spec       communityv1alpha1.WeeklySpec `yaml:"spec"`
}

type WeeklyMetadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

func (w Weekly) ToYaml() ([]byte, error) {
	return yaml.Marshal(w)
}
