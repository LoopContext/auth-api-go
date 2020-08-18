package migrations

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/loopcontext/auth-api-go/gen"
	"gopkg.in/gormigrate.v1"
)

// InitRoles initial roles of the api
func InitRoles(aid string) *gormigrate.Migration {
	if aid == "" {
		return nil
	}
	return &gormigrate.Migration{
		ID: "0003_INIT_ROLES",
		Migrate: func(tx *gorm.DB) error {
			var permissions []*gen.Permission
			au := &gen.User{ID: aid}
			// Assign all permisions to admin role
			if err := tx.Model(gen.Permission{}).Find(&permissions).Error; err != nil {
				return err
			}
			adminRole := &gen.Role{
				ID:        uuid.Must(uuid.NewV4()).String(),
				Name:      "admin",
				CreatedBy: &aid,
			}
			tx.Create(adminRole).Association("Permissions").Append(permissions)
			tx.Preload("Permissions").First(adminRole)
			// Give the role to our first user
			tx.Model(au).Association("Roles").Append(adminRole)
			tx.Model(au).Association("Permissions").Append(adminRole.Permissions)

			tx = tx.Create(&gen.Role{
				ID:        uuid.Must(uuid.NewV4()).String(),
				Name:      "user",
				CreatedBy: &aid,
			})
			if len(tx.GetErrors()) > 0 {
				errmsg := "[Role.Errors]: "
				for _, e := range tx.GetErrors() {
					errmsg += " - " + e.Error()
				}
				return errors.New(errmsg)
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Raw("TRUNCATE TABLE roles").Error
		},
	}
}
