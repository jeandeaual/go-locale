// +build !windows,!darwin,!js,!android

package locale

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func resetEnv(t *testing.T) {
	t.Log("Resetting LANGUAGE, LC_ALL, LC_MESSAGES and LANG")

	for _, env := range [...]string{"LANGUAGE", "LC_ALL", "LC_MESSAGES", "LANG"} {
		err := os.Setenv(env, "")
		assert.Equal(t, nil, err)
	}
}

func TestMultipleLocales(t *testing.T) {
	resetEnv(t)

	err := os.Setenv("LANGUAGE", "en_US:fr:ja")
	assert.Equal(t, nil, err, "err should be nil")
	err = os.Setenv("LC_ALL", "")
	assert.Equal(t, nil, err, "err should be nil")
	err = os.Setenv("LC_MESSAGES", "")
	assert.Equal(t, nil, err, "err should be nil")
	err = os.Setenv("LANG", "en_US.UTF-8")
	assert.Equal(t, nil, err, "err should be nil")

	locales, err := GetLocales()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, []string{"en-US", "fr", "ja"}, locales)

	locale, err := GetLocale()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "en-US", locale)

	lang, err := GetLanguage()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "en", lang)

	region, err := GetRegion()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "US", region)

	// If "C" is set, we should ignore LANGUAGE
	err = os.Setenv("LC_ALL", "C")
	assert.Equal(t, nil, err, "err should be nil")

	var nilStringSlice []string

	locales, err = GetLocales()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, nilStringSlice, locales)

	locale, err = GetLocale()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "", locale)

	lang, err = GetLanguage()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "", lang)

	region, err = GetRegion()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "", region)
}

func TestSingleLocale(t *testing.T) {
	resetEnv(t)

	err := os.Setenv("LANGUAGE", "ja_JP")
	assert.Equal(t, nil, err, "err should be nil")
	err = os.Setenv("LC_ALL", "en_US")
	assert.Equal(t, nil, err, "err should be nil")
	err = os.Setenv("LC_MESSAGES", "ko_KR")
	assert.Equal(t, nil, err, "err should be nil")
	err = os.Setenv("LANG", "fr_FR")
	assert.Equal(t, nil, err, "err should be nil")

	locales, err := GetLocales()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, []string{"ja-JP"}, locales)

	locale, err := GetLocale()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "ja-JP", locale)

	lang, err := GetLanguage()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "ja", lang)

	region, err := GetRegion()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "JP", region)

	err = os.Setenv("LANGUAGE", "")
	assert.Equal(t, nil, err, "err should be nil")

	locale, err = GetLocale()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "en-US", locale)

	lang, err = GetLanguage()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "en", lang)

	region, err = GetRegion()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "US", region)

	err = os.Setenv("LC_ALL", "")
	assert.Equal(t, nil, err, "err should be nil")

	locale, err = GetLocale()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "ko-KR", locale)

	lang, err = GetLanguage()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "ko", lang)

	region, err = GetRegion()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "KR", region)

	err = os.Setenv("LC_MESSAGES", "")
	assert.Equal(t, nil, err, "err should be nil")

	locale, err = GetLocale()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "fr-FR", locale)

	lang, err = GetLanguage()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "fr", lang)

	region, err = GetRegion()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "FR", region)
}

func TestLocaleNoRegion(t *testing.T) {
	resetEnv(t)

	err := os.Setenv("LANG", "fr")
	assert.Equal(t, nil, err, "err should be nil")

	locales, err := GetLocales()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, []string{"fr"}, locales)

	locale, err := GetLocale()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "fr", locale)

	lang, err := GetLanguage()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "fr", lang)

	region, err := GetRegion()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "", region)
}

func TestNoLocale(t *testing.T) {
	resetEnv(t)

	var nilStringSlice []string

	locales, err := GetLocales()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, nilStringSlice, locales)

	locale, err := GetLocale()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "", locale)

	lang, err := GetLanguage()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "", lang)

	region, err := GetRegion()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "", region)
}
