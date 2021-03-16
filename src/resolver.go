package src

import (
	"context"
	"fmt"

	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src/utils"
	"github.com/loopcontext/go-graphql-orm/events"
	"github.com/rs/zerolog/log"
)

// New ...
func New(db *gen.DB, ec *gen.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	resolver.Handlers.OnEvent = func(ctx context.Context, r *gen.GeneratedResolver, e *events.Event) (err error) {
		// After save
		// log.Debug().Msgf("event: %#v", e)
		return nil
	}

	resolver.Handlers.CreateUser = func(ctx context.Context, r *gen.GeneratedResolver, input map[string]interface{}) (item *gen.User, err error) {
		// Before create
		if passw, ok := input["password"].(string); ok && passw != "" {
			if newpassw, err := utils.HashPassword(passw); err != nil {
				return nil, err
			} else {
				input["password"] = newpassw
			}
		}
		item, err = gen.CreateUserHandler(ctx, r, input)
		if err != nil {
			return item, err
		}
		// After create
		if err = roleChanges(ctx, r, item.ID, input); err != nil {
			return item, err
		}
		return item, err
	}

	resolver.Handlers.UpdateUser = func(ctx context.Context, r *gen.GeneratedResolver, id string, input map[string]interface{}) (item *gen.User, err error) {
		// Before update
		if email, ok := input["email"].(string); ok && email != "" {
			if err = utils.EmailCheck(email); err != nil {
				return nil, err
			}
		}
		if passw, ok := input["password"].(string); ok && passw != "" {
			if newpassw, err := utils.HashPassword(passw); err != nil {
				return nil, err
			} else {
				input["password"] = newpassw
			}
		}
		item, err = gen.UpdateUserHandler(ctx, r, id, input)
		if err != nil {
			return nil, err // boo
		}
		// After update
		if err = roleChanges(ctx, r, id, input); err != nil {
			return nil, err
		}
		return item, err
	}

	resolver.Handlers.UpdateRole = func(ctx context.Context, r *gen.GeneratedResolver, id string, input map[string]interface{}) (item *gen.Role, err error) {
		// Before update
		item, err = gen.UpdateRoleHandler(ctx, r, id, input)
		// After update
		// Check role users, and update permissions as should
		for _, u := range item.Users {
			roleIds := []string{}
			for _, ur := range u.Roles {
				roleIds = append(roleIds, ur.ID)
			}
			if err = roleChanges(ctx, r, id, map[string]interface{}{"roleIds": roleIds}); err != nil {
				log.Error().Msg(err.Error())
				return item, fmt.Errorf("UpdateRole Error - User[%s]: %s", u.ID, err.Error())
			}
		}
		return item, err
	}

	return resolver
}

// =============================================================================
// ================================= Helpers ===================================
// =============================================================================

func roleChanges(ctx context.Context, r *gen.GeneratedResolver, userID string, userInput map[string]interface{}) (err error) {
	// Check its roles, and update permissions as should
	u, err := gen.QueryUserHandler(ctx, r, gen.QueryUserHandlerOptions{ID: &userID})
	if err != nil {
		log.Error().Msg(err.Error())
	}
	// Deal with role changes
	if roles, ok := userInput["rolesIds"].([]interface{}); ok {
		r.GetDB(ctx).Model(u).Association("Permissions").Clear()
		for _, roleID := range roles {
			if rID, ok := roleID.(string); ok {
				role := gen.Role{ID: rID}
				if err := r.GetDB(ctx).Model(role).Preload("Permissions").First(&role).Error; err != nil {
					log.Error().Msg(err.Error())
					continue
					// return err
				}
				if len(role.Permissions) > 0 {
					if err := r.GetDB(ctx).Model(u).Association("Permissions").Append(role.Permissions).Error; err != nil {
						log.Error().Msg(err.Error())
						continue
					}
				}
			}
		}
	}
	return nil
}
