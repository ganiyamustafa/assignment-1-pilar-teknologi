package connections

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SQLite *gorm.DB

func ConnectSqlite() error {
	db, err := gorm.Open(sqlite.Open("chat.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	SQLite = db

	return nil
}
