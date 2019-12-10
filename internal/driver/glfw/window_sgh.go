package glfw

import (
	"fyne.io/fyne"
	"fyne.io/fyne/internal/painter/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func (d *gLDriver) CreateSghWindow(title string, borderless bool) fyne.Window {
	var ret *window
	runOnMain(func() {
		initOnce.Do(d.initGLFW)

		// make the window hidden, we will set it up and then show it later
		glfw.WindowHint(glfw.Visible, 0)
		if borderless {
			glfw.WindowHint(glfw.Decorated, 0)
		}
		initWindowHints()

		win, err := glfw.CreateWindow(10, 10, title, nil, nil)
		if err != nil {
			fyne.LogError("window creation error", err)
			return
		}
		win.MakeContextCurrent()

		ret = &window{viewport: win, title: title}

		// This channel will be closed when the window is closed.
		ret.eventQueue = make(chan func(), 1024)
		go ret.runEventQueue()

		ret.canvas = newCanvas()
		ret.canvas.painter = gl.NewPainter(ret.canvas, ret)
		ret.canvas.painter.Init()
		ret.canvas.context = ret
		ret.canvas.detectedScale = ret.detectScale()
		ret.canvas.scale = ret.selectScale()
		ret.SetIcon(ret.icon)
		d.windows = append(d.windows, ret)

		win.SetCloseCallback(ret.closed)
		win.SetPosCallback(ret.moved)
		win.SetSizeCallback(ret.resized)
		win.SetFramebufferSizeCallback(ret.frameSized)
		win.SetRefreshCallback(ret.refresh)
		win.SetCursorPosCallback(ret.mouseMoved)
		win.SetMouseButtonCallback(ret.mouseClicked)
		win.SetScrollCallback(ret.mouseScrolled)
		win.SetKeyCallback(ret.keyPressed)
		win.SetCharModsCallback(ret.charModInput)
		win.SetFocusCallback(ret.focused)
		glfw.DetachCurrentContext()
	})
	return ret
}

func (w *window) GetSghGlfwWindow() *glfw.Window {
	return w.viewport
}
