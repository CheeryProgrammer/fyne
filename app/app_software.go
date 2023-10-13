//go:build ci
// +build ci

package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/cheeryprogrammer/fyne/v2/internal/painter/software"
)

// NewWithID returns a new app instance using the test (headless) driver.
// The ID string should be globally unique to this app.
func NewWithID(id string) fyne.App {
	return newAppWithDriver(test.NewDriverWithPainter(software.NewPainter()), id)
}
