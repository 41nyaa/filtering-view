/*
 * Copyright (c) 2021 41nyaa
 * All rights reserved.
 */
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type FilterView struct {
	filter []Filter
	ctr *fyne.Container
	callback func(map[string]bool)
}

func NewFilterView(filter []Filter, callback func(map[string]bool)) *FilterView {
	view := new(FilterView)
	view.filter = filter
	view.ctr = container.New(layout.NewVBoxLayout())
	view.callback = callback
	return view
}

func (view *FilterView) Display() {
	for i, _ := range view.filter {
		i := i
		check := widget.NewCheck(view.filter[i].Name, func(value bool) {
			view.filter[i].Displayed = value
			SaveFilter(view.filter)
			view.callback(ToMapFilter(view.filter))
		})
		check.SetChecked(view.filter[i].Displayed)
		view.ctr.Add(check)
	}
}
