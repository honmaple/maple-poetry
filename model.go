/*********************************************************************************
Copyright © 2018 jianglin
File Name: model.go
Author: jianglin
Email: xiyang0807@gmail.com
Created: 2018-01-30 14:13:43 (CST)
Last Update: 星期三 2018-01-31 01:45:44 (CST)
         By:
Description:
*********************************************************************************/
package main

import (
    "github.com/go-pg/pg/orm"
    "strings"
    // "strconv"
)

// type Objects struct {
// }

// func (self *Objects) Filter(params map[string]string) Objects {

// }

// func (self *Objects) Pagination(_page int, _number string) {
//     number, err := strconv.ParseInt(_number, 10, 64)
//     if err != nil || number > 100 {
//         return "a", false
//     }
//     page, err := strconv.ParseInt(_page, 10, 64)
//     if err != nil {
//         return "a", false
//     }
// }

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

func (self PoemModel) ToJson(poems []PoemModel) []P {
    var new_poems []P
    for _, poem := range poems {
        strains := []string{}
        for _, strain := range strings.Split(poem.Strains, "。") {
            strains = append(strains, strings.Replace(strain, "，", " ", -1))
        }
        paragraphs := []string{}
        for _, paragraph := range strings.Split(poem.Paragraphs, "。") {
            paragraphs = append(paragraphs, strings.Replace(paragraph, "，", " ", -1))
        }
        new_poem := P{
            Title:      poem.Title,
            Author:     poem.Author,
            Strains:    strains,
            Paragraphs: paragraphs,
        }
        new_poems = append(new_poems, new_poem)
    }
    return new_poems
}

type AuthorModel struct {
    BaseModel
    TableName struct{} `sql:"author"`

    Id          int64
    Name        string
    Description string
}
