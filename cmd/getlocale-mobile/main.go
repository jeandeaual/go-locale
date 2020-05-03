package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"github.com/jeandeaual/go-locale"
)

func main() {
	a := app.New()

	locales, err := locale.GetLocales()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the locales: %v\n", err)
	}

	userLocale, err := locale.GetLocale()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the locale: %v\n", err)
	}

	language, err := locale.GetLanguage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the language: %v\n", err)
	}

	region, err := locale.GetRegion()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the region: %v\n", err)
	}

	w := a.NewWindow("Locale Information")
	w.SetContent(widget.NewVBox(
		widget.NewLabel(fmt.Sprintf("Locale: %s", userLocale)),
		widget.NewLabel(fmt.Sprintf("Language: %s", language)),
		widget.NewLabel(fmt.Sprintf("Region: %s", region)),
		widget.NewLabel(fmt.Sprintf("Locales: %v", locales)),
	))

	w.ShowAndRun()
}
