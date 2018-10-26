/*********************************************************************************
Copyright © 2018 jianglin
File Name: main.go
Author: jianglin
Email: xiyang0807@gmail.com
Created: 2018-01-30 13:39:49 (CST)
Last Update: Friday 2018-10-26 13:13:20 (CST)
		 By:
Description:
*********************************************************************************/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"io/ioutil"
	"os"
)

// VERSION ..
const VERSION = "0.1.0"

// PGConfiguration ..
type PGConfiguration struct {
	User     string
	Password string
	Database string
}

// CORSConfiguration ..
type CORSConfiguration struct {
	AllowOrigin string `json:"allow_origin"`
	AllowMethod string `json:"allow_method"`
	AllowHeader string `json:"allow_header"`
}

// Configuration ..
type Configuration struct {
	PG     *PGConfiguration
	CORS   *CORSConfiguration
	Server string
	Debug  bool
}

var (
	db     *pg.DB
	config Configuration
)

// CORSMiddleware ..
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.CORS.AllowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Headers", config.CORS.AllowHeader)
		c.Writer.Header().Set("Access-Control-Allow-Methods", config.CORS.AllowMethod)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// main ..
func main() {
	if config.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	poem := new(PoemAPI)
	poem.Init(r, "/api/poem")

	author := new(AuthorAPI)
	author.Init(r, "/api/author")

	db = pg.Connect(&pg.Options{
		User:     config.PG.User,
		Password: config.PG.Password,
		Database: config.PG.Database,
	})

	r.Run(config.Server)
}

// initDB ..
func initDB() {
	for _, model := range []interface{}{&Poem{}, &Author{}} {
		err := db.CreateTable(model, nil)
		if err != nil {
			panic(err)
		}
	}
	return
}

// init ..
func init() {
	var (
		conf    string
		version bool
		initdb  bool
	)
	flag.StringVar(&conf, "c", "config.json", "config file")
	flag.BoolVar(&version, "v", false, "get version")
	flag.BoolVar(&initdb, "initdb", false, "init db")
	flag.Parse()

	if version {
		fmt.Println("version:", VERSION)
		os.Exit(0)
	}

	raw, err := ioutil.ReadFile(conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(raw, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if initdb {
		initDB()
	}
}
