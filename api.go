/*********************************************************************************
 Copyright © 2018 jianglin
 File Name: api.go
 Author: jianglin
 Email: mail@honmaple.com
 Created: 2018-09-03 16:00:48 (CST)
 Last Update: Wednesday 2018-09-12 15:45:32 (CST)
		  By:
 Description:
 *********************************************************************************/
package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// PoemAPI ..
type PoemAPI struct {
}

// GETHTML ..
func (api *PoemAPI) GETHTML(c *gin.Context) {
	ins := make([]*Poem, 0)
	err := db.Model(&ins).Limit(10).Select()
	if err != nil {
		HTTP{c}.ServerError("", nil)
		return
	}
	serializer := PoemSerializer{Instances: ins}
	c.HTML(http.StatusOK, "poem.html", gin.H{
		"title": "Poem",
		"poems": serializer.Data(),
	})
}

// GET ..
func (api *PoemAPI) GET(c *gin.Context) {
	ins := make([]*Poem, 0)

	offset, err := strconv.ParseInt(c.DefaultQuery("offset", "10"), 10, 64)
	if err != nil || offset > 100 {
		HTTP{c}.BadRequest("", nil)
		return
	}
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		HTTP{c}.BadRequest("", nil)
		return
	}
	rand.Seed(time.Now().UnixNano())
	// 311828
	err = db.Model(&ins).Offset(int((page-1)*offset) + rand.Intn(1000)).Limit(int(offset)).Select()
	if err != nil {
		panic(err)
	}

	serializer := PoemSerializer{Instances: ins}
	HTTP{c}.OK("", serializer.Data())
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

	serializer := PoemSerializer{Instance: &ins}
	HTTP{c}.OK("", serializer.Data())
}

// AuthorAPI ..
type AuthorAPI struct {
}

// GETHTML ..
func (api *AuthorAPI) GETHTML(c *gin.Context) {
	ins := make([]*Author, 0)
	err := db.Model(&ins).Limit(10).Select()
	if err != nil {
		HTTP{c}.ServerError("", nil)
	}
	serializer := AuthorSerializer{Instances: ins}
	c.HTML(http.StatusOK, "author.html", gin.H{
		"title":   "Author",
		"authors": serializer.Data(),
	})
}

// GET ..
func (api *AuthorAPI) GET(c *gin.Context) {
	ins := make([]*Author, 0)
	offset, err := strconv.ParseInt(c.DefaultQuery("offset", "10"), 10, 64)
	if err != nil || offset > 100 {
		HTTP{c}.BadRequest("offset param error", nil)
		return
	}

	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		HTTP{c}.BadRequest("page param error", nil)
		return
	}
	err = db.Model(&ins).Offset(int((page - 1) * offset)).Limit(int(offset)).Select()
	if err != nil {
		panic(err)
	}

	serializer := AuthorSerializer{Instances: ins}
	HTTP{c}.OK("", serializer.Data())
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

	serializer := AuthorSerializer{Instance: &ins}
	HTTP{c}.OK("", serializer.Data())
}

// Init ..
func (api *PoemAPI) Init(router *gin.Engine, prefix string) {
	r := router.Group(prefix)
	{
		r.GET("", api.GET)
		r.GET("/:pk", api.GETITEM)
	}
	router.GET("/", api.GETHTML)
	router.GET("/poem", api.GETHTML)
}

// Init ..
func (api *AuthorAPI) Init(router *gin.Engine, prefix string) {
	r := router.Group(prefix)
	{
		r.GET("", api.GET)
		r.GET("/:pk", api.GETITEM)
	}
	router.GET("/author", api.GETHTML)
}
