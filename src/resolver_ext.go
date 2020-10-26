package src

import (
	"context"
	"fmt"

	"github.com/loopcontext/auth-api-go/gen"
)

const (
	jwtTokenPermissionErrMsg = "You don't have permission to %s the entity %s"
)

// Users method
func (r *QueryResolver) Users(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.UserSortType, filter *gen.UserFilterType) (*gen.UserResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "users", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "users")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Users to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Users(ctx, offset, limit, q, sort, filter)
}

// CreateUser method
func (r *MutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *gen.User, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "users", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "users")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Users to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateUser(ctx, input)
}

// ReadUser method
func (r *QueryResolver) User(ctx context.Context, id *string, q *string, filter *gen.UserFilterType) (*gen.User, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "users", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "users")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Users to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.User(ctx, id, q, filter)
}

// UpdateUser method
func (r *MutationResolver) UpdateUser(ctx context.Context, id string, input map[string]interface{}) (item *gen.User, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "users", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "users")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Users to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateUser(ctx, id, input)
}

// DeleteUser method
func (r *MutationResolver) DeleteUser(ctx context.Context, id string) (item *gen.User, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "users", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "users")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Users to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteUser(ctx, id)
}

// DeleteAllUsers method
func (r *MutationResolver) DeleteAllUsers(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "users", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "users")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Users to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllUsers(ctx)
}

// UserAPIKeys method
func (r *QueryResolver) UserAPIKeys(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.UserAPIKeySortType, filter *gen.UserAPIKeyFilterType) (*gen.UserAPIKeyResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "user_api_keys", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "user_api_keys")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope UserAPIKeys to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.UserAPIKeys(ctx, offset, limit, q, sort, filter)
}

// CreateUserAPIKey method
func (r *MutationResolver) CreateUserAPIKey(ctx context.Context, input map[string]interface{}) (item *gen.UserAPIKey, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "user_api_keys", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "user_api_keys")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope UserAPIKeys to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateUserAPIKey(ctx, input)
}

// ReadUserAPIKey method
func (r *QueryResolver) UserAPIKey(ctx context.Context, id *string, q *string, filter *gen.UserAPIKeyFilterType) (*gen.UserAPIKey, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "user_api_keys", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "user_api_keys")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope UserAPIKeys to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.UserAPIKey(ctx, id, q, filter)
}

// UpdateUserAPIKey method
func (r *MutationResolver) UpdateUserAPIKey(ctx context.Context, id string, input map[string]interface{}) (item *gen.UserAPIKey, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "user_api_keys", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "user_api_keys")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope UserAPIKeys to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateUserAPIKey(ctx, id, input)
}

// DeleteUserAPIKey method
func (r *MutationResolver) DeleteUserAPIKey(ctx context.Context, id string) (item *gen.UserAPIKey, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "user_api_keys", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "user_api_keys")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope UserAPIKeys to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteUserAPIKey(ctx, id)
}

// DeleteAllUserAPIKeys method
func (r *MutationResolver) DeleteAllUserAPIKeys(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "user_api_keys", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "user_api_keys")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope UserAPIKeys to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllUserAPIKeys(ctx)
}

// Profiles method
func (r *QueryResolver) Profiles(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.ProfileSortType, filter *gen.ProfileFilterType) (*gen.ProfileResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "profiles", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "profiles")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Profiles to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Profiles(ctx, offset, limit, q, sort, filter)
}

// CreateProfile method
func (r *MutationResolver) CreateProfile(ctx context.Context, input map[string]interface{}) (item *gen.Profile, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "profiles", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "profiles")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Profiles to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateProfile(ctx, input)
}

// ReadProfile method
func (r *QueryResolver) Profile(ctx context.Context, id *string, q *string, filter *gen.ProfileFilterType) (*gen.Profile, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "profiles", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "profiles")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Profiles to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Profile(ctx, id, q, filter)
}

// UpdateProfile method
func (r *MutationResolver) UpdateProfile(ctx context.Context, id string, input map[string]interface{}) (item *gen.Profile, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "profiles", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "profiles")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Profiles to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateProfile(ctx, id, input)
}

// DeleteProfile method
func (r *MutationResolver) DeleteProfile(ctx context.Context, id string) (item *gen.Profile, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "profiles", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "profiles")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Profiles to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteProfile(ctx, id)
}

// DeleteAllProfiles method
func (r *MutationResolver) DeleteAllProfiles(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "profiles", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "profiles")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Profiles to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllProfiles(ctx)
}

// Roles method
func (r *QueryResolver) Roles(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.RoleSortType, filter *gen.RoleFilterType) (*gen.RoleResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "roles", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "roles")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Roles to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Roles(ctx, offset, limit, q, sort, filter)
}

// CreateRole method
func (r *MutationResolver) CreateRole(ctx context.Context, input map[string]interface{}) (item *gen.Role, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "roles", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "roles")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Roles to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateRole(ctx, input)
}

// ReadRole method
func (r *QueryResolver) Role(ctx context.Context, id *string, q *string, filter *gen.RoleFilterType) (*gen.Role, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "roles", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "roles")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Roles to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Role(ctx, id, q, filter)
}

// UpdateRole method
func (r *MutationResolver) UpdateRole(ctx context.Context, id string, input map[string]interface{}) (item *gen.Role, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "roles", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "roles")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Roles to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateRole(ctx, id, input)
}

// DeleteRole method
func (r *MutationResolver) DeleteRole(ctx context.Context, id string) (item *gen.Role, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "roles", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "roles")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Roles to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteRole(ctx, id)
}

// DeleteAllRoles method
func (r *MutationResolver) DeleteAllRoles(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "roles", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "roles")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Roles to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllRoles(ctx)
}

// Permissions method
func (r *QueryResolver) Permissions(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PermissionSortType, filter *gen.PermissionFilterType) (*gen.PermissionResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "permissions", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "permissions")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Permissions to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Permissions(ctx, offset, limit, q, sort, filter)
}

// CreatePermission method
func (r *MutationResolver) CreatePermission(ctx context.Context, input map[string]interface{}) (item *gen.Permission, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "permissions", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "permissions")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Permissions to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreatePermission(ctx, input)
}

// ReadPermission method
func (r *QueryResolver) Permission(ctx context.Context, id *string, q *string, filter *gen.PermissionFilterType) (*gen.Permission, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "permissions", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "permissions")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Permissions to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Permission(ctx, id, q, filter)
}

// UpdatePermission method
func (r *MutationResolver) UpdatePermission(ctx context.Context, id string, input map[string]interface{}) (item *gen.Permission, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "permissions", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "permissions")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Permissions to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdatePermission(ctx, id, input)
}

// DeletePermission method
func (r *MutationResolver) DeletePermission(ctx context.Context, id string) (item *gen.Permission, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "permissions", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "permissions")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Permissions to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeletePermission(ctx, id)
}

// DeleteAllPermissions method
func (r *MutationResolver) DeleteAllPermissions(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "permissions", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "permissions")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Permissions to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllPermissions(ctx)
}
