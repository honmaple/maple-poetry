/*********************************************************************************
Copyright © 2018 jianglin
File Name: model.go
Author: jianglin
Email: xiyang0807@gmail.com
Created: 2018-01-30 14:13:43 (CST)
Last Update: 星期一 2018-02-05 00:17:53 (CST)
         By:
Description:
*********************************************************************************/
package main

import (
    "github.com/go-pg/pg/orm"
    "strings"
)

type BaseModel struct {
}

func (cls BaseModel) Query() *orm.Query {
    return db.Model(&cls)
}

type PoemModel struct {
    BaseModel
    TableName struct{} `sql:"poet"`

    Id         int64
    Title      string
    Author     string
    Strains    string
    Paragraphs string
}

type P struct {
    Title      string
    Author     string
    Strains    []string
    Paragraphs []string
}

func (self *PoemModel) ToJSON() P {
    strains := []string{}
    for _, strain := range strings.Split(self.Strains, "。") {
        strains = append(strains, strings.Replace(strain, "，", " ", -1))
    }
    paragraphs := []string{}
    for _, paragraph := range strings.Split(self.Paragraphs, "。") {
        paragraphs = append(paragraphs, strings.Replace(paragraph, "，", " ", -1))
    }
    poem := P{
        Title:      self.Title,
        Author:     self.Author,
        Strains:    strains,
        Paragraphs: paragraphs,
    }
    return poem
}

type AuthorModel struct {
    BaseModel
    TableName struct{} `sql:"author"`

    Id          int64
    Name        string
    Description string
}

type A struct {
    Name        string
    Description string
}

func (self *AuthorModel) ToJSON() A {
    return A{
        Name:        self.Name,
        Description: self.Description,
    }
}
