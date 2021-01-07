package gen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/vektah/gqlparser/v2/ast"
)

// GeneratedQueryResolver struct
type GeneratedQueryResolver struct{ *GeneratedResolver }

// QueryUserHandlerOptions struct
type QueryUserHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *UserFilterType
}

// User ...
func (r *GeneratedQueryResolver) User(ctx context.Context, id *string, q *string, filter *UserFilterType) (*User, error) {
	opts := QueryUserHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryUser(ctx, r.GeneratedResolver, opts)
}

// QueryUserHandler handler
func QueryUserHandler(ctx context.Context, r *GeneratedResolver, opts QueryUserHandlerOptions) (*User, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := UserQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &UserResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("users")+".id = ?", *opts.ID)
	}

	var items []*User
	giOpts := GetItemsOptions{
		Alias: TableName("users"),
		Preloaders: []string{
			"Apikeys",
			"Roles",
			"Profiles",
			"Permissions",
		},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryUsersHandlerOptions struct
type QueryUsersHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*UserSortType
	Filter *UserFilterType
}

// Users handler options
func (r *GeneratedQueryResolver) Users(ctx context.Context, offset *int, limit *int, q *string, sort []*UserSortType, filter *UserFilterType) (*UserResultType, error) {
	opts := QueryUsersHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryUsers(ctx, r.GeneratedResolver, opts)
}

