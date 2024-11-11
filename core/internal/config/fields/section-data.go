package cfgfields

type ConfigData []SectionData

type SectionData struct {
	Name   string      `json:"name"`
	Fields []FieldData `json:"fields"`
}
