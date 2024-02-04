// Package locale implements functions to retrieve the current locale(s)
// of the operating system.
package locale

import (
	"testing"
)

var (
	result  string
	results []string
	e       error
)

func BenchmarkGetLocale(b *testing.B) {
	var (
		locale string
		err    error
	)

	for n := 0; n < b.N; n++ {
		locale, err = GetLocale()
	}

	result = locale
	e = err
}

func BenchmarkGetLocales(b *testing.B) {
	var (
		locales []string
		err     error
	)

	for n := 0; n < b.N; n++ {
		locales, err = GetLocales()
	}

	results = locales
	e = err
}

func BenchmarkGetLanguage(b *testing.B) {
	var (
		language string
		err      error
	)

	for n := 0; n < b.N; n++ {
		language, err = GetLanguage()
	}

	result = language
	e = err
}

func BenchmarkGetRegion(b *testing.B) {
	var (
		region string
		err    error
	)

	for n := 0; n < b.N; n++ {
		region, err = GetRegion()
	}

	result = region
	e = err
}
