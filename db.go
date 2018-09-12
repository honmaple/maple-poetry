/*********************************************************************************
 Copyright © 2018 jianglin
 File Name: db.go
 Author: jianglin
 Email: mail@honmaple.com
 Created: 2018-09-12 14:20:11 (CST)
 Last Update: Wednesday 2018-09-12 15:44:03 (CST)
		  By:
 Description:
 *********************************************************************************/
package main

import (
	"strings"
)

// Poem ..
type Poem struct {
	TableName struct{} `sql:"poem"`

	ID         int64
	Title      string
	Author     string
	Strains    string
	Paragraphs string
}

// Serializer ..
func (self *Poem) Serializer() map[string]interface{} {
	strains := make([]string, 0)
	for _, strain := range strings.Split(self.Strains, "。") {
		strains = append(strains, strings.Replace(strain, "，", " ", -1))
	}
	paragraphs := make([]string, 0)
	for _, paragraph := range strings.Split(self.Paragraphs, "。") {
		paragraphs = append(paragraphs, strings.Replace(paragraph, "，", " ", -1))
	}
	return map[string]interface{}{
		"title":      self.Title,
		"author":     self.Author,
		"strains":    strains,
		"paragraphs": paragraphs,
	}
}

// Author ..
type Author struct {
	TableName   struct{} `sql:"author"`
	ID          int64
	Name        string
	Description string
}

// Serializer ..
func (self *Author) Serializer() map[string]interface{} {
	return map[string]interface{}{
		"name":        self.Name,
		"description": self.Description,
	}
}
