/*********************************************************************************
Copyright © 2018 jianglin
File Name: poem.go
Author: jianglin
Email: xiyang0807@gmail.com
Created: 2018-01-30 13:39:49 (CST)
Last Update: 星期日 2018-02-04 23:56:25 (CST)
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

type PGConfiguration struct {
    User     string
    Password string
    Database string
}

type Configuration struct {
    PG      *PGConfiguration
    Version string
    Server  string
    Debug   bool
}

var (
    db     *pg.DB
    config Configuration
)

func main() {
    if config.Debug == false {
        gin.SetMode(gin.ReleaseMode)
    }
    route := gin.Default()
    route.Static("/static", "./static")
    route.LoadHTMLGlob("templates/*")
    poem := Poem{Router: route}
    poem.Init("/api/poem")

    author := Author{Router: route}
    author.Init("/api/author")

    db = pg.Connect(&pg.Options{
        User:     config.PG.User,
        Password: config.PG.Password,
        Database: config.PG.Database,
    })

    route.Run(config.Server)
}

func initdb() {
    for _, model := range []interface{}{&PoemModel{}, &AuthorModel{}} {
        err := db.CreateTable(model, nil)
        if err != nil {
            panic(err)
        }
    }
    os.Exit(0)
}

func init() {
    var (
        conf      string
        print_ver bool
        db_init   bool
    )
    flag.StringVar(&conf, "c", "config.json", "config file")
    flag.BoolVar(&print_ver, "v", false, "get version")
    flag.BoolVar(&db_init, "initdb", false, "init db")
    flag.Parse()

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

    if print_ver {
        fmt.Println("version:", config.Version)
        os.Exit(0)
    }
    if db_init {
        initdb()
    }
}
