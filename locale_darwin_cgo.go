//go:build darwin && !ios && cgo

package locale

// Non-CGO implementation is in locale_darwin_nocgo.go

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework CoreFoundation

#include <stdlib.h>

const char * preferredLocalization();
const char * preferredLocalizations();
char * systemLocale();
char * systemLanguages();
*/
import "C"
import (
	"strings"
	"unsafe"
)

// takeCString converts a malloc'd C string to Go and frees it.
func takeCString(s *C.char) string {
	if s == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(s))
	return C.GoString(s)
}

// normalizeDarwinLocale turns a CoreFoundation locale identifier such as
// "en_US@rg=gbzzzz" into a BCP 47 tag such as "en-US".
func normalizeDarwinLocale(id string) string {
	if idx := strings.Index(id, "@"); idx != -1 {
		id = id[:idx]
	}
	return strings.ReplaceAll(strings.TrimSpace(id), "_", "-")
}

// GetLocale retrieves the IETF BCP 47 language tag set on the system.
func GetLocale() (string, error) {
	if output := C.GoString(C.preferredLocalization()); output != "" {
		return strings.Replace(output, "_", "-", 1), nil
	}

	if output := normalizeDarwinLocale(takeCString(C.systemLocale())); output != "" {
		return output, nil
	}

	return getLocaleCli()
}

// GetLocales retrieves the IETF BCP 47 language tags set on the system.
func GetLocales() ([]string, error) {
	if output := C.GoString(C.preferredLocalizations()); output != "" {
		r := []string{}
		for _, s := range strings.Split(output, ",") {
			r = append(r, strings.Replace(s, "_", "-", 1))
		}
		return r, nil
	}

	if output := takeCString(C.systemLanguages()); output != "" {
		r := []string{}
		for _, s := range strings.Split(output, ",") {
			if s = normalizeDarwinLocale(s); s != "" {
				r = append(r, s)
			}
		}
		if len(r) > 0 {
			return r, nil
		}
	}

	return getLocalesCli()
}
