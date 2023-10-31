package fci

import "github.com/flarehotspot/core/sdk/api/fci"

func NewFieldLsRow(cfg *FciConfig, fl *FciFieldList, m []map[string]any) *FciFieldLsRow {
	return &FciFieldLsRow{
		cfg:    cfg,
		fl:     fl,
		flmap:  m,
		Fields: []*FciInputField{},
	}
}

type FciFieldLsRow struct {
	cfg    *FciConfig
	fl     *FciFieldList
	flmap  []map[string]any
	Fields []*FciInputField `json:"fields"`
}

func (flrow *FciFieldLsRow) Field(col string, name string) fci.IFciInputField {
	var field *FciInputField

	input, ok := flrow.GetField(col)
	if !ok {
		field = input.(*FciInputField)
	} else {
		field = NewFciInputField(flrow.cfg, map[string]any{})
		colidx := flrow.fl.GetColIdx(col)
		flrow.Fields[colidx] = field
	}

	field.Fname = name

	return nil
}

func (flrow *FciFieldLsRow) GetField(col string) (input fci.IFciInputField, ok bool) {
	colidx := flrow.fl.GetColIdx(col)
	return flrow.Fields[colidx], colidx != -1
}

func (flrow *FciFieldLsRow) Values() map[string]string {
	m := map[string]string{}
	for _, field := range flrow.Fields {
		m[field.Fname] = field.Fvalue
	}
	return m
}

func (flrow *FciFieldLsRow) Value(col string) (value string, ok bool) {
	input, ok := flrow.GetField(col)
	if ok {
		return input.Value(), true
	}
	return "", false
}

func (flrow *FciFieldLsRow) Parse() error {
	m := flrow.flmap

	for _, field := range m {
		f := NewFciInputField(flrow.cfg, field)
		err := f.Parse()
		if err == nil {
			flrow.Fields = append(flrow.Fields, f)
		}
	}

	return nil
}
