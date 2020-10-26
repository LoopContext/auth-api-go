module github.com/loopcontext/auth-api-go

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/akrylysov/algnhsa v0.12.1
	github.com/cloudevents/sdk-go v1.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/graph-gophers/dataloader v5.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/loopcontext/checkmail v0.0.1
	github.com/loopcontext/cloudevents-aws-transport v1.0.9
	github.com/loopcontext/go-graphql-orm v0.0.0-20201026025740-27c64b1123fb
	github.com/markbates/goth v1.64.2
	github.com/mitchellh/mapstructure v1.3.3
	github.com/rs/cors v1.7.0
	github.com/rs/zerolog v1.20.0
	github.com/urfave/cli v1.22.4
	github.com/vektah/gqlparser/v2 v2.1.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	gopkg.in/gormigrate.v1 v1.6.0
)

replace github.com/loopcontext/go-graphql-orm v0.0.0-20201026025740-27c64b1123fb => ../go-graphql-orm
