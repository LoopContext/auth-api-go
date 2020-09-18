package migrations

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/loopcontext/auth-api-go/gen"
	"gopkg.in/gormigrate.v1"
)

// InitPermissions initial permissions of the api
func InitPermissions(aid string) *gormigrate.Migration {
	if aid == "" {
		return nil
	}
	ps := []struct {
		Entity      string
		Permissions []string
	}{
		{
			Entity:      "users",
			Permissions: []string{"create", "read", "update", "delete", "list"},
		},
		{
			Entity:      "roles",
			Permissions: []string{"create", "read", "update", "delete", "list"},
		},
		{
			Entity:      "permissions",
			Permissions: []string{"create", "read", "update", "delete", "list"},
		},
	}
	var permissions []*gen.Permission
	for _, p := range ps {
		for _, pd := range p.Permissions {
			permissions = append(permissions, &gen.Permission{
				ID:          uuid.Must(uuid.NewV4()).String(),
				Domain:      "app",
				Description: "Permission to [" + pd + "] for entity [" + p.Entity + "]",
				Tag:         p.Entity + ":" + pd,
				CreatedBy:   &aid,
			})
		}
	}
	return &gormigrate.Migration{
		ID: "0002_INIT_PERMISSIONS",
		Migrate: func(tx *gorm.DB) error {
			for _, p := range permissions {
				tx = tx.Create(p)
			}
			if len(tx.GetErrors()) > 0 {
				errmsg := "[Permission.Errors]: "
				for _, e := range tx.GetErrors() {
					errmsg += " - " + e.Error()
				}
				return errors.New(errmsg)
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			for _, p := range permissions {
				tx = tx.Delete(p)
			}
			if len(tx.GetErrors()) > 0 {
				errmsg := "[Permission.Errors]: "
				for _, e := range tx.GetErrors() {
					errmsg += " - " + e.Error()
				}
				return errors.New(errmsg)
			}
			return nil
		},
	}
}
