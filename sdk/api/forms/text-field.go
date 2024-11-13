package sdkforms

type TextField struct {
	Name       string
	DefaultVal string
}

func (f TextField) GetName() string {
	return f.Name
}

func (f TextField) GetType() string {
	return FormFieldTypeText
}

func (f TextField) GetDefaultVal() interface{} {
	return f.DefaultVal
}
