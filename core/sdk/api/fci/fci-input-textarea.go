package sdkfci

type IFciTextarea interface {
	Type() IFciInputTypes
	SetAttr(name string, value string)
	DependsOn(name string, value string)
	Attrs() map[string]string
	Value() string
}
