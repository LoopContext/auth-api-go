package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
)

// UserResultType struct
type UserResultType struct {
	EntityResultType
}

// User struct
type User struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Active      bool       `json:"active" gorm:"column:active;default:false"`
	Email       string     `json:"email" gorm:"column:email;unique"`
	Password    *string    `json:"password" gorm:"column:password"`
	AvatarURL   *string    `json:"avatarURL" gorm:"column:avatarURL;type:text"`
	DisplayName *string    `json:"displayName" gorm:"column:displayName"`
	FirstName   *string    `json:"firstName" gorm:"column:firstName"`
	LastName    *string    `json:"lastName" gorm:"column:lastName"`
	NickName    *string    `json:"nickName" gorm:"column:nickName"`
	Location    *string    `json:"location" gorm:"column:location"`
	Description *string    `json:"description" gorm:"column:description;type:text"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	Apikeys          []*UserAPIKey `json:"apikeys" gorm:"foreignkey:UserID"`
	ApikeysPreloaded bool          `gorm:"-"`

	Roles          []*Role `json:"roles" gorm:"many2many:role_users;jointable_foreignkey:userId;association_jointable_foreignkey:roleId"`
	RolesPreloaded bool    `gorm:"-"`

	Profiles          []*Profile `json:"profiles" gorm:"many2many:profile_users;jointable_foreignkey:userId;association_jointable_foreignkey:profileId"`
	ProfilesPreloaded bool       `gorm:"-"`

	Permissions          []*Permission `json:"permissions" gorm:"many2many:permission_users;jointable_foreignkey:userId;association_jointable_foreignkey:permissionId"`
	PermissionsPreloaded bool          `gorm:"-"`
}

// IsEntity ...
func (m *User) IsEntity() {}

// UserChanges struct
type UserChanges struct {
	ID          string
	Active      bool
	Email       string
	Password    *string
	AvatarURL   *string
	DisplayName *string
	FirstName   *string
	LastName    *string
	NickName    *string
	Location    *string
	Description *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string

	ApikeysIDs     []*string
	RolesIDs       []*string
	ProfilesIDs    []*string
	PermissionsIDs []*string
}

// UserAPIKeyResultType struct
type UserAPIKeyResultType struct {
	EntityResultType
}

// UserAPIKey struct
type UserAPIKey struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Key         string     `json:"key" gorm:"column:key;unique"`
	Description *string    `json:"description" gorm:"column:description;type:text"`
	UserID      *string    `json:"userId" gorm:"column:userId"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	User          *User `json:"user"`
	UserPreloaded bool  `gorm:"-"`

	Permissions          []*Permission `json:"permissions" gorm:"many2many:userAPIKey_permissions;jointable_foreignkey:apikeyId;association_jointable_foreignkey:permissionId"`
	PermissionsPreloaded bool          `gorm:"-"`
}

// IsEntity ...
func (m *UserAPIKey) IsEntity() {}

// UserAPIKeyChanges struct
type UserAPIKeyChanges struct {
	ID          string
	Key         string
	Description *string
	UserID      *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string

	PermissionsIDs []*string
}

// UserAPIKeyPermissions struct
type UserAPIKeyPermissions struct {
	ApikeyID     string
	PermissionID string
}

// TableName ...
func (UserAPIKeyPermissions) TableName() string {
	return TableName("userAPIKey_permissions")
}

// ProfileResultType struct
type ProfileResultType struct {
	EntityResultType
}

// Profile struct
type Profile struct {
	ID             string     `json:"id" gorm:"column:id;primary_key"`
	Email          string     `json:"email" gorm:"column:email"`
	ExternalUserID *string    `json:"externalUserId" gorm:"column:externalUserId"`
	Provider       *string    `json:"provider" gorm:"column:provider"`
	AvatarURL      *string    `json:"avatarURL" gorm:"column:avatarURL;type:text"`
	Name           *string    `json:"name" gorm:"column:name"`
	FirstName      *string    `json:"firstName" gorm:"column:firstName"`
	LastName       *string    `json:"lastName" gorm:"column:lastName"`
	NickName       *string    `json:"nickName" gorm:"column:nickName"`
	Description    *string    `json:"description" gorm:"column:description;type:text"`
	Location       *string    `json:"location" gorm:"column:location"`
	UpdatedAt      *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt      time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy      *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy      *string    `json:"createdBy" gorm:"column:createdBy"`

	Users []*User `json:"users" gorm:"many2many:profile_users;jointable_foreignkey:profileId;association_jointable_foreignkey:userId"`
}

