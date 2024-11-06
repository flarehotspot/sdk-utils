package sdkfields

type Section struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Fields      []ConfigField `json:"fields"`
}
