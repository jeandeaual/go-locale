// +build darwin

package locale

import (
	"fmt"
	"regexp"
	"strings"
)

// GetLocale retrieves the IETF BCP 47 language tag set on the system.
func GetLocale() (string, error) {
	_, output, err := execCommand("defaults", "read", "-g", "AppleLocale")
	if err != nil {
		return "", fmt.Errorf("cannot determine locale: %w (output: %s)", err, output)
	}

	return strings.TrimRight(strings.Replace(string(output), "_", "-", 1), "\n"), nil
}

var appleLanguagesRegex = regexp.MustCompile(`"([a-z]{2}-[A-Z]{2})"`)

// GetLocales retrieves the IETF BCP 47 language tags set on the system.
func GetLocales() ([]string, error) {
	_, output, err := execCommand("defaults", "read", "-g", "AppleLanguages")
	if err != nil {
		return nil, fmt.Errorf("cannot determine locale: %w", err)
	}

	matches := appleLanguagesRegex.FindAllStringSubmatch(string(output), -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("invalid output from \"defaults read -g AppleLanguages\": %s", output)
	}

	locales := make([]string, 0, len(matches))

	for _, match := range matches {
		locales = append(locales, match[1])
	}

	return locales, nil
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
