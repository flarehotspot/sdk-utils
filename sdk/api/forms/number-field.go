package sdkforms

type NumberField struct {
	Name       string
	Label      string
	DefaultVal int
}

func (f NumberField) GetName() string {
	return f.Name
}

func (f NumberField) GetLabel() string {
	return f.Label
}

func (f NumberField) GetType() string {
	return FormFieldTypeNumber
}

func (f NumberField) GetDefaultVal() interface{} {
	return f.DefaultVal
}
