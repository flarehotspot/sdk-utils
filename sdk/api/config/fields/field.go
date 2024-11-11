package sdkfields

const (
	FieldTypeText   = "text"
	FieldTypeNumber = "number"
	FieldTypeMulti  = "multi"
)

type ColField struct {
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Default interface{} `json:"default"`
}

type Field struct {
	Name      string      `json:"name"`
	Label     string      `json:"label"`
	InputType string      `json:"type"`    // string, number, multi
	Columns   []ColField  `json:"columns"` // for multi-field input
	Default   interface{} `json:"default"`
}
