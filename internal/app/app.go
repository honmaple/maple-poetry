package app

import (
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"poetry/internal/app/config"
	"poetry/internal/app/model"
	"poetry/internal/app/router"

	"github.com/honmaple/forest"
	"github.com/honmaple/forest/middleware"
)

type App struct {
	DB     *model.DB
	Log    *config.Logger
	Config *config.Config
}

func (app *App) Run(webFS fs.FS) error {
	conf := app.Config

	srv := forest.New()
	if conf.GetString("server.mode") == "dev" {
		srv.SetOptions(forest.Debug())
	}

	corsConfig := middleware.CorsConfig{
		AllowOrigins: conf.GetStringSlice("server.cors.allow_origins"),
		AllowMethods: conf.GetStringSlice("server.cors.allow_methods"),
		AllowHeaders: conf.GetStringSlice("server.cors.allow_headers"),
	}
	if len(corsConfig.AllowOrigins) == 0 {
		corsConfig.AllowOrigins = middleware.DefaultCorsConfig.AllowOrigins
	}
	if len(corsConfig.AllowMethods) == 0 {
		corsConfig.AllowMethods = middleware.DefaultCorsConfig.AllowMethods
	}
	if len(corsConfig.AllowHeaders) == 0 {
		corsConfig.AllowHeaders = middleware.DefaultCorsConfig.AllowHeaders
	}

	srv.Use(middleware.Logger(), middleware.CorsWithConfig(corsConfig))
	{
		srv.GET("/", func(c forest.Context) error {
			return c.FileFromFS("index.html", http.FS(webFS))
		})
		srv.GET("/icons/*", func(c forest.Context) error {
			path := filepath.Join("icons", c.Param("*"))
			return c.FileFromFS(path, http.FS(webFS))
		})
		srv.GET("/assets/*", func(c forest.Context) error {
			path := filepath.Join("assets", c.Param("*"))
			return c.FileFromFS(path, http.FS(webFS))
		})
		srv.GET("/robots.txt", func(c forest.Context) error {
			return c.FileFromFS("robots.txt", http.FS(webFS))
		})
		srv.GET("/favicon.ico", func(c forest.Context) error {
			return c.FileFromFS("favicon.ico", http.FS(webFS))
		})
	}

	r := router.New(app.Config, app.Log, app.DB)
	bp := srv.Group(forest.WithPrefix("/api"))
	{
		bp.GET("/tags", r.GetTags)
		bp.GET("/tags/:id", r.GetTag)
	}
	{
		bp.GET("/poems", r.GetPoems)
		bp.GET("/poems/:id", r.GetPoem)
		bp.GET("/poems/random", r.GetRandomPoem)
	}
	{
		bp.GET("/authors", r.GetAuthors)
		bp.GET("/authors/:id", r.GetAuthor)
		bp.GET("/authors/random", r.GetRandomAuthor)
	}
	{
		bp.GET("/dynasties", r.GetDynasties)
		bp.GET("/dynasties/:id", r.GetDynasty)
	}
	{
		bp.GET("/collections", r.GetCollections)
		bp.GET("/collections/:id", r.GetCollection)
	}
	return srv.Start(conf.GetString("server.addr"))
}

func (app *App) Init(file string, strs ...string) error {
	conf := app.Config

	if _, err := os.Stat(file); err == nil || os.IsExist(err) {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		conf.SetConfigFile(file)
		if err := conf.ReadConfig(strings.NewReader(os.ExpandEnv(string(content)))); err != nil {
			return err
		}
	}

	for _, str := range strs {
		c := strings.SplitN(str, "=", 2)
		if len(c) < 2 {
			conf.Set(c[0], "")
		} else {
			conf.Set(c[0], c[1])
		}
	}

	db, err := config.NewDB(conf)
	if err != nil {
		return err
	}
	log, err := config.NewLogger(conf)
	if err != nil {
		return err
	}
	app.DB = model.New(conf, db)
	app.Log = log
	return nil
}

func New() *App {
	return &App{
		Config: config.Default(),
	}
}
