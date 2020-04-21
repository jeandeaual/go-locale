package locale

import (
	"strings"
)

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
