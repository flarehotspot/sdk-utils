//go:build dev
package cmd

import "log"

func ExecAsh(command string) error {
  log.Println("/bin/ash " + command)
  return nil
}
