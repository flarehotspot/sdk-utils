package fci

type IFciFieldList interface {
	Type() IFciInputTypes
	Cols(cols ...string)
	Row(index int) (row IFciInputLsRow, ok bool)
	Rows() []IFciInputLsRow
	DependsOn(name string, value string)
	Values() []map[string]string
}

type IFciInputLsRow interface {
	Field(col string, name string) IFciInputField
	GetField(col string) (input IFciInputField, ok bool)
	Values() map[string]string
	Value(col string) (value string, ok bool)
}
