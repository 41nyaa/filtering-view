/*
 * Copyright (c) 2021 41nyaa
 * All rights reserved.
 */
package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
)

type DataView struct {
	datas []string
	filter map[string]bool
	ctr *fyne.Container
}

func NewDataView(datas []string, filter map[string]bool) *DataView {
	view := new(DataView)
	view.datas = datas
	view.filter = filter
//	view.editor = widget.NewTextGrid()
//	view.editor.Resize(fyne.NewSize(700, 450))
	row :=  canvas.NewText("", color.White)
	view.ctr = container.NewVBox(row)

	return view
}

func (view *DataView)Display() {
	// content := ""
	// for _, l := range view.logs {
	// 	if !view.filter[l.kind] {
	// 		content += l.raw
	// 		content += "\n"
	// // 		// ld := calcLineDiff(view.tmax, l.tdiff)
	// // 		// for i := 0; i < ld; i++ {
	// // 		// 	view.editor.SetRow(row, empty)
	// // 		// }
	// 	}
	// }
	// if 0 != len(content) {
	// 	content = content[:len(content)-2]
	// }
	// view.editor.SetText(content)

	view.ctr.Objects = nil
//	rows := []fyne.CanvasObject{}
	for _, l := range view.datas {
		if !view.filter[l] {
			row :=  canvas.NewText(l, color.White)
			view.ctr.Add(row)
		}
	}
}

func (view *DataView)Update(datas []string) {
	view.datas = datas
	view.Display()
}

func (view *DataView)Execute(filter map[string]bool) {
	view.filter = filter
	view.Display()
}
