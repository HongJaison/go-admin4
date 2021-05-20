package example

import (
	"github.com/HongJaison/go-admin4/context"
	"github.com/HongJaison/go-admin4/modules/auth"
	"github.com/HongJaison/go-admin4/modules/db"
	"github.com/HongJaison/go-admin4/modules/service"
)

func (e *Example) initRouter(prefix string, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/example", auth.Middleware(db.GetConnection(srv)), e.TestHandler)

	return app
}
