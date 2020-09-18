package main

import (
	"github.com/akrylysov/algnhsa"
	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src"
)

func main() {
	db := gen.NewDBFromEnvVars()

	eventController, err := gen.NewEventController()
	if err != nil {
		panic(err)
	}

	handler := gen.GetHTTPServeMux(src.New(db, &eventController), db, src.GetMigrations(db))
	algnhsa.ListenAndServe(handler, &algnhsa.Options{
		UseProxyPath: true,
	})
}
