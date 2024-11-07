package cfgfields

type SectionData struct {
	Title  string      `json:"title"`
	Fields []FieldData `json:"fields"`
}
