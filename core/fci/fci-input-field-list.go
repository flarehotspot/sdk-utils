package fci

import "github.com/flarehotspot/core/sdk/api/fci"

func NewFciFieldList(cfg *FciConfig, m [][]map[string]any) *FciFieldList {
	return &FciFieldList{
		cfg:    cfg,
		fmap:   m,
		Fltype: fci.FciInputFieldList,
	}
}

type FciFieldList struct {
	cfg       *FciConfig
	fmap      [][]map[string]any
	Fltype    fci.IFciInputTypes `json:"type"`
	Flname    string             `json:"name"`
	Fllabel   string             `json:"label"`
	Flcols    []string           `json:"cols"`
	Flrows    []*FciFieldLsRow   `json:"rows"`
	Fldepends map[string]string  `json:"depends"`
}

func (fl *FciFieldList) Type() fci.IFciInputTypes {
	return fl.Fltype
}

func (fl *FciFieldList) Name() string {
	return fl.Flname
}

func (fl *FciFieldList) Label() string {
	return fl.Fllabel
}

func (fl *FciFieldList) Cols(cols ...string) {
	fl.Flcols = cols
}

func (fl *FciFieldList) GetCols() []string {
	return fl.Flcols
}

func (fl *FciFieldList) Row(index int) (row fci.IFciInputLsRow) {
	if index < len(fl.Flrows) {
		return fl.Flrows[index]
	}
	m := make([]map[string]any, 0)
	r := NewFieldLsRow(fl.cfg, fl, m)
	fl.Flrows = append(fl.Flrows, r)
	return r
}

func (fl *FciFieldList) Rows() []fci.IFciInputLsRow {
	rows := make([]fci.IFciInputLsRow, len(fl.Flrows))
	for i, row := range fl.Flrows {
		rows[i] = row
	}
	return rows
}

func (fl *FciFieldList) DependsOn(name string, value string) {
	fl.Fldepends[name] = value
}

func (fl *FciFieldList) Values() []map[string]string {
	values := make([]map[string]string, len(fl.Flrows))
	for i, row := range fl.Flrows {
		values[i] = row.Values()
	}
	return values
}

func (fl *FciFieldList) Parse() error {
	m := fl.fmap
	fl.Flrows = []*FciFieldLsRow{}

	for i, row := range m {
		flrow := NewFieldLsRow(fl.cfg, fl, row)
		err := flrow.Parse()
		if err != nil {
			return err
		}

		fl.Flrows[i] = flrow
	}

	return nil
}

func (fl *FciFieldList) GetColIdx(col string) int {
	for i, c := range fl.Flcols {
		if c == col {
			return i
		}
	}
	return -1
}
