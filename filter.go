/*
 * Copyright (c) 2021 41nyaa
 * All rights reserved.
 */
package main

import (
	"encoding/json"
	"io/ioutil"
)

type Filter struct {
	Name      string
	Displayed bool
}

func ReadFilter(name string) []Filter {
	fil := []Filter{}

	jsonString, err := ioutil.ReadFile("filter.json")
	if err != nil {
		return fil
	}

	err = json.Unmarshal(jsonString, &fil)
	if err != nil {
		return fil
	}

	return fil
}

func SaveFilter(filter []Filter) {
	b, err := json.Marshal(filter)
	if err != nil {
		return
	}
	err = ioutil.WriteFile("filter.json", b, 0644)
	if err != nil {
		return
	}
}

func ToMapFilter(filter []Filter) map[string]bool {
	m := map[string]bool{}
	for _, f := range filter {
		m[f.Name] = f.Displayed
	}
	return m
}