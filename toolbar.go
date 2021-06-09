/*
 * Copyright (c) 2021 41nyaa
 * All rights reserved.
 */
 package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewToolbar(window fyne.Window, callback func([]string)) *widget.Toolbar {

	items := []widget.ToolbarItem{}
	file := widget.NewToolbarAction(theme.FileIcon(), func() {
		dialog.ShowFileOpen(func(path fyne.URIReadCloser, err error) {
			p := path.URI().Path()
			names := []string{p}
			file := NewData(names)
			callback(file)
		}, window)
	})
	items = append(items, file)
	folder := widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		dialog.ShowFolderOpen(func(path fyne.ListableURI, err error) {
			paths, _ := path.List()
			names := []string{}
			for _, d := range paths {
				if d.Extension() == ".txt" {
					names = append(names, d.Path())
				}
			}
			file := NewData(names)
			callback(file)
		}, window)
	})
	items = append(items, folder)
	return widget.NewToolbar(items...)
}
