package gofiber

import (
	// add fasthttp adapter
	ada "github.com/HongJaison/go-admin4/adapter/gofiber"
	// add mysql driver
	_ "github.com/HongJaison/go-admin4/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/HongJaison/go-admin4/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/HongJaison/go-admin4/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/HongJaison/go-admin4/modules/db/drivers/mssql"
	// add adminlte ui theme
	_ "github.com/HongJaison/themes4/adminlte"

	"os"

	"github.com/HongJaison/go-admin4/engine"
	"github.com/HongJaison/go-admin4/modules/config"
	"github.com/HongJaison/go-admin4/modules/language"
	"github.com/HongJaison/go-admin4/plugins/admin"
	"github.com/HongJaison/go-admin4/plugins/admin/modules/table"
	"github.com/HongJaison/go-admin4/template"
	"github.com/HongJaison/go-admin4/template/chartjs"
	"github.com/HongJaison/go-admin4/tests/tables"
	"github.com/HongJaison/themes4/adminlte"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func newHandler() fasthttp.RequestHandler {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
	})

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators).AddDisplayFilterXssJsFilter()
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app.Handler()
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) fasthttp.RequestHandler {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
	})

	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfig(&config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddAdapter(new(ada.Gofiber)).
		AddGenerators(gens).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app.Handler()
}
