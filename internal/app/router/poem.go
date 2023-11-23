package router

import (
	"poetry/internal/app/config"
	"poetry/internal/app/model"

	"github.com/honmaple/forest"
	"github.com/honmaple/forest-contrib/response"
)

type Router struct {
	orm  *model.DB
	log  *config.Logger
	conf *config.Config
}

func New(conf *config.Config, log *config.Logger, orm *model.DB) *Router {
	return &Router{
		orm:  orm,
		log:  log,
		conf: conf,
	}
}

func (r *Router) GetPoems(c forest.Context) error {
	m := make(map[string]interface{})
	if err := c.Bind(&m); err != nil {
		return err
	}

	ins, pageinfo, err := r.orm.GetPoems(model.NewOption(m))
	if err != nil {
		return err
	}
	return response.OK(c, "", model.List{pageinfo, ins})
}

func (r *Router) GetPoem(c forest.Context) error {
	id := c.Param("id")

	ins, err := r.orm.GetPoem(id)
	if err != nil {
		return err
	}
	return response.OK(c, "", ins)
}

func (r *Router) GetRandomPoem(c forest.Context) error {
	ins, err := r.orm.GetRandomPoem()
	if err != nil {
		return err
	}
	return response.OK(c, "", ins)
}
