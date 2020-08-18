package src

import (
	"context"
	"fmt"
	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/checkmail"
	"github.com/loopcontext/graphql-orm/events"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// New ...
func New(db *gen.DB, ec *events.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	resolver.Handlers.CreateUser = func(ctx context.Context, r *gen.GeneratedResolver, input map[string]interface{}) (item *gen.User, err error) {
		// Before create
		if err = passwordCheck(input); err != nil {
			return nil, err
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
		if err = emailCheck(input); err != nil {
			return nil, err
		}
		if err = passwordCheck(input); err != nil {
			return nil, err
		}
		item, err = gen.UpdateUserHandler(ctx, r, id, input)
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
				log.Err(err)
				return item, fmt.Errorf("UpdateRole Error - User[%s]: %s", u.ID, err.Error())
			}
		}
		return item, err
	}

	resolver.Handlers.OnEvent = func(ctx context.Context, r *gen.GeneratedResolver, e *events.Event) (err error) {
		// After save
		log.Debug().Msgf("event: %#v", e)
		return nil
	}

	return resolver
}

// Login logs the user in
func (r *QueryResolver) Login(ctx context.Context) (string, error) {
	return "logged in", nil
}

// =============================================================================
// ================================= Helpers ===================================
// =============================================================================

func emailCheck(input map[string]interface{}) (err error) {
	if email, ok := input["email"].(string); ok {
		if email != "" {
			err = checkmail.ValidateFormat(email)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func passwordCheck(input map[string]interface{}) (err error) {
	p, ok := input["password"].(string)
	if input["password"] != nil && p == "" {
		return fmt.Errorf("The password should not be empty")
	} else if ok {
		input["password"], err = hashPassword(input["password"].(string))
		if err != nil {
			return err
		}
	}
	return nil

}

func hashPassword(passw string) (string, error) {
	if passw != "" {
		if pw, err := bcrypt.GenerateFromPassword([]byte(passw), 11); err == nil {
			return string(pw), nil
		}
	}
	return "", fmt.Errorf("If password is set, it cannot be empty")
}

func roleChanges(ctx context.Context, r *gen.GeneratedResolver, userID string, userInput map[string]interface{}) (err error) {
	// Check its roles, and update permissions as should
	u, err := gen.QueryUserHandler(ctx, r, gen.QueryUserHandlerOptions{ID: &userID})
	if err != nil {
		log.Err(err)
	}
	// Deal with role changes
	if roles, ok := userInput["rolesIds"].([]interface{}); ok {
		r.GetDB(ctx).Model(u).Association("Permissions").Clear()
		for _, roleID := range roles {
			if rID, ok := roleID.(string); ok {
				role := gen.Role{ID: rID}
				if err := r.GetDB(ctx).Model(role).Preload("Permissions").First(&role).Error; err != nil {
					log.Err(err)
					continue
					// return err
				}
				if len(role.Permissions) > 0 {
					if err := r.GetDB(ctx).Model(u).Association("Permissions").Append(role.Permissions).Error; err != nil {
						log.Err(err)
						continue
					}
				}
			}
		}
	}
	return nil
}