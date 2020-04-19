// +build !windows,!darwin

package locale

import (
	"errors"
	"os"
	"strings"
)

func splitLocales(locales string) []string {
	// If the user set different locales, they might be set in $LANGUAGE,
	// separated by a colon
	return strings.Split(locales, ":")
}

func getLangFromEnv() string {
	locale := ""

	// Check the following environment variables for the language information
	// See https://www.gnu.org/software/gettext/manual/html_node/Locale-Environment-Variables.html
	for _, env := range [...]string{"LANGUAGE", "LC_ALL", "LC_MESSAGES", "LANG"} {
		locale = os.Getenv(env)
		if len(locale) > 0 {
			return locale
		}
	}

	return locale
}

func getUnixLocales() ([]string, error) {
	locale := getLangFromEnv()
	if len(locale) == 0 {
		return nil, errors.New("cannot determine locale")
	}

	return splitLocales(locale), nil
}

// GetLocale retrieves the IETF BCP 47 language tag set on the system.
func GetLocale() (string, error) {
	unixLocales, err := getUnixLocales()
	if err != nil {
		return "", err
	}

	language, region := splitLocale(unixLocales[0])
	locale := language
	if len(region) > 0 {
		locale = strings.Join([]string{language, region}, "-")
	}

	return locale, err
}

// GetLocales retrieves the IETF BCP 47 language tags set on the system.
func GetLocales() ([]string, error) {
	unixLocales, err := getUnixLocales()
	if err != nil {
		return nil, err
	}

	locales := make([]string, 0, len(unixLocales))

	for _, unixLocale := range unixLocales {
		language, region := splitLocale(unixLocale)
		locale := language
		if len(region) > 0 {
			locale = strings.Join([]string{language, region}, "-")
		}
		locales = append(locales, locale)
	}

	return locales, nil
}

// GetLanguage retrieves the IETF BCP 47 language tag set on the system and
// returns the language part of the tag.
func GetLanguage() (string, error) {
	language := ""

	unixLocales, err := getUnixLocales()
	if err == nil {
		language, _ = splitLocale(unixLocales[0])
	}

	return language, err
}

// GetRegion retrieves the IETF BCP 47 language tag set on the system and
// returns the region part of the tag.
func GetRegion() (string, error) {
	region := ""

	unixLocales, err := getUnixLocales()
	if err == nil {
		_, region = splitLocale(unixLocales[0])
	}

	return region, err
}
