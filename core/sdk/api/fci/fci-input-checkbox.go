package fci

type IFciCheckbox interface {
	Type() IFciInputTypes
	SetAttr(name string, value string)
	DependsOn(name string, value string)
	Attrs() map[string]string
	Depends() map[string]string
	Checked() bool
	Value() string
}
