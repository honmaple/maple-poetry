package model

import (
	"fmt"
)

type (
	Collection struct {
		BaseModel
		Name string `gorm:"size:36;not null;unique_index:idx_name" json:"name"`
		Desc string `                                              json:"desc"`

		Dynasty   *Dynasty `gorm:"foreignKey:DynastyId;"           json:"dynasty"`
		DynastyId uint     `gorm:"not null;unique_index:idx_name"  json:"dynasty_id"`
	}
	Collections []*Collection
)

func (db *DB) GetCollections(opt *Option) (Collections, PageInfo, error) {
	q := db.Model(Collection{}).Preload("Dynasty")

	if name := opt.GetString("name"); name != "" {
		q = q.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}

	if dynasty := opt.GetInt("dynasty"); dynasty > 0 {
		q = q.Where("dynasty_id = ?", dynasty)
	} else if dynasty := opt.GetString("dynasty"); dynasty != "" {
		q = q.Where("dynasty_id in (?)", db.Model(Dynasty{}).Select("id").Where("name = ?", dynasty))
	}

	if sort := opt.GetString("sort"); sort == "random" {
		q = q.Order("RANDOM()")
	} else if sort == "desc" {
		q = q.Order("id DESC")
	}

	pageinfo := PageInfo{
		Page:  opt.GetInt("page"),
		Limit: opt.GetInt("limit"),
	}
	q = q.Count(&pageinfo.Total)

	// 不限制数量时默认显示全部
	if pageinfo.Limit > 0 {
		offset, limit := db.GetLimit(&pageinfo)
		q = q.Offset(offset).Limit(limit)
	}

	ins := make(Collections, 0)
	result := q.Find(&ins)
	return ins, pageinfo, result.Error
}

func (db *DB) GetCollection(id string) (*Collection, error) {
	ins := new(Collection)
	result := db.Preload("Dynasty").First(ins, "id = ?", id)
	return ins, result.Error
}
