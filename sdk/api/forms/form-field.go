package sdkforms

const (
	FormFieldTypeText    = "text"
	FormFieldTypeNumber  = "number"
	FormFieldTypeBoolean = "bool"
	FormFieldTypeList    = "list"
	FormFieldTypeMulti   = "multi"
)

type FormField interface {
	GetName() string
	GetType() string
	GetDefaultVal() interface{}
}
