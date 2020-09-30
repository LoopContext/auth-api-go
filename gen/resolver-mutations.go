package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/loopcontext/go-graphql-orm/events"
)

// GeneratedMutationResolver ...
type GeneratedMutationResolver struct{ *GeneratedResolver }

// MutationEvents ...
type MutationEvents struct {
	Events []events.Event
}

// EnrichContextWithMutations ...
func EnrichContextWithMutations(ctx context.Context, r *GeneratedResolver) context.Context {
	_ctx := context.WithValue(ctx, KeyMutationTransaction, r.DB.db.Begin())
	_ctx = context.WithValue(_ctx, KeyMutationEvents, &MutationEvents{})
	return _ctx
}

// FinishMutationContext ...
func FinishMutationContext(ctx context.Context, r *GeneratedResolver) (err error) {
	s := GetMutationEventStore(ctx)

	for _, event := range s.Events {
		err = r.Handlers.OnEvent(ctx, r, &event)
		if err != nil {
			return
		}
	}

	tx := r.GetDB(ctx)
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, event := range s.Events {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}

// RollbackMutationContext ...
func RollbackMutationContext(ctx context.Context, r *GeneratedResolver) error {
	tx := r.GetDB(ctx)
	return tx.Rollback().Error
}

// GetMutationEventStore ...
func GetMutationEventStore(ctx context.Context) *MutationEvents {
	return ctx.Value(KeyMutationEvents).(*MutationEvents)
}

// AddMutationEvent ...
func AddMutationEvent(ctx context.Context, e events.Event) {
	s := GetMutationEventStore(ctx)
	s.Events = append(s.Events, e)
}

// CreateUser ...
func (r *GeneratedMutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *User, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateUser(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateUserHandler ...
func CreateUserHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *User, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &User{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "User",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes UserChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["active"]; ok && (item.Active != changes.Active) {
		item.Active = changes.Active

		event.AddNewValue("active", changes.Active)
	}

	if _, ok := input["email"]; ok && (item.Email != changes.Email) {
		item.Email = changes.Email

		event.AddNewValue("email", changes.Email)
	}

	if _, ok := input["password"]; ok && (item.Password != changes.Password) && (item.Password == nil || changes.Password == nil || *item.Password != *changes.Password) {
		item.Password = changes.Password

		event.AddNewValue("password", changes.Password)
	}

	if _, ok := input["avatarURL"]; ok && (item.AvatarURL != changes.AvatarURL) && (item.AvatarURL == nil || changes.AvatarURL == nil || *item.AvatarURL != *changes.AvatarURL) {
		item.AvatarURL = changes.AvatarURL

		event.AddNewValue("avatarURL", changes.AvatarURL)
	}

	if _, ok := input["displayName"]; ok && (item.DisplayName != changes.DisplayName) && (item.DisplayName == nil || changes.DisplayName == nil || *item.DisplayName != *changes.DisplayName) {
		item.DisplayName = changes.DisplayName

		event.AddNewValue("displayName", changes.DisplayName)
	}

	if _, ok := input["firstName"]; ok && (item.FirstName != changes.FirstName) && (item.FirstName == nil || changes.FirstName == nil || *item.FirstName != *changes.FirstName) {
		item.FirstName = changes.FirstName

		event.AddNewValue("firstName", changes.FirstName)
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		item.LastName = changes.LastName

		event.AddNewValue("lastName", changes.LastName)
	}

	if _, ok := input["nickName"]; ok && (item.NickName != changes.NickName) && (item.NickName == nil || changes.NickName == nil || *item.NickName != *changes.NickName) {
		item.NickName = changes.NickName

		event.AddNewValue("nickName", changes.NickName)
	}

	if _, ok := input["location"]; ok && (item.Location != changes.Location) && (item.Location == nil || changes.Location == nil || *item.Location != *changes.Location) {
		item.Location = changes.Location

		event.AddNewValue("location", changes.Location)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["apikeysIds"]; exists {
		items := []UserAPIKey{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Apikeys")
		association.Replace(items)
	}

	if ids, exists := input["rolesIds"]; exists {
		items := []Role{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Roles")
		association.Replace(items)
	}

	if ids, exists := input["profilesIds"]; exists {
		items := []Profile{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Profiles")
		association.Replace(items)
	}

	if ids, exists := input["permissionsIds"]; exists {
		items := []Permission{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Permissions")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateUser ...
func (r *GeneratedMutationResolver) UpdateUser(ctx context.Context, id string, input map[string]interface{}) (item *User, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateUser(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateUserHandler ...
func UpdateUserHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *User, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &User{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "User",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes UserChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["active"]; ok && (item.Active != changes.Active) {
		event.AddOldValue("active", item.Active)
		event.AddNewValue("active", changes.Active)
		item.Active = changes.Active
	}

	if _, ok := input["email"]; ok && (item.Email != changes.Email) {
		event.AddOldValue("email", item.Email)
		event.AddNewValue("email", changes.Email)
		item.Email = changes.Email
	}

	if _, ok := input["password"]; ok && (item.Password != changes.Password) && (item.Password == nil || changes.Password == nil || *item.Password != *changes.Password) {
		event.AddOldValue("password", item.Password)
		event.AddNewValue("password", changes.Password)
		item.Password = changes.Password
	}

	if _, ok := input["avatarURL"]; ok && (item.AvatarURL != changes.AvatarURL) && (item.AvatarURL == nil || changes.AvatarURL == nil || *item.AvatarURL != *changes.AvatarURL) {
		event.AddOldValue("avatarURL", item.AvatarURL)
		event.AddNewValue("avatarURL", changes.AvatarURL)
		item.AvatarURL = changes.AvatarURL
	}

	if _, ok := input["displayName"]; ok && (item.DisplayName != changes.DisplayName) && (item.DisplayName == nil || changes.DisplayName == nil || *item.DisplayName != *changes.DisplayName) {
		event.AddOldValue("displayName", item.DisplayName)
		event.AddNewValue("displayName", changes.DisplayName)
		item.DisplayName = changes.DisplayName
	}

	if _, ok := input["firstName"]; ok && (item.FirstName != changes.FirstName) && (item.FirstName == nil || changes.FirstName == nil || *item.FirstName != *changes.FirstName) {
		event.AddOldValue("firstName", item.FirstName)
		event.AddNewValue("firstName", changes.FirstName)
		item.FirstName = changes.FirstName
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		event.AddOldValue("lastName", item.LastName)
		event.AddNewValue("lastName", changes.LastName)
		item.LastName = changes.LastName
	}

	if _, ok := input["nickName"]; ok && (item.NickName != changes.NickName) && (item.NickName == nil || changes.NickName == nil || *item.NickName != *changes.NickName) {
		event.AddOldValue("nickName", item.NickName)
		event.AddNewValue("nickName", changes.NickName)
		item.NickName = changes.NickName
	}

	if _, ok := input["location"]; ok && (item.Location != changes.Location) && (item.Location == nil || changes.Location == nil || *item.Location != *changes.Location) {
		event.AddOldValue("location", item.Location)
		event.AddNewValue("location", changes.Location)
		item.Location = changes.Location
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["apikeysIds"]; exists {
		items := []UserAPIKey{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Apikeys")
		association.Replace(items)
	}

	if ids, exists := input["rolesIds"]; exists {
		items := []Role{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Roles")
		association.Replace(items)
	}

	if ids, exists := input["profilesIds"]; exists {
		items := []Profile{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Profiles")
		association.Replace(items)
	}

	if ids, exists := input["permissionsIds"]; exists {
		items := []Permission{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Permissions")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteUser ...
func (r *GeneratedMutationResolver) DeleteUser(ctx context.Context, id string) (item *User, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteUser(ctx, r.GeneratedResolver, id)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteUserHandler handler
func DeleteUserHandler(ctx context.Context, r *GeneratedResolver, id string) (item *User, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &User{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "User",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("users")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllUsers ...
func (r *GeneratedMutationResolver) DeleteAllUsers(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllUsers(ctx, r.GeneratedResolver)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllUsersHandler handler
func DeleteAllUsersHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&User{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateUserAPIKey ...
func (r *GeneratedMutationResolver) CreateUserAPIKey(ctx context.Context, input map[string]interface{}) (item *UserAPIKey, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateUserAPIKey(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateUserAPIKeyHandler ...
func CreateUserAPIKeyHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *UserAPIKey, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &UserAPIKey{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "UserAPIKey",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes UserAPIKeyChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["key"]; ok && (item.Key != changes.Key) {
		item.Key = changes.Key

		event.AddNewValue("key", changes.Key)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	if _, ok := input["userId"]; ok && (item.UserID != changes.UserID) && (item.UserID == nil || changes.UserID == nil || *item.UserID != *changes.UserID) {
		item.UserID = changes.UserID

		event.AddNewValue("userId", changes.UserID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["permissionsIds"]; exists {
		items := []Permission{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Permissions")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateUserAPIKey ...
func (r *GeneratedMutationResolver) UpdateUserAPIKey(ctx context.Context, id string, input map[string]interface{}) (item *UserAPIKey, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateUserAPIKey(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateUserAPIKeyHandler ...
func UpdateUserAPIKeyHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *UserAPIKey, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &UserAPIKey{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "UserAPIKey",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes UserAPIKeyChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["key"]; ok && (item.Key != changes.Key) {
		event.AddOldValue("key", item.Key)
		event.AddNewValue("key", changes.Key)
		item.Key = changes.Key
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	if _, ok := input["userId"]; ok && (item.UserID != changes.UserID) && (item.UserID == nil || changes.UserID == nil || *item.UserID != *changes.UserID) {
		event.AddOldValue("userId", item.UserID)
		event.AddNewValue("userId", changes.UserID)
		item.UserID = changes.UserID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["permissionsIds"]; exists {
		items := []Permission{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Permissions")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteUserAPIKey ...
func (r *GeneratedMutationResolver) DeleteUserAPIKey(ctx context.Context, id string) (item *UserAPIKey, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteUserAPIKey(ctx, r.GeneratedResolver, id)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteUserAPIKeyHandler handler
func DeleteUserAPIKeyHandler(ctx context.Context, r *GeneratedResolver, id string) (item *UserAPIKey, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &UserAPIKey{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "UserAPIKey",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("user_api_keys")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllUserAPIKeys ...
func (r *GeneratedMutationResolver) DeleteAllUserAPIKeys(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllUserAPIKeys(ctx, r.GeneratedResolver)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllUserAPIKeysHandler handler
func DeleteAllUserAPIKeysHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&UserAPIKey{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateProfile ...
func (r *GeneratedMutationResolver) CreateProfile(ctx context.Context, input map[string]interface{}) (item *Profile, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateProfile(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateProfileHandler ...
func CreateProfileHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Profile, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Profile{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Profile",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ProfileChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["email"]; ok && (item.Email != changes.Email) {
		item.Email = changes.Email

		event.AddNewValue("email", changes.Email)
	}

	if _, ok := input["externalUserId"]; ok && (item.ExternalUserID != changes.ExternalUserID) && (item.ExternalUserID == nil || changes.ExternalUserID == nil || *item.ExternalUserID != *changes.ExternalUserID) {
		item.ExternalUserID = changes.ExternalUserID

		event.AddNewValue("externalUserId", changes.ExternalUserID)
	}

	if _, ok := input["provider"]; ok && (item.Provider != changes.Provider) && (item.Provider == nil || changes.Provider == nil || *item.Provider != *changes.Provider) {
		item.Provider = changes.Provider

		event.AddNewValue("provider", changes.Provider)
	}

	if _, ok := input["avatarURL"]; ok && (item.AvatarURL != changes.AvatarURL) && (item.AvatarURL == nil || changes.AvatarURL == nil || *item.AvatarURL != *changes.AvatarURL) {
		item.AvatarURL = changes.AvatarURL

		event.AddNewValue("avatarURL", changes.AvatarURL)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["firstName"]; ok && (item.FirstName != changes.FirstName) && (item.FirstName == nil || changes.FirstName == nil || *item.FirstName != *changes.FirstName) {
		item.FirstName = changes.FirstName

		event.AddNewValue("firstName", changes.FirstName)
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		item.LastName = changes.LastName

		event.AddNewValue("lastName", changes.LastName)
	}

	if _, ok := input["nickName"]; ok && (item.NickName != changes.NickName) && (item.NickName == nil || changes.NickName == nil || *item.NickName != *changes.NickName) {
		item.NickName = changes.NickName

		event.AddNewValue("nickName", changes.NickName)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	if _, ok := input["location"]; ok && (item.Location != changes.Location) && (item.Location == nil || changes.Location == nil || *item.Location != *changes.Location) {
		item.Location = changes.Location

		event.AddNewValue("location", changes.Location)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["usersIds"]; exists {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Users")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateProfile ...
func (r *GeneratedMutationResolver) UpdateProfile(ctx context.Context, id string, input map[string]interface{}) (item *Profile, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateProfile(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateProfileHandler ...
func UpdateProfileHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Profile, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Profile{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Profile",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ProfileChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["email"]; ok && (item.Email != changes.Email) {
		event.AddOldValue("email", item.Email)
		event.AddNewValue("email", changes.Email)
		item.Email = changes.Email
	}

	if _, ok := input["externalUserId"]; ok && (item.ExternalUserID != changes.ExternalUserID) && (item.ExternalUserID == nil || changes.ExternalUserID == nil || *item.ExternalUserID != *changes.ExternalUserID) {
		event.AddOldValue("externalUserId", item.ExternalUserID)
		event.AddNewValue("externalUserId", changes.ExternalUserID)
		item.ExternalUserID = changes.ExternalUserID
	}

	if _, ok := input["provider"]; ok && (item.Provider != changes.Provider) && (item.Provider == nil || changes.Provider == nil || *item.Provider != *changes.Provider) {
		event.AddOldValue("provider", item.Provider)
		event.AddNewValue("provider", changes.Provider)
		item.Provider = changes.Provider
	}

	if _, ok := input["avatarURL"]; ok && (item.AvatarURL != changes.AvatarURL) && (item.AvatarURL == nil || changes.AvatarURL == nil || *item.AvatarURL != *changes.AvatarURL) {
		event.AddOldValue("avatarURL", item.AvatarURL)
		event.AddNewValue("avatarURL", changes.AvatarURL)
		item.AvatarURL = changes.AvatarURL
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["firstName"]; ok && (item.FirstName != changes.FirstName) && (item.FirstName == nil || changes.FirstName == nil || *item.FirstName != *changes.FirstName) {
		event.AddOldValue("firstName", item.FirstName)
		event.AddNewValue("firstName", changes.FirstName)
		item.FirstName = changes.FirstName
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		event.AddOldValue("lastName", item.LastName)
		event.AddNewValue("lastName", changes.LastName)
		item.LastName = changes.LastName
	}

	if _, ok := input["nickName"]; ok && (item.NickName != changes.NickName) && (item.NickName == nil || changes.NickName == nil || *item.NickName != *changes.NickName) {
		event.AddOldValue("nickName", item.NickName)
		event.AddNewValue("nickName", changes.NickName)
		item.NickName = changes.NickName
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	if _, ok := input["location"]; ok && (item.Location != changes.Location) && (item.Location == nil || changes.Location == nil || *item.Location != *changes.Location) {
		event.AddOldValue("location", item.Location)
		event.AddNewValue("location", changes.Location)
		item.Location = changes.Location
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["usersIds"]; exists {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Users")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteProfile ...
func (r *GeneratedMutationResolver) DeleteProfile(ctx context.Context, id string) (item *Profile, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteProfile(ctx, r.GeneratedResolver, id)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteProfileHandler handler
func DeleteProfileHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Profile, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Profile{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Profile",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("profiles")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllProfiles ...
func (r *GeneratedMutationResolver) DeleteAllProfiles(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllProfiles(ctx, r.GeneratedResolver)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllProfilesHandler handler
func DeleteAllProfilesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&Profile{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateRole ...
func (r *GeneratedMutationResolver) CreateRole(ctx context.Context, input map[string]interface{}) (item *Role, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateRole(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateRoleHandler ...
func CreateRoleHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Role, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Role{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Role",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes RoleChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["domain"]; ok && (item.Domain != changes.Domain) {
		item.Domain = changes.Domain

		event.AddNewValue("domain", changes.Domain)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["usersIds"]; exists {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Users")
		association.Replace(items)
	}

	if ids, exists := input["parentsIds"]; exists {
		items := []Role{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Parents")
		association.Replace(items)
	}

	if ids, exists := input["childrenIds"]; exists {
		items := []Role{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Children")
		association.Replace(items)
	}

	if ids, exists := input["permissionsIds"]; exists {
		items := []Permission{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Permissions")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateRole ...
func (r *GeneratedMutationResolver) UpdateRole(ctx context.Context, id string, input map[string]interface{}) (item *Role, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateRole(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateRoleHandler ...
func UpdateRoleHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Role, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Role{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Role",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes RoleChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["domain"]; ok && (item.Domain != changes.Domain) {
		event.AddOldValue("domain", item.Domain)
		event.AddNewValue("domain", changes.Domain)
		item.Domain = changes.Domain
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["usersIds"]; exists {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Users")
		association.Replace(items)
	}

	if ids, exists := input["parentsIds"]; exists {
		items := []Role{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Parents")
		association.Replace(items)
	}

	if ids, exists := input["childrenIds"]; exists {
		items := []Role{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Children")
		association.Replace(items)
	}

	if ids, exists := input["permissionsIds"]; exists {
		items := []Permission{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Permissions")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteRole ...
func (r *GeneratedMutationResolver) DeleteRole(ctx context.Context, id string) (item *Role, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteRole(ctx, r.GeneratedResolver, id)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteRoleHandler handler
func DeleteRoleHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Role, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Role{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Role",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("roles")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllRoles ...
func (r *GeneratedMutationResolver) DeleteAllRoles(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllRoles(ctx, r.GeneratedResolver)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllRolesHandler handler
func DeleteAllRolesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&Role{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreatePermission ...
func (r *GeneratedMutationResolver) CreatePermission(ctx context.Context, input map[string]interface{}) (item *Permission, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreatePermission(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreatePermissionHandler ...
func CreatePermissionHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Permission, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Permission{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Permission",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PermissionChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["domain"]; ok && (item.Domain != changes.Domain) {
		item.Domain = changes.Domain

		event.AddNewValue("domain", changes.Domain)
	}

	if _, ok := input["tag"]; ok && (item.Tag != changes.Tag) {
		item.Tag = changes.Tag

		event.AddNewValue("tag", changes.Tag)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["usersIds"]; exists {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Users")
		association.Replace(items)
	}

	if ids, exists := input["rolesIds"]; exists {
		items := []Role{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Roles")
		association.Replace(items)
	}

	if ids, exists := input["apikeysIds"]; exists {
		items := []UserAPIKey{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Apikeys")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdatePermission ...
func (r *GeneratedMutationResolver) UpdatePermission(ctx context.Context, id string, input map[string]interface{}) (item *Permission, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdatePermission(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdatePermissionHandler ...
func UpdatePermissionHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Permission, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Permission{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Permission",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PermissionChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["domain"]; ok && (item.Domain != changes.Domain) {
		event.AddOldValue("domain", item.Domain)
		event.AddNewValue("domain", changes.Domain)
		item.Domain = changes.Domain
	}

	if _, ok := input["tag"]; ok && (item.Tag != changes.Tag) {
		event.AddOldValue("tag", item.Tag)
		event.AddNewValue("tag", changes.Tag)
		item.Tag = changes.Tag
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["usersIds"]; exists {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Users")
		association.Replace(items)
	}

	if ids, exists := input["rolesIds"]; exists {
		items := []Role{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Roles")
		association.Replace(items)
	}

	if ids, exists := input["apikeysIds"]; exists {
		items := []UserAPIKey{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Apikeys")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeletePermission ...
func (r *GeneratedMutationResolver) DeletePermission(ctx context.Context, id string) (item *Permission, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeletePermission(ctx, r.GeneratedResolver, id)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeletePermissionHandler handler
func DeletePermissionHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Permission, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Permission{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Permission",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("permissions")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllPermissions ...
func (r *GeneratedMutationResolver) DeleteAllPermissions(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllPermissions(ctx, r.GeneratedResolver)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllPermissionsHandler handler
func DeleteAllPermissionsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&Permission{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}
