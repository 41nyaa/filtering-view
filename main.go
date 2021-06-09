/*
 * Copyright (c) 2021 41nyaa
 * All rights reserved.
 */
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("filtering-view")
	myWindow.Resize(fyne.Size{Width: 800, Height: 500})

	names := make([]string, 0)
	data := NewData(names)
	filter := ReadFilter("filter.json")

	dataview := NewDataView(data, ToMapFilter(filter))
	dataview.Display()
	scrollview := container.NewScroll(dataview.ctr)
	scrollview.SetMinSize(fyne.NewSize(700, 450))

	filterview := NewFilterView(filter, dataview.Execute)
	filterview.Display()

	toolbar := NewToolbar(myWindow, dataview.Update)

	content := container.NewBorder(toolbar, nil, scrollview, filterview.ctr)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
