package formsutl

import sdkforms "sdk/api/forms"

type SectionData struct {
	Name   string               `json:"name"`
	Fields []sdkforms.FieldData `json:"fields"`
}
