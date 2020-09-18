package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// IsEmpty ...
func (f *UserFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *UserFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("users"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *UserFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Apikeys != nil {
		_alias := alias + "_apikeys"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("user_api_keys"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("userId")+" = "+dialect.Quote(alias)+".id")
		err := f.Apikeys.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Roles != nil {
		_alias := alias + "_roles"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("role_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" LEFT JOIN "+dialect.Quote(TableName("roles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("roleId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Roles.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Profiles != nil {
		_alias := alias + "_profiles"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("profile_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" LEFT JOIN "+dialect.Quote(TableName("profiles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("profileId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Profiles.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Permissions != nil {
		_alias := alias + "_permissions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("permission_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" LEFT JOIN "+dialect.Quote(TableName("permissions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Permissions.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *UserFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Active != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("active")+" = ?")
		values = append(values, f.Active)
	}

	if f.ActiveNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("active")+" != ?")
		values = append(values, f.ActiveNe)
	}

	if f.ActiveGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("active")+" > ?")
		values = append(values, f.ActiveGt)
	}

	if f.ActiveLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("active")+" < ?")
		values = append(values, f.ActiveLt)
	}

	if f.ActiveGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("active")+" >= ?")
		values = append(values, f.ActiveGte)
	}

	if f.ActiveLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("active")+" <= ?")
		values = append(values, f.ActiveLte)
	}

	if f.ActiveIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("active")+" IN (?)")
		values = append(values, f.ActiveIn)
	}

	if f.ActiveNull != nil {
		if *f.ActiveNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("active")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("active")+" IS NOT NULL")
		}
	}

	if f.Email != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" = ?")
		values = append(values, f.Email)
	}

	if f.EmailNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" != ?")
		values = append(values, f.EmailNe)
	}

	if f.EmailGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" > ?")
		values = append(values, f.EmailGt)
	}

	if f.EmailLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" < ?")
		values = append(values, f.EmailLt)
	}

	if f.EmailGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" >= ?")
		values = append(values, f.EmailGte)
	}

	if f.EmailLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" <= ?")
		values = append(values, f.EmailLte)
	}

	if f.EmailIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IN (?)")
		values = append(values, f.EmailIn)
	}

	if f.EmailLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailPrefix))
	}

	if f.EmailSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailSuffix))
	}

	if f.EmailNull != nil {
		if *f.EmailNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IS NOT NULL")
		}
	}

	if f.Password != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("password")+" = ?")
		values = append(values, f.Password)
	}

	if f.PasswordNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("password")+" != ?")
		values = append(values, f.PasswordNe)
	}

	if f.PasswordGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("password")+" > ?")
		values = append(values, f.PasswordGt)
	}

	if f.PasswordLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("password")+" < ?")
		values = append(values, f.PasswordLt)
	}

	if f.PasswordGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("password")+" >= ?")
		values = append(values, f.PasswordGte)
	}

	if f.PasswordLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("password")+" <= ?")
		values = append(values, f.PasswordLte)
	}

	if f.PasswordIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("password")+" IN (?)")
		values = append(values, f.PasswordIn)
	}

	if f.PasswordLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("password")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PasswordLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PasswordPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("password")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PasswordPrefix))
	}

	if f.PasswordSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("password")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PasswordSuffix))
	}

	if f.PasswordNull != nil {
		if *f.PasswordNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("password")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("password")+" IS NOT NULL")
		}
	}

	if f.AvatarURL != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" = ?")
		values = append(values, f.AvatarURL)
	}

	if f.AvatarURLNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" != ?")
		values = append(values, f.AvatarURLNe)
	}

	if f.AvatarURLGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" > ?")
		values = append(values, f.AvatarURLGt)
	}

	if f.AvatarURLLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" < ?")
		values = append(values, f.AvatarURLLt)
	}

	if f.AvatarURLGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" >= ?")
		values = append(values, f.AvatarURLGte)
	}

	if f.AvatarURLLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" <= ?")
		values = append(values, f.AvatarURLLte)
	}

	if f.AvatarURLIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" IN (?)")
		values = append(values, f.AvatarURLIn)
	}

	if f.AvatarURLLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AvatarURLLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AvatarURLPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AvatarURLPrefix))
	}

	if f.AvatarURLSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AvatarURLSuffix))
	}

	if f.AvatarURLNull != nil {
		if *f.AvatarURLNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" IS NOT NULL")
		}
	}

	if f.DisplayName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" = ?")
		values = append(values, f.DisplayName)
	}

	if f.DisplayNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" != ?")
		values = append(values, f.DisplayNameNe)
	}

	if f.DisplayNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" > ?")
		values = append(values, f.DisplayNameGt)
	}

	if f.DisplayNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" < ?")
		values = append(values, f.DisplayNameLt)
	}

	if f.DisplayNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" >= ?")
		values = append(values, f.DisplayNameGte)
	}

	if f.DisplayNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" <= ?")
		values = append(values, f.DisplayNameLte)
	}

	if f.DisplayNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" IN (?)")
		values = append(values, f.DisplayNameIn)
	}

	if f.DisplayNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DisplayNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DisplayNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DisplayNamePrefix))
	}

	if f.DisplayNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DisplayNameSuffix))
	}

	if f.DisplayNameNull != nil {
		if *f.DisplayNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("displayName")+" IS NOT NULL")
		}
	}

	if f.FirstName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" = ?")
		values = append(values, f.FirstName)
	}

	if f.FirstNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" != ?")
		values = append(values, f.FirstNameNe)
	}

	if f.FirstNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" > ?")
		values = append(values, f.FirstNameGt)
	}

	if f.FirstNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" < ?")
		values = append(values, f.FirstNameLt)
	}

	if f.FirstNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" >= ?")
		values = append(values, f.FirstNameGte)
	}

	if f.FirstNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" <= ?")
		values = append(values, f.FirstNameLte)
	}

	if f.FirstNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IN (?)")
		values = append(values, f.FirstNameIn)
	}

	if f.FirstNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FirstNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNamePrefix))
	}

	if f.FirstNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameSuffix))
	}

	if f.FirstNameNull != nil {
		if *f.FirstNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IS NOT NULL")
		}
	}

	if f.LastName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" = ?")
		values = append(values, f.LastName)
	}

	if f.LastNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" != ?")
		values = append(values, f.LastNameNe)
	}

	if f.LastNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" > ?")
		values = append(values, f.LastNameGt)
	}

	if f.LastNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" < ?")
		values = append(values, f.LastNameLt)
	}

	if f.LastNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" >= ?")
		values = append(values, f.LastNameGte)
	}

	if f.LastNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" <= ?")
		values = append(values, f.LastNameLte)
	}

	if f.LastNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IN (?)")
		values = append(values, f.LastNameIn)
	}

	if f.LastNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LastNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNamePrefix))
	}

	if f.LastNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameSuffix))
	}

	if f.LastNameNull != nil {
		if *f.LastNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IS NOT NULL")
		}
	}

	if f.NickName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" = ?")
		values = append(values, f.NickName)
	}

	if f.NickNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" != ?")
		values = append(values, f.NickNameNe)
	}

	if f.NickNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" > ?")
		values = append(values, f.NickNameGt)
	}

	if f.NickNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" < ?")
		values = append(values, f.NickNameLt)
	}

	if f.NickNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" >= ?")
		values = append(values, f.NickNameGte)
	}

	if f.NickNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" <= ?")
		values = append(values, f.NickNameLte)
	}

	if f.NickNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" IN (?)")
		values = append(values, f.NickNameIn)
	}

	if f.NickNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NickNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NickNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NickNamePrefix))
	}

	if f.NickNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NickNameSuffix))
	}

	if f.NickNameNull != nil {
		if *f.NickNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" IS NOT NULL")
		}
	}

	if f.Location != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" = ?")
		values = append(values, f.Location)
	}

	if f.LocationNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" != ?")
		values = append(values, f.LocationNe)
	}

	if f.LocationGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" > ?")
		values = append(values, f.LocationGt)
	}

	if f.LocationLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" < ?")
		values = append(values, f.LocationLt)
	}

	if f.LocationGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" >= ?")
		values = append(values, f.LocationGte)
	}

	if f.LocationLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" <= ?")
		values = append(values, f.LocationLte)
	}

	if f.LocationIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" IN (?)")
		values = append(values, f.LocationIn)
	}

	if f.LocationLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LocationLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LocationPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LocationPrefix))
	}

	if f.LocationSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LocationSuffix))
	}

	if f.LocationNull != nil {
		if *f.LocationNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *UserFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.ActiveMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("active")+") = ?")
		values = append(values, f.ActiveMin)
	}

	if f.ActiveMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("active")+") = ?")
		values = append(values, f.ActiveMax)
	}

	if f.ActiveMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("active")+") != ?")
		values = append(values, f.ActiveMinNe)
	}

	if f.ActiveMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("active")+") != ?")
		values = append(values, f.ActiveMaxNe)
	}

	if f.ActiveMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("active")+") > ?")
		values = append(values, f.ActiveMinGt)
	}

	if f.ActiveMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("active")+") > ?")
		values = append(values, f.ActiveMaxGt)
	}

	if f.ActiveMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("active")+") < ?")
		values = append(values, f.ActiveMinLt)
	}

	if f.ActiveMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("active")+") < ?")
		values = append(values, f.ActiveMaxLt)
	}

	if f.ActiveMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("active")+") >= ?")
		values = append(values, f.ActiveMinGte)
	}

	if f.ActiveMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("active")+") >= ?")
		values = append(values, f.ActiveMaxGte)
	}

	if f.ActiveMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("active")+") <= ?")
		values = append(values, f.ActiveMinLte)
	}

	if f.ActiveMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("active")+") <= ?")
		values = append(values, f.ActiveMaxLte)
	}

	if f.ActiveMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("active")+") IN (?)")
		values = append(values, f.ActiveMinIn)
	}

	if f.ActiveMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("active")+") IN (?)")
		values = append(values, f.ActiveMaxIn)
	}

	if f.EmailMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") = ?")
		values = append(values, f.EmailMin)
	}

	if f.EmailMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") = ?")
		values = append(values, f.EmailMax)
	}

	if f.EmailMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") != ?")
		values = append(values, f.EmailMinNe)
	}

	if f.EmailMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") != ?")
		values = append(values, f.EmailMaxNe)
	}

	if f.EmailMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") > ?")
		values = append(values, f.EmailMinGt)
	}

	if f.EmailMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") > ?")
		values = append(values, f.EmailMaxGt)
	}

	if f.EmailMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") < ?")
		values = append(values, f.EmailMinLt)
	}

	if f.EmailMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") < ?")
		values = append(values, f.EmailMaxLt)
	}

	if f.EmailMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") >= ?")
		values = append(values, f.EmailMinGte)
	}

	if f.EmailMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") >= ?")
		values = append(values, f.EmailMaxGte)
	}

	if f.EmailMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") <= ?")
		values = append(values, f.EmailMinLte)
	}

	if f.EmailMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") <= ?")
		values = append(values, f.EmailMaxLte)
	}

	if f.EmailMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") IN (?)")
		values = append(values, f.EmailMinIn)
	}

	if f.EmailMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") IN (?)")
		values = append(values, f.EmailMaxIn)
	}

	if f.EmailMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailMinPrefix))
	}

	if f.EmailMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailMaxPrefix))
	}

	if f.EmailMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailMinSuffix))
	}

	if f.EmailMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailMaxSuffix))
	}

	if f.PasswordMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("password")+") = ?")
		values = append(values, f.PasswordMin)
	}

	if f.PasswordMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("password")+") = ?")
		values = append(values, f.PasswordMax)
	}

	if f.PasswordMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("password")+") != ?")
		values = append(values, f.PasswordMinNe)
	}

	if f.PasswordMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("password")+") != ?")
		values = append(values, f.PasswordMaxNe)
	}

	if f.PasswordMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("password")+") > ?")
		values = append(values, f.PasswordMinGt)
	}

	if f.PasswordMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("password")+") > ?")
		values = append(values, f.PasswordMaxGt)
	}

	if f.PasswordMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("password")+") < ?")
		values = append(values, f.PasswordMinLt)
	}

	if f.PasswordMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("password")+") < ?")
		values = append(values, f.PasswordMaxLt)
	}

	if f.PasswordMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("password")+") >= ?")
		values = append(values, f.PasswordMinGte)
	}

	if f.PasswordMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("password")+") >= ?")
		values = append(values, f.PasswordMaxGte)
	}

	if f.PasswordMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("password")+") <= ?")
		values = append(values, f.PasswordMinLte)
	}

	if f.PasswordMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("password")+") <= ?")
		values = append(values, f.PasswordMaxLte)
	}

	if f.PasswordMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("password")+") IN (?)")
		values = append(values, f.PasswordMinIn)
	}

	if f.PasswordMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("password")+") IN (?)")
		values = append(values, f.PasswordMaxIn)
	}

	if f.PasswordMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("password")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PasswordMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PasswordMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("password")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PasswordMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PasswordMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("password")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PasswordMinPrefix))
	}

	if f.PasswordMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("password")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PasswordMaxPrefix))
	}

	if f.PasswordMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("password")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PasswordMinSuffix))
	}

	if f.PasswordMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("password")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PasswordMaxSuffix))
	}

	if f.AvatarURLMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") = ?")
		values = append(values, f.AvatarURLMin)
	}

	if f.AvatarURLMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") = ?")
		values = append(values, f.AvatarURLMax)
	}

	if f.AvatarURLMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") != ?")
		values = append(values, f.AvatarURLMinNe)
	}

	if f.AvatarURLMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") != ?")
		values = append(values, f.AvatarURLMaxNe)
	}

	if f.AvatarURLMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") > ?")
		values = append(values, f.AvatarURLMinGt)
	}

	if f.AvatarURLMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") > ?")
		values = append(values, f.AvatarURLMaxGt)
	}

	if f.AvatarURLMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") < ?")
		values = append(values, f.AvatarURLMinLt)
	}

	if f.AvatarURLMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") < ?")
		values = append(values, f.AvatarURLMaxLt)
	}

	if f.AvatarURLMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") >= ?")
		values = append(values, f.AvatarURLMinGte)
	}

	if f.AvatarURLMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") >= ?")
		values = append(values, f.AvatarURLMaxGte)
	}

	if f.AvatarURLMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") <= ?")
		values = append(values, f.AvatarURLMinLte)
	}

	if f.AvatarURLMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") <= ?")
		values = append(values, f.AvatarURLMaxLte)
	}

	if f.AvatarURLMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") IN (?)")
		values = append(values, f.AvatarURLMinIn)
	}

	if f.AvatarURLMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") IN (?)")
		values = append(values, f.AvatarURLMaxIn)
	}

	if f.AvatarURLMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AvatarURLMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AvatarURLMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AvatarURLMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AvatarURLMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AvatarURLMinPrefix))
	}

	if f.AvatarURLMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AvatarURLMaxPrefix))
	}

	if f.AvatarURLMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AvatarURLMinSuffix))
	}

	if f.AvatarURLMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AvatarURLMaxSuffix))
	}

	if f.DisplayNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") = ?")
		values = append(values, f.DisplayNameMin)
	}

	if f.DisplayNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") = ?")
		values = append(values, f.DisplayNameMax)
	}

	if f.DisplayNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") != ?")
		values = append(values, f.DisplayNameMinNe)
	}

	if f.DisplayNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") != ?")
		values = append(values, f.DisplayNameMaxNe)
	}

	if f.DisplayNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") > ?")
		values = append(values, f.DisplayNameMinGt)
	}

	if f.DisplayNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") > ?")
		values = append(values, f.DisplayNameMaxGt)
	}

	if f.DisplayNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") < ?")
		values = append(values, f.DisplayNameMinLt)
	}

	if f.DisplayNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") < ?")
		values = append(values, f.DisplayNameMaxLt)
	}

	if f.DisplayNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") >= ?")
		values = append(values, f.DisplayNameMinGte)
	}

	if f.DisplayNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") >= ?")
		values = append(values, f.DisplayNameMaxGte)
	}

	if f.DisplayNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") <= ?")
		values = append(values, f.DisplayNameMinLte)
	}

	if f.DisplayNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") <= ?")
		values = append(values, f.DisplayNameMaxLte)
	}

	if f.DisplayNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") IN (?)")
		values = append(values, f.DisplayNameMinIn)
	}

	if f.DisplayNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") IN (?)")
		values = append(values, f.DisplayNameMaxIn)
	}

	if f.DisplayNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DisplayNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DisplayNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DisplayNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DisplayNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DisplayNameMinPrefix))
	}

	if f.DisplayNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DisplayNameMaxPrefix))
	}

	if f.DisplayNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("displayName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DisplayNameMinSuffix))
	}

	if f.DisplayNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("displayName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DisplayNameMaxSuffix))
	}

	if f.FirstNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") = ?")
		values = append(values, f.FirstNameMin)
	}

	if f.FirstNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") = ?")
		values = append(values, f.FirstNameMax)
	}

	if f.FirstNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") != ?")
		values = append(values, f.FirstNameMinNe)
	}

	if f.FirstNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") != ?")
		values = append(values, f.FirstNameMaxNe)
	}

	if f.FirstNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") > ?")
		values = append(values, f.FirstNameMinGt)
	}

	if f.FirstNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") > ?")
		values = append(values, f.FirstNameMaxGt)
	}

	if f.FirstNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") < ?")
		values = append(values, f.FirstNameMinLt)
	}

	if f.FirstNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") < ?")
		values = append(values, f.FirstNameMaxLt)
	}

	if f.FirstNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") >= ?")
		values = append(values, f.FirstNameMinGte)
	}

	if f.FirstNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") >= ?")
		values = append(values, f.FirstNameMaxGte)
	}

	if f.FirstNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") <= ?")
		values = append(values, f.FirstNameMinLte)
	}

	if f.FirstNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") <= ?")
		values = append(values, f.FirstNameMaxLte)
	}

	if f.FirstNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") IN (?)")
		values = append(values, f.FirstNameMinIn)
	}

	if f.FirstNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") IN (?)")
		values = append(values, f.FirstNameMaxIn)
	}

	if f.FirstNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FirstNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FirstNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNameMinPrefix))
	}

	if f.FirstNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNameMaxPrefix))
	}

	if f.FirstNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameMinSuffix))
	}

	if f.FirstNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameMaxSuffix))
	}

	if f.LastNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") = ?")
		values = append(values, f.LastNameMin)
	}

	if f.LastNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") = ?")
		values = append(values, f.LastNameMax)
	}

	if f.LastNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") != ?")
		values = append(values, f.LastNameMinNe)
	}

	if f.LastNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") != ?")
		values = append(values, f.LastNameMaxNe)
	}

	if f.LastNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") > ?")
		values = append(values, f.LastNameMinGt)
	}

	if f.LastNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") > ?")
		values = append(values, f.LastNameMaxGt)
	}

	if f.LastNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") < ?")
		values = append(values, f.LastNameMinLt)
	}

	if f.LastNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") < ?")
		values = append(values, f.LastNameMaxLt)
	}

	if f.LastNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") >= ?")
		values = append(values, f.LastNameMinGte)
	}

	if f.LastNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") >= ?")
		values = append(values, f.LastNameMaxGte)
	}

	if f.LastNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") <= ?")
		values = append(values, f.LastNameMinLte)
	}

	if f.LastNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") <= ?")
		values = append(values, f.LastNameMaxLte)
	}

	if f.LastNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") IN (?)")
		values = append(values, f.LastNameMinIn)
	}

	if f.LastNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") IN (?)")
		values = append(values, f.LastNameMaxIn)
	}

	if f.LastNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LastNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LastNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNameMinPrefix))
	}

	if f.LastNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNameMaxPrefix))
	}

	if f.LastNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameMinSuffix))
	}

	if f.LastNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameMaxSuffix))
	}

	if f.NickNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") = ?")
		values = append(values, f.NickNameMin)
	}

	if f.NickNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") = ?")
		values = append(values, f.NickNameMax)
	}

	if f.NickNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") != ?")
		values = append(values, f.NickNameMinNe)
	}

	if f.NickNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") != ?")
		values = append(values, f.NickNameMaxNe)
	}

	if f.NickNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") > ?")
		values = append(values, f.NickNameMinGt)
	}

	if f.NickNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") > ?")
		values = append(values, f.NickNameMaxGt)
	}

	if f.NickNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") < ?")
		values = append(values, f.NickNameMinLt)
	}

	if f.NickNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") < ?")
		values = append(values, f.NickNameMaxLt)
	}

	if f.NickNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") >= ?")
		values = append(values, f.NickNameMinGte)
	}

	if f.NickNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") >= ?")
		values = append(values, f.NickNameMaxGte)
	}

	if f.NickNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") <= ?")
		values = append(values, f.NickNameMinLte)
	}

	if f.NickNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") <= ?")
		values = append(values, f.NickNameMaxLte)
	}

	if f.NickNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") IN (?)")
		values = append(values, f.NickNameMinIn)
	}

	if f.NickNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") IN (?)")
		values = append(values, f.NickNameMaxIn)
	}

	if f.NickNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NickNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NickNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NickNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NickNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NickNameMinPrefix))
	}

	if f.NickNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NickNameMaxPrefix))
	}

	if f.NickNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NickNameMinSuffix))
	}

	if f.NickNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NickNameMaxSuffix))
	}

	if f.LocationMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") = ?")
		values = append(values, f.LocationMin)
	}

	if f.LocationMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") = ?")
		values = append(values, f.LocationMax)
	}

	if f.LocationMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") != ?")
		values = append(values, f.LocationMinNe)
	}

	if f.LocationMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") != ?")
		values = append(values, f.LocationMaxNe)
	}

	if f.LocationMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") > ?")
		values = append(values, f.LocationMinGt)
	}

	if f.LocationMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") > ?")
		values = append(values, f.LocationMaxGt)
	}

	if f.LocationMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") < ?")
		values = append(values, f.LocationMinLt)
	}

	if f.LocationMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") < ?")
		values = append(values, f.LocationMaxLt)
	}

	if f.LocationMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") >= ?")
		values = append(values, f.LocationMinGte)
	}

	if f.LocationMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") >= ?")
		values = append(values, f.LocationMaxGte)
	}

	if f.LocationMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") <= ?")
		values = append(values, f.LocationMinLte)
	}

	if f.LocationMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") <= ?")
		values = append(values, f.LocationMaxLte)
	}

	if f.LocationMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") IN (?)")
		values = append(values, f.LocationMinIn)
	}

	if f.LocationMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") IN (?)")
		values = append(values, f.LocationMaxIn)
	}

	if f.LocationMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LocationMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LocationMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LocationMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LocationMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LocationMinPrefix))
	}

	if f.LocationMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LocationMaxPrefix))
	}

	if f.LocationMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LocationMinSuffix))
	}

	if f.LocationMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LocationMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *UserFilterType) AndWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *UserFilterType) OrWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *UserAPIKeyFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *UserAPIKeyFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("user_api_keys"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *UserAPIKeyFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.User != nil {
		_alias := alias + "_user"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("userId"))
		err := f.User.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Permissions != nil {
		_alias := alias + "_permissions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("userAPIKey_permissions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("apikeyId")+" LEFT JOIN "+dialect.Quote(TableName("permissions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Permissions.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *UserAPIKeyFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Key != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("key")+" = ?")
		values = append(values, f.Key)
	}

	if f.KeyNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("key")+" != ?")
		values = append(values, f.KeyNe)
	}

	if f.KeyGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("key")+" > ?")
		values = append(values, f.KeyGt)
	}

	if f.KeyLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("key")+" < ?")
		values = append(values, f.KeyLt)
	}

	if f.KeyGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("key")+" >= ?")
		values = append(values, f.KeyGte)
	}

	if f.KeyLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("key")+" <= ?")
		values = append(values, f.KeyLte)
	}

	if f.KeyIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("key")+" IN (?)")
		values = append(values, f.KeyIn)
	}

	if f.KeyLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("key")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.KeyLike, "?", "_", -1), "*", "%", -1))
	}

	if f.KeyPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("key")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.KeyPrefix))
	}

	if f.KeySuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("key")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.KeySuffix))
	}

	if f.KeyNull != nil {
		if *f.KeyNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("key")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("key")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.UserID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" = ?")
		values = append(values, f.UserID)
	}

	if f.UserIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" != ?")
		values = append(values, f.UserIDNe)
	}

	if f.UserIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" > ?")
		values = append(values, f.UserIDGt)
	}

	if f.UserIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" < ?")
		values = append(values, f.UserIDLt)
	}

	if f.UserIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" >= ?")
		values = append(values, f.UserIDGte)
	}

	if f.UserIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" <= ?")
		values = append(values, f.UserIDLte)
	}

	if f.UserIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" IN (?)")
		values = append(values, f.UserIDIn)
	}

	if f.UserIDNull != nil {
		if *f.UserIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *UserAPIKeyFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.KeyMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("key")+") = ?")
		values = append(values, f.KeyMin)
	}

	if f.KeyMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("key")+") = ?")
		values = append(values, f.KeyMax)
	}

	if f.KeyMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("key")+") != ?")
		values = append(values, f.KeyMinNe)
	}

	if f.KeyMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("key")+") != ?")
		values = append(values, f.KeyMaxNe)
	}

	if f.KeyMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("key")+") > ?")
		values = append(values, f.KeyMinGt)
	}

	if f.KeyMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("key")+") > ?")
		values = append(values, f.KeyMaxGt)
	}

	if f.KeyMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("key")+") < ?")
		values = append(values, f.KeyMinLt)
	}

	if f.KeyMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("key")+") < ?")
		values = append(values, f.KeyMaxLt)
	}

	if f.KeyMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("key")+") >= ?")
		values = append(values, f.KeyMinGte)
	}

	if f.KeyMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("key")+") >= ?")
		values = append(values, f.KeyMaxGte)
	}

	if f.KeyMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("key")+") <= ?")
		values = append(values, f.KeyMinLte)
	}

	if f.KeyMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("key")+") <= ?")
		values = append(values, f.KeyMaxLte)
	}

	if f.KeyMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("key")+") IN (?)")
		values = append(values, f.KeyMinIn)
	}

	if f.KeyMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("key")+") IN (?)")
		values = append(values, f.KeyMaxIn)
	}

	if f.KeyMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("key")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.KeyMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.KeyMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("key")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.KeyMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.KeyMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("key")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.KeyMinPrefix))
	}

	if f.KeyMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("key")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.KeyMaxPrefix))
	}

	if f.KeyMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("key")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.KeyMinSuffix))
	}

	if f.KeyMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("key")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.KeyMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.UserIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") = ?")
		values = append(values, f.UserIDMin)
	}

	if f.UserIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") = ?")
		values = append(values, f.UserIDMax)
	}

	if f.UserIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") != ?")
		values = append(values, f.UserIDMinNe)
	}

	if f.UserIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") != ?")
		values = append(values, f.UserIDMaxNe)
	}

	if f.UserIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") > ?")
		values = append(values, f.UserIDMinGt)
	}

	if f.UserIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") > ?")
		values = append(values, f.UserIDMaxGt)
	}

	if f.UserIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") < ?")
		values = append(values, f.UserIDMinLt)
	}

	if f.UserIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") < ?")
		values = append(values, f.UserIDMaxLt)
	}

	if f.UserIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") >= ?")
		values = append(values, f.UserIDMinGte)
	}

	if f.UserIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") >= ?")
		values = append(values, f.UserIDMaxGte)
	}

	if f.UserIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") <= ?")
		values = append(values, f.UserIDMinLte)
	}

	if f.UserIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") <= ?")
		values = append(values, f.UserIDMaxLte)
	}

	if f.UserIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") IN (?)")
		values = append(values, f.UserIDMinIn)
	}

	if f.UserIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") IN (?)")
		values = append(values, f.UserIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *UserAPIKeyFilterType) AndWith(f2 ...*UserAPIKeyFilterType) *UserAPIKeyFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserAPIKeyFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *UserAPIKeyFilterType) OrWith(f2 ...*UserAPIKeyFilterType) *UserAPIKeyFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserAPIKeyFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *ProfileFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *ProfileFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("profiles"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *ProfileFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Users != nil {
		_alias := alias + "_users"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("profile_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("profileId")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Users.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *ProfileFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Email != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" = ?")
		values = append(values, f.Email)
	}

	if f.EmailNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" != ?")
		values = append(values, f.EmailNe)
	}

	if f.EmailGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" > ?")
		values = append(values, f.EmailGt)
	}

	if f.EmailLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" < ?")
		values = append(values, f.EmailLt)
	}

	if f.EmailGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" >= ?")
		values = append(values, f.EmailGte)
	}

	if f.EmailLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" <= ?")
		values = append(values, f.EmailLte)
	}

	if f.EmailIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IN (?)")
		values = append(values, f.EmailIn)
	}

	if f.EmailLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailPrefix))
	}

	if f.EmailSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailSuffix))
	}

	if f.EmailNull != nil {
		if *f.EmailNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IS NOT NULL")
		}
	}

	if f.ExternalUserID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("externalUserId")+" = ?")
		values = append(values, f.ExternalUserID)
	}

	if f.ExternalUserIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("externalUserId")+" != ?")
		values = append(values, f.ExternalUserIDNe)
	}

	if f.ExternalUserIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("externalUserId")+" > ?")
		values = append(values, f.ExternalUserIDGt)
	}

	if f.ExternalUserIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("externalUserId")+" < ?")
		values = append(values, f.ExternalUserIDLt)
	}

	if f.ExternalUserIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("externalUserId")+" >= ?")
		values = append(values, f.ExternalUserIDGte)
	}

	if f.ExternalUserIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("externalUserId")+" <= ?")
		values = append(values, f.ExternalUserIDLte)
	}

	if f.ExternalUserIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("externalUserId")+" IN (?)")
		values = append(values, f.ExternalUserIDIn)
	}

	if f.ExternalUserIDLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("externalUserId")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ExternalUserIDLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ExternalUserIDPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("externalUserId")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ExternalUserIDPrefix))
	}

	if f.ExternalUserIDSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("externalUserId")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ExternalUserIDSuffix))
	}

	if f.ExternalUserIDNull != nil {
		if *f.ExternalUserIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("externalUserId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("externalUserId")+" IS NOT NULL")
		}
	}

	if f.Provider != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("provider")+" = ?")
		values = append(values, f.Provider)
	}

	if f.ProviderNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("provider")+" != ?")
		values = append(values, f.ProviderNe)
	}

	if f.ProviderGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("provider")+" > ?")
		values = append(values, f.ProviderGt)
	}

	if f.ProviderLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("provider")+" < ?")
		values = append(values, f.ProviderLt)
	}

	if f.ProviderGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("provider")+" >= ?")
		values = append(values, f.ProviderGte)
	}

	if f.ProviderLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("provider")+" <= ?")
		values = append(values, f.ProviderLte)
	}

	if f.ProviderIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("provider")+" IN (?)")
		values = append(values, f.ProviderIn)
	}

	if f.ProviderLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("provider")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ProviderLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ProviderPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("provider")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ProviderPrefix))
	}

	if f.ProviderSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("provider")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ProviderSuffix))
	}

	if f.ProviderNull != nil {
		if *f.ProviderNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("provider")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("provider")+" IS NOT NULL")
		}
	}

	if f.AvatarURL != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" = ?")
		values = append(values, f.AvatarURL)
	}

	if f.AvatarURLNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" != ?")
		values = append(values, f.AvatarURLNe)
	}

	if f.AvatarURLGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" > ?")
		values = append(values, f.AvatarURLGt)
	}

	if f.AvatarURLLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" < ?")
		values = append(values, f.AvatarURLLt)
	}

	if f.AvatarURLGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" >= ?")
		values = append(values, f.AvatarURLGte)
	}

	if f.AvatarURLLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" <= ?")
		values = append(values, f.AvatarURLLte)
	}

	if f.AvatarURLIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" IN (?)")
		values = append(values, f.AvatarURLIn)
	}

	if f.AvatarURLLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AvatarURLLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AvatarURLPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AvatarURLPrefix))
	}

	if f.AvatarURLSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AvatarURLSuffix))
	}

	if f.AvatarURLNull != nil {
		if *f.AvatarURLNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("avatarURL")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.FirstName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" = ?")
		values = append(values, f.FirstName)
	}

	if f.FirstNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" != ?")
		values = append(values, f.FirstNameNe)
	}

	if f.FirstNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" > ?")
		values = append(values, f.FirstNameGt)
	}

	if f.FirstNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" < ?")
		values = append(values, f.FirstNameLt)
	}

	if f.FirstNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" >= ?")
		values = append(values, f.FirstNameGte)
	}

	if f.FirstNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" <= ?")
		values = append(values, f.FirstNameLte)
	}

	if f.FirstNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IN (?)")
		values = append(values, f.FirstNameIn)
	}

	if f.FirstNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FirstNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNamePrefix))
	}

	if f.FirstNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameSuffix))
	}

	if f.FirstNameNull != nil {
		if *f.FirstNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IS NOT NULL")
		}
	}

	if f.LastName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" = ?")
		values = append(values, f.LastName)
	}

	if f.LastNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" != ?")
		values = append(values, f.LastNameNe)
	}

	if f.LastNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" > ?")
		values = append(values, f.LastNameGt)
	}

	if f.LastNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" < ?")
		values = append(values, f.LastNameLt)
	}

	if f.LastNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" >= ?")
		values = append(values, f.LastNameGte)
	}

	if f.LastNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" <= ?")
		values = append(values, f.LastNameLte)
	}

	if f.LastNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IN (?)")
		values = append(values, f.LastNameIn)
	}

	if f.LastNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LastNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNamePrefix))
	}

	if f.LastNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameSuffix))
	}

	if f.LastNameNull != nil {
		if *f.LastNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IS NOT NULL")
		}
	}

	if f.NickName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" = ?")
		values = append(values, f.NickName)
	}

	if f.NickNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" != ?")
		values = append(values, f.NickNameNe)
	}

	if f.NickNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" > ?")
		values = append(values, f.NickNameGt)
	}

	if f.NickNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" < ?")
		values = append(values, f.NickNameLt)
	}

	if f.NickNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" >= ?")
		values = append(values, f.NickNameGte)
	}

	if f.NickNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" <= ?")
		values = append(values, f.NickNameLte)
	}

	if f.NickNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" IN (?)")
		values = append(values, f.NickNameIn)
	}

	if f.NickNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NickNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NickNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NickNamePrefix))
	}

	if f.NickNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NickNameSuffix))
	}

	if f.NickNameNull != nil {
		if *f.NickNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("nickName")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.Location != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" = ?")
		values = append(values, f.Location)
	}

	if f.LocationNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" != ?")
		values = append(values, f.LocationNe)
	}

	if f.LocationGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" > ?")
		values = append(values, f.LocationGt)
	}

	if f.LocationLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" < ?")
		values = append(values, f.LocationLt)
	}

	if f.LocationGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" >= ?")
		values = append(values, f.LocationGte)
	}

	if f.LocationLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" <= ?")
		values = append(values, f.LocationLte)
	}

	if f.LocationIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" IN (?)")
		values = append(values, f.LocationIn)
	}

	if f.LocationLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LocationLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LocationPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LocationPrefix))
	}

	if f.LocationSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LocationSuffix))
	}

	if f.LocationNull != nil {
		if *f.LocationNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("location")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *ProfileFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.EmailMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") = ?")
		values = append(values, f.EmailMin)
	}

	if f.EmailMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") = ?")
		values = append(values, f.EmailMax)
	}

	if f.EmailMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") != ?")
		values = append(values, f.EmailMinNe)
	}

	if f.EmailMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") != ?")
		values = append(values, f.EmailMaxNe)
	}

	if f.EmailMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") > ?")
		values = append(values, f.EmailMinGt)
	}

	if f.EmailMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") > ?")
		values = append(values, f.EmailMaxGt)
	}

	if f.EmailMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") < ?")
		values = append(values, f.EmailMinLt)
	}

	if f.EmailMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") < ?")
		values = append(values, f.EmailMaxLt)
	}

	if f.EmailMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") >= ?")
		values = append(values, f.EmailMinGte)
	}

	if f.EmailMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") >= ?")
		values = append(values, f.EmailMaxGte)
	}

	if f.EmailMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") <= ?")
		values = append(values, f.EmailMinLte)
	}

	if f.EmailMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") <= ?")
		values = append(values, f.EmailMaxLte)
	}

	if f.EmailMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") IN (?)")
		values = append(values, f.EmailMinIn)
	}

	if f.EmailMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") IN (?)")
		values = append(values, f.EmailMaxIn)
	}

	if f.EmailMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailMinPrefix))
	}

	if f.EmailMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailMaxPrefix))
	}

	if f.EmailMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailMinSuffix))
	}

	if f.EmailMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailMaxSuffix))
	}

	if f.ExternalUserIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("externalUserId")+") = ?")
		values = append(values, f.ExternalUserIDMin)
	}

	if f.ExternalUserIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("externalUserId")+") = ?")
		values = append(values, f.ExternalUserIDMax)
	}

	if f.ExternalUserIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("externalUserId")+") != ?")
		values = append(values, f.ExternalUserIDMinNe)
	}

	if f.ExternalUserIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("externalUserId")+") != ?")
		values = append(values, f.ExternalUserIDMaxNe)
	}

	if f.ExternalUserIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("externalUserId")+") > ?")
		values = append(values, f.ExternalUserIDMinGt)
	}

	if f.ExternalUserIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("externalUserId")+") > ?")
		values = append(values, f.ExternalUserIDMaxGt)
	}

	if f.ExternalUserIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("externalUserId")+") < ?")
		values = append(values, f.ExternalUserIDMinLt)
	}

	if f.ExternalUserIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("externalUserId")+") < ?")
		values = append(values, f.ExternalUserIDMaxLt)
	}

	if f.ExternalUserIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("externalUserId")+") >= ?")
		values = append(values, f.ExternalUserIDMinGte)
	}

	if f.ExternalUserIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("externalUserId")+") >= ?")
		values = append(values, f.ExternalUserIDMaxGte)
	}

	if f.ExternalUserIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("externalUserId")+") <= ?")
		values = append(values, f.ExternalUserIDMinLte)
	}

	if f.ExternalUserIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("externalUserId")+") <= ?")
		values = append(values, f.ExternalUserIDMaxLte)
	}

	if f.ExternalUserIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("externalUserId")+") IN (?)")
		values = append(values, f.ExternalUserIDMinIn)
	}

	if f.ExternalUserIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("externalUserId")+") IN (?)")
		values = append(values, f.ExternalUserIDMaxIn)
	}

	if f.ExternalUserIDMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("externalUserId")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ExternalUserIDMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ExternalUserIDMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("externalUserId")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ExternalUserIDMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ExternalUserIDMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("externalUserId")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ExternalUserIDMinPrefix))
	}

	if f.ExternalUserIDMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("externalUserId")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ExternalUserIDMaxPrefix))
	}

	if f.ExternalUserIDMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("externalUserId")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ExternalUserIDMinSuffix))
	}

	if f.ExternalUserIDMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("externalUserId")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ExternalUserIDMaxSuffix))
	}

	if f.ProviderMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("provider")+") = ?")
		values = append(values, f.ProviderMin)
	}

	if f.ProviderMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("provider")+") = ?")
		values = append(values, f.ProviderMax)
	}

	if f.ProviderMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("provider")+") != ?")
		values = append(values, f.ProviderMinNe)
	}

	if f.ProviderMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("provider")+") != ?")
		values = append(values, f.ProviderMaxNe)
	}

	if f.ProviderMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("provider")+") > ?")
		values = append(values, f.ProviderMinGt)
	}

	if f.ProviderMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("provider")+") > ?")
		values = append(values, f.ProviderMaxGt)
	}

	if f.ProviderMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("provider")+") < ?")
		values = append(values, f.ProviderMinLt)
	}

	if f.ProviderMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("provider")+") < ?")
		values = append(values, f.ProviderMaxLt)
	}

	if f.ProviderMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("provider")+") >= ?")
		values = append(values, f.ProviderMinGte)
	}

	if f.ProviderMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("provider")+") >= ?")
		values = append(values, f.ProviderMaxGte)
	}

	if f.ProviderMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("provider")+") <= ?")
		values = append(values, f.ProviderMinLte)
	}

	if f.ProviderMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("provider")+") <= ?")
		values = append(values, f.ProviderMaxLte)
	}

	if f.ProviderMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("provider")+") IN (?)")
		values = append(values, f.ProviderMinIn)
	}

	if f.ProviderMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("provider")+") IN (?)")
		values = append(values, f.ProviderMaxIn)
	}

	if f.ProviderMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("provider")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ProviderMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ProviderMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("provider")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ProviderMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ProviderMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("provider")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ProviderMinPrefix))
	}

	if f.ProviderMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("provider")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ProviderMaxPrefix))
	}

	if f.ProviderMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("provider")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ProviderMinSuffix))
	}

	if f.ProviderMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("provider")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ProviderMaxSuffix))
	}

	if f.AvatarURLMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") = ?")
		values = append(values, f.AvatarURLMin)
	}

	if f.AvatarURLMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") = ?")
		values = append(values, f.AvatarURLMax)
	}

	if f.AvatarURLMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") != ?")
		values = append(values, f.AvatarURLMinNe)
	}

	if f.AvatarURLMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") != ?")
		values = append(values, f.AvatarURLMaxNe)
	}

	if f.AvatarURLMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") > ?")
		values = append(values, f.AvatarURLMinGt)
	}

	if f.AvatarURLMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") > ?")
		values = append(values, f.AvatarURLMaxGt)
	}

	if f.AvatarURLMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") < ?")
		values = append(values, f.AvatarURLMinLt)
	}

	if f.AvatarURLMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") < ?")
		values = append(values, f.AvatarURLMaxLt)
	}

	if f.AvatarURLMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") >= ?")
		values = append(values, f.AvatarURLMinGte)
	}

	if f.AvatarURLMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") >= ?")
		values = append(values, f.AvatarURLMaxGte)
	}

	if f.AvatarURLMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") <= ?")
		values = append(values, f.AvatarURLMinLte)
	}

	if f.AvatarURLMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") <= ?")
		values = append(values, f.AvatarURLMaxLte)
	}

	if f.AvatarURLMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") IN (?)")
		values = append(values, f.AvatarURLMinIn)
	}

	if f.AvatarURLMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") IN (?)")
		values = append(values, f.AvatarURLMaxIn)
	}

	if f.AvatarURLMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AvatarURLMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AvatarURLMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AvatarURLMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AvatarURLMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AvatarURLMinPrefix))
	}

	if f.AvatarURLMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AvatarURLMaxPrefix))
	}

	if f.AvatarURLMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AvatarURLMinSuffix))
	}

	if f.AvatarURLMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("avatarURL")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AvatarURLMaxSuffix))
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.FirstNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") = ?")
		values = append(values, f.FirstNameMin)
	}

	if f.FirstNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") = ?")
		values = append(values, f.FirstNameMax)
	}

	if f.FirstNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") != ?")
		values = append(values, f.FirstNameMinNe)
	}

	if f.FirstNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") != ?")
		values = append(values, f.FirstNameMaxNe)
	}

	if f.FirstNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") > ?")
		values = append(values, f.FirstNameMinGt)
	}

	if f.FirstNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") > ?")
		values = append(values, f.FirstNameMaxGt)
	}

	if f.FirstNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") < ?")
		values = append(values, f.FirstNameMinLt)
	}

	if f.FirstNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") < ?")
		values = append(values, f.FirstNameMaxLt)
	}

	if f.FirstNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") >= ?")
		values = append(values, f.FirstNameMinGte)
	}

	if f.FirstNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") >= ?")
		values = append(values, f.FirstNameMaxGte)
	}

	if f.FirstNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") <= ?")
		values = append(values, f.FirstNameMinLte)
	}

	if f.FirstNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") <= ?")
		values = append(values, f.FirstNameMaxLte)
	}

	if f.FirstNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") IN (?)")
		values = append(values, f.FirstNameMinIn)
	}

	if f.FirstNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") IN (?)")
		values = append(values, f.FirstNameMaxIn)
	}

	if f.FirstNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FirstNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FirstNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNameMinPrefix))
	}

	if f.FirstNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNameMaxPrefix))
	}

	if f.FirstNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameMinSuffix))
	}

	if f.FirstNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameMaxSuffix))
	}

	if f.LastNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") = ?")
		values = append(values, f.LastNameMin)
	}

	if f.LastNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") = ?")
		values = append(values, f.LastNameMax)
	}

	if f.LastNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") != ?")
		values = append(values, f.LastNameMinNe)
	}

	if f.LastNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") != ?")
		values = append(values, f.LastNameMaxNe)
	}

	if f.LastNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") > ?")
		values = append(values, f.LastNameMinGt)
	}

	if f.LastNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") > ?")
		values = append(values, f.LastNameMaxGt)
	}

	if f.LastNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") < ?")
		values = append(values, f.LastNameMinLt)
	}

	if f.LastNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") < ?")
		values = append(values, f.LastNameMaxLt)
	}

	if f.LastNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") >= ?")
		values = append(values, f.LastNameMinGte)
	}

	if f.LastNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") >= ?")
		values = append(values, f.LastNameMaxGte)
	}

	if f.LastNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") <= ?")
		values = append(values, f.LastNameMinLte)
	}

	if f.LastNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") <= ?")
		values = append(values, f.LastNameMaxLte)
	}

	if f.LastNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") IN (?)")
		values = append(values, f.LastNameMinIn)
	}

	if f.LastNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") IN (?)")
		values = append(values, f.LastNameMaxIn)
	}

	if f.LastNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LastNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LastNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNameMinPrefix))
	}

	if f.LastNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNameMaxPrefix))
	}

	if f.LastNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameMinSuffix))
	}

	if f.LastNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameMaxSuffix))
	}

	if f.NickNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") = ?")
		values = append(values, f.NickNameMin)
	}

	if f.NickNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") = ?")
		values = append(values, f.NickNameMax)
	}

	if f.NickNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") != ?")
		values = append(values, f.NickNameMinNe)
	}

	if f.NickNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") != ?")
		values = append(values, f.NickNameMaxNe)
	}

	if f.NickNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") > ?")
		values = append(values, f.NickNameMinGt)
	}

	if f.NickNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") > ?")
		values = append(values, f.NickNameMaxGt)
	}

	if f.NickNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") < ?")
		values = append(values, f.NickNameMinLt)
	}

	if f.NickNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") < ?")
		values = append(values, f.NickNameMaxLt)
	}

	if f.NickNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") >= ?")
		values = append(values, f.NickNameMinGte)
	}

	if f.NickNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") >= ?")
		values = append(values, f.NickNameMaxGte)
	}

	if f.NickNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") <= ?")
		values = append(values, f.NickNameMinLte)
	}

	if f.NickNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") <= ?")
		values = append(values, f.NickNameMaxLte)
	}

	if f.NickNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") IN (?)")
		values = append(values, f.NickNameMinIn)
	}

	if f.NickNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") IN (?)")
		values = append(values, f.NickNameMaxIn)
	}

	if f.NickNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NickNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NickNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NickNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NickNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NickNameMinPrefix))
	}

	if f.NickNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NickNameMaxPrefix))
	}

	if f.NickNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NickNameMinSuffix))
	}

	if f.NickNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("nickName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NickNameMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.LocationMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") = ?")
		values = append(values, f.LocationMin)
	}

	if f.LocationMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") = ?")
		values = append(values, f.LocationMax)
	}

	if f.LocationMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") != ?")
		values = append(values, f.LocationMinNe)
	}

	if f.LocationMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") != ?")
		values = append(values, f.LocationMaxNe)
	}

	if f.LocationMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") > ?")
		values = append(values, f.LocationMinGt)
	}

	if f.LocationMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") > ?")
		values = append(values, f.LocationMaxGt)
	}

	if f.LocationMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") < ?")
		values = append(values, f.LocationMinLt)
	}

	if f.LocationMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") < ?")
		values = append(values, f.LocationMaxLt)
	}

	if f.LocationMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") >= ?")
		values = append(values, f.LocationMinGte)
	}

	if f.LocationMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") >= ?")
		values = append(values, f.LocationMaxGte)
	}

	if f.LocationMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") <= ?")
		values = append(values, f.LocationMinLte)
	}

	if f.LocationMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") <= ?")
		values = append(values, f.LocationMaxLte)
	}

	if f.LocationMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") IN (?)")
		values = append(values, f.LocationMinIn)
	}

	if f.LocationMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") IN (?)")
		values = append(values, f.LocationMaxIn)
	}

	if f.LocationMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LocationMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LocationMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LocationMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LocationMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LocationMinPrefix))
	}

	if f.LocationMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LocationMaxPrefix))
	}

	if f.LocationMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LocationMinSuffix))
	}

	if f.LocationMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("location")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LocationMaxSuffix))
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *ProfileFilterType) AndWith(f2 ...*ProfileFilterType) *ProfileFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ProfileFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ProfileFilterType) OrWith(f2 ...*ProfileFilterType) *ProfileFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ProfileFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *RoleFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *RoleFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("roles"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *RoleFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Users != nil {
		_alias := alias + "_users"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("role_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("roleId")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Users.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Parents != nil {
		_alias := alias + "_parents"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("role_parents"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("childId")+" LEFT JOIN "+dialect.Quote(TableName("roles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("parentId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Parents.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Children != nil {
		_alias := alias + "_children"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("role_parents"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("parentId")+" LEFT JOIN "+dialect.Quote(TableName("roles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("childId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Children.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Permissions != nil {
		_alias := alias + "_permissions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("permission_roles"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("roleId")+" LEFT JOIN "+dialect.Quote(TableName("permissions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Permissions.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *RoleFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Domain != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" = ?")
		values = append(values, f.Domain)
	}

	if f.DomainNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" != ?")
		values = append(values, f.DomainNe)
	}

	if f.DomainGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" > ?")
		values = append(values, f.DomainGt)
	}

	if f.DomainLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" < ?")
		values = append(values, f.DomainLt)
	}

	if f.DomainGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" >= ?")
		values = append(values, f.DomainGte)
	}

	if f.DomainLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" <= ?")
		values = append(values, f.DomainLte)
	}

	if f.DomainIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" IN (?)")
		values = append(values, f.DomainIn)
	}

	if f.DomainLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DomainLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DomainPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DomainPrefix))
	}

	if f.DomainSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DomainSuffix))
	}

	if f.DomainNull != nil {
		if *f.DomainNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *RoleFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.DomainMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") = ?")
		values = append(values, f.DomainMin)
	}

	if f.DomainMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") = ?")
		values = append(values, f.DomainMax)
	}

	if f.DomainMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") != ?")
		values = append(values, f.DomainMinNe)
	}

	if f.DomainMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") != ?")
		values = append(values, f.DomainMaxNe)
	}

	if f.DomainMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") > ?")
		values = append(values, f.DomainMinGt)
	}

	if f.DomainMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") > ?")
		values = append(values, f.DomainMaxGt)
	}

	if f.DomainMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") < ?")
		values = append(values, f.DomainMinLt)
	}

	if f.DomainMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") < ?")
		values = append(values, f.DomainMaxLt)
	}

	if f.DomainMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") >= ?")
		values = append(values, f.DomainMinGte)
	}

	if f.DomainMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") >= ?")
		values = append(values, f.DomainMaxGte)
	}

	if f.DomainMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") <= ?")
		values = append(values, f.DomainMinLte)
	}

	if f.DomainMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") <= ?")
		values = append(values, f.DomainMaxLte)
	}

	if f.DomainMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") IN (?)")
		values = append(values, f.DomainMinIn)
	}

	if f.DomainMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") IN (?)")
		values = append(values, f.DomainMaxIn)
	}

	if f.DomainMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DomainMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DomainMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DomainMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DomainMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DomainMinPrefix))
	}

	if f.DomainMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DomainMaxPrefix))
	}

	if f.DomainMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DomainMinSuffix))
	}

	if f.DomainMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DomainMaxSuffix))
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *RoleFilterType) AndWith(f2 ...*RoleFilterType) *RoleFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &RoleFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *RoleFilterType) OrWith(f2 ...*RoleFilterType) *RoleFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &RoleFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *PermissionFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *PermissionFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("permissions"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *PermissionFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Users != nil {
		_alias := alias + "_users"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("permission_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Users.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Roles != nil {
		_alias := alias + "_roles"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("permission_roles"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" LEFT JOIN "+dialect.Quote(TableName("roles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("roleId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Roles.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Apikeys != nil {
		_alias := alias + "_apikeys"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("userAPIKey_permissions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" LEFT JOIN "+dialect.Quote(TableName("user_api_keys"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("apikeyId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Apikeys.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *PermissionFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Domain != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" = ?")
		values = append(values, f.Domain)
	}

	if f.DomainNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" != ?")
		values = append(values, f.DomainNe)
	}

	if f.DomainGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" > ?")
		values = append(values, f.DomainGt)
	}

	if f.DomainLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" < ?")
		values = append(values, f.DomainLt)
	}

	if f.DomainGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" >= ?")
		values = append(values, f.DomainGte)
	}

	if f.DomainLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" <= ?")
		values = append(values, f.DomainLte)
	}

	if f.DomainIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" IN (?)")
		values = append(values, f.DomainIn)
	}

	if f.DomainLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DomainLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DomainPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DomainPrefix))
	}

	if f.DomainSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DomainSuffix))
	}

	if f.DomainNull != nil {
		if *f.DomainNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("domain")+" IS NOT NULL")
		}
	}

	if f.Tag != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("tag")+" = ?")
		values = append(values, f.Tag)
	}

	if f.TagNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("tag")+" != ?")
		values = append(values, f.TagNe)
	}

	if f.TagGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("tag")+" > ?")
		values = append(values, f.TagGt)
	}

	if f.TagLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("tag")+" < ?")
		values = append(values, f.TagLt)
	}

	if f.TagGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("tag")+" >= ?")
		values = append(values, f.TagGte)
	}

	if f.TagLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("tag")+" <= ?")
		values = append(values, f.TagLte)
	}

	if f.TagIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("tag")+" IN (?)")
		values = append(values, f.TagIn)
	}

	if f.TagLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("tag")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TagLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TagPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("tag")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TagPrefix))
	}

	if f.TagSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("tag")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TagSuffix))
	}

	if f.TagNull != nil {
		if *f.TagNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("tag")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("tag")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *PermissionFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.DomainMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") = ?")
		values = append(values, f.DomainMin)
	}

	if f.DomainMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") = ?")
		values = append(values, f.DomainMax)
	}

	if f.DomainMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") != ?")
		values = append(values, f.DomainMinNe)
	}

	if f.DomainMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") != ?")
		values = append(values, f.DomainMaxNe)
	}

	if f.DomainMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") > ?")
		values = append(values, f.DomainMinGt)
	}

	if f.DomainMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") > ?")
		values = append(values, f.DomainMaxGt)
	}

	if f.DomainMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") < ?")
		values = append(values, f.DomainMinLt)
	}

	if f.DomainMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") < ?")
		values = append(values, f.DomainMaxLt)
	}

	if f.DomainMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") >= ?")
		values = append(values, f.DomainMinGte)
	}

	if f.DomainMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") >= ?")
		values = append(values, f.DomainMaxGte)
	}

	if f.DomainMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") <= ?")
		values = append(values, f.DomainMinLte)
	}

	if f.DomainMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") <= ?")
		values = append(values, f.DomainMaxLte)
	}

	if f.DomainMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") IN (?)")
		values = append(values, f.DomainMinIn)
	}

	if f.DomainMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") IN (?)")
		values = append(values, f.DomainMaxIn)
	}

	if f.DomainMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DomainMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DomainMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DomainMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DomainMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DomainMinPrefix))
	}

	if f.DomainMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DomainMaxPrefix))
	}

	if f.DomainMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("domain")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DomainMinSuffix))
	}

	if f.DomainMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("domain")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DomainMaxSuffix))
	}

	if f.TagMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("tag")+") = ?")
		values = append(values, f.TagMin)
	}

	if f.TagMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("tag")+") = ?")
		values = append(values, f.TagMax)
	}

	if f.TagMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("tag")+") != ?")
		values = append(values, f.TagMinNe)
	}

	if f.TagMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("tag")+") != ?")
		values = append(values, f.TagMaxNe)
	}

	if f.TagMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("tag")+") > ?")
		values = append(values, f.TagMinGt)
	}

	if f.TagMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("tag")+") > ?")
		values = append(values, f.TagMaxGt)
	}

	if f.TagMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("tag")+") < ?")
		values = append(values, f.TagMinLt)
	}

	if f.TagMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("tag")+") < ?")
		values = append(values, f.TagMaxLt)
	}

	if f.TagMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("tag")+") >= ?")
		values = append(values, f.TagMinGte)
	}

	if f.TagMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("tag")+") >= ?")
		values = append(values, f.TagMaxGte)
	}

	if f.TagMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("tag")+") <= ?")
		values = append(values, f.TagMinLte)
	}

	if f.TagMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("tag")+") <= ?")
		values = append(values, f.TagMaxLte)
	}

	if f.TagMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("tag")+") IN (?)")
		values = append(values, f.TagMinIn)
	}

	if f.TagMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("tag")+") IN (?)")
		values = append(values, f.TagMaxIn)
	}

	if f.TagMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("tag")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TagMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TagMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("tag")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TagMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TagMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("tag")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TagMinPrefix))
	}

	if f.TagMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("tag")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TagMaxPrefix))
	}

	if f.TagMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("tag")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TagMinSuffix))
	}

	if f.TagMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("tag")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TagMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *PermissionFilterType) AndWith(f2 ...*PermissionFilterType) *PermissionFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &PermissionFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *PermissionFilterType) OrWith(f2 ...*PermissionFilterType) *PermissionFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &PermissionFilterType{
		Or: append(_f2, f),
	}
}
