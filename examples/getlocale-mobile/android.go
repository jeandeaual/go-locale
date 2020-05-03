// +build android

package main

import (
	"github.com/fyne-io/mobile/app"

	"github.com/jeandeaual/go-locale"
)

func init() {
	locale.SetRunOnJVM(app.RunOnJVM)
}
