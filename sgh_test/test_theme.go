package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"os"
	"path/filepath"
)

var (
	content = "阅读是一件非常愉快的事情，哈哈。早睡早起可以养生!o(╯□╰)o \r\nHappy Friday"
)

func init() {
	os.Setenv("FYNE_FONT", filepath.Join("sgh_test", "assets", "static", "PingFang Regular.ttf"))
	os.Setenv("FYNE_SCALE", "1")
}

func main() {
	app := app.New()
	lightTheme := theme.LightTheme()
	app.Settings().SetTheme(lightTheme)
	w := app.NewWindow("")
	label := widget.NewLabel(content)
	vbox := widget.NewVBox(
		label,
	)
	w.SetContent(vbox)
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.ShowAndRun()
}
