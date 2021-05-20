package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/HongJaison/go-admin4/adapter/beego"
	_ "github.com/HongJaison/go-admin4/modules/db/drivers/mssql"

	"github.com/HongJaison/go-admin4/engine"
	"github.com/HongJaison/go-admin4/examples/datamodel"
	"github.com/HongJaison/go-admin4/modules/config"
	"github.com/HongJaison/go-admin4/modules/language"
	"github.com/HongJaison/go-admin4/plugins/example"
	"github.com/HongJaison/go-admin4/template"
	"github.com/HongJaison/go-admin4/template/chartjs"
	"github.com/HongJaison/themes4/adminlte"
	"github.com/astaxie/beego"
)

func main() {
	app := beego.NewApp()

	eng := engine.Default()

	cfg := config.Config{
		Env: config.EnvLocal,
		Databases: config.DatabaseList{
			"default": {
				Host: "172.16.74.222",
				Name: "Ninja",
				// Host:       "127.0.0.1",
				// Name:       "Casino",
				Port:       "1433",
				User:       "sa",
				Pwd:        "Akduifwkro1988",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     config.DriverMssql,
			},
		},
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		UrlPrefix:   "admin",
		IndexUrl:    "/",
		Debug:       true,
		Language:    language.EN,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}

	template.AddComp(chartjs.NewChart())

	// customize a plugin

	examplePlugin := example.NewExample()

	// load from golang.Plugin
	//
	// examplePlugin := plugins.LoadFromPlugin("../datamodel/example.so")

	// customize the login page
	// example: https://github.com/HongJaison/demo.go-admin.cn/blob/master/main.go#L39
	//
	// template.AddComp("login", datamodel.LoginPage)

	// load config from json file
	//
	// eng.AddConfigFromJSON("../datamodel/config.json")

	beego.SetStaticPath("/uploads", "uploads")

	if err := eng.AddConfig(&cfg).
		AddGenerators(datamodel.Generators).
		AddDisplayFilterXssJsFilter().
		// add generator, first parameter is the url prefix of table when visit.
		// example:
		//
		// "user" => http://localhost:9033/admin/info/user
		//
		AddGenerator("user", datamodel.GetUserTable).
		AddPlugins(examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	// you can custom your pages like:

	eng.HTML("GET", "/admin", datamodel.GetContent)

	beego.BConfig.Listen.HTTPAddr = ""
	// beego.BConfig.Listen.HTTPPort = 9087
	beego.BConfig.Listen.HTTPPort = 13301
	go app.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}
