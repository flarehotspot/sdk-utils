package sdkutils

import (
	"errors"
	"fmt"
	"path/filepath"
)

type BuildOutput struct {
	OutputDirName string
	Files         []string
	CustomFiles   []BuildOutputCustomEntry
}

type BuildOutputCustomEntry struct {
	Src  string
	Dest string
}

type BuildOutputMeta struct {
	GoVersion string   `json:"go_version"`
	GoArch    string   `json:"go_arch"`
	OutputDir string   `json:"output_dir"`
	OutputZip string   `json:"output_zip"`
	Files     []string `json:"files"`
}

func ReadBuildOutput(outputDir string) (meta BuildOutputMeta, err error) {
	err = JsonRead(filepath.Join(outputDir, "metadata.json"), &meta)
	return
}

func (b *BuildOutput) Run() error {
	if err := FsEmptyDir(b.outputPath()); err != nil {
		return err
	}

	contentList := []string{}
	for _, entry := range b.Files {
		srcPath := filepath.Join(PathAppDir, entry)
		destPath := filepath.Join(b.outputPath(), entry)
		if err := b.copy(srcPath, destPath); err != nil {
			panic(err)
		}
		contentList = append(contentList, entry)
	}

	for _, entry := range b.CustomFiles {
		srcPath := filepath.Join(PathAppDir, entry.Src)
		destPath := filepath.Join(b.outputPath(), entry.Dest)
		if err := b.copy(srcPath, destPath); err != nil {
			panic(err)
		}
		contentList = append(contentList, entry.Dest)
	}

	// new implementation using tar.gz
	if err := CompressTar(b.outputPath(), b.targzFilePath()); err != nil {
		return err
	}

	md := BuildOutputMeta{
		GoVersion: GO_VERSION,
		GoArch:    GOARCH,
		OutputDir: b.outputPath(),
		OutputZip: b.targzFilePath(),
		Files:     contentList,
	}

	if err := JsonWrite(b.metadataPath(), md); err != nil {
		return err
	}

	return nil
}

func (b *BuildOutput) copy(srcPath string, destPath string) error {
	fmt.Printf("Copying '%s' -> '%s'\n", srcPath, destPath)

	if !FsExists(srcPath) {
		return errors.New("File does not exist: " + srcPath)
	}

	return FsCopy(srcPath, destPath)
}

func (b *BuildOutput) outputPath() string {
	return filepath.Join(PathAppDir, "output", b.OutputDirName)
}

func (b *BuildOutput) targzFilePath() string {
	return filepath.Join(b.outputPath() + ".tar.gz")
}

func (b *BuildOutput) metadataPath() string {
	return filepath.Join(PathAppDir, "output/metadata.json")
}
