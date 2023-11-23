package model

import (
	"fmt"
)

type (
	Poem struct {
		BaseModel
		Title   string `json:"title"`
		Chapter string `json:"chapter"`
		Section string `json:"section"`
		Content string `json:"content"`
		Note    string `json:"note"`

		Author         *Author       `json:"author"`
		AuthorId       uint        `json:"author_id"`
		Dynasty      *Dynasty    `json:"dynasty"`
		DynastyId    uint        `json:"dynasty_id"`
		Collection   *Collection `json:"collection"`
		CollectionId uint        `json:"collection_id"`
	}
	Poems []*Poem
)

func (db *DB) GetPoems(opt *Option) (Poems, PageInfo, error) {
	q := db.Model(Poem{}).Preload("Author")

	if title := opt.GetString("title"); title != "" {
		q = q.Where("title LIKE ?", fmt.Sprintf("%%%s%%", title))
	}

	if author := opt.GetInt("author"); author > 0 {
		q = q.Where("author_id = ?", author)
	} else if author := opt.GetString("author"); author != "" {
		q = q.Where("author_id in (?)", db.Model(Author{}).Select("id").Where("name = ?", author))
	}

	if dynasty := opt.GetInt("dynasty"); dynasty > 0 {
		q = q.Where("dynasty_id = ?", dynasty)
	} else if dynasty := opt.GetString("dynasty"); dynasty != "" {
		q = q.Where("dynasty_id in (?)", db.Model(Dynasty{}).Select("id").Where("name = ?", dynasty))
	}

	if collection := opt.GetInt("collection"); collection > 0 {
		q = q.Where("collection_id = ?", collection)
	} else if collection := opt.GetString("collection"); collection != "" {
		q = q.Where("collection_id in (?)", db.Model(Collection{}).Select("id").Where("name = ?", collection))
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

	ins := make(Poems, 0)
	result := q.Find(&ins)
	return ins, pageinfo, result.Error
}

func (db *DB) GetPoem(id string) (*Poem, error) {
	ins := new(Poem)
	result := db.Preload("Author").Preload("Dynasty").First(ins, "id = ?", id)
	return ins, result.Error
}

func (db *DB) GetRandomPoem() (*Poem, error) {
	ins := new(Poem)
	result := db.Preload("Author").Preload("Dynasty").Order("RANDOM()").Limit(1).Find(&ins)
	return ins, result.Error
}
