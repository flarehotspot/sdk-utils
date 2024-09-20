package sdkextract

import (
	"bytes"
	"errors"
	"log"
	"os"

	"sdk/utils/targz"
	"sdk/utils/unzip"
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
	UnknownCompressionFormatError = errors.New("unknown compression format")
)

var (
	compressionFormats = map[string]CompressionFormat{
		// targz
		"gzip": {
			Format:   "gzip",
			MagicNum: []byte{0x1F, 0x8B},
		},
		"zip": {
			Format:   "zip",
			MagicNum: []byte{0x50, 0x4B, 0x03, 0x04},
		},
	}
)

// Extracts the content of the given file without specifying
// the compression type
func Extract(filePath string, destPath string) error {
	var fileExtract = FileExtract{
		FilePath: filePath,
		DestPath: destPath,
	}

	setCompressionFormat(&fileExtract)

	err := fileExtract.extract()
	if err != nil {
		log.Fatal("Error:", err)
		return err
	}

	return nil
}

// identifies the compression format based on the specified file's magic number
// and updates the FileExtract format
func setCompressionFormat(fe *FileExtract) error {
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
	for _, cf := range compressionFormats {
		if bytes.HasPrefix(buf, cf.MagicNum) {
			fe.CompFormat = cf
			return nil
		}
	}

	return UnknownCompressionFormatError
}

func (f *FileExtract) extract() error {
	switch f.CompFormat.Format {
	case compressionFormats["gzip"].Format:
		err := sdktargz.UntarGz(f.FilePath, f.DestPath)
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		return nil
	case compressionFormats["zip"].Format:
		err := sdkunzip.Unzip(f.FilePath, f.DestPath)
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		return nil
	}

	return UnknownCompressionFormatError
}
