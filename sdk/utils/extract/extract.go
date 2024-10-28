package sdkextract

import (
	"bytes"
	"errors"
	"log"
	"os"

	sdktargz "github.com/flarehotspot/go-utils/targz"
	sdkunzip "github.com/flarehotspot/go-utils/unzip"
)

type FileExtract struct {
	FilePath   string
	DestPath   string
	CompFormat CompressionFormat
}

type CompressionFormat struct {
	Format   string
	MagicNum []byte
}

var (
	ErrUnknownCompressionFormat = errors.New("unknown compression format")
)

var (
	CompressionZip = CompressionFormat{
		Format:   "zip",
		MagicNum: []byte{0x50, 0x4B, 0x03, 0x04},
	}
	CompressionGzip = CompressionFormat{
		Format:   "gzip",
		MagicNum: []byte{0x1F, 0x8B},
	}
)

// Extracts the content of the given file without specifying
// the compression type
func Extract(filePath string, destPath string) error {
	var fileExtract = FileExtract{
		FilePath: filePath,
		DestPath: destPath,
	}

	setCompFormat(&fileExtract)

	err := fileExtract.extract()
	if err != nil {
		log.Fatal("Error:", err)
		return err
	}

	return nil
}

// identifies the compression format based on the specified file's magic number
// and sets it compression format
func setCompFormat(fe *FileExtract) error {
	// open file
	f, err := os.Open(fe.FilePath)
	if err != nil {
		log.Fatal("Error:", err)
		return err
	}

	// Read the first 4 bytes (or more if needed for other formats)
	buf := make([]byte, 4)
	if _, err := f.Read(buf); err != nil {
		return err
	}

	// identify compression format
	switch {
	case bytes.HasPrefix(buf, CompressionZip.MagicNum):
		fe.CompFormat = CompressionZip
		return nil
	case bytes.HasPrefix(buf, CompressionGzip.MagicNum):
		fe.CompFormat = CompressionGzip
		return nil
	}

	return ErrUnknownCompressionFormat
}

func (f *FileExtract) extract() error {
	switch f.CompFormat.Format {
	case CompressionGzip.Format:
		err := sdktargz.UntarGz(f.FilePath, f.DestPath)
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		return nil
	case CompressionZip.Format:
		err := sdkunzip.Unzip(f.FilePath, f.DestPath)
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		return nil
	}

	return ErrUnknownCompressionFormat
}
