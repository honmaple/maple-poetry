package model

const (
	TAG_KIND_STYLE = "体裁"
)

type (
	Tag struct {
		BaseModel
		Name  string   `gorm:"size:36;not null;unique_index:idx_name"    json:"name"`
		Kind  string   `gorm:"size:36;not null;unique_index:idx_name"    json:"kind"`
		Poems []*Poems `gorm:"many2many:poem_tags;"                      json:"-"`
	}
	Tags []*Tag
)

func (db *DB) GetTags(opt *Option) (Tags, PageInfo, error) {
	q := db.Model(Tag{})

	if name := opt.GetString("name"); name != "" {
		q = q.Where("name = ?", name)
	}

	if kind := opt.GetString("kind"); kind != "" {
		q = q.Where("kind = ?", kind)
	}

	if sort := opt.GetString("sort"); sort == "desc" {
		q = q.Order("id DESC")
	}

	pageinfo := PageInfo{
		Page:  opt.GetInt("page"),
		Limit: opt.GetInt("limit"),
	}
	offset, limit := db.GetLimit(&pageinfo)
	q = q.Count(&pageinfo.Total).Offset(offset).Limit(limit)

	ins := make(Tags, 0)
	result := q.Find(&ins)
	return ins, pageinfo, result.Error
}

func (db *DB) GetTag(id string) (*Tag, error) {
	ins := new(Tag)
	result := db.First(ins, "id = ?", id)
	return ins, result.Error
}
