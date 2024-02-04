//go:build !windows && !darwin && !js && !android
// +build !windows,!darwin,!js,!android

package locale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleLocales(t *testing.T) {
	t.Setenv("LANGUAGE", "en_US:fr:ja")
	t.Setenv("LC_ALL", "")
	t.Setenv("LC_MESSAGES", "")
	t.Setenv("LANG", "en_US.UTF-8")

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
	t.Setenv("LC_ALL", "C")

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
	t.Setenv("LANGUAGE", "ja_JP")
	t.Setenv("LC_ALL", "en_US")
	t.Setenv("LC_MESSAGES", "ko_KR")
	t.Setenv("LANG", "fr_FR")

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

	t.Setenv("LANGUAGE", "")

	locale, err = GetLocale()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "en-US", locale)

	lang, err = GetLanguage()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "en", lang)

	region, err = GetRegion()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "US", region)

	t.Setenv("LC_ALL", "")

	locale, err = GetLocale()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "ko-KR", locale)

	lang, err = GetLanguage()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "ko", lang)

	region, err = GetRegion()
	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "KR", region)

	t.Setenv("LC_MESSAGES", "")

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
	t.Setenv("LANG", "fr")

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
	var nilStringSlice []string

	for _, env := range [...]string{"LANGUAGE", "LC_ALL", "LC_MESSAGES", "LANG"} {
		t.Setenv(env, "")
	}

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
