package gorilla

import (
	// add gorilla adapter
	_ "github.com/HongJaison/go-admin4/adapter/gorilla"
	"github.com/HongJaison/go-admin4/modules/config"
	"github.com/HongJaison/go-admin4/modules/language"
	"github.com/HongJaison/go-admin4/plugins/admin/modules/table"
	"github.com/HongJaison/themes4/adminlte"

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

	"net/http"
	"os"

	"github.com/HongJaison/go-admin4/engine"
	"github.com/HongJaison/go-admin4/plugins/admin"
	"github.com/HongJaison/go-admin4/plugins/example"
	"github.com/HongJaison/go-admin4/template"
	"github.com/HongJaison/go-admin4/template/chartjs"
	"github.com/HongJaison/go-admin4/tests/tables"
	"github.com/gorilla/mux"
)

func newHandler() http.Handler {
	app := mux.NewRouter()
	eng := engine.Default()

	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(admin.NewAdmin(tables.Generators).
			AddGenerator("user", tables.GetUserTable), examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	app := mux.NewRouter()
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
		AddPlugins(admin.NewAdmin(gens)).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}
