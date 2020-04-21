# go-locale

GoLang library used to retrieve the current locale(s) of the operating system.

## OS Support

* Windows, using Powershell's `Get-Culture`.
* macOS, using `defaults read -g AppleLocale` and `defaults read -g AppleLanguages` (since environment variables like `LANG` are not usually set on macOS).
* WASM (JavaScript), using [`navigator.language`](https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage/language) and [`navigator.languages`](https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage/languages).
* Unix-like systems (Linux, BSD, etc.), using the `LANGUAGE`, `LC_ALL`, `LC_MESSAGES` and `LANG` environment variables.

## Usage

## GetLocale

`GetLocale` returns the current locale as defined in [BCP 47](https://tools.ietf.org/rfc/bcp/bcp47.txt) (e.g. `"en-US"`).

```golang
userLocale, err := locale.GetLocale()
if err == nil {
    fmt.Println("Locale:", userLocale)
}
```

## GetLanguage

`GetLanguage` returns the current language as an [ISO 639](http://en.wikipedia.org/wiki/ISO_639) language code (e.g. `"en"`).

```golang
userLanguage, err := locale.GetLanguage()
if err == nil {
    fmt.Println("Language:", userLocale)
}
```

## GetRegion

`GetRegion` returns the current language as an [ISO 3166](http://en.wikipedia.org/wiki/ISO_3166-1) country code (e.g. `"US"`).

```golang
userRegion, err := locale.GetRegion()
if err == nil {
    fmt.Println("Region:", userRegion)
}
```

## GetLocales

`GetLocales` returns the user's preferred locales, by order of preference, as a slice of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag) string (e.g. `[]string{"en-US", "fr-FR", "ja-JP"}`).

This works if the user set multiple languages on macOS and other Unix systems.
Otherwise, it returns a slice with a single locale.

```golang
userLocales, err := locale.GetLocales()
if err == nil {
    fmt.Println("Locales:", userLocales)
}
```

This can be used with [golang.org/x/text](https://godoc.org/golang.org/x/text) or [go-i18n](https://github.com/nicksnyder/go-i18n) to set the localizer's language preferences:

```golang
userLocales, _ := locale.GetLocales()
localizer := i18n.NewLocalizer(bundle, userLocales...)
```

## Aknowledgements

Inspired by [jibber_jabber](https://github.com/cloudfoundry-attic/jibber_jabber).
