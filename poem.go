/*********************************************************************************
Copyright © 2018 jianglin
File Name: poem.go
Author: jianglin
Email: xiyang0807@gmail.com
Created: 2018-01-30 13:39:49 (CST)
Last Update: 星期三 2018-01-31 01:55:01 (CST)
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
}

var (
    db     *pg.DB
    config Configuration
)

func main() {
    route := gin.Default()
    route.LoadHTMLGlob("templates/*")
    poem := Poem{Router: route}
    poem.Init("/api/poem")

    author := Author{Router: route}
    author.Init("/api/author")

    // err := DB.CreateTable(&InventoryModel{},nil)
    // if err != nil {
    //     panic(err)
    // }
    db = pg.Connect(&pg.Options{
        User:     config.PG.User,
        Password: config.PG.Password,
        Database: config.PG.Database,
    })

    route.Run("127.0.0.1:8000") // listen and serve on 0.0.0.0:8080
}

func init() {
    var (
        conf_file string
        print_ver bool
    )
    flag.StringVar(&conf_file, "c", "config.json", "config file")
    flag.BoolVar(&print_ver, "v", false, "get version")
    flag.Parse()

    raw, err := ioutil.ReadFile(conf_file)
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
}
