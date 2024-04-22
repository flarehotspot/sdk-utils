//go:build !dev

package cmd

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Exec(command string) error {
	cmdarr := strings.Fields(command)
	bin := cmdarr[0]
	args := cmdarr[1:]

	log.Println(bin, strings.Join(args, " "))

	var (
		errBuilder strings.Builder
		outBuilder strings.Builder
	)

	cmd := exec.Command(bin, args...)
	cmd.Stderr = &errBuilder
	cmd.Stdout = &outBuilder

	err := cmd.Run()
	if err != nil {
		var msg string

		if errBuilder.Len() > 0 {
			msg = errBuilder.String()
		} else if outBuilder.Len() > 0 {
			msg = outBuilder.String()
		} else {
			msg = err.Error()
		}

		log.Printf("\nError executing command: %v\n\t%s\n\n", command, msg)
		return errors.New(msg)
	}
	return nil
}

func ExecAll(commands []string) error {
	for _, c := range commands {
		err := Exec(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func ExecOutput(command string, out io.Writer) (err error) {
	cmdarr := strings.Fields(command)
	bin := cmdarr[0]
	args := cmdarr[1:]

	log.Println(bin, strings.Join(args, " "))

	var stderr bytes.Buffer

	cmd := exec.Command(bin, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = out

	err = cmd.Run()
	if err != nil {
		log.Println(err, stderr)
		if stderr.String() != "" {
			return errors.New(stderr.String())
		}
		return err
	}

	return nil
}
