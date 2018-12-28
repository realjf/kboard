package api

import (
	"kboard/template"
	"kboard/config"
	"net/http"
	"kboard/k8s/resource"
	"kboard/k8s"
)

type IOrder struct {
	Api
}

func NewIOrder(config *config.Config, w http.ResponseWriter, r *http.Request) *IOrder {
	return &IOrder{
		Api{
			Config: config,
			TplEngine: template.NewTplEngine(w, r),
			Module: "index",
			Actions: map[string]func(){},
			R: r,
			W: w,
		},
	}
}

func (this *IOrder) Index() {

	this.TplEngine.Response(100, "", "")
}


// @todo 创建工单
func (this *IOrder) Edit() {

}

func (this *IOrder) Save() {

}

func (this *IOrder) List() {
	resReplicaSet := resource.NewResReplicaSet()
	resReplicaSet.SetMetadataName("hello")
	resReplicaSet.SetReplicas(3)
	resReplicaSet.SetNamespace("myapp")

	labels := map[string]string{
		"app":"nginx",
	}

	resReplicaSet.SetSelector(resource.Selector{
		MatchLabels:labels,
	})

	container := resource.NewContainer("mycontainer", "image")
	container.Resources = &resource.Resource{
		Limits: &resource.Limits{
			Cpu: "0.5",
			Memory:"100Mi",
		},
		Requests: &resource.Request{
			Cpu: "0.1",
			Memory: "50Mi",
		},
	}
	container.LivenessProbe = &resource.LivenessProbe{
		ProbeAction: resource.ProbeAction{
			Exec: &resource.ExecAction{
				Command: []string{
					"/bin/sh",
					"-c",
				},
			},
		},
		InitialDelaySeconds: 50,
		PeriodSeconds: 10,
		TimeoutSeconds: 10,
		SuccessThreshold: 1,
		FailureThreshold: 10,
	}
	resReplicaSet.AddContainer(container)
	resReplicaSet.SetTemplateLabel(labels)
	resReplicaSet.SetLabels(labels)
	yamlData, err := resReplicaSet.ToYamlFile()
	if err != nil {
		this.TplEngine.Response(99, err, "错误")
	}
	lib := k8s.NewReplicaSet(this.Config)

	res := lib.WriteToEtcd("myapp", "hello", yamlData)
	this.TplEngine.Response(100, res, "数据")
}


