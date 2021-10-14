//go:build windows || (darwin && !ios)
// +build windows darwin,!ios

package locale

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	localeRegex   = regexp.MustCompile(`^[a-z]{2}(?:-[A-Z]{2})?$`)
	languageRegex = regexp.MustCompile(`^[a-z]{2}$`)
	regionRegex   = regexp.MustCompile(`^[A-Z]{2}$`)
)

func TestGetLocale(t *testing.T) {
	locale, err := GetLocale()
	assert.Equal(t, nil, err, "err should be nil")
	assert.True(t, localeRegex.MatchString(locale), "\"%s\" should match %v", locale, localeRegex)
}

func TestGetLanguage(t *testing.T) {
	language, err := GetLanguage()
	assert.Equal(t, nil, err, "err should be nil")
	assert.True(t, languageRegex.MatchString(language), "\"%s\" should match %v", language, languageRegex)
}

func TestGetRegion(t *testing.T) {
	region, err := GetRegion()
	assert.Equal(t, nil, err, "err should be nil")
	assert.True(t, regionRegex.MatchString(region), "\"%s\" should match %v", region, regionRegex)
}

func TestGetLocales(t *testing.T) {
	locales, err := GetLocales()
	assert.Equal(t, nil, err, "err should be nil")
	assert.NotZero(t, locales)

	for _, locale := range locales {
		assert.True(t, localeRegex.MatchString(locale), "\"%s\" should match %v", locale, localeRegex)
	}
}
