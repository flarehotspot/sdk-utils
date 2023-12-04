package views

import (
	"io"
	"os"

	"github.com/flarehotspot/core/sdk/utils/fs"
)

func NewJetLoader() *JetLoader {
	return &JetLoader{}
}

type JetLoader struct{}

func (jl *JetLoader) Exists(f string) bool {
	return fs.Exists(f)
}

func (jl *JetLoader) Open(f string) (io.ReadCloser, error) {
	return os.Open(f)
}
