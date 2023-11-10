package fci

import "html/template"

type IFciInputField interface {
	Type() IFciInputTypes
	SetAttr(name string, value string)
	DependsOn(name string, value string)
	Attrs() []template.HTMLAttr
	Value() string
	Label() string
	Help() string
}
