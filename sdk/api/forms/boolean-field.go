package sdkforms

type BooleanField struct {
	Name    string
	Label   string
	ValueFn func() bool
}

func (f BooleanField) GetName() string {
	return f.Name
}

func (f BooleanField) GetLabel() string {
	return f.Label
}

func (f BooleanField) GetType() string {
	return FormFieldTypeBoolean
}

func (f BooleanField) GetValue() interface{} {
	if f.ValueFn != nil {
		return f.ValueFn()
	}
	return false
}
