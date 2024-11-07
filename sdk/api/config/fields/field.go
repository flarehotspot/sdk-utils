package sdkfields

const (
	FieldTypeText   = "text"
	FieldTypeNumber = "number"
	FieldTypeMulti  = "multi"
)

type ConfigField interface {
	GetType() string
	GetName() string
	GetDefaultValue() interface{}
}
