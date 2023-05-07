package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func init() {
	type chats struct {
		Id         uuid.UUID `gorm:"type:uuid;primary_key"`
		ChatRoomId string    `gorm:"type:uuid;foreign_key"`
		Message    string    `gorm:"type:text;"`
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}

	newMigration := gormigrate.Migration{
		ID: "0003",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(chats{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(chats{})
		},
	}
	migrations = append(migrations, &newMigration)
}
