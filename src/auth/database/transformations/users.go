package transformations

import (
	"errors"
	"github.com/loopcontext/auth-api-go/gen"

	"github.com/markbates/goth"

	"github.com/gofrs/uuid"
)

// GothUserToDBUser transforms [user] goth to db model
func GothUserToDBUser(i *goth.User, update bool, ids ...string) (o *gen.User, err error) {
	if i.Email == "" && !update {
		return nil, errors.New("field [Email] is required")
	}
	o = &gen.User{
		Email:       i.Email,
		DisplayName: &i.Name,
		FirstName:   &i.FirstName,
		LastName:    &i.LastName,
		NickName:    &i.NickName,
		Location:    &i.Location,
		AvatarURL:   &i.AvatarURL,
		Description: &i.Description,
	}
	if !update {
		uid, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		o.ID = uid.String()
	}
	if len(ids) > 0 {
		updID, err := uuid.FromString(ids[0])
		if err != nil {
			return nil, err
		}
		o.ID = updID.String()
	}
	return o, err
}

// GothUserToDBUserProfile transforms [user] goth to db model
func GothUserToDBUserProfile(i *goth.User, update bool, ids ...string) (o *gen.Profile, err error) {
	if i.UserID == "" && !update {
		return nil, errors.New("field [UserID] is required")
	}
	if i.Email == "" && !update {
		return nil, errors.New("field [Email] is required")
	}
	o = &gen.Profile{
		ExternalUserID: &i.UserID,
		Provider:       &i.Provider,
		Email:          i.Email,
		Name:           &i.Name,
		FirstName:      &i.FirstName,
		LastName:       &i.LastName,
		NickName:       &i.NickName,
		Location:       &i.Location,
		AvatarURL:      &i.AvatarURL,
		Description:    &i.Description,
	}
	if !update {
		uid, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		o.ID = uid.String()
	}
	if len(ids) > 0 {
		updID, err := uuid.FromString(ids[0])
		if err != nil {
			return nil, err
		}
		o.ID = updID.String()
	}
	return o, err
}
