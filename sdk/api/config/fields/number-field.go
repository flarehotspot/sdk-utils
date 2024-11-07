package sdkfields

type INumberField interface {
	Value() string
}

type NumberField struct {
	Name    string `json:"name"`
	Label   string `json:"label"`
	Min     int    `json:"min"`
	Max     int    `json:"max"`
	Default int    `json:"default"`
}

func (f NumberField) GetType() string {
	return FieldTypeNumber
}

func (f NumberField) GetName() string {
	return f.Name
}

func (f NumberField) GetDefaultValue() interface{} {
	return f.Default
}
