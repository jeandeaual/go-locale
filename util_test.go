//go:build !android
// +build !android

package locale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	globalLanguage string
	globalRegion   string
)

func BenchmarkSplitLocale(b *testing.B) {
	language := ""
	region := ""
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		language, region = splitLocale("en_US.UTF-8")
	}

	globalLanguage = language
	globalRegion = region
}

func TestSplitLocale(t *testing.T) {
	language, region := splitLocale("en_US.UTF-8")
	assert.Equal(t, "en", language)
	assert.Equal(t, "US", region)

	language, region = splitLocale("fr-FR")
	assert.Equal(t, "fr", language)
	assert.Equal(t, "FR", region)

	language, region = splitLocale("ja")
	assert.Equal(t, "ja", language)
	assert.Equal(t, "", region)

	language, region = splitLocale("test")
	assert.Equal(t, "test", language)
	assert.Equal(t, "", region)
}
