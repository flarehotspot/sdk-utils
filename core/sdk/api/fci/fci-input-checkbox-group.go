package fci

type IFciCheckboxGrp interface {
	Type() IFciInputTypes
	CheckboxItem(name string, value string, label string) IFciCheckbox
	DependsOn(name string, value string)
	Values() map[string]string
}
