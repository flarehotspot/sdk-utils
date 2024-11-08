package sdkfields

type Section struct {
	Name        string        `json:"name"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Fields      []ConfigField `json:"fields"`
}
