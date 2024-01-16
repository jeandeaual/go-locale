//go:build android
// +build android

package main

import (
	"fyne.io/fyne/v2/driver"

	"github.com/jeandeaual/go-locale"
)

func init() {
	locale.SetRunOnJVM(func(fn func(vm, env, ctx uintptr) error) {
		driver.RunNative(func(ctx interface{}) error {
			and := env.(driver.AndroidContext)
			fn(and.VM, and.Env, and.Ctx)
		})
	})
}
