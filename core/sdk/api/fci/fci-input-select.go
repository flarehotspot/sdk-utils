package fci

type IFciSelect interface {
	Type() IFciInputTypes
	SetAttr(name string, value string, multi bool)
	Option(value string, text string, selected bool)
	DependsOn(name string, value string)
	Attrs() map[string]string
	MultiValue() []string
	Value() string
}
