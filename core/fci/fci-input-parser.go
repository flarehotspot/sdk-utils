package fci

import (
	"errors"

	"github.com/flarehotspot/core/sdk/api/fci"
)

func NewInputParser(cfg *FciConfig, im map[string]any) *InputParser {
	return &InputParser{cfg, im}
}

type InputParser struct {
	cfg    *FciConfig
	iptmap map[string]any
}

func (ip *InputParser) Parse() (fci.IFciInput, error) {
	m := ip.iptmap
	typ := m["type"].(fci.IFciInputTypes)
	switch typ {
	case fci.FciInputCheckboxGroup:
		ckg := NewFciInputCheckboxGrp(ip.cfg, m)
		err := ckg.Parse()
		return ckg, err
	}

	return nil, errors.New("Invalid input type")
}
