package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/jeandeaual/go-locale"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	userLocales, err := locale.GetLocales()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the locales: %v\n", err)
	}

	userLocale, err := locale.GetLocale()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the locale: %v\n", err)
	}

	lang, err := locale.GetLanguage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the language: %v\n", err)
	}

	region, err := locale.GetRegion()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't retrieve the region: %v\n", err)
	}

	bundle := i18n.NewBundle(language.English)
	bundle.AddMessages(
		language.French,
		&i18n.Message{
			ID:    "Language",
			Other: "Langue: {{.Language}}",
		},
		&i18n.Message{
			ID:    "Locale",
			Other: "Locale: {{.Locale}}",
		},
		&i18n.Message{
			ID:    "Locales",
			Other: "Locales: {{.Locales}}",
		},
		&i18n.Message{
			ID:    "Region",
			Other: "Région: {{.Region}}",
		},
	)
	bundle.AddMessages(
		language.Japanese,
		&i18n.Message{
			ID:    "Language",
			Other: "言語： {{.Language}}",
		},
		&i18n.Message{
			ID:    "Locale",
			Other: "ロケール： {{.Locale}}",
		},
		&i18n.Message{
			ID:    "Locales",
			Other: "ロケール： {{.Locales}}",
		},
		&i18n.Message{
			ID:    "Region",
			Other: "領域： {{.Region}}",
		},
	)
	localizer := i18n.NewLocalizer(bundle, userLocales...)

	a := app.New()

	w := a.NewWindow("Locale Information")
	w.SetContent(container.NewVBox(
		widget.NewLabel(localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "Locales",
				Other: "Locales: {{.Locales}}",
			},
			TemplateData: map[string]interface{}{
				"Locales": userLocales,
			},
		})),
		widget.NewLabel(localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "Language",
				Other: "Language: {{.Language}}",
			},
			TemplateData: map[string]interface{}{
				"Language": lang,
			},
		})),
		widget.NewLabel(localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "Region",
				Other: "Region: {{.Region}}",
			},
			TemplateData: map[string]interface{}{
				"Region": region,
			},
		})),
		widget.NewLabel(localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "Locale",
				Other: "Locale: {{.Locale}}",
			},
			TemplateData: map[string]interface{}{
				"Locale": userLocale,
			},
		})),
	))

	w.ShowAndRun()
}
