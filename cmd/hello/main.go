// Package main loads a very basic Hello World graphical application.
package main

import (
	"image/color"

	"github.com/cheeryprogrammer/fyne/v2"
	"github.com/cheeryprogrammer/fyne/v2/app"
	"github.com/cheeryprogrammer/fyne/v2/container"
	"github.com/cheeryprogrammer/fyne/v2/theme"
	"github.com/cheeryprogrammer/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w := a.NewWindow("Hello", false)

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome 😀")
		}),
	))

	w.ShowAndRun()
}
