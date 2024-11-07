package sdkfields

type ITextField interface {
	Value() string
}

type TextField struct {
	Name    string `json:"name"`
	Label   string `json:"label"`
	MinLen  int    `json:"min"`
	MaxLen  int    `json:"max"`
	Default string `json:"default"`
	Value   string `json:"-"`
}

func (f TextField) GetType() string {
	return FieldTypeText
}

func (f TextField) GetName() string {
	return f.Name
}
