package sdkfields

type IMultiField interface {
}

type MultiField struct {
	Name    string          `json:"name"`
	Columns []string        `json:"columns"`
	Fields  [][]ConfigField `json:"fields"`
	Default [][]ConfigField `json:"default"`
}

func (f MultiField) GetType() string {
	return FieldTypeMulti
}

func (f MultiField) GetName() string {
	return f.Name
}

func (f MultiField) GetDefaultValue() interface{} {
	return f.Name
}
