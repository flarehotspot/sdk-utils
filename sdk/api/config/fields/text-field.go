package sdkfields

type TextField struct {
	Name    string `json:"name"`
	Label   string `json:"label"`
	Min     int    `json:"min"`
	Max     int    `json:"max"`
	Value   string `json:"-"`
	Default string `json:"-"`
}

func (f *TextField) GetType() string {
	return FieldTypeText
}

func (f *TextField) GetName() string {
	return f.Name
}

func (f *TextField) GetValue() interface{} {
	return f.Value
}
