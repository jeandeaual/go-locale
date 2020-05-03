package main

import (
	"fmt"
	"os"

	"github.com/jeandeaual/go-locale"
)

func main() {
	locales, err := locale.GetLocales()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the locales: %v\n", err)
	} else {
		fmt.Println("Locales:", locales)
	}

	userLocale, err := locale.GetLocale()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the locale: %v\n", err)
	} else {
		fmt.Println("Locale:", userLocale)
	}

	language, err := locale.GetLanguage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the language: %v\n", err)
	} else {
		fmt.Println("Language:", language)
	}

	region, err := locale.GetRegion()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the region: %v\n", err)
	} else {
		fmt.Println("Region:", region)
	}
}
