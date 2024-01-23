package fci

type FciApi struct {
	pkg string
}

func (fci *FciApi) Config(name string) *FciConfig {
	return NewFciConfig(fci.pkg, name)
}
