package sdkforms

const (
	FormFieldTypeText    string = "text"
	FormFieldTypeNumber  string = "number"
	FormFieldTypeBoolean string = "bool"
	FormFieldTypeList    string = "list"
	FormFieldTypeMulti   string = "multi"
)

type FieldData struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type IFormField interface {
	GetName() string
	GetLabel() string
	GetType() string
	GetDefaultVal() interface{}
}

type FormSection struct {
	Name   string
	Fields []IFormField
}

type Form struct {
	Name          string
	CallbackRoute string
	Sections      []FormSection
}
