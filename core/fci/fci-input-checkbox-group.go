package fci

import (
	"errors"

	"github.com/flarehotspot/core/sdk/api/fci"
)

func NewFciInputCheckboxGrp(cfg *FciConfig, ckgmap map[string]any) *FciInputCheckboxGrp {
	return &FciInputCheckboxGrp{
		cfg:       cfg,
		ckgmap:    ckgmap,
		Cktype:    fci.FciInputCheckboxGroup,
		Ckdepends: make(map[string]string),
		Ckboxes:   []*FciInputCheckbox{},
	}
}

type FciInputCheckboxGrp struct {
	cfg       *FciConfig
	ckgmap    map[string]any
	Cktype    fci.IFciInputTypes  `json:"type"`
	Ckname    string              `json:"name"`
	Cklabel   string              `json:"label"`
	Ckdepends map[string]string   `json:"depends"`
	Ckboxes   []*FciInputCheckbox `json:"checkboxes"`
}

func (ckg *FciInputCheckboxGrp) Name() string {
	return ckg.Ckname
}

func (ckg *FciInputCheckboxGrp) Type() fci.IFciInputTypes {
	return fci.FciInputCheckboxGroup
}

func (ckg *FciInputCheckboxGrp) CheckboxItem(name string, value string, label string) fci.IFciCheckbox {
	for _, ckitem := range ckg.Ckboxes {
		if ckitem.Ckname == name {
			return ckitem
		}
	}

	m := map[string]any{}
	cki := NewFciInputCheckbox(ckg.cfg, m)
	ckg.Ckboxes = append(ckg.Ckboxes, cki)

	return cki
}

func (ckg *FciInputCheckboxGrp) DependsOn(name string, value string) {
	ckg.Ckdepends[name] = value
}

func (ckg *FciInputCheckboxGrp) Values() map[string]string {
	values := make(map[string]string)
	for _, ckitem := range ckg.Ckboxes {
		if ckitem.Ckchecked {
			values[ckitem.Ckname] = ckitem.Ckvalue
		}
	}
	return values
}

func (ckg *FciInputCheckboxGrp) Parse() error {
	m := ckg.ckgmap
	nameval, nok := m["name"]
	labelval, lok := m["label"]
	dependsval, dok := m["depends"]
	itemapval, iok := m["items"]
	ckgErr := errors.New("Invalid checkbox group")

	ok := nok && lok && dok && iok
	if !ok {
		return ckgErr
	}

	name, nok := nameval.(string)
	label, lok := labelval.(string)
	depends, dok := dependsval.(map[string]string)
	itemap, iok := itemapval.([]map[string]any)

	ok = nok && lok && dok && iok
	if !ok {
		return ckgErr
	}

	items := []*FciInputCheckbox{}
	for _, im := range itemap {
		item := NewFciInputCheckbox(ckg.cfg, im)
		err := item.Parse()
		if err == nil {
			items = append(items, item)
		}
	}

	ckg.Ckname = name
	ckg.Cklabel = label
	ckg.Ckdepends = depends
	ckg.Ckboxes = items

	return nil
}
