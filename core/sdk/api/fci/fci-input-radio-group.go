package fci

type IFciRadioGrp interface {
	Type() IFciInputTypes
	Radio(value string, text string)
	DependsOn(name string, value string)
	Attrs() map[string]string
}
