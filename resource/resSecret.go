package resource

import (
	"encoding/base64"
	"errors"
	"gopkg.in/yaml.v2"
)

type IResSecret interface {
	IResource
	SetMetaDataName(name string) bool
	GetMetaDataName() string
	SetType(typeName string) bool
	SetData([]map[string]string) error
	GetData() (map[string]string, error)
	SetNamespace(string) bool
}

type ResSecret struct {
	Kind       string
	ApiVersion string `yaml:"apiVersion"`
	Type       string
	MetaData   struct {
		Name      string
		Namespace string
	}
	Data map[string]string
}

func NewSecret() *ResSecret {
	return &ResSecret{
		Kind:       "Secret",
		ApiVersion: "v1",
		Type:       "Opaque",
		MetaData: struct {
			Name      string
			Namespace string
		}{Name: "", Namespace: ""},
		Data: map[string]string{},
	}
}

func (r *ResSecret) SetNamespace(ns string) bool {
	r.MetaData.Namespace = ns
	return true
}

func (r *ResSecret) SetMetaDataName(name string) bool {
	r.MetaData.Name = name
	return true
}

func (r *ResSecret) GetMetaDataName() string {
	return r.MetaData.Name
}

func (r *ResSecret) SetType(typeName string) bool {
	r.Type = typeName
	return true
}

func (r *ResSecret) SetData(data []map[string]string) error {
	if len(data) > 0 {
		for _, v := range data {
			if v["key"] == "" || v["val"] == "" {
				return errors.New("key or val is empty")
			}
			// base64编码存储
			r.Data[v["key"]] = base64.StdEncoding.EncodeToString([]byte(v["val"]))
		}
		return nil
	} else {
		return errors.New("no data will be set")
	}
}

func (r *ResSecret) GetData() (map[string]string, error) {
	data := map[string]string{}
	if len(r.Data) > 0 {
		for k, v := range r.Data {
			// base64解码
			val, err := base64.StdEncoding.DecodeString(v)
			if err != nil {
				return map[string]string{}, err
			}
			data[k] = string(val)
		}
	}
	return data, nil
}

func (r *ResSecret) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}