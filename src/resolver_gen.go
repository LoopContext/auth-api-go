package src

import (
	"context"

	"github.com/loopcontext/auth-api-go/gen"
)

// NewResolver ...
func NewResolver(db *gen.DB, ec *gen.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{gen.NewGeneratedResolver(handlers, db, ec)}
}

// Resolver ...
type Resolver struct {
	*gen.GeneratedResolver
}

// MutationResolver ...
type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

// BeginTransaction ...
func (r *MutationResolver) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	ctx = gen.EnrichContextWithMutations(ctx, r.GeneratedResolver)
	err := fn(ctx)
	if err != nil {
		tx := r.GeneratedResolver.GetDB(ctx)
		tx.Rollback()
		return err
	}
	return gen.FinishMutationContext(ctx, r.GeneratedResolver)
}

// QueryResolver ...
type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

// Mutation ...
func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{
		GeneratedMutationResolver: &gen.GeneratedMutationResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// Query ...
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{
		GeneratedQueryResolver: &gen.GeneratedQueryResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// UserResultTypeResolver struct
type UserResultTypeResolver struct {
	*gen.GeneratedUserResultTypeResolver
}

// UserResultType ...
func (r *Resolver) UserResultType() gen.UserResultTypeResolver {
	return &UserResultTypeResolver{
		GeneratedUserResultTypeResolver: &gen.GeneratedUserResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// UserResolver struct
type UserResolver struct {
	*gen.GeneratedUserResolver
}

// User ...
func (r *Resolver) User() gen.UserResolver {
	return &UserResolver{
		GeneratedUserResolver: &gen.GeneratedUserResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// UserAPIKeyResultTypeResolver struct
type UserAPIKeyResultTypeResolver struct {
	*gen.GeneratedUserAPIKeyResultTypeResolver
}

// UserAPIKeyResultType ...
func (r *Resolver) UserAPIKeyResultType() gen.UserAPIKeyResultTypeResolver {
	return &UserAPIKeyResultTypeResolver{
		GeneratedUserAPIKeyResultTypeResolver: &gen.GeneratedUserAPIKeyResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// UserAPIKeyResolver struct
type UserAPIKeyResolver struct {
	*gen.GeneratedUserAPIKeyResolver
}

// UserAPIKey ...
func (r *Resolver) UserAPIKey() gen.UserAPIKeyResolver {
	return &UserAPIKeyResolver{
		GeneratedUserAPIKeyResolver: &gen.GeneratedUserAPIKeyResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// ProfileResultTypeResolver struct
type ProfileResultTypeResolver struct {
	*gen.GeneratedProfileResultTypeResolver
}

// ProfileResultType ...
func (r *Resolver) ProfileResultType() gen.ProfileResultTypeResolver {
	return &ProfileResultTypeResolver{
		GeneratedProfileResultTypeResolver: &gen.GeneratedProfileResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// ProfileResolver struct
type ProfileResolver struct {
	*gen.GeneratedProfileResolver
}

// Profile ...
func (r *Resolver) Profile() gen.ProfileResolver {
	return &ProfileResolver{
		GeneratedProfileResolver: &gen.GeneratedProfileResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// RoleResultTypeResolver struct
type RoleResultTypeResolver struct {
	*gen.GeneratedRoleResultTypeResolver
}

// RoleResultType ...
func (r *Resolver) RoleResultType() gen.RoleResultTypeResolver {
	return &RoleResultTypeResolver{
		GeneratedRoleResultTypeResolver: &gen.GeneratedRoleResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// RoleResolver struct
type RoleResolver struct {
	*gen.GeneratedRoleResolver
}

// Role ...
func (r *Resolver) Role() gen.RoleResolver {
	return &RoleResolver{
		GeneratedRoleResolver: &gen.GeneratedRoleResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// PermissionResultTypeResolver struct
type PermissionResultTypeResolver struct {
	*gen.GeneratedPermissionResultTypeResolver
}

// PermissionResultType ...
func (r *Resolver) PermissionResultType() gen.PermissionResultTypeResolver {
	return &PermissionResultTypeResolver{
		GeneratedPermissionResultTypeResolver: &gen.GeneratedPermissionResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// PermissionResolver struct
type PermissionResolver struct {
	*gen.GeneratedPermissionResolver
}

// Permission ...
func (r *Resolver) Permission() gen.PermissionResolver {
	return &PermissionResolver{
		GeneratedPermissionResolver: &gen.GeneratedPermissionResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}
