package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func init() {
	type users struct {
		Id        uuid.UUID `gorm:"type:uuid;primary_key"`
		Name      string    `gorm:"type:varchar(20);not null"`
		Email     string    `gorm:"type:varchar(100);not null;unique"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	newMigration := gormigrate.Migration{
		ID: "0001",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(users{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(users{})
		},
	}
	migrations = append(migrations, &newMigration)
}
