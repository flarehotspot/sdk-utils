package sdkextract

import (
	"bytes"
	"errors"
	"os"

	sdktargz "github.com/flarehotspot/go-utils/targz"
	sdkunzip "github.com/flarehotspot/go-utils/unzip"
)

var (
	MagicNumZip                 = []byte{0x50, 0x4B, 0x03, 0x04}
	MagicNumGzip                = []byte{0x1F, 0x8B}
	ErrUnknownCompressionFormat = errors.New("unknown compression format")
)

func Extract(file string, dest string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	// Read the first 4 bytes (or more if needed for other formats)
	buf := make([]byte, 4)
	if _, err := f.Read(buf); err != nil {
		return err
	}

	// identify compression format
	switch {
	case bytes.HasPrefix(buf, MagicNumZip):
		return sdkunzip.Unzip(file, dest)
	case bytes.HasPrefix(buf, MagicNumGzip):
		return sdktargz.UntarGz(file, dest)
	}

	return ErrUnknownCompressionFormat
}
