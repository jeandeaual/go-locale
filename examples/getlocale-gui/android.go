//go:build android
// +build android

package main

import (
	"fyne.io/fyne/v2/driver"

	"github.com/jeandeaual/go-locale"
)

func init() {
	locale.SetRunOnJVM(func(fn func(vm, env, ctx uintptr) error) error {
		driver.RunNative(func(ctx interface{}) error {
			and := ctx.(*driver.AndroidContext)
			return fn(and.VM, and.Env, and.Ctx)
		})
		return nil
	})
}
