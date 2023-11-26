package model

type (
	Dynasty struct {
		BaseModel
		Name string `gorm:"size:32;not null;unique" json:"name"`
		Desc string `                               json:"desc"`
	}
	Dynasties []*Dynasty
)

func (db *DB) GetDynasties(opt *Option) (Dynasties, PageInfo, error) {
	q := db.Model(Dynasty{})

	if name := opt.GetString("name"); name != "" {
		q = q.Where("name = ?", name)
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

	ins := make(Dynasties, 0)
	result := q.Find(&ins)
	return ins, pageinfo, result.Error
}

func (db *DB) GetDynasty(id string) (*Dynasty, error) {
	ins := new(Dynasty)
	result := db.First(ins, "id = ?", id)
	return ins, result.Error
}
