package fci

import (
	"encoding/json"
	"os"
	"path"

	"github.com/flarehotspot/core/sdk/api/fci"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

func NewFciConfig(pkg string, name string) *FciConfig {
	return &FciConfig{
		file:     path.Join(paths.AppDir, "plugins", pkg, name+".json"),
		cfg:      make(map[string]any),
		Sections: []*FciSection{},
	}
}

type FciConfig struct {
	file     string
	cfg      map[string]any
	Sections []*FciSection `json:"sections"`
}

func (cfg *FciConfig) Section(name string, desc string) fci.IFciSection {
	var sec *FciSection

	isec, ok := cfg.GetSection(name)
	if !ok {
		m := map[string]any{}
		sec = NewFciSection(cfg, m)
		cfg.Sections = append(cfg.Sections, sec)
	} else {
		sec = isec.(*FciSection)
	}

	sec.Secname = name
	sec.Secdesc = desc

	return sec
}

func (cfg *FciConfig) GetSection(name string) (sec fci.IFciSection, ok bool) {
	for _, s := range cfg.Sections {
		if s.Secname == name {
			return s, true
		}
	}
	return nil, false
}

func (cfg *FciConfig) Save() error {
	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(cfg.file, b, 0644)
}

func (cfg *FciConfig) FromFile() {
	b, err := os.ReadFile(cfg.file)
	if err != nil {
		return
	}

	var config []map[string]any
	err = json.Unmarshal(b, &config)
	if err != nil {
		return
	}

	for _, secmap := range config {
		sec := NewFciSection(cfg, secmap)
		err := sec.Parse()
		if err != nil {
			continue
		}

		cfg.Sections = append(cfg.Sections, sec)
	}
}

func (cfg *FciConfig) FromPostForm(m map[string][]string) {
	for _, sec := range cfg.Sections {
		for _, input := range sec.Inputs {
			t := input.Type()
			if t == fci.FciInputCheckboxGroup {

			}
		}
	}
}

// func (cfg *FciConfig) ParseForm(req *http.Request)
