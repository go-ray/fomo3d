package database

import (
	"github.com/go-ray/fomo3d/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbs = map[string]*gorm.DB{}

func InitDatabaseConfig(name string, c conf.DBConf) error {
	db, err := gorm.Open(c.DriverName, c.DataSource)
	if err != nil {
		return err
	}
	db.DB().SetMaxIdleConns(c.MaxIdleConns)
	db.DB().SetMaxOpenConns(c.MaxOpenConns)
	dbs[name] = db
	return nil
}

func DefaultDB() *gorm.DB {
	return dbs["default"]
}