// IsEntity ...
func (m *Profile) IsEntity() {}

// ProfileChanges struct
type ProfileChanges struct {
	ID             string
	Email          string
	ExternalUserID *string
	Provider       *string
	AvatarURL      *string
	Name           *string
	FirstName      *string
	LastName       *string
	NickName       *string
	Description    *string
	Location       *string
	UpdatedAt      *time.Time
	CreatedAt      time.Time
	UpdatedBy      *string
	CreatedBy      *string

	UsersIDs []*string
}

// ProfileUsers struct
type ProfileUsers struct {
	ProfileID string
	UserID    string
}

// TableName ...
func (ProfileUsers) TableName() string {
	return TableName("profile_users")
}

// RoleResultType struct
type RoleResultType struct {
	EntityResultType
}

// Role struct
type Role struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Name        string     `json:"name" gorm:"column:name"`
	Description *string    `json:"description" gorm:"column:description;type:text"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	Users []*User `json:"users" gorm:"many2many:role_users;jointable_foreignkey:roleId;association_jointable_foreignkey:userId"`

	Parents []*Role `json:"parents" gorm:"many2many:role_parents;jointable_foreignkey:childId;association_jointable_foreignkey:parentId"`

	Children []*Role `json:"children" gorm:"many2many:role_parents;jointable_foreignkey:parentId;association_jointable_foreignkey:childId"`

	Permissions []*Permission `json:"permissions" gorm:"many2many:permission_roles;jointable_foreignkey:roleId;association_jointable_foreignkey:permissionId"`
}

// IsEntity ...
func (m *Role) IsEntity() {}

// RoleChanges struct
type RoleChanges struct {
	ID          string
	Name        string
	Description *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string

	UsersIDs       []*string
	ParentsIDs     []*string
	ChildrenIDs    []*string
	PermissionsIDs []*string
}

// RoleUsers struct
type RoleUsers struct {
	RoleID string
	UserID string
}

// TableName ...
func (RoleUsers) TableName() string {
	return TableName("role_users")
}

// RoleParents struct
type RoleParents struct {
	ChildID  string
	ParentID string
}

// TableName ...
func (RoleParents) TableName() string {
	return TableName("role_parents")
}

// PermissionResultType struct
type PermissionResultType struct {
	EntityResultType
}

// Permission struct
type Permission struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Tag         string     `json:"tag" gorm:"column:tag;unique"`
	Description string     `json:"description" gorm:"column:description;type:text"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	Users []*User `json:"users" gorm:"many2many:permission_users;jointable_foreignkey:permissionId;association_jointable_foreignkey:userId"`

	Roles []*Role `json:"roles" gorm:"many2many:permission_roles;jointable_foreignkey:permissionId;association_jointable_foreignkey:roleId"`

	Apikeys []*UserAPIKey `json:"apikeys" gorm:"many2many:userAPIKey_permissions;jointable_foreignkey:permissionId;association_jointable_foreignkey:apikeyId"`
}

// IsEntity ...
func (m *Permission) IsEntity() {}

// PermissionChanges struct
type PermissionChanges struct {
	ID          string
	Tag         string
	Description string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string

	UsersIDs   []*string
	RolesIDs   []*string
	ApikeysIDs []*string
}

// PermissionUsers struct
type PermissionUsers struct {
	PermissionID string
	UserID       string
}

// TableName ...
func (PermissionUsers) TableName() string {
	return TableName("permission_users")
}

// PermissionRoles struct
type PermissionRoles struct {
	PermissionID string
	RoleID       string
}

// TableName ...
func (PermissionRoles) TableName() string {
	return TableName("permission_roles")
}

// ApplyChanges used to convert map[string]interface{} to EntityChanges struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {

			if b == reflect.TypeOf(time.Time{}) {
				switch a.Kind() {
				case reflect.String:
					return time.Parse(time.RFC3339, v.(string))
				case reflect.Float64:
					return time.Unix(0, int64(v.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, v.(int64)*int64(time.Millisecond)), nil
				default:
					return v, fmt.Errorf("Unable to parse date from %v", v)
				}
			}

			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}
