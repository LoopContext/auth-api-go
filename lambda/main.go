package main

import (
	"github.com/akrylysov/algnhsa"
	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src"
	"github.com/loopcontext/graphql-orm/events"
)

func main() {
	db := gen.NewDBFromEnvVars()

	eventController, err := events.NewEventController()
	if err != nil {
		panic(err)
	}

	handler := gen.GetHTTPServeMux(src.New(db, &eventController), db, src.GetMigrations(db))
	algnhsa.ListenAndServe(handler, &algnhsa.Options{
		UseProxyPath: true,
	})
}
