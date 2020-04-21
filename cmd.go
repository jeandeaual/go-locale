// +build windows darwin

package locale

import (
	"os/exec"
	"syscall"
)

func execCommand(cmd string, args ...string) (status int, out []byte, err error) {
	status = -1
	command := exec.Command(cmd, args...)

	// Execute the command and get the standard and error outputs
	out, err = command.CombinedOutput()
	if err != nil {
		return
	}

	// Check the status code
	if w, ok := command.ProcessState.Sys().(syscall.WaitStatus); ok {
		status = w.ExitStatus()
	}

	return
}
