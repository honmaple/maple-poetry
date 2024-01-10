package router

import (
	"poetry/internal/app/config"
	"poetry/internal/app/model"
)

type Router struct {
	orm  *model.DB
	log  *config.Logger
	conf *config.Config
}

func New(conf *config.Config, log *config.Logger, db *config.DB) *Router {
	return &Router{
		orm:  model.New(conf, db),
		log:  log,
		conf: conf,
	}
}
