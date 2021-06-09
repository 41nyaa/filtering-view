/*
 * Copyright (c) 2021 41nyaa
 * All rights reserved.
 */
 package main

import (
	"io/ioutil"
	"log"
	"regexp"
)

func NewData(names []string) []string {
	datas := make([]string, 0)
	if 0 == len(names) {
		return datas
	}
	var content string
	for _, name := range names {
		data, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}
		content += string(data)
	}

	sep := regexp.MustCompile(`\r|\n`)
	//	sep := regexp.MustCompile(`\r|\n\[[0-9]{4}\/`)
	datas = sep.Split(content, -1)
	return datas
}
