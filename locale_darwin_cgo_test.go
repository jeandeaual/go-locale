//go:build darwin && !ios && cgo

package locale

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormalizeDarwinLocale(t *testing.T) {
	for in, want := range map[string]string{
		"en_US":           "en-US",
		"en_US@rg=gbzzzz": "en-US",
		"fr-FR":           "fr-FR",
		"zh-Hans-CN":      "zh-Hans-CN",
		"zh_Hans_CN":      "zh-Hans-CN",
		"ja":              "ja",
		"":                "",
	} {
		assert.Equal(t, want, normalizeDarwinLocale(in), in)
	}
}

// The native lookup should agree with the old "defaults read" path.
func TestNativeMatchesDefaults(t *testing.T) {
	native, err := GetLocale()
	require.NoError(t, err)
	cli, err := getLocaleCli()
	if err != nil {
		t.Skip("defaults read failed, nothing to compare against:", err)
	}
	assert.Equal(t, cli, native)

	natives, err := GetLocales()
	require.NoError(t, err)
	require.NotEmpty(t, natives)
	clis, err := getLocalesCli()
	if assert.NoError(t, err) && assert.Len(t, clis, len(natives)) {
		lang, _ := splitLocale(natives[0])
		cliLang, _ := splitLocale(clis[0])
		assert.Equal(t, cliLang, lang)
	}
}
