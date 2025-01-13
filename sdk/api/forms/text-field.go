package sdkforms

type TextField struct {
	Name    string
	Label   string
	ValueFn func() string
}

func (f TextField) GetName() string {
	return f.Name
}

func (f TextField) GetLabel() string {
	return f.Label
}

func (f TextField) GetType() string {
	return FormFieldTypeText
}

func (f TextField) GetValue() interface{} {
	if f.ValueFn != nil {
		return f.ValueFn()
	}
	return ""
}
