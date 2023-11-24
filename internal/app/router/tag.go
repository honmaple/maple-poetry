package router

import (
	"poetry/internal/app/model"

	"github.com/honmaple/forest"
	"github.com/honmaple/forest-contrib/response"
)

func (r *Router) GetTags(c forest.Context) error {
	m := make(map[string]interface{})
	if err := c.Bind(&m); err != nil {
		return err
	}

	ins, pageinfo, err := r.orm.GetTags(model.NewOption(m))
	if err != nil {
		return err
	}
	return response.OK(c, "", model.List{pageinfo, ins})
}

func (r *Router) GetTag(c forest.Context) error {
	id := c.Param("id")

	ins, err := r.orm.GetTag(id)
	if err != nil {
		return err
	}
	return response.OK(c, "", ins)
}
