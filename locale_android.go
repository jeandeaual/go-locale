// +build android

package main

/*
#cgo LDFLAGS: -landroid -llog

#include <stdlib.h>

const char *getLocales(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx);
const char *getLocale(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx);
const char *getLanguage(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx);
const char *getRegion(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx);
*/
import "C"
import (
	"strings"
	"unsafe"

	"golang.org/x/mobile/app"
)

// GetLocale retrieves the IETF BCP 47 language tag set on the system.
func GetLocale() (string, error) {
	locale := ""

	err := app.RunOnJVM(func(vm, env, ctx uintptr) error {
		chars := C.getLocale(C.uintptr_t(vm), C.uintptr_t(env), C.uintptr_t(ctx))
		locale = C.GoString(chars)
		C.free(unsafe.Pointer(chars))
		return nil
	})

	return locale, err
}

// GetLocales retrieves the IETF BCP 47 language tags set on the system.
func GetLocales() ([]string, error) {
	locales := ""

	err := app.RunOnJVM(func(vm, env, ctx uintptr) error {
		chars := C.getLocales(C.uintptr_t(vm), C.uintptr_t(env), C.uintptr_t(ctx))
		locales = C.GoString(chars)
		C.free(unsafe.Pointer(chars))
		return nil
	})

	return strings.Split(locales, ","), err
}

// GetLanguage retrieves the IETF BCP 47 language tag set on the system and
// returns the language part of the tag.
func GetLanguage() (string, error) {
	language := ""

	err := app.RunOnJVM(func(vm, env, ctx uintptr) error {
		chars := C.getLanguage(C.uintptr_t(vm), C.uintptr_t(env), C.uintptr_t(ctx))
		language = C.GoString(chars)
		C.free(unsafe.Pointer(chars))
		return nil
	})

	return language, err
}

// GetRegion retrieves the IETF BCP 47 language tag set on the system and
// returns the region part of the tag.
func GetRegion() (string, error) {
	region := ""

	err := app.RunOnJVM(func(vm, env, ctx uintptr) error {
		chars := C.getRegion(C.uintptr_t(vm), C.uintptr_t(env), C.uintptr_t(ctx))
		region = C.GoString(chars)
		C.free(unsafe.Pointer(chars))
		return nil
	})

	return region, err
}
