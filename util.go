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

func splitLocale(locale string) (string, string) {
	// Remove the encoding, if present
	formattedLocale := strings.Split(locale, ".")[0]
	// Normalize by replacing the hyphens with underscores
	formattedLocale = strings.Replace(formattedLocale, "-", "_", -1)

	// Split at the underscore
	split := strings.Split(formattedLocale, "_")
	language := split[0]
	territory := ""
	if len(split) > 1 {
		territory = split[1]
	}

	return language, territory
}
