package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/loopcontext/auth-api-go/gen"
	"gopkg.in/gormigrate.v1"
	"os"
	"strings"
)

// InitUsers initial users of the api
func InitUsers(aid string) *gormigrate.Migration {
	if aid == "" {
		return nil
	}
	fname := os.Getenv("AUTH_API_USER_FIRSTNAME") // "Christian"
	lname := os.Getenv("AUTH_API_USER_LASTNAME")  // "Melgarejo"
	dname := fname + " " + lname
	nname := strings.ToLower(string(fname[0]) + lname)
	au := &gen.User{
		ID:          aid,
		Email:       os.Getenv("AUTH_API_USER_EMAIL"), // "cmelgarejo.dev@gmail.com"
		FirstName:   &fname,
		LastName:    &lname,
		DisplayName: &dname,
		NickName:    &nname,
		CreatedBy:   &aid, // self-reference
	}
	return &gormigrate.Migration{
		ID: "0001_INIT_USERS",
		Migrate: func(tx *gorm.DB) error {
			return tx.Create(au).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Delete(au).Error
		},
	}

}
