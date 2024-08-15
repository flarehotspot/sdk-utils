package sdktargz

import (
    "archive/tar"
    "compress/gzip"
    "io"
    "os"
    "path/filepath"
)

func TarGz(sourceDir, outputFile string) error {
    // Create the output file
    file, err := os.Create(outputFile)
    if err != nil {
        return err
    }
    defer file.Close()

    // Create a gzip writer
    gw := gzip.NewWriter(file)
    defer gw.Close()

    // Create a tar writer
    tw := tar.NewWriter(gw)
    defer tw.Close()

    // Walk through the directory
    err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Create a header for the current file
        header, err := tar.FileInfoHeader(info, info.Name())
        if err != nil {
            return err
        }

        // Update the name to reflect the correct path in the archive
        header.Name, err = filepath.Rel(sourceDir, path)
        if err != nil {
            return err
        }

        // Write the header
        if err := tw.WriteHeader(header); err != nil {
            return err
        }

        // If the file is not a directory, write the file content
        if !info.IsDir() {
            file, err := os.Open(path)
            if err != nil {
                return err
            }
            defer file.Close()

            if _, err := io.Copy(tw, file); err != nil {
                return err
            }
        }

        return nil
    })

    return err
}

