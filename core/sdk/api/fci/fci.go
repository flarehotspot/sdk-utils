// Flarehotspot Configuration Interface (FCI)

package fci

type IFciInputTypes string

const (
	FciInputField         IFciInputTypes = "field"
	FciInputFieldList                    = "field-list"
	FciInputTextarea                     = "textarea"
	FciInputSelect                       = "select"
	FciInputCheckbox                     = "checkbox"
	FciInputCheckboxGroup                = "checkbox-group"
	FciInputRadioGroup                   = "radio-group"
)

type IFciApi interface {
	Config(name string) IFciConfig
}

type IFciConfig interface {
	Section(name string, desc string) IFciSection
	GetSection(name string) (section IFciSection, ok bool)
	GetInput(name string) (input IFciInput, ok bool)
}

type IFciSection interface {
	Name() string
	Description() string

	CheckboxGroup(name string, label string) IFciCheckboxGrp
	GetCheckboxGroup(name string) (ckgrp IFciCheckboxGrp, ok bool)

	Checkbox(name string, label string, help string) IFciCheckbox
	GetCheckbox(name string) (ck IFciCheckbox, ok bool)

	FieldList(name string, label string) IFciFieldList
	GetFieldList(name string) (fl IFciFieldList, ok bool)

	Field(name string, label string, help string) IFciInputField
	GetField(name string) (input IFciInputField, ok bool)

	// Select(name string, label string, help string) IFciSelect
	// GetSelect(name string) (sel IFciSelect, ok bool)

	// RadioGroup(name string, label string) IFciRadioGrp
	// GetRadioGroup(name string) (rg IFciRadioGrp, ok bool)
}

type IFciInput interface {
	Type() IFciInputTypes
	Name() string
}
