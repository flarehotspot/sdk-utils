package cfgfields

type SectionData struct {
	Name   string      `json:"name"`
	Fields []FieldData `json:"fields"`
}
