package fci

type IFciFieldList interface {
	Name() string
	Type() IFciInputTypes
	Label() string
	Cols(cols ...string)
	GetCols() []string
	Row(index int) IFciInputLsRow
	GetRows() []IFciInputLsRow
	DependsOn(name string, value string)
	Values() []map[string]string
}

type IFciInputLsRow interface {
	Field(col string, name string) IFciInputField
	GetField(col string) (input IFciInputField, ok bool)
	GetFields() []IFciInputField
	Values() map[string]string
	Value(col string) (value string, ok bool)
}
