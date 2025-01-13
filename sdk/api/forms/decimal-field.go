package sdkforms

type DecimalField struct {
	Name      string
	Label     string
	Step      float64
	Precision int
	ValueFn   func() float64
}

func (f DecimalField) GetName() string {
	return f.Name
}

func (f DecimalField) GetLabel() string {
	return f.Label
}

func (f DecimalField) GetType() string {
	return FormFieldTypeDecimal
}

func (f DecimalField) GetValue() interface{} {
	if f.ValueFn != nil {
		return f.ValueFn()
	}
	return 0.0
}
