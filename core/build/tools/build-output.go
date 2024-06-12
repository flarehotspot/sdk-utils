package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	sdkfs "sdk/utils/fs"
	sdkpaths "sdk/utils/paths"
	sdkruntime "sdk/utils/runtime"
	sdkzip "sdk/utils/zip"
)

type BuildOutput struct {
	OutputDirName string
	Files         []string
	ExtraFiles    []ExtraFiles
}

type ExtraFiles struct {
	Src  string
	Dest string
}

type metajson struct {
	GoVersion string   `json:"go_version"`
	GoArch    string   `json:"go_arch"`
	OutputDir string   `json:"output_dir"`
	OutputZip string   `json:"output_zip"`
	Files     []string `json:"files"`
}

func (b *BuildOutput) Run() error {
	if err := sdkfs.EmptyDir(b.outputPath()); err != nil {
		return err
	}

	files := []string{}
	for _, entry := range b.Files {
		srcPath := filepath.Join(sdkpaths.AppDir, entry)
		destPath := filepath.Join(b.outputPath(), entry)
		if err := b.copy(srcPath, destPath); err != nil {
			panic(err)
		}
		files = append(files, entry)
	}

	for _, entry := range b.ExtraFiles {
		srcPath := filepath.Join(sdkpaths.AppDir, entry.Src)
		destPath := filepath.Join(b.outputPath(), entry.Dest)
		if err := b.copy(srcPath, destPath); err != nil {
			panic(err)
		}
		files = append(files, entry.Dest)
	}

	if err := sdkzip.Zip(b.outputPath(), b.zipFilePath()); err != nil {
		return err
	}

	md := metajson{
		GoVersion: sdkruntime.GOVERSION,
		GoArch:    sdkruntime.GOARCH,
		OutputDir: b.outputPath(),
		OutputZip: b.zipFilePath(),
		Files:     files,
	}

	metadata, err := json.MarshalIndent(&md, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("Metadata: %s\n", metadata)

	if err := os.WriteFile(b.metadataPath(), metadata, sdkfs.PermFile); err != nil {
		return err
	}

	return nil
}

func (b *BuildOutput) copy(srcPath string, destPath string) error {
	fmt.Printf("Copying '%s' -> '%s'\n", srcPath, destPath)

	if !sdkfs.Exists(srcPath) {
		return errors.New("File does not exist: " + srcPath)
	}

	if sdkfs.IsFile(srcPath) {
		if err := sdkfs.CopyFile(srcPath, destPath); err != nil {
			return err
		}
	} else if sdkfs.IsDir(srcPath) {
		if err := sdkfs.CopyDir(srcPath, destPath, nil); err != nil {
			return err
		}
	} else {
		return errors.New("Unknown file type: " + srcPath)
	}
	return nil
}

func (b *BuildOutput) outputPath() string {
	return filepath.Join(sdkpaths.AppDir, "output", b.OutputDirName)
}

func (b *BuildOutput) zipFilePath() string {
	return filepath.Join(b.outputPath() + ".zip")
}

func (b *BuildOutput) metadataPath() string {
	return filepath.Join(sdkpaths.AppDir, "output/metadata.json")
}
