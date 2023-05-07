package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func init() {
	type chat_rooms struct {
		Id         uuid.UUID `gorm:"type:uuid;primary_key"`
		SenderId   string    `gorm:"type:uuid;foreign_key"`
		ReceiverId string    `gorm:"type:uuid;foreign_key"`
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}

	newMigration := gormigrate.Migration{
		ID: "0002",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(chat_rooms{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(chat_rooms{})
		},
	}
	migrations = append(migrations, &newMigration)
}