// UsersItems handler
func (r *GeneratedResolver) UsersItems(ctx context.Context, opts QueryUsersHandlerOptions) (res []*User, err error) {
	resultType, err := r.Handlers.QueryUsers(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("users"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// UsersCount handler
func (r *GeneratedResolver) UsersCount(ctx context.Context, opts QueryUsersHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryUsers(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("users"),
	}, &User{})
}

// QueryUsersHandler handler
func QueryUsersHandler(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions) (*UserResultType, error) {
	query := UserQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &UserResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedUserResultTypeResolver struct
type GeneratedUserResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedUserResultTypeResolver) Items(ctx context.Context, obj *UserResultType) (items []*User, err error) {
	otps := GetItemsOptions{
		Alias: TableName("users"),
		Preloaders: []string{
			"Apikeys",
			"Roles",
			"Profiles",
			"Permissions",
		},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), otps, &items)

	for _, item := range items {

		item.ApikeysPreloaded = true
		item.RolesPreloaded = true
		item.ProfilesPreloaded = true
		item.PermissionsPreloaded = true
	}

	uniqueItems := []*User{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedUserResultTypeResolver) Count(ctx context.Context, obj *UserResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias: TableName("users"),
		Preloaders: []string{
			"Apikeys",
			"Roles",
			"Profiles",
			"Permissions",
		},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &User{})
}

// GeneratedUserResolver struct
type GeneratedUserResolver struct{ *GeneratedResolver }

// Apikeys ...
func (r *GeneratedUserResolver) Apikeys(ctx context.Context, obj *User) (res []*UserAPIKey, err error) {
	return r.Handlers.UserApikeys(ctx, r.GeneratedResolver, obj)
}

// UserApikeysHandler handler
func UserApikeysHandler(ctx context.Context, r *GeneratedResolver, obj *User) (res []*UserAPIKey, err error) {

	if obj.ApikeysPreloaded {
		res = obj.Apikeys
	} else {

		items := []*UserAPIKey{}
		db := r.GetDB(ctx)
		err = db.Model(obj).Related(&items, "Apikeys").Error
		res = items

	}

	return
}

// ApikeysIds ...
func (r *GeneratedUserResolver) ApikeysIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*UserAPIKey{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("user_api_keys")+".id").Related(&items, "Apikeys").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// ApikeysConnection method
func (r *GeneratedUserResolver) ApikeysConnection(ctx context.Context, obj *User, offset *int, limit *int, q *string, sort []*UserAPIKeySortType, filter *UserAPIKeyFilterType) (res *UserAPIKeyResultType, err error) {
	f := &UserAPIKeyFilterType{
		User: &UserFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &UserAPIKeyFilterType{
			And: []*UserAPIKeyFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryUserAPIKeysHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryUserAPIKeys(ctx, r.GeneratedResolver, opts)
}

// Roles ...
func (r *GeneratedUserResolver) Roles(ctx context.Context, obj *User) (res []*Role, err error) {
	return r.Handlers.UserRoles(ctx, r.GeneratedResolver, obj)
}

// UserRolesHandler handler
func UserRolesHandler(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Role, err error) {

	if obj.RolesPreloaded {
		res = obj.Roles
	} else {

		items := []*Role{}
		db := r.GetDB(ctx)
		err = db.Model(obj).Related(&items, "Roles").Error
		res = items

	}

	return
}

// RolesIds ...
func (r *GeneratedUserResolver) RolesIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*Role{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("roles")+".id").Related(&items, "Roles").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// RolesConnection method
func (r *GeneratedUserResolver) RolesConnection(ctx context.Context, obj *User, offset *int, limit *int, q *string, sort []*RoleSortType, filter *RoleFilterType) (res *RoleResultType, err error) {
	f := &RoleFilterType{
		Users: &UserFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &RoleFilterType{
			And: []*RoleFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryRolesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryRoles(ctx, r.GeneratedResolver, opts)
}

// Profiles ...
func (r *GeneratedUserResolver) Profiles(ctx context.Context, obj *User) (res []*Profile, err error) {
	return r.Handlers.UserProfiles(ctx, r.GeneratedResolver, obj)
}

// UserProfilesHandler handler
func UserProfilesHandler(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Profile, err error) {

	if obj.ProfilesPreloaded {
		res = obj.Profiles
	} else {

		items := []*Profile{}
		db := r.GetDB(ctx)
		err = db.Model(obj).Related(&items, "Profiles").Error
		res = items

	}

	return
}

// ProfilesIds ...
func (r *GeneratedUserResolver) ProfilesIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*Profile{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("profiles")+".id").Related(&items, "Profiles").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// ProfilesConnection method
func (r *GeneratedUserResolver) ProfilesConnection(ctx context.Context, obj *User, offset *int, limit *int, q *string, sort []*ProfileSortType, filter *ProfileFilterType) (res *ProfileResultType, err error) {
	f := &ProfileFilterType{
		Users: &UserFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &ProfileFilterType{
			And: []*ProfileFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryProfilesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryProfiles(ctx, r.GeneratedResolver, opts)
}

// Permissions ...
func (r *GeneratedUserResolver) Permissions(ctx context.Context, obj *User) (res []*Permission, err error) {
	return r.Handlers.UserPermissions(ctx, r.GeneratedResolver, obj)
}

// UserPermissionsHandler handler
func UserPermissionsHandler(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Permission, err error) {

	if obj.PermissionsPreloaded {
		res = obj.Permissions
	} else {

		items := []*Permission{}
		db := r.GetDB(ctx)
		err = db.Model(obj).Related(&items, "Permissions").Error
		res = items

	}

	return
}

// PermissionsIds ...
func (r *GeneratedUserResolver) PermissionsIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*Permission{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("permissions")+".id").Related(&items, "Permissions").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// PermissionsConnection method
func (r *GeneratedUserResolver) PermissionsConnection(ctx context.Context, obj *User, offset *int, limit *int, q *string, sort []*PermissionSortType, filter *PermissionFilterType) (res *PermissionResultType, err error) {
	f := &PermissionFilterType{
		Users: &UserFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &PermissionFilterType{
			And: []*PermissionFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryPermissionsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPermissions(ctx, r.GeneratedResolver, opts)
}

// QueryUserAPIKeyHandlerOptions struct
type QueryUserAPIKeyHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *UserAPIKeyFilterType
}

// UserAPIKey ...
func (r *GeneratedQueryResolver) UserAPIKey(ctx context.Context, id *string, q *string, filter *UserAPIKeyFilterType) (*UserAPIKey, error) {
	opts := QueryUserAPIKeyHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryUserAPIKey(ctx, r.GeneratedResolver, opts)
}

// QueryUserAPIKeyHandler handler
func QueryUserAPIKeyHandler(ctx context.Context, r *GeneratedResolver, opts QueryUserAPIKeyHandlerOptions) (*UserAPIKey, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := UserAPIKeyQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &UserAPIKeyResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("user_api_keys")+".id = ?", *opts.ID)
	}

	var items []*UserAPIKey
	giOpts := GetItemsOptions{
		Alias: TableName("user_api_keys"),
		Preloaders: []string{
			"User",
			"Permissions",
		},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryUserAPIKeysHandlerOptions struct
type QueryUserAPIKeysHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*UserAPIKeySortType
	Filter *UserAPIKeyFilterType
}

// UserAPIKeys handler options
func (r *GeneratedQueryResolver) UserAPIKeys(ctx context.Context, offset *int, limit *int, q *string, sort []*UserAPIKeySortType, filter *UserAPIKeyFilterType) (*UserAPIKeyResultType, error) {
	opts := QueryUserAPIKeysHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryUserAPIKeys(ctx, r.GeneratedResolver, opts)
}

// UserAPIKeysItems handler
func (r *GeneratedResolver) UserAPIKeysItems(ctx context.Context, opts QueryUserAPIKeysHandlerOptions) (res []*UserAPIKey, err error) {
	resultType, err := r.Handlers.QueryUserAPIKeys(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("user_api_keys"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// UserAPIKeysCount handler
func (r *GeneratedResolver) UserAPIKeysCount(ctx context.Context, opts QueryUserAPIKeysHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryUserAPIKeys(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("user_api_keys"),
	}, &UserAPIKey{})
}

// QueryUserAPIKeysHandler handler
func QueryUserAPIKeysHandler(ctx context.Context, r *GeneratedResolver, opts QueryUserAPIKeysHandlerOptions) (*UserAPIKeyResultType, error) {
	query := UserAPIKeyQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &UserAPIKeyResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedUserAPIKeyResultTypeResolver struct
type GeneratedUserAPIKeyResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedUserAPIKeyResultTypeResolver) Items(ctx context.Context, obj *UserAPIKeyResultType) (items []*UserAPIKey, err error) {
	otps := GetItemsOptions{
		Alias: TableName("user_api_keys"),
		Preloaders: []string{
			"User",
			"Permissions",
		},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), otps, &items)

	for _, item := range items {

		item.UserPreloaded = true
		item.PermissionsPreloaded = true
	}

	uniqueItems := []*UserAPIKey{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedUserAPIKeyResultTypeResolver) Count(ctx context.Context, obj *UserAPIKeyResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias: TableName("user_api_keys"),
		Preloaders: []string{
			"User",
			"Permissions",
		},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &UserAPIKey{})
}

// GeneratedUserAPIKeyResolver struct
type GeneratedUserAPIKeyResolver struct{ *GeneratedResolver }

// User ...
func (r *GeneratedUserAPIKeyResolver) User(ctx context.Context, obj *UserAPIKey) (res *User, err error) {
	return r.Handlers.UserAPIKeyUser(ctx, r.GeneratedResolver, obj)
}

// UserAPIKeyUserHandler handler
func UserAPIKeyUserHandler(ctx context.Context, r *GeneratedResolver, obj *UserAPIKey) (res *User, err error) {

	if obj.UserPreloaded {
		res = obj.User
	} else {

		loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
		if obj.UserID != nil {
			item, _err := loaders["User"].Load(ctx, dataloader.StringKey(*obj.UserID))()
			res, _ = item.(*User)

			err = _err
		}

	}

	return
}

// Permissions ...
func (r *GeneratedUserAPIKeyResolver) Permissions(ctx context.Context, obj *UserAPIKey) (res []*Permission, err error) {
	return r.Handlers.UserAPIKeyPermissions(ctx, r.GeneratedResolver, obj)
}

// UserAPIKeyPermissionsHandler handler
func UserAPIKeyPermissionsHandler(ctx context.Context, r *GeneratedResolver, obj *UserAPIKey) (res []*Permission, err error) {

	if obj.PermissionsPreloaded {
		res = obj.Permissions
	} else {

		items := []*Permission{}
		db := r.GetDB(ctx)
		err = db.Model(obj).Related(&items, "Permissions").Error
		res = items

	}

	return
}

// PermissionsIds ...
func (r *GeneratedUserAPIKeyResolver) PermissionsIds(ctx context.Context, obj *UserAPIKey) (ids []string, err error) {
	ids = []string{}

	items := []*Permission{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("permissions")+".id").Related(&items, "Permissions").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// PermissionsConnection method
func (r *GeneratedUserAPIKeyResolver) PermissionsConnection(ctx context.Context, obj *UserAPIKey, offset *int, limit *int, q *string, sort []*PermissionSortType, filter *PermissionFilterType) (res *PermissionResultType, err error) {
	f := &PermissionFilterType{
		Apikeys: &UserAPIKeyFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &PermissionFilterType{
			And: []*PermissionFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryPermissionsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPermissions(ctx, r.GeneratedResolver, opts)
}

// QueryProfileHandlerOptions struct
type QueryProfileHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ProfileFilterType
}

// Profile ...
func (r *GeneratedQueryResolver) Profile(ctx context.Context, id *string, q *string, filter *ProfileFilterType) (*Profile, error) {
	opts := QueryProfileHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryProfile(ctx, r.GeneratedResolver, opts)
}

// QueryProfileHandler handler
func QueryProfileHandler(ctx context.Context, r *GeneratedResolver, opts QueryProfileHandlerOptions) (*Profile, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := ProfileQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &ProfileResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("profiles")+".id = ?", *opts.ID)
	}

	var items []*Profile
	giOpts := GetItemsOptions{
		Alias:      TableName("profiles"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryProfilesHandlerOptions struct
type QueryProfilesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*ProfileSortType
	Filter *ProfileFilterType
}

// Profiles handler options
func (r *GeneratedQueryResolver) Profiles(ctx context.Context, offset *int, limit *int, q *string, sort []*ProfileSortType, filter *ProfileFilterType) (*ProfileResultType, error) {
	opts := QueryProfilesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryProfiles(ctx, r.GeneratedResolver, opts)
}

// ProfilesItems handler
func (r *GeneratedResolver) ProfilesItems(ctx context.Context, opts QueryProfilesHandlerOptions) (res []*Profile, err error) {
	resultType, err := r.Handlers.QueryProfiles(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("profiles"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// ProfilesCount handler
func (r *GeneratedResolver) ProfilesCount(ctx context.Context, opts QueryProfilesHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryProfiles(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("profiles"),
	}, &Profile{})
}

// QueryProfilesHandler handler
func QueryProfilesHandler(ctx context.Context, r *GeneratedResolver, opts QueryProfilesHandlerOptions) (*ProfileResultType, error) {
	query := ProfileQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &ProfileResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedProfileResultTypeResolver struct
type GeneratedProfileResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedProfileResultTypeResolver) Items(ctx context.Context, obj *ProfileResultType) (items []*Profile, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("profiles"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), otps, &items)

	uniqueItems := []*Profile{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedProfileResultTypeResolver) Count(ctx context.Context, obj *ProfileResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("profiles"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &Profile{})
}

// GeneratedProfileResolver struct
type GeneratedProfileResolver struct{ *GeneratedResolver }

// Users ...
func (r *GeneratedProfileResolver) Users(ctx context.Context, obj *Profile) (res []*User, err error) {
	return r.Handlers.ProfileUsers(ctx, r.GeneratedResolver, obj)
}

// ProfileUsersHandler handler
func ProfileUsersHandler(ctx context.Context, r *GeneratedResolver, obj *Profile) (res []*User, err error) {

	items := []*User{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Related(&items, "Users").Error
	res = items

	return
}

// UsersIds ...
func (r *GeneratedProfileResolver) UsersIds(ctx context.Context, obj *Profile) (ids []string, err error) {
	ids = []string{}

	items := []*User{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("users")+".id").Related(&items, "Users").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// UsersConnection method
func (r *GeneratedProfileResolver) UsersConnection(ctx context.Context, obj *Profile, offset *int, limit *int, q *string, sort []*UserSortType, filter *UserFilterType) (res *UserResultType, err error) {
	f := &UserFilterType{
		Profiles: &ProfileFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &UserFilterType{
			And: []*UserFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryUsersHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryUsers(ctx, r.GeneratedResolver, opts)
}

// QueryRoleHandlerOptions struct
type QueryRoleHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *RoleFilterType
}

// Role ...
func (r *GeneratedQueryResolver) Role(ctx context.Context, id *string, q *string, filter *RoleFilterType) (*Role, error) {
	opts := QueryRoleHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryRole(ctx, r.GeneratedResolver, opts)
}

// QueryRoleHandler handler
func QueryRoleHandler(ctx context.Context, r *GeneratedResolver, opts QueryRoleHandlerOptions) (*Role, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := RoleQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &RoleResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("roles")+".id = ?", *opts.ID)
	}

	var items []*Role
	giOpts := GetItemsOptions{
		Alias:      TableName("roles"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryRolesHandlerOptions struct
type QueryRolesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*RoleSortType
	Filter *RoleFilterType
}

// Roles handler options
func (r *GeneratedQueryResolver) Roles(ctx context.Context, offset *int, limit *int, q *string, sort []*RoleSortType, filter *RoleFilterType) (*RoleResultType, error) {
	opts := QueryRolesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryRoles(ctx, r.GeneratedResolver, opts)
}

// RolesItems handler
func (r *GeneratedResolver) RolesItems(ctx context.Context, opts QueryRolesHandlerOptions) (res []*Role, err error) {
	resultType, err := r.Handlers.QueryRoles(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("roles"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// RolesCount handler
func (r *GeneratedResolver) RolesCount(ctx context.Context, opts QueryRolesHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryRoles(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("roles"),
	}, &Role{})
}

// QueryRolesHandler handler
func QueryRolesHandler(ctx context.Context, r *GeneratedResolver, opts QueryRolesHandlerOptions) (*RoleResultType, error) {
	query := RoleQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &RoleResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedRoleResultTypeResolver struct
type GeneratedRoleResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedRoleResultTypeResolver) Items(ctx context.Context, obj *RoleResultType) (items []*Role, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("roles"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), otps, &items)

	uniqueItems := []*Role{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedRoleResultTypeResolver) Count(ctx context.Context, obj *RoleResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("roles"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &Role{})
}

// GeneratedRoleResolver struct
type GeneratedRoleResolver struct{ *GeneratedResolver }

// Users ...
func (r *GeneratedRoleResolver) Users(ctx context.Context, obj *Role) (res []*User, err error) {
	return r.Handlers.RoleUsers(ctx, r.GeneratedResolver, obj)
}

// RoleUsersHandler handler
func RoleUsersHandler(ctx context.Context, r *GeneratedResolver, obj *Role) (res []*User, err error) {

	items := []*User{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Related(&items, "Users").Error
	res = items

	return
}

// UsersIds ...
func (r *GeneratedRoleResolver) UsersIds(ctx context.Context, obj *Role) (ids []string, err error) {
	ids = []string{}

	items := []*User{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("users")+".id").Related(&items, "Users").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// UsersConnection method
func (r *GeneratedRoleResolver) UsersConnection(ctx context.Context, obj *Role, offset *int, limit *int, q *string, sort []*UserSortType, filter *UserFilterType) (res *UserResultType, err error) {
	f := &UserFilterType{
		Roles: &RoleFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &UserFilterType{
			And: []*UserFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryUsersHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryUsers(ctx, r.GeneratedResolver, opts)
}

// Parents ...
func (r *GeneratedRoleResolver) Parents(ctx context.Context, obj *Role) (res []*Role, err error) {
	return r.Handlers.RoleParents(ctx, r.GeneratedResolver, obj)
}

// RoleParentsHandler handler
func RoleParentsHandler(ctx context.Context, r *GeneratedResolver, obj *Role) (res []*Role, err error) {

	items := []*Role{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Related(&items, "Parents").Error
	res = items

	return
}

// ParentsIds ...
func (r *GeneratedRoleResolver) ParentsIds(ctx context.Context, obj *Role) (ids []string, err error) {
	ids = []string{}

	items := []*Role{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("roles")+".id").Related(&items, "Parents").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// ParentsConnection method
func (r *GeneratedRoleResolver) ParentsConnection(ctx context.Context, obj *Role, offset *int, limit *int, q *string, sort []*RoleSortType, filter *RoleFilterType) (res *RoleResultType, err error) {
	f := &RoleFilterType{
		Children: &RoleFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &RoleFilterType{
			And: []*RoleFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryRolesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryRoles(ctx, r.GeneratedResolver, opts)
}

// Children ...
func (r *GeneratedRoleResolver) Children(ctx context.Context, obj *Role) (res []*Role, err error) {
	return r.Handlers.RoleChildren(ctx, r.GeneratedResolver, obj)
}

// RoleChildrenHandler handler
func RoleChildrenHandler(ctx context.Context, r *GeneratedResolver, obj *Role) (res []*Role, err error) {

	items := []*Role{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Related(&items, "Children").Error
	res = items

	return
}

// ChildrenIds ...
func (r *GeneratedRoleResolver) ChildrenIds(ctx context.Context, obj *Role) (ids []string, err error) {
	ids = []string{}

	items := []*Role{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("roles")+".id").Related(&items, "Children").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// ChildrenConnection method
func (r *GeneratedRoleResolver) ChildrenConnection(ctx context.Context, obj *Role, offset *int, limit *int, q *string, sort []*RoleSortType, filter *RoleFilterType) (res *RoleResultType, err error) {
	f := &RoleFilterType{
		Parents: &RoleFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &RoleFilterType{
			And: []*RoleFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryRolesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryRoles(ctx, r.GeneratedResolver, opts)
}

// Permissions ...
func (r *GeneratedRoleResolver) Permissions(ctx context.Context, obj *Role) (res []*Permission, err error) {
	return r.Handlers.RolePermissions(ctx, r.GeneratedResolver, obj)
}

// RolePermissionsHandler handler
func RolePermissionsHandler(ctx context.Context, r *GeneratedResolver, obj *Role) (res []*Permission, err error) {

	items := []*Permission{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Related(&items, "Permissions").Error
	res = items

	return
}

// PermissionsIds ...
func (r *GeneratedRoleResolver) PermissionsIds(ctx context.Context, obj *Role) (ids []string, err error) {
	ids = []string{}

	items := []*Permission{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("permissions")+".id").Related(&items, "Permissions").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// PermissionsConnection method
func (r *GeneratedRoleResolver) PermissionsConnection(ctx context.Context, obj *Role, offset *int, limit *int, q *string, sort []*PermissionSortType, filter *PermissionFilterType) (res *PermissionResultType, err error) {
	f := &PermissionFilterType{
		Roles: &RoleFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &PermissionFilterType{
			And: []*PermissionFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryPermissionsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPermissions(ctx, r.GeneratedResolver, opts)
}

// QueryPermissionHandlerOptions struct
type QueryPermissionHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *PermissionFilterType
}

// Permission ...
func (r *GeneratedQueryResolver) Permission(ctx context.Context, id *string, q *string, filter *PermissionFilterType) (*Permission, error) {
	opts := QueryPermissionHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryPermission(ctx, r.GeneratedResolver, opts)
}

// QueryPermissionHandler handler
func QueryPermissionHandler(ctx context.Context, r *GeneratedResolver, opts QueryPermissionHandlerOptions) (*Permission, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := PermissionQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &PermissionResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("permissions")+".id = ?", *opts.ID)
	}

	var items []*Permission
	giOpts := GetItemsOptions{
		Alias:      TableName("permissions"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryPermissionsHandlerOptions struct
type QueryPermissionsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*PermissionSortType
	Filter *PermissionFilterType
}

// Permissions handler options
func (r *GeneratedQueryResolver) Permissions(ctx context.Context, offset *int, limit *int, q *string, sort []*PermissionSortType, filter *PermissionFilterType) (*PermissionResultType, error) {
	opts := QueryPermissionsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPermissions(ctx, r.GeneratedResolver, opts)
}

// PermissionsItems handler
func (r *GeneratedResolver) PermissionsItems(ctx context.Context, opts QueryPermissionsHandlerOptions) (res []*Permission, err error) {
	resultType, err := r.Handlers.QueryPermissions(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("permissions"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// PermissionsCount handler
func (r *GeneratedResolver) PermissionsCount(ctx context.Context, opts QueryPermissionsHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryPermissions(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("permissions"),
	}, &Permission{})
}

// QueryPermissionsHandler handler
func QueryPermissionsHandler(ctx context.Context, r *GeneratedResolver, opts QueryPermissionsHandlerOptions) (*PermissionResultType, error) {
	query := PermissionQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &PermissionResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedPermissionResultTypeResolver struct
type GeneratedPermissionResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedPermissionResultTypeResolver) Items(ctx context.Context, obj *PermissionResultType) (items []*Permission, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("permissions"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), otps, &items)

	uniqueItems := []*Permission{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedPermissionResultTypeResolver) Count(ctx context.Context, obj *PermissionResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("permissions"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &Permission{})
}

// GeneratedPermissionResolver struct
type GeneratedPermissionResolver struct{ *GeneratedResolver }

// Users ...
func (r *GeneratedPermissionResolver) Users(ctx context.Context, obj *Permission) (res []*User, err error) {
	return r.Handlers.PermissionUsers(ctx, r.GeneratedResolver, obj)
}

// PermissionUsersHandler handler
func PermissionUsersHandler(ctx context.Context, r *GeneratedResolver, obj *Permission) (res []*User, err error) {

	items := []*User{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Related(&items, "Users").Error
	res = items

	return
}

// UsersIds ...
func (r *GeneratedPermissionResolver) UsersIds(ctx context.Context, obj *Permission) (ids []string, err error) {
	ids = []string{}

	items := []*User{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("users")+".id").Related(&items, "Users").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// UsersConnection method
func (r *GeneratedPermissionResolver) UsersConnection(ctx context.Context, obj *Permission, offset *int, limit *int, q *string, sort []*UserSortType, filter *UserFilterType) (res *UserResultType, err error) {
	f := &UserFilterType{
		Permissions: &PermissionFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &UserFilterType{
			And: []*UserFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryUsersHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryUsers(ctx, r.GeneratedResolver, opts)
}

// Roles ...
func (r *GeneratedPermissionResolver) Roles(ctx context.Context, obj *Permission) (res []*Role, err error) {
	return r.Handlers.PermissionRoles(ctx, r.GeneratedResolver, obj)
}

// PermissionRolesHandler handler
func PermissionRolesHandler(ctx context.Context, r *GeneratedResolver, obj *Permission) (res []*Role, err error) {

	items := []*Role{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Related(&items, "Roles").Error
	res = items

	return
}

// RolesIds ...
func (r *GeneratedPermissionResolver) RolesIds(ctx context.Context, obj *Permission) (ids []string, err error) {
	ids = []string{}

	items := []*Role{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("roles")+".id").Related(&items, "Roles").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// RolesConnection method
func (r *GeneratedPermissionResolver) RolesConnection(ctx context.Context, obj *Permission, offset *int, limit *int, q *string, sort []*RoleSortType, filter *RoleFilterType) (res *RoleResultType, err error) {
	f := &RoleFilterType{
		Permissions: &PermissionFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &RoleFilterType{
			And: []*RoleFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryRolesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryRoles(ctx, r.GeneratedResolver, opts)
}

// Apikeys ...
func (r *GeneratedPermissionResolver) Apikeys(ctx context.Context, obj *Permission) (res []*UserAPIKey, err error) {
	return r.Handlers.PermissionApikeys(ctx, r.GeneratedResolver, obj)
}

// PermissionApikeysHandler handler
func PermissionApikeysHandler(ctx context.Context, r *GeneratedResolver, obj *Permission) (res []*UserAPIKey, err error) {

	items := []*UserAPIKey{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Related(&items, "Apikeys").Error
	res = items

	return
}

// ApikeysIds ...
func (r *GeneratedPermissionResolver) ApikeysIds(ctx context.Context, obj *Permission) (ids []string, err error) {
	ids = []string{}

	items := []*UserAPIKey{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("user_api_keys")+".id").Related(&items, "Apikeys").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// ApikeysConnection method
func (r *GeneratedPermissionResolver) ApikeysConnection(ctx context.Context, obj *Permission, offset *int, limit *int, q *string, sort []*UserAPIKeySortType, filter *UserAPIKeyFilterType) (res *UserAPIKeyResultType, err error) {
	f := &UserAPIKeyFilterType{
		Permissions: &PermissionFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &UserAPIKeyFilterType{
			And: []*UserAPIKeyFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryUserAPIKeysHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryUserAPIKeys(ctx, r.GeneratedResolver, opts)
}
