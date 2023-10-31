package fci

import (
	"errors"

	"github.com/flarehotspot/core/sdk/api/fci"
)

func NewFciSection(cfg *FciConfig, secmap map[string]any) *FciSection {
	return &FciSection{
		cfg:    cfg,
		secmap: secmap,
	}
}

type FciSection struct {
	cfg     *FciConfig
	secmap  map[string]any
	Secname string          `json:"name"`
	Secdesc string          `json:"desc"`
	Inputs  []fci.IFciInput `json:"inputs"`
}

func (sec *FciSection) Name() string {
	return sec.Secname
}

func (sec *FciSection) Description() string {
	return sec.Secdesc
}

func (sec *FciSection) CheckboxGroup(name string, label string) fci.IFciCheckboxGrp {
	var ckgrp *FciInputCheckboxGrp

	ickgrp, ok := sec.GetCheckboxGroup(name)
	if !ok {
		m := map[string]any{}
		ckgrp = NewFciInputCheckboxGrp(sec.cfg, m)
		sec.Inputs = append(sec.Inputs, ckgrp)
	} else {
		ckgrp = ickgrp.(*FciInputCheckboxGrp)
	}

	ckgrp.Ckname = name
	ckgrp.Cklabel = label

	return ckgrp
}

func (sec *FciSection) GetCheckboxGroup(name string) (ckg fci.IFciCheckboxGrp, ok bool) {
	for _, input := range sec.Inputs {
		if input.Type() == fci.FciInputCheckboxGroup && input.Name() == name {
			return input.(fci.IFciCheckboxGrp), true
		}
	}

	return nil, false
}

func (sec *FciSection) Checkbox(name string, label string, help string) fci.IFciCheckbox {
	var ck *FciInputCheckbox

	ick, ok := sec.GetCheckbox(name)
	if !ok {
		m := map[string]any{}
		ck = NewFciInputCheckbox(sec.cfg, m)
		sec.Inputs = append(sec.Inputs, ck)
	} else {
		ck = ick.(*FciInputCheckbox)
	}

	ck.Ckname = name
	ck.Cklabel = label
	ck.Ckhelp = help

	return ck
}

func (sec *FciSection) GetCheckbox(name string) (ck fci.IFciCheckbox, ok bool) {
	for _, input := range sec.Inputs {
		if input.Type() == fci.FciInputCheckbox && input.Name() == name {
			return input.(fci.IFciCheckbox), true
		}
	}

	return nil, false
}

func (sec *FciSection) FieldList(name string, label string) fci.IFciFieldList {
	var fl *FciFieldList

	ifl, ok := sec.GetFieldList(name)
	if !ok {
		m := [][]map[string]any{}
		fl = NewFciFieldList(sec.cfg, m)
	} else {
		fl = ifl.(*FciFieldList)
	}

	fl.Flname = name
	fl.Fllabel = label
	return fl
}

func (sec *FciSection) GetFieldList(name string) (fl fci.IFciFieldList, ok bool) {
	for _, input := range sec.Inputs {
		if input.Type() == fci.FciInputFieldLIst && input.Name() == name {
			return input.(fci.IFciFieldList), true
		}
	}

	return nil, false
}

func (sec *FciSection) Field(name string, label string, help string) fci.IFciInputField {
	var field *FciInputField

	ifield, ok := sec.GetField(name)
	if !ok {
		m := map[string]any{}
		field = NewFciInputField(sec.cfg, m)
		sec.Inputs = append(sec.Inputs, field)
	} else {
		field = ifield.(*FciInputField)
	}

	return field
}

func (sec *FciSection) GetField(name string) (input fci.IFciInputField, ok bool) {
	for _, input := range sec.Inputs {
		if input.Type() == fci.FciInputField && input.Name() == name {
			return input.(fci.IFciInputField), true
		}
	}
	return nil, false
}

func (sec *FciSection) Parse() error {
	m := sec.secmap
	nameval, nok := m["name"]
	descval, dok := m["desc"]
	iptmapval, iok := m["inputs"]
	defErr := errors.New("section definition error")

	ok := nok && dok && iok
	if !ok {
		return defErr
	}

	name, nok := nameval.(string)
	desc, dok := descval.(string)
	iptmap, iok := iptmapval.([]map[string]any)

	ok = nok && dok && iok
	if !ok {
		return defErr
	}

	inputs := []fci.IFciInput{}
	for _, m := range iptmap {
		parser := NewInputParser(sec.cfg, m)
		input, err := parser.Parse()

		if err == nil {
			inputs = append(inputs, input)
		}
	}

	sec.Secname = name
	sec.Secdesc = desc
	sec.Inputs = inputs

	return nil
}
