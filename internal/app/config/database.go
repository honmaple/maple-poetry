package config

import (
	"errors"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DB = gorm.DB

func NewDB(conf *Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
		dsn = conf.GetString("database.dsn")
	)
	gcfg := &gorm.Config{
		Logger: logger.Default,
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	switch {
	case strings.HasPrefix(dsn, "mysql://"):
		db, err = gorm.Open(mysql.Open(dsn[8:]), gcfg)
	case strings.HasPrefix(dsn, "sqlite://"):
		db, err = gorm.Open(sqlite.Open(dsn[9:]), gcfg)
	case strings.HasPrefix(dsn, "postgres://"):
		db, err = gorm.Open(postgres.Open(dsn[11:]), gcfg)
	default:
		return nil, errors.New("unknown or unsupported sql driver")
	}
	if err != nil {
		return nil, err
	}

	if conf.GetBool("database.echo") {
		db = db.Debug()
	}
	return db, nil
}
