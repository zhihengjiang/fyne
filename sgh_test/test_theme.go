package main

import (
	"fmt"
	"fyne.io/fyne"
	fyneApp "fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"os"
	"path/filepath"
	"time"
)

var (
	content      = "阅读是一件非常愉快的事情，哈哈。早睡早起可以养生!o(╯□╰)o \r\nHappy Friday"
	defaultTitle = "Reader"
	app          fyne.App
)

type win struct {
	fw         fyne.Window
	bfw        fyne.Window
	label      *widget.Label
	borderless bool
}

func NewWin() *win {
	label := widget.NewLabel("请通过快捷键进行操作")
	return &win{
		label: label,
	}
}

func CreateWin(title string, borderless bool) fyne.Window {
	return createWinCore(title, borderless)
}

func centerOnScreen(w fyne.Window) {

	viewWidth, viewHeight := w.GetSghGlfwWindow().GetSize()

	// get window dimensions in pixels
	monitor := w.GetSghGlfwWindow().GetMonitor()
	if monitor == nil {
		return
	}
	monMode := monitor.GetVideoMode()

	// these come into play when dealing with multiple monitors
	monX, monY := monitor.GetPos()

	// math them to the middle
	newX := (monMode.Width / 2) - (viewWidth / 2) + monX
	newY := (monMode.Height / 2) - (viewHeight / 2) + monY

	// set new window coordinates
	w.GetSghGlfwWindow().SetPos(newX, newY)
}

func createWinCore(title string, borderless bool) fyne.Window {
	w := app.Driver().CreateSghWindow(title, borderless)
	//w.CenterOnScreen()
	w.SetFixedSize(true)
	return w
}

func (w *win) Exit() {
	app.Quit()
	os.Exit(0)
}

func (w *win) initWin(borderless bool) {
	nw := CreateWin(defaultTitle, borderless)
	nw.SetContent(w.label)
	nw.SetOnClosed(func() {
		w.Exit()
	})
	if borderless {
		w.bfw = nw
	} else {
		w.fw = nw
	}
}

func (w *win) show(borderless bool) {
	if w.fw == nil {
		w.initWin(false)
	}
	if w.bfw == nil {
		w.initWin(true)
	}

	w.borderless = borderless
	if borderless {
		x1, y1 := w.fw.GetSghGlfwWindow().GetPos()
		x2, y2 := w.bfw.GetSghGlfwWindow().GetPos()
		if x1 != x2 || y1 != y2 {
			w.bfw.GetSghGlfwWindow().SetPos(x1, y1)
		}
		w.fw.Hide()
		w.bfw.Show()
	} else {
		w.bfw.Hide()
		w.fw.Show()
	}
}

func (w *win) SetBorderless(borderless bool) {
	w.show(borderless)
}

func init() {
	os.Setenv("FYNE_FONT", filepath.Join("sgh_test", "assets", "static", "PingFang Regular.ttf"))
	os.Setenv("FYNE_SCALE", "1")
	lightTheme := theme.SghLightTheme()
	app = fyneApp.New()
	app.Settings().SetTheme(lightTheme)
}

func main() {
	time.Sleep(time.Second * 1)
	w := NewWin()
	go func() {
		for {
			time.Sleep(time.Second * 5)
			fmt.Printf("\r\nnow setting borderless %t", !w.borderless)
			w.SetBorderless(!w.borderless)
		}
	}()
	go func() {
		cnt := 0
		for {
			time.Sleep(time.Second * 5)
			w.label.SetText(fmt.Sprintf("%s %d", content, cnt))
			cnt++
		}
	}()
	go func() {
		w.show(false)
	}()
	app.Run()
}
