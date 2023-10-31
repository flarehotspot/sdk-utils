package fci

import (
	"errors"

	"github.com/flarehotspot/core/sdk/api/fci"
)

func NewFciInputField(cfg *FciConfig, m map[string]any) *FciInputField {
	return &FciInputField{
		cfg:   cfg,
		ifmap: m,
		Ftype: fci.FciInputField,
	}
}

type FciInputField struct {
	cfg      *FciConfig
	ifmap    map[string]any
	Ftype    fci.IFciInputTypes `json:"type"`
	Fname    string             `json:"name"`
	Flabel   string             `json:"label"`
	Fattrs   map[string]string  `json:"attrs"`
	Fdepends map[string]string  `json:"depends"`
	Fvalue   string             `json:"value"`
}

func (f *FciInputField) Type() fci.IFciInputTypes {
	return f.Ftype
}

func (f *FciInputField) Name() string {
	return f.Fname
}

func (f *FciInputField) SetAttr(name string, value string) {
	f.Fattrs[name] = value
}

func (f *FciInputField) Attrs() map[string]string {
	return f.Fattrs
}

func (f *FciInputField) Value() string {
	return f.Fvalue
}

func (f *FciInputField) DependsOn(name string, value string) {
	f.Fdepends[name] = value
}

func (f *FciInputField) Parse() error {
	m := f.ifmap
	nameval, nok := m["name"]
	labelval, lok := m["label"]
	attrsval, aok := m["attrs"]
	dependsval, dok := m["depends"]
	valval, vok := m["value"]
	fErr := errors.New("fci: input field parse error")

	ok := nok && lok && aok && dok && vok
	if !ok {
		return fErr
	}

	name, nok := nameval.(string)
	label, lok := labelval.(string)
	attrs, aok := attrsval.(map[string]string)
	depends, dok := dependsval.(map[string]string)
	value, vok := valval.(string)

	ok = nok && lok && aok && dok && vok
	if !ok {
		return fErr
	}

	f.Fname = name
	f.Flabel = label
	f.Fattrs = attrs
	f.Fdepends = depends
	f.Fvalue = value

	return nil
}
