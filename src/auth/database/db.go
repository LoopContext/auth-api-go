package database

import (
	"github.com/markbates/goth"

	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src/auth/database/transformations"
)

// FindUserByJWT finds the user by JWT information
func FindUserByJWT(db *gen.DB, id string, email string, provider string) (*gen.User, error) {
	tx := db.Query()
	u := &gen.User{}
	if tx := tx.Preload("Roles").Preload("Roles.Permissions").Preload("Permissions").
		Joins(`JOIN `+gen.TableName("profiles")+` p ON p."externalUserId" = ? AND p.provider = ?`, id, provider).
		Joins(`JOIN `+gen.TableName("profile_users")+` pu ON pu."userId" = `+gen.TableName("users")+`.id AND pu."profileId" = p.id`).
		Where(gen.TableName("users")+".email = ?", email).
		First(u); tx.RecordNotFound() || tx.Error != nil {
		return nil, tx.Error
	}
	return u, nil
}

// FindUserByAPIKey finds the user by JWT information
func FindUserByAPIKey(db *gen.DB, apiKey string) (*gen.User, error) {
	tx := db.Query()
	u := &gen.User{}
	if tx := tx.Preload("Apikeys").Preload("Roles").
		Preload("Roles.Permissions").Preload("Permissions").
		Joins(`JOIN `+gen.TableName("user_api_keys")+` uak ON "userId" = u.id`).
		Where(`uak.key = ?`, apiKey).
		First(u); tx.RecordNotFound() || tx.Error != nil {
		return nil, tx.Error
	}
	return u, nil
}

// UpsertUserProfile saves the user if doesn't exists and adds the OAuth profile
func UpsertUserProfile(db *gen.DB, input *goth.User) (u *gen.User, err error) {
	tx := db.Query()
	u = &gen.User{}
	if fu := tx.Where("email = ?", input.Email).First(u); fu.RecordNotFound() {
		u, err = transformations.GothUserToDBUser(input, false)
		if err != nil {
			return nil, err
		}
		if cu := tx.Create(u); cu.Error != nil {
			return nil, err
		}
	} else if tx.Error != nil {
		return nil, err
	}

	for _, p := range u.Profiles {
		if p.ExternalUserID == &input.UserID &&
			p.Email == input.Email &&
			p.Provider == &input.Provider {
			return u, nil
		}
	}

	p, err := transformations.GothUserToDBUserProfile(input, false)
	if tx := tx.Where(`"externalUserId" = ? AND email = ? AND provider = ?`,
		input.UserID, input.Email, input.Provider).FirstOrCreate(p); tx.Error != nil {
		return nil, err
	}

	tx.Model(u).Association("Profiles").Append(p)

	if tx := tx.Save(u).First(u); tx.Error != nil {
		return nil, err
	}

	return u, nil
}
