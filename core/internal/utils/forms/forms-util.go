package formsutl

type FieldData struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type ConfigData []SectionData

type SectionData struct {
	Name   string      `json:"name"`
	Fields []FieldData `json:"fields"`
}
