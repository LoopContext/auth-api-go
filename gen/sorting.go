package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Apply method
func (s UserSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("users"), sorts, joins)
}

// ApplyWithAlias method
func (s UserSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Active != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("active"), Direction: s.Active.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ActiveMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("active") + ")", Direction: s.ActiveMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ActiveMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("active") + ")", Direction: s.ActiveMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Email != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("email"), Direction: s.Email.String()}
		*sorts = append(*sorts, sort)
	}

	if s.EmailMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("email") + ")", Direction: s.EmailMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.EmailMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("email") + ")", Direction: s.EmailMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Password != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("password"), Direction: s.Password.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PasswordMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("password") + ")", Direction: s.PasswordMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PasswordMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("password") + ")", Direction: s.PasswordMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AvatarURL != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("avatarURL"), Direction: s.AvatarURL.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AvatarURLMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("avatarURL") + ")", Direction: s.AvatarURLMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AvatarURLMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("avatarURL") + ")", Direction: s.AvatarURLMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DisplayName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("displayName"), Direction: s.DisplayName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DisplayNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("displayName") + ")", Direction: s.DisplayNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DisplayNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("displayName") + ")", Direction: s.DisplayNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FirstName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("firstName"), Direction: s.FirstName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.FirstNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("firstName") + ")", Direction: s.FirstNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FirstNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("firstName") + ")", Direction: s.FirstNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LastName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("lastName"), Direction: s.LastName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.LastNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("lastName") + ")", Direction: s.LastNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LastNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("lastName") + ")", Direction: s.LastNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NickName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("nickName"), Direction: s.NickName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NickNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("nickName") + ")", Direction: s.NickNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NickNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("nickName") + ")", Direction: s.NickNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Location != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("location"), Direction: s.Location.String()}
		*sorts = append(*sorts, sort)
	}

	if s.LocationMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("location") + ")", Direction: s.LocationMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LocationMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("location") + ")", Direction: s.LocationMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Apikeys != nil {
		_alias := alias + "_apikeys"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("user_api_keys"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("userId")+" = "+dialect.Quote(alias)+".id")
		err := s.Apikeys.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Roles != nil {
		_alias := alias + "_roles"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("role_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" LEFT JOIN "+dialect.Quote(TableName("roles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("roleId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Roles.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Profiles != nil {
		_alias := alias + "_profiles"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("profile_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" LEFT JOIN "+dialect.Quote(TableName("profiles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("profileId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Profiles.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Permissions != nil {
		_alias := alias + "_permissions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("permission_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" LEFT JOIN "+dialect.Quote(TableName("permissions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Permissions.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s UserAPIKeySortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("user_api_keys"), sorts, joins)
}

// ApplyWithAlias method
func (s UserAPIKeySortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Key != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("key"), Direction: s.Key.String()}
		*sorts = append(*sorts, sort)
	}

	if s.KeyMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("key") + ")", Direction: s.KeyMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.KeyMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("key") + ")", Direction: s.KeyMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UserID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("userId"), Direction: s.UserID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UserIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("userId") + ")", Direction: s.UserIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UserIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("userId") + ")", Direction: s.UserIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.User != nil {
		_alias := alias + "_user"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("userId"))
		err := s.User.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Permissions != nil {
		_alias := alias + "_permissions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("userAPIKey_permissions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("apikeyId")+" LEFT JOIN "+dialect.Quote(TableName("permissions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Permissions.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s ProfileSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("profiles"), sorts, joins)
}

// ApplyWithAlias method
func (s ProfileSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Email != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("email"), Direction: s.Email.String()}
		*sorts = append(*sorts, sort)
	}

	if s.EmailMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("email") + ")", Direction: s.EmailMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.EmailMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("email") + ")", Direction: s.EmailMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ExternalUserID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("externalUserId"), Direction: s.ExternalUserID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ExternalUserIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("externalUserId") + ")", Direction: s.ExternalUserIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ExternalUserIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("externalUserId") + ")", Direction: s.ExternalUserIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Provider != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("provider"), Direction: s.Provider.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ProviderMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("provider") + ")", Direction: s.ProviderMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ProviderMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("provider") + ")", Direction: s.ProviderMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AvatarURL != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("avatarURL"), Direction: s.AvatarURL.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AvatarURLMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("avatarURL") + ")", Direction: s.AvatarURLMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AvatarURLMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("avatarURL") + ")", Direction: s.AvatarURLMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FirstName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("firstName"), Direction: s.FirstName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.FirstNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("firstName") + ")", Direction: s.FirstNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FirstNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("firstName") + ")", Direction: s.FirstNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LastName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("lastName"), Direction: s.LastName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.LastNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("lastName") + ")", Direction: s.LastNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LastNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("lastName") + ")", Direction: s.LastNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NickName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("nickName"), Direction: s.NickName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NickNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("nickName") + ")", Direction: s.NickNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NickNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("nickName") + ")", Direction: s.NickNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Location != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("location"), Direction: s.Location.String()}
		*sorts = append(*sorts, sort)
	}

	if s.LocationMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("location") + ")", Direction: s.LocationMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LocationMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("location") + ")", Direction: s.LocationMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Users != nil {
		_alias := alias + "_users"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("profile_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("profileId")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Users.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s RoleSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("roles"), sorts, joins)
}

// ApplyWithAlias method
func (s RoleSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Domain != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("domain"), Direction: s.Domain.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DomainMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("domain") + ")", Direction: s.DomainMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DomainMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("domain") + ")", Direction: s.DomainMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Users != nil {
		_alias := alias + "_users"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("role_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("roleId")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Users.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Parents != nil {
		_alias := alias + "_parents"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("role_parents"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("childId")+" LEFT JOIN "+dialect.Quote(TableName("roles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("parentId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Parents.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Children != nil {
		_alias := alias + "_children"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("role_parents"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("parentId")+" LEFT JOIN "+dialect.Quote(TableName("roles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("childId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Children.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Permissions != nil {
		_alias := alias + "_permissions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("permission_roles"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("roleId")+" LEFT JOIN "+dialect.Quote(TableName("permissions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Permissions.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s PermissionSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("permissions"), sorts, joins)
}

// ApplyWithAlias method
func (s PermissionSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Domain != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("domain"), Direction: s.Domain.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DomainMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("domain") + ")", Direction: s.DomainMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DomainMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("domain") + ")", Direction: s.DomainMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Tag != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("tag"), Direction: s.Tag.String()}
		*sorts = append(*sorts, sort)
	}

	if s.TagMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("tag") + ")", Direction: s.TagMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.TagMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("tag") + ")", Direction: s.TagMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Users != nil {
		_alias := alias + "_users"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("permission_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Users.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Roles != nil {
		_alias := alias + "_roles"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("permission_roles"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" LEFT JOIN "+dialect.Quote(TableName("roles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("roleId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Roles.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Apikeys != nil {
		_alias := alias + "_apikeys"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("userAPIKey_permissions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" LEFT JOIN "+dialect.Quote(TableName("user_api_keys"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("apikeyId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Apikeys.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
