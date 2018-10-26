/*********************************************************************************
 Copyright © 2018 jianglin
 File Name: db.go
 Author: jianglin
 Email: mail@honmaple.com
 Created: 2018-09-12 14:20:11 (CST)
 Last Update: Monday 2018-10-22 19:02:09 (CST)
		  By:
 Description:
 *********************************************************************************/
package main

import (
	"github.com/go-pg/pg"
	"strings"
)

// DBClient ..
func DBClient() *pg.DB {
	client := pg.Connect(&pg.Options{
		User:     config.PG.User,
		Password: config.PG.Password,
		Database: config.PG.Database,
	})
	return client
}

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
		strain = strings.Replace(strain, "，", " ", -1)
		strains = append(strains, strain)
	}
	paragraphs := make([]string, 0)
	for _, paragraph := range strings.Split(self.Paragraphs, "。") {
		paragraph = strings.Replace(paragraph, "，", " ", -1)
		paragraphs = append(paragraphs, paragraph)
	}
	if strains[len(strains)-1] == "" {
		strains = strains[:len(strains)-1]
	}
	if paragraphs[len(paragraphs)-1] == "" {
		paragraphs = paragraphs[:len(paragraphs)-1]
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
		"author": self.Name,
		"desc":   self.Description,
	}
}
