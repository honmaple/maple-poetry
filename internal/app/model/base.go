package model

import (
	"poetry/internal/app/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type (
	DB struct {
		*gorm.DB
		config *config.Config
	}
	BaseModel struct {
		Id uint `gorm:"primary_key;auto_increment;" json:"id"`
	}
)

type (
	PageInfo struct {
		Page     int   `json:"page"  query:"page"`
		Limit    int   `json:"limit" query:"limit"`
		Total    int64 `json:"total" query:"-"`
		NotLimit bool  `json:"-"     query:"-"`
	}
	List struct {
		PageInfo
		List interface{} `json:"list,omitempty"`
	}
	Option struct {
		*viper.Viper
	}
)

func (db *DB) GetLimit(info *PageInfo) (int, int) {
	if info.Page <= 0 {
		info.Page = 1
	}

	pagesize := db.config.GetInt("server.pagesize")
	if pagesize == 0 {
		pagesize = 10
	}

	if info.Limit <= 0 || info.Limit > pagesize {
		info.Limit = pagesize
	}
	offset := (info.Page - 1) * info.Limit
	return offset, info.Limit
}

func NewOption(m map[string]interface{}) *Option {
	cf := viper.New()
	for k, v := range m {
		cf.Set(k, v)
	}
	return &Option{cf}
}

func New(conf *config.Config, db *gorm.DB) *DB {
	return &DB{DB: db, config: conf}
}

func Init(conf *config.Config, db *gorm.DB) error {
	return db.Debug().AutoMigrate(
		new(Tag),
		new(Poem),
		new(Author),
		new(Dynasty),
		new(Collection),
	)
}
