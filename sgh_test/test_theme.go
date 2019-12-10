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

func createWinCore(title string, borderless bool) fyne.Window {
	w := app.Driver().CreateSghWindow(title, borderless)
	w.CenterOnScreen()
	w.SetFixedSize(true)
	return w
}

func (w *win) Exit() {
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
			time.Sleep(time.Second * 3)
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
