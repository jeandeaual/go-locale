// +build windows

package locale

import (
	"fmt"
	"strings"
)

// GetLocale retrieves the IETF BCP 47 language tag set on the system.
func GetLocale() (string, error) {
	_, output, err := execCommand("powershell", "Get-Culture | select -exp IetfLanguageTag")
	if err != nil {
		return "", fmt.Errorf("cannot determine locale: %w", err)
	}

	return strings.TrimRight(string(output), "\r\n"), nil
}

// GetLocales retrieves the IETF BCP 47 language tags set on the system.
func GetLocales() ([]string, error) {
	locale, err := GetLocale()
	if err != nil {
		return nil, err
	}

	return []string{locale}, nil
}

// GetLanguage retrieves the IETF BCP 47 language tag set on the system and
// returns the language part of the tag.
func GetLanguage() (string, error) {
	language := ""

	locale, err := GetLocale()
	if err == nil {
		language, _ = splitLocale(locale)
	}

	return language, err
}

// GetRegion retrieves the IETF BCP 47 language tag set on the system and
// returns the region part of the tag.
func GetRegion() (string, error) {
	region := ""

	locale, err := GetLocale()
	if err == nil {
		_, region = splitLocale(locale)
	}

	return region, err
}
