package main

import (
	"atlas/ent"
	"atlas/routes"
	"atlas/service/impl"
	"database/sql"
	"io"
	"os"
	"text/template"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func main() {
	e := echo.New()
	client, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static Routes
	e.Static("/static", "static")

	//Templates
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	//Others
	p := impl.InitPaymentController(client)
	r := impl.InitRouterController()
	// Routes
	routes.InitRoutes(e, p, r)
	// e.GET("/something", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "qr.html", map[string]interface{}{
	// 		"qrUrl": "https://www.cashfree.com",
	// 	})
	// })

	// Start server
	e.Logger.Fatal(e.Start(":9811"))
}

func ConnectDB() (*ent.Client, error) {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
