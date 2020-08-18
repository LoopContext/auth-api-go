package gen

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/loopcontext/graphql-orm/events"
)

// ResolutionHandlers struct
type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateUser     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *User, err error)
	UpdateUser     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *User, err error)
	DeleteUser     func(ctx context.Context, r *GeneratedResolver, id string) (item *User, err error)
	DeleteAllUsers func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryUser      func(ctx context.Context, r *GeneratedResolver, opts QueryUserHandlerOptions) (*User, error)
	QueryUsers     func(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions) (*UserResultType, error)

	UserApikeys func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*UserAPIKey, err error)

	UserRoles func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Role, err error)

	UserProfiles func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Profile, err error)

	UserPermissions func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Permission, err error)

	CreateUserAPIKey     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *UserAPIKey, err error)
	UpdateUserAPIKey     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *UserAPIKey, err error)
	DeleteUserAPIKey     func(ctx context.Context, r *GeneratedResolver, id string) (item *UserAPIKey, err error)
	DeleteAllUserAPIKeys func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryUserAPIKey      func(ctx context.Context, r *GeneratedResolver, opts QueryUserAPIKeyHandlerOptions) (*UserAPIKey, error)
	QueryUserAPIKeys     func(ctx context.Context, r *GeneratedResolver, opts QueryUserAPIKeysHandlerOptions) (*UserAPIKeyResultType, error)

	UserAPIKeyUser func(ctx context.Context, r *GeneratedResolver, obj *UserAPIKey) (res *User, err error)

	UserAPIKeyPermissions func(ctx context.Context, r *GeneratedResolver, obj *UserAPIKey) (res []*Permission, err error)

	CreateProfile     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Profile, err error)
	UpdateProfile     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Profile, err error)
	DeleteProfile     func(ctx context.Context, r *GeneratedResolver, id string) (item *Profile, err error)
	DeleteAllProfiles func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryProfile      func(ctx context.Context, r *GeneratedResolver, opts QueryProfileHandlerOptions) (*Profile, error)
	QueryProfiles     func(ctx context.Context, r *GeneratedResolver, opts QueryProfilesHandlerOptions) (*ProfileResultType, error)

	ProfileUsers func(ctx context.Context, r *GeneratedResolver, obj *Profile) (res []*User, err error)

	CreateRole     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Role, err error)
	UpdateRole     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Role, err error)
	DeleteRole     func(ctx context.Context, r *GeneratedResolver, id string) (item *Role, err error)
	DeleteAllRoles func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryRole      func(ctx context.Context, r *GeneratedResolver, opts QueryRoleHandlerOptions) (*Role, error)
	QueryRoles     func(ctx context.Context, r *GeneratedResolver, opts QueryRolesHandlerOptions) (*RoleResultType, error)

	RoleUsers func(ctx context.Context, r *GeneratedResolver, obj *Role) (res []*User, err error)

	RoleParents func(ctx context.Context, r *GeneratedResolver, obj *Role) (res []*Role, err error)

	RoleChildren func(ctx context.Context, r *GeneratedResolver, obj *Role) (res []*Role, err error)

	RolePermissions func(ctx context.Context, r *GeneratedResolver, obj *Role) (res []*Permission, err error)

	CreatePermission     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Permission, err error)
	UpdatePermission     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Permission, err error)
	DeletePermission     func(ctx context.Context, r *GeneratedResolver, id string) (item *Permission, err error)
	DeleteAllPermissions func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPermission      func(ctx context.Context, r *GeneratedResolver, opts QueryPermissionHandlerOptions) (*Permission, error)
	QueryPermissions     func(ctx context.Context, r *GeneratedResolver, opts QueryPermissionsHandlerOptions) (*PermissionResultType, error)

	PermissionUsers func(ctx context.Context, r *GeneratedResolver, obj *Permission) (res []*User, err error)

	PermissionRoles func(ctx context.Context, r *GeneratedResolver, obj *Permission) (res []*Role, err error)

	PermissionApikeys func(ctx context.Context, r *GeneratedResolver, obj *Permission) (res []*UserAPIKey, err error)
}

// DefaultResolutionHandlers ...
func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

		CreateUser:     CreateUserHandler,
		UpdateUser:     UpdateUserHandler,
		DeleteUser:     DeleteUserHandler,
		DeleteAllUsers: DeleteAllUsersHandler,
		QueryUser:      QueryUserHandler,
		QueryUsers:     QueryUsersHandler,

		UserApikeys: UserApikeysHandler,

		UserRoles: UserRolesHandler,

		UserProfiles: UserProfilesHandler,

		UserPermissions: UserPermissionsHandler,

		CreateUserAPIKey:     CreateUserAPIKeyHandler,
		UpdateUserAPIKey:     UpdateUserAPIKeyHandler,
		DeleteUserAPIKey:     DeleteUserAPIKeyHandler,
		DeleteAllUserAPIKeys: DeleteAllUserAPIKeysHandler,
		QueryUserAPIKey:      QueryUserAPIKeyHandler,
		QueryUserAPIKeys:     QueryUserAPIKeysHandler,

		UserAPIKeyUser: UserAPIKeyUserHandler,

		UserAPIKeyPermissions: UserAPIKeyPermissionsHandler,

		CreateProfile:     CreateProfileHandler,
		UpdateProfile:     UpdateProfileHandler,
		DeleteProfile:     DeleteProfileHandler,
		DeleteAllProfiles: DeleteAllProfilesHandler,
		QueryProfile:      QueryProfileHandler,
		QueryProfiles:     QueryProfilesHandler,

		ProfileUsers: ProfileUsersHandler,

		CreateRole:     CreateRoleHandler,
		UpdateRole:     UpdateRoleHandler,
		DeleteRole:     DeleteRoleHandler,
		DeleteAllRoles: DeleteAllRolesHandler,
		QueryRole:      QueryRoleHandler,
		QueryRoles:     QueryRolesHandler,

		RoleUsers: RoleUsersHandler,

		RoleParents: RoleParentsHandler,

		RoleChildren: RoleChildrenHandler,

		RolePermissions: RolePermissionsHandler,

		CreatePermission:     CreatePermissionHandler,
		UpdatePermission:     UpdatePermissionHandler,
		DeletePermission:     DeletePermissionHandler,
		DeleteAllPermissions: DeleteAllPermissionsHandler,
		QueryPermission:      QueryPermissionHandler,
		QueryPermissions:     QueryPermissionsHandler,

		PermissionUsers: PermissionUsersHandler,

		PermissionRoles: PermissionRolesHandler,

		PermissionApikeys: PermissionApikeysHandler,
	}
	return handlers
}

// GeneratedResolver struct
type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *events.EventController
}

// GetDB returns database connection or transaction for given context (if exists)
func (r *GeneratedResolver) GetDB(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(KeyMutationTransaction).(*gorm.DB)
	if db == nil {
		db = r.DB.Query()
	}
	return db
}
