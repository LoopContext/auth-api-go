package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/vektah/gqlparser/v2/ast"
)

// UserQueryFilter struct
type UserQueryFilter struct {
	Query *string
}

// Apply ...
func (qf *UserQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("users"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *UserQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["email"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("email")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["password"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("password")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["avatarURL"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("avatarURL")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["displayName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("displayName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["description"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("description")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["firstName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("firstName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["lastName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("lastName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["nickName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("nickName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["location"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("location")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["apikeys"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_apikeys"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("user_api_keys"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("userId")+" = "+dialect.Quote(alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := UserAPIKeyQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["roles"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_roles"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("role_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" LEFT JOIN "+dialect.Quote(TableName("roles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("roleId")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := RoleQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["profiles"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_profiles"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("profile_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" LEFT JOIN "+dialect.Quote(TableName("profiles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("profileId")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := ProfileQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["permissions"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_permissions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("permission_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" LEFT JOIN "+dialect.Quote(TableName("permissions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := PermissionQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// UserAPIKeyQueryFilter struct
type UserAPIKeyQueryFilter struct {
	Query *string
}

// Apply ...
func (qf *UserAPIKeyQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("user_api_keys"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *UserAPIKeyQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["key"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("key")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["description"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("description")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["user"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_user"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("userId"))

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := UserQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["permissions"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_permissions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("userAPIKey_permissions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("apikeyId")+" LEFT JOIN "+dialect.Quote(TableName("permissions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := PermissionQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// ProfileQueryFilter struct
type ProfileQueryFilter struct {
	Query *string
}

// Apply ...
func (qf *ProfileQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("profiles"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *ProfileQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["email"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("email")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["externalUserId"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("externalUserId")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["provider"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("provider")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["avatarURL"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("avatarURL")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["name"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("name")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["firstName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("firstName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["lastName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("lastName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["nickName"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("nickName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["description"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("description")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["location"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("location")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["users"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_users"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("profile_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("profileId")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := UserQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// RoleQueryFilter struct
type RoleQueryFilter struct {
	Query *string
}

// Apply ...
func (qf *RoleQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("roles"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *RoleQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["domain"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("domain")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["name"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("name")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["description"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("description")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["users"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_users"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("role_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("roleId")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := UserQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["parents"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_parents"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("role_parents"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("childId")+" LEFT JOIN "+dialect.Quote(TableName("roles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("parentId")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := RoleQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["children"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_children"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("role_parents"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("parentId")+" LEFT JOIN "+dialect.Quote(TableName("roles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("childId")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := RoleQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["permissions"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_permissions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("permission_roles"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("roleId")+" LEFT JOIN "+dialect.Quote(TableName("permissions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := PermissionQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// PermissionQueryFilter struct
type PermissionQueryFilter struct {
	Query *string
}

// Apply ...
func (qf *PermissionQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("permissions"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *PermissionQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["domain"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("domain")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["tag"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("tag")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["description"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("description")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if fs, ok := fieldsMap["users"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_users"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("permission_users"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("userId")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := UserQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["roles"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_roles"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("permission_roles"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" LEFT JOIN "+dialect.Quote(TableName("roles"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("roleId")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := RoleQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	if fs, ok := fieldsMap["apikeys"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_apikeys"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("userAPIKey_permissions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("permissionId")+" LEFT JOIN "+dialect.Quote(TableName("user_api_keys"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("apikeyId")+" = "+dialect.Quote(_alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := UserAPIKeyQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
