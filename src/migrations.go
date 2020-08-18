package src

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src/migrations"
	"gopkg.in/gormigrate.v1"
)

// GetMigrations gets migrations
func GetMigrations(db *gen.DB) []*gormigrate.Migration {
	// Admin user ID
	aid := uuid.Must(uuid.NewV4()).String()
	return []*gormigrate.Migration{
		&gormigrate.Migration{
			ID: "0000_INIT",
			Migrate: func(tx *gorm.DB) error {
				return db.AutoMigrate()
			},
			Rollback: func(tx *gorm.DB) error {
				// there's not much we can do if initialization/automigration failed
				return nil
			},
		},
		migrations.InitUsers(aid),
		migrations.InitPermissions(aid),
		migrations.InitRoles(aid),
	}
}
