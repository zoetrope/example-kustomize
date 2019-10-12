package main

import (
	"errors"
	"io/ioutil"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
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
	RegoFiles []string `json:"regos"`
	BaseFile  string   `json:"base"`
}

type constraintTemplate struct {
	types.TypeMeta   `json:",inline"`
	types.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec             constraintTemplateSpec `json:"spec,omitempty"`
}

type constraintTemplateSpec struct {
	CRD     crd      `json:"crd,omitempty"`
	Targets []target `json:"targets,omitempty"`
}

type crd struct {
	Spec crdSpec `json:"spec,omitempty"`
}

type crdSpec struct {
	Names      names       `json:"names,omitempty"`
	Validation *validation `json:"validation,omitempty"`
}

type names struct {
	Kind string `json:"kind,omitempty"`
}

type validation struct {
	OpenAPIV3Schema *apiextensionsv1beta1.JSONSchemaProps `json:"openAPIV3Schema,omitempty"`
}

type target struct {
	Target string   `json:"target,omitempty"`
	Rego   string   `json:"rego,omitempty"`
	Libs   []string `json:"libs,omitempty"`
}

func (p *plugin) Config(ldr ifc.Loader, rf *resmap.Factory, config []byte) error {
	p.rf = rf
	p.ldr = ldr
	return yaml.Unmarshal(config, p)
}

func (p *plugin) Generate() (resmap.ResMap, error) {
	baseData, err := ioutil.ReadFile(p.BaseFile)
	if err != nil {
		return nil, err
	}
	var tmpl constraintTemplate
	err = yaml.Unmarshal(baseData, &tmpl)
	if err != nil {
		return nil, err
	}

	if len(tmpl.Spec.Targets) != len(p.RegoFiles) {
		return nil, errors.New("length mismatch")
	}

	for i, regoFile := range p.RegoFiles {
		regoData, err := ioutil.ReadFile(regoFile)
		if err != nil {
			return nil, err
		}
		tmpl.Spec.Targets[i].Rego = string(regoData)
	}

	buf, err := yaml.Marshal(&tmpl)
	if err != nil {
		return nil, err
	}
	return p.rf.NewResMapFromBytes(buf)
}
