package locale

import (
	"os/exec"
	"strings"
	"syscall"
)

func execCommand(cmd string, args ...string) (status int, out []byte, err error) {
	status = -1
	command := exec.Command(cmd, args...)

	// Execute the command and get the standard and error outputs
	out, err = command.Output()
	if err != nil {
		return
	}

	// Check the status code
	if w, ok := command.ProcessState.Sys().(syscall.WaitStatus); ok {
		status = w.ExitStatus()
	}

	return
}

func splitLocale(locale string) (string, string) {
	formattedLocale := strings.Split(locale, ".")[0]
	formattedLocale = strings.Replace(formattedLocale, "-", "_", -1)

	pieces := strings.Split(formattedLocale, "_")
	language := pieces[0]
	territory := ""
	if len(pieces) > 1 {
		territory = strings.Split(formattedLocale, "_")[1]
	}

	return language, territory
}
