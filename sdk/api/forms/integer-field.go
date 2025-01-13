package sdkforms

type IntegerField struct {
	Name    string
	Label   string
	ValueFn func() int64
}

func (f IntegerField) GetName() string {
	return f.Name
}

func (f IntegerField) GetLabel() string {
	return f.Label
}

func (f IntegerField) GetType() string {
	return FormFieldTypeInteger
}

func (f IntegerField) GetValue() interface{} {
	if f.ValueFn != nil {
		return f.ValueFn()
	}
	return 0
}
