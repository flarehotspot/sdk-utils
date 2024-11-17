package sdkforms

const (
	FormFieldTypeText    string = "text"
	FormFieldTypeNumber  string = "number"
	FormFieldTypeBoolean string = "bool"
	FormFieldTypeList    string = "list"
	FormFieldTypeMulti   string = "multi"
)

type FormField interface {
	GetName() string
	GetLabel() string
	GetType() string
	GetDefaultVal() interface{}
}
