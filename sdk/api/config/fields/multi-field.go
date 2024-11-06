package sdkfields

type MultiField struct {
	Name    string
	Columns []string
	Fields  [][]ConfigField
}

func (f *MultiField) GetType() string {
	return FieldTypeMulti
}

func (f *MultiField) GetName() string {
	return f.Name
}

func (f *MultiField) GetValue() interface{} {
	// TODO: return value
	return nil
}
