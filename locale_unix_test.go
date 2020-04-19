// +build !windows,!darwin

package locale

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func resetEnv(t *testing.T) {
	t.Log("Resetting LANGUAGE, LC_ALL, LC_MESSAGES and LANG")

	os.Setenv("LANGUAGE", "")
	os.Setenv("LC_ALL", "")
	os.Setenv("LC_MESSAGES", "")
	os.Setenv("LANG", "")
}

func TestMultipleLocales(t *testing.T) {
	resetEnv(t)

	os.Setenv("LANGUAGE", "en_US:fr:ja")
	os.Setenv("LC_ALL", "")
	os.Setenv("LC_MESSAGES", "")
	os.Setenv("LANG", "en_US.UTF-8")

	locales, err := GetLocales()
	assert.Equal(t, nil, err)
	assert.Equal(t, []string{"en-US", "fr", "ja"}, locales)

	locale, err := GetLocale()
	assert.Equal(t, nil, err)
	assert.Equal(t, "en-US", locale)

	lang, err := GetLanguage()
	assert.Equal(t, nil, err)
	assert.Equal(t, "en", lang)

	region, err := GetRegion()
	assert.Equal(t, nil, err)
	assert.Equal(t, "US", region)
}

func TestSingleLocale(t *testing.T) {
	resetEnv(t)

	os.Setenv("LANGUAGE", "ja_JP")
	os.Setenv("LC_ALL", "en_US")
	os.Setenv("LC_MESSAGES", "ko_KR")
	os.Setenv("LANG", "fr_FR")

	locales, err := GetLocales()
	assert.Equal(t, nil, err)
	assert.Equal(t, []string{"ja-JP"}, locales)

	locale, err := GetLocale()
	assert.Equal(t, nil, err)
	assert.Equal(t, "ja-JP", locale)

	lang, err := GetLanguage()
	assert.Equal(t, nil, err)
	assert.Equal(t, "ja", lang)

	region, err := GetRegion()
	assert.Equal(t, nil, err)
	assert.Equal(t, "JP", region)

	os.Setenv("LANGUAGE", "")

	locale, err = GetLocale()
	assert.Equal(t, nil, err)
	assert.Equal(t, "en-US", locale)

	lang, err = GetLanguage()
	assert.Equal(t, nil, err)
	assert.Equal(t, "en", lang)

	region, err = GetRegion()
	assert.Equal(t, nil, err)
	assert.Equal(t, "US", region)

	os.Setenv("LC_ALL", "")

	locale, err = GetLocale()
	assert.Equal(t, nil, err)
	assert.Equal(t, "ko-KR", locale)

	lang, err = GetLanguage()
	assert.Equal(t, nil, err)
	assert.Equal(t, "ko", lang)

	region, err = GetRegion()
	assert.Equal(t, nil, err)
	assert.Equal(t, "KR", region)

	os.Setenv("LC_MESSAGES", "")

	locale, err = GetLocale()
	assert.Equal(t, nil, err)
	assert.Equal(t, "fr-FR", locale)

	lang, err = GetLanguage()
	assert.Equal(t, nil, err)
	assert.Equal(t, "fr", lang)

	region, err = GetRegion()
	assert.Equal(t, nil, err)
	assert.Equal(t, "FR", region)
}

func TestLocaleNoRegion(t *testing.T) {
	resetEnv(t)

	os.Setenv("LANG", "fr")

	locales, err := GetLocales()
	assert.Equal(t, nil, err)
	assert.Equal(t, []string{"fr"}, locales)

	locale, err := GetLocale()
	assert.Equal(t, nil, err)
	assert.Equal(t, "fr", locale)

	lang, err := GetLanguage()
	assert.Equal(t, nil, err)
	assert.Equal(t, "fr", lang)

	region, err := GetRegion()
	assert.Equal(t, nil, err)
	assert.Equal(t, "", region)
}

func TestNoLocale(t *testing.T) {
	resetEnv(t)

	var nilStringSlice []string

	locales, err := GetLocales()
	assert.NotEqual(t, nil, err)
	assert.Equal(t, nilStringSlice, locales)

	locale, err := GetLocale()
	assert.NotEqual(t, nil, err)
	assert.Equal(t, "", locale)

	lang, err := GetLanguage()
	assert.NotEqual(t, nil, err)
	assert.Equal(t, "", lang)

	region, err := GetRegion()
	assert.NotEqual(t, nil, err)
	assert.Equal(t, "", region)
}
