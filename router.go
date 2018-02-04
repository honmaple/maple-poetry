/*********************************************************************************
Copyright © 2018 jianglin
File Name: router.go
Author: jianglin
Email: xiyang0807@gmail.com
Created: 2018-01-30 13:41:18 (CST)
Last Update: 星期五 2018-02-02 19:04:36 (CST)
         By:
Description:
*********************************************************************************/
package main

import (
    "github.com/gin-gonic/gin"
    // "github.com/go-pg/pg/orm"
    "math/rand"
    "net/http"
    "strconv"
    "time"
)

type Poem struct {
    Router *gin.Engine
}

func (self *Poem) Init(prefix string) {
    api := self.Router.Group(prefix)
    {
        api.GET("", self.GET)
        api.GET("/:pk", self.GETITEM)
    }
    self.Router.GET("/poem", self.GETHTML)
}
func (self *Poem) GETHTML(c *gin.Context) {
    var poems []PoemModel
    err := db.Model(&poems).Limit(10).Select()
    if err != nil {
        panic(err)
    }
    new_poems := []P{}
    for _, poem := range poems {
        new_poems = append(new_poems, poem.ToJSON())
    }
    c.HTML(http.StatusOK, "poem.html", gin.H{
        "title": "Poem",
        "poems": new_poems,
    })
}

func (self *Poem) GET(c *gin.Context) {
    var poems []PoemModel

    offset, err := strconv.ParseInt(c.DefaultQuery("offset", "10"), 10, 64)
    if err != nil || offset > 100 {
        c.JSON(200, gin.H{
            "message": err,
            "data":    nil,
        })
        return
    }
    page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
    if err != nil {
        c.JSON(200, gin.H{
            "message": err,
            "data":    nil,
        })
        return
    }
    rand.Seed(time.Now().UnixNano())
    // 311828
    err = db.Model(&poems).Offset(int((page-1)*offset) + rand.Intn(1000)).Limit(int(offset)).Select()
    if err != nil {
        panic(err)
    }
    new_poems := []P{}
    for _, poem := range poems {
        new_poems = append(new_poems, poem.ToJSON())
    }

    c.JSON(200, gin.H{
        "message": "pong",
        "data":    new_poems,
    })
}

func (self *Poem) GETITEM(c *gin.Context) {
    pk, err := strconv.ParseInt(c.Param("pk"), 10, 64)
    if err != nil {
        c.JSON(200, gin.H{
            "message": err,
            "data":    nil,
        })
        return
    }
    poem := PoemModel{Id: pk}
    err = db.Select(&poem)
    if err != nil {
        panic(err)
    }
    c.JSON(200, gin.H{
        "message": "pong",
        "data":    poem.ToJSON(),
    })
}

type Author struct {
    Router *gin.Engine
}

func (self *Author) Init(prefix string) {
    api := self.Router.Group(prefix)
    {
        api.GET("", self.GET)
        api.GET("/:pk", self.GETITEM)
    }
    self.Router.GET("/author", self.GETHTML)
}

func (self *Author) GETHTML(c *gin.Context) {
    c.HTML(http.StatusOK, "author.html", gin.H{
        "title": "Main website",
    })
}

func (self *Author) GET(c *gin.Context) {
    var authors []AuthorModel
    offset, err := strconv.ParseInt(c.DefaultQuery("offset", "10"), 10, 64)
    if err != nil || offset > 100 {
        c.JSON(200, gin.H{
            "message": err,
            "data":    nil,
        })
        return
    }
    page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
    if err != nil {
        c.JSON(200, gin.H{
            "message": err,
            "data":    nil,
        })
        return
    }
    err = db.Model(&authors).Offset(int((page - 1) * offset)).Limit(int(offset)).Select()
    if err != nil {
        panic(err)
    }
    c.JSON(200, gin.H{
        "message": "pong",
        "data":    authors,
    })
}

func (self *Author) GETITEM(c *gin.Context) {
    pk, err := strconv.ParseInt(c.Param("pk"), 10, 64)
    if err != nil {
        c.JSON(200, gin.H{
            "message": err,
            "data":    nil,
        })
        return
    }
    author := AuthorModel{Id: pk}
    err = db.Select(&author)
    if err != nil {
        panic(err)
    }
    c.JSON(200, gin.H{
        "message": "pong",
        "data":    author,
    })
}
