package sdkforms

type ListOption struct {
	Label string
	Value interface{}
}

type ListField struct {
	Name       string
	Type       string
	Multiple   bool
	Options    func() []ListOption
	DefaultVal interface{}
}

func (f ListField) GetName() string {
	return f.Name
}

func (f ListField) GetType() string {
	return FormFieldTypeList
}

func (f ListField) GetDefaultVal() interface{} {
	return f.DefaultVal
}
