module github.com/loopcontext/auth-api-go

go 1.14

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/akrylysov/algnhsa v0.12.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gorilla/mux v1.7.4
	github.com/graph-gophers/dataloader v5.0.0+incompatible
	github.com/jinzhu/gorm v1.9.15
	github.com/loopcontext/checkmail v0.0.1
	github.com/loopcontext/graphql-orm v0.0.0-20200818223329-bac92a3d17ae
	github.com/markbates/goth v1.64.2
	github.com/mitchellh/mapstructure v1.3.3
	github.com/rs/cors v1.7.0
	github.com/rs/zerolog v1.19.0
	github.com/urfave/cli v1.22.4
	github.com/vektah/gqlparser/v2 v2.0.1
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de
	gopkg.in/gormigrate.v1 v1.6.0
)

// replace github.com/loopcontext/graphql-orm master => ../graphql-orm
