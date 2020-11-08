package tutorials

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func bindingScreen(_ fyne.Window) fyne.CanvasObject {
	data := binding.NewFloat()
	label := widget.NewLabelWithData(binding.FloatToString(data))

	slide := widget.NewSliderWithData(0, 1, data)
	slide.Step = 0.1
	bar := widget.NewProgressBarWithData(data)

	buttons := container.NewGridWithColumns(4,
		widget.NewButton("0%", func() {
			data.Set(0)
		}),
		widget.NewButton("30%", func() {
			data.Set(0.3)
		}),
		widget.NewButton("70%", func() {
			data.Set(0.7)
		}),
		widget.NewButton("100%", func() {
			data.Set(1)
		}))

	item := container.NewVBox(container.NewHBox(widget.NewLabel("Float current value:"), label),
		slide, bar, buttons)

	dataList := binding.NewFloatList()
	dataList.Append(0.1)
	dataList.Append(0.2)
	dataList.Append(0.3)

	button := widget.NewButton("Append", func() {
		dataList.Append(float64(dataList.Length()+1) / 10)
	})

	list := container.NewBorder(nil, button, nil, nil, widget.NewListWithData(dataList,
		func() fyne.CanvasObject {
			return container.NewBorder(nil, nil, nil, widget.NewButton("+", nil),
				widget.NewLabel("item x.y"))
		},
		func(item binding.DataItem, obj fyne.CanvasObject) {
			btn := obj.(*fyne.Container).Objects[1].(*widget.Button)
			text := obj.(*fyne.Container).Objects[0].(*widget.Label)
			btn.OnTapped = func() {
				item.(binding.Float).Set(item.(binding.Float).Get() + 1)
			}
			text.SetText(fmt.Sprintf("item %0.1f", item.(binding.Float).Get()))
		}))

	return container.NewBorder(item, nil, nil, nil, list)
}
