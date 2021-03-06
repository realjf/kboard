package resource

import "kboard/internal"

type IResClusterRoleBinding interface {
	internal.IResource
	SetMetadataName(string) error
}

type ResClusterRoleBinding struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name string
	}
	Subjects []internal.RoleBindingSubject
	RoleRef  internal.RoleRef
}

func NewResClusterRoleBinding() *ResClusterRoleBinding {
	return &ResClusterRoleBinding{
		ApiVersion: "rbac.authorization.k8s.io/v1",
		Kind:       internal.RESOURCE_CLUSTER_ROLE_BINDING,
		Metadata: struct {
			Name string
		}{Name: ""},
		Subjects: []internal.RoleBindingSubject{},
		RoleRef:  internal.RoleRef{},
	}
}
