package fci

import (
	"errors"
	"fmt"
	"html/template"

	fci "github.com/flarehotspot/core/sdk/api/fci"
)

func NewFciInputField(cfg *FciConfig, m map[string]any) *FciInputField {
	return &FciInputField{
		cfg:      cfg,
		ifmap:    m,
		Ftype:    fci.FciInputField,
		Fattrs:   make(map[string]string),
		Fdepends: make(map[string]string),
	}
}

type FciInputField struct {
	cfg      *FciConfig
	ifmap    map[string]any
	Ftype    fci.IFciInputTypes `json:"type"`
	Flabel   string             `json:"label"`
	Fhelp    string             `json:"help"`
	Fattrs   map[string]string  `json:"attrs"`
	Fdepends map[string]string  `json:"depends"`
}

func (f *FciInputField) Type() fci.IFciInputTypes {
	return f.Ftype
}

func (f *FciInputField) Name() string {
	name, ok := f.Fattrs["name"]
	if !ok {
		return ""
	}
	return name
}

func (f *FciInputField) SetAttr(name string, value string) {
	f.Fattrs[name] = value
}

func (f *FciInputField) Attrs() []template.HTMLAttr {
	attrs := make([]template.HTMLAttr, len(f.Fattrs))
	for k, v := range f.Fattrs {
		s := fmt.Sprintf("%s=\"%s\"", k, v)
		attrs = append(attrs, template.HTMLAttr(s))
	}

	return attrs
}

func (f *FciInputField) Value() string {
	value, ok := f.Fattrs["value"]
	if !ok {
		return ""
	}
	return value
}

func (f *FciInputField) Label() string {
	return f.Flabel
}

func (f *FciInputField) Help() string {
	return f.Fhelp
}

func (f *FciInputField) DependsOn(name string, value string) {
	f.Fdepends[name] = value
}

func (f *FciInputField) Parse() error {
	m := f.ifmap
	labelval, lok := m["label"]
	attrsval, aok := m["attrs"]
	dependsval, dok := m["depends"]
	fErr := errors.New("fci: input field parse error")

	ok := lok && aok && dok
	if !ok {
		return fErr
	}

	label, lok := labelval.(string)
	attrs, aok := attrsval.(map[string]string)
	depends, dok := dependsval.(map[string]string)

	ok = lok && aok && dok
	if !ok {
		return fErr
	}

	f.Flabel = label
	f.Fattrs = attrs
	f.Fdepends = depends

	return nil
}
