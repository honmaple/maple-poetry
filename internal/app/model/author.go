package model

import (
	"fmt"
)

type (
	Author struct {
		BaseModel
		Name      string `gorm:"size:36;not null;unique_index:idx_name" json:"name"`
		Desc      string `                                              json:"desc"`
		ShortDesc string `                                              json:"short_desc"`

		Dynasty   *Dynasty `gorm:"foreignKey:DynastyId;"                json:"dynasty"`
		DynastyId uint     `gorm:"not null;unique_index:idx_name"       json:"dynasty_id"`

		Prev *Author `gorm:"-"                                          json:"prev"`
		Next *Author `gorm:"-"                                          json:"next"`
	}
	Authors []*Author
)

func (db *DB) GetAuthors(opt *Option) (Authors, PageInfo, error) {
	q := db.Model(Author{}).Preload("Dynasty")
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
	offset, limit := db.GetLimit(&pageinfo)
	q = q.Count(&pageinfo.Total).Offset(offset).Limit(limit)

	ins := make(Authors, 0)
	result := q.Find(&ins)
	return ins, pageinfo, result.Error
}

func (db *DB) GetAuthor(id string) (*Author, error) {
	ins := new(Author)
	if err := db.Preload("Dynasty").First(ins, "id = ?", id).Error; err != nil {
		return nil, err
	}

	prev := new(Author)
	if err := db.Order("id DESC").First(prev, "id < ?", ins.Id).Error; err == nil {
		ins.Prev = prev
	}

	next := new(Author)
	if err := db.Order("id").First(next, "id > ?", ins.Id).Error; err == nil {
		ins.Next = next
	}
	return ins, nil
}

func (db *DB) GetRandomAuthor() (*Author, error) {
	ins := new(Author)
	result := db.Preload("Dynasty").Order("RANDOM()").Limit(1).Find(&ins)
	return ins, result.Error
}
