package dao

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func Database(confRead string, confWrite string) {
	var ormLogger logger.Interface

	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      confRead,
		DefaultStringSize:        256,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex:   true, //重命名索引
		DontSupportRenameColumn:  true,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(confWrite)},
		Replicas: []gorm.Dialector{mysql.Open(confRead)},
		Policy:   dbresolver.RandomPolicy{},
	}))

	migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
