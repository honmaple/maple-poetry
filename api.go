/*********************************************************************************
 Copyright © 2018 jianglin
 File Name: api.go
 Author: jianglin
 Email: mail@honmaple.com
 Created: 2018-09-03 16:00:48 (CST)
 Last Update: Tuesday 2018-12-25 13:38:09 (CST)
		  By:
 Description:
 *********************************************************************************/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// PoemAPI ..
type PoemAPI struct {
}

// GETHTML ..
func (api *PoemAPI) GETHTML(c *gin.Context) {
	ins := make([]*Poem, 0)
	err := db.Model(&ins).Where("random() < ?", 0.01).Limit(10).Select()
	if err != nil {
		HTTP{c}.ServerError("", nil)
		return
	}

	c.HTML(http.StatusOK, "poem.html", gin.H{
		"title": "Poem",
		"poems": Serializer(ins),
	})
}

// RANDOM ..
func (api *PoemAPI) RANDOM(c *gin.Context) {
	ins := new(Poem)
	err := db.Model(ins).Where("random() < ?", 0.01).Limit(1).Select()
	if err != nil {
		HTTP{c}.ServerError("", nil)
	}
	HTTP{c}.OK("", ins.Serializer())
}

// GET ..
func (api *PoemAPI) GET(c *gin.Context) {
	ins := make([]*Poem, 0)

	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil || limit > 100 {
		HTTP{c}.BadRequest("", nil)
		return
	}
	err = db.Model(&ins).Where("random() < ?", 0.01).Limit(int(limit)).Select()
	if err != nil {
		HTTP{c}.ServerError("", nil)
	}

	HTTP{c}.OK("", Serializer(ins))
}

// GETITEM ..
func (api *PoemAPI) GETITEM(c *gin.Context) {
	pk, err := strconv.ParseInt(c.Param("pk"), 10, 64)
	if err != nil {
		HTTP{c}.BadRequest("pk param error", nil)
		return
	}
	ins := Poem{ID: pk}
	err = db.Select(&ins)
	if err != nil {
		HTTP{c}.ServerError("", nil)
	}

	HTTP{c}.OK("", Serializer(ins))
}

// AuthorAPI ..
type AuthorAPI struct {
}

// GETHTML ..
func (api *AuthorAPI) GETHTML(c *gin.Context) {
	ins := make([]*Author, 0)
	err := db.Model(&ins).Where("random() < ?", 0.01).Limit(10).Select()
	if err != nil {
		HTTP{c}.ServerError("", nil)
	}
	c.HTML(http.StatusOK, "author.html", gin.H{
		"title":   "Author",
		"authors": Serializer(ins),
	})
}

// RANDOM ..
func (api *AuthorAPI) RANDOM(c *gin.Context) {
	ins := new(Author)
	err := db.Model(ins).Where("random() < ?", 0.01).Limit(1).Select()
	if err != nil {
		HTTP{c}.ServerError("", nil)
	}
	HTTP{c}.OK("", ins.Serializer())
}

// GET ..
func (api *AuthorAPI) GET(c *gin.Context) {
	ins := make([]*Author, 0)
	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil || limit > 100 {
		HTTP{c}.BadRequest("", nil)
		return
	}
	err = db.Model(&ins).Where("random() < ?", 0.01).Limit(int(limit)).Select()
	if err != nil {
		HTTP{c}.ServerError("", nil)
	}

	HTTP{c}.OK("", Serializer(ins))
}

// GETITEM ..
func (api *AuthorAPI) GETITEM(c *gin.Context) {
	pk, err := strconv.ParseInt(c.Param("pk"), 10, 64)
	if err != nil {
		HTTP{c}.BadRequest("pk param error", nil)
		return
	}
	ins := Author{ID: pk}
	err = db.Select(&ins)
	if err != nil {
		HTTP{c}.ServerError("", nil)
	}

	HTTP{c}.OK("", Serializer(ins))
}

// Init ..
func (api *PoemAPI) Init(router *gin.Engine, prefix string) {
	r := router.Group(prefix)
	{
		r.GET("", api.GET)
		r.GET("/:pk", func(c *gin.Context) {
			if strings.HasPrefix(c.Request.RequestURI, prefix+"/random") {
				api.RANDOM(c)
				return
			}
			api.GETITEM(c)
		})
	}
	router.GET("/", api.GETHTML)
	router.GET("/poem", api.GETHTML)
}

// Init ..
func (api *AuthorAPI) Init(router *gin.Engine, prefix string) {
	r := router.Group(prefix)
	{
		r.GET("", api.GET)
		r.GET("/:pk", func(c *gin.Context) {
			if strings.HasPrefix(c.Request.RequestURI, prefix+"/random") {
				api.RANDOM(c)
				return
			}
			api.GETITEM(c)
		})
	}
	router.GET("/author", api.GETHTML)
}
