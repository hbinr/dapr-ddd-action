package database

import (
	"database/sql"
	"log"

	"github.com/dapr-ddd-action/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	sqlDB *sql.DB
)

// Init 初始化 MySQL, 基于 gorm + mysql
func Init(conf *conf.Database) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.Source), gormConfig(conf.LogMode))
	if err != nil {
		log.Fatal(err)
	}

	if sqlDB, err = db.DB(); err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(int(conf.MaxIdleConns))
	sqlDB.SetMaxOpenConns(int(conf.MaxOpenConns))

	return db
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(mod bool) (c *gorm.Config) {
	if mod {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	}
	return
}
