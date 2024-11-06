package sdkfields

const (
	FieldTypeText   = "text"
	FieldTypeNumber = "number"
	FieldTypeMulti  = "multi"
)

type ConfigField interface {
	GetType() string
	GetName() string
	GetValue() interface{}
}
