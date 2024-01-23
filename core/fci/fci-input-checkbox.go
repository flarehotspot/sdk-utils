package fci

import (
	"errors"

	fci "github.com/flarehotspot/core/sdk/api/fci"
)

func NewFciInputCheckbox(cfg *FciConfig, ckmap map[string]any) *FciInputCheckbox {
	return &FciInputCheckbox{
		cfg:    cfg,
		ckmap:  ckmap,
		Cktype: fci.FciInputCheckbox,
	}
}

type FciInputCheckbox struct {
	cfg       *FciConfig
	ckmap     map[string]any
	Cktype    fci.IFciInputTypes `json:"type"`
	Ckname    string             `json:"name"`
	Cklabel   string             `json:"label"`
	Ckvalue   string             `json:"value"`
	Ckhelp    string             `json:"help"`
	Ckchecked bool               `json:"checked"`
	Ckattrs   map[string]string  `json:"attrs"`
	Ckdepends map[string]string  `json:"depends"`
}

func (ck *FciInputCheckbox) Type() fci.IFciInputTypes {
	return fci.FciInputCheckbox
}

func (ck *FciInputCheckbox) Name() string {
	return ck.Ckname
}

func (ck *FciInputCheckbox) SetAttr(name string, value string) {
	ck.Ckattrs[name] = value
}

func (ck *FciInputCheckbox) DependsOn(name string, value string) {
	ck.Ckdepends[name] = value
}

func (ck *FciInputCheckbox) Attrs() map[string]string {
	return ck.Ckattrs
}

func (ck *FciInputCheckbox) Depends() map[string]string {
	return ck.Ckdepends
}

func (ck *FciInputCheckbox) Checked() bool {
	return ck.Ckchecked
}

func (ck *FciInputCheckbox) Value() string {
	return ck.Ckvalue
}

func (ck *FciInputCheckbox) Parse() error {
	im := ck.ckmap
	nameval, nok := im["name"]
	valueval, vok := im["value"]
	labelval, lok := im["label"]
	checkedval, cok := im["checked"]
	attrsval, aok := im["attrs"]
	dependsval, dok := im["depends"]
	ckErr := errors.New("Invalid checkbox")

	ok := nok && vok && lok && cok && aok && dok
	if !ok {
		return ckErr
	}

	name, nok := nameval.(string)
	value, vok := valueval.(string)
	label, lok := labelval.(string)
	checked, cok := checkedval.(bool)
	attrs, aok := attrsval.(map[string]string)
	depends, dok := dependsval.(map[string]string)

	ok = nok && vok && lok && cok && aok && dok
	if !ok {
		return ckErr
	}

	ck.Ckname = name
	ck.Ckvalue = value
	ck.Cklabel = label
	ck.Ckchecked = checked
	ck.Ckattrs = attrs
	ck.Ckdepends = depends

	return nil
}
