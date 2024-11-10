package sdkfields

type MultiField struct {
	Name    string          `json:"name"`
	Columns []ConfigField   `json:"columns"`
	Default [][]interface{} `json:"default"`
}

func (f MultiField) GetType() string {
	return FieldTypeMulti
}

func (f MultiField) GetName() string {
	return f.Name
}

func (f MultiField) GetDefaultValue() interface{} {
	return f.Default
}

type IMultiField interface {
	GetStringValue(row int, name string) (string, error)
	GetIntValue(row int, name string) (int, error)
}
