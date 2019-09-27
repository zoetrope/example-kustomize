package main

import (
	"sigs.k8s.io/kustomize/v3/pkg/ifc"
	"sigs.k8s.io/kustomize/v3/pkg/resmap"
	"sigs.k8s.io/kustomize/v3/pkg/types"
	"sigs.k8s.io/yaml"
)

//nolint: golint
//noinspection GoUnusedGlobalVariable
var KustomizePlugin plugin

type plugin struct {
	rf        *resmap.Factory
	ldr       ifc.Loader
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	RegoFile  string `json:"regofile,omitempty" file:"name,omitempty"`
}

type constraintTemplate struct {
	types.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	targets          []constraintTemplateTarget
}

type constraintTemplateSpec struct {
	targets []constraintTemplateTarget
}

type constraintTemplateTarget struct {
	target string
	rego   string
}

func (p *plugin) Config(ldr ifc.Loader, rf *resmap.Factory, config []byte) error {
	p.rf = rf
	p.ldr = ldr
	return yaml.Unmarshal(config, p)
}

func (p *plugin) Generate() (resmap.ResMap, error) {
	target := constraintTemplate{
		ObjectMeta: types.ObjectMeta{
			Name:      "hoge",
			Namespace: "piyo",
		},
	}
	buf, err := yaml.Marshal(&target)
	if err != nil {
		return nil, err
	}
	return p.rf.NewResMapFromBytes(buf)
}
