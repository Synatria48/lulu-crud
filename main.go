// main.go
package main

import (
	"io"
	"net/http"
	"path/filepath"
	"text/template"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	crud "github.com/lulu-crud/CRUD"
)

func init() {
	crud.InitDB()
}

// Render HTML
func renderHTML(c echo.Context, code int, name string, data interface{}) error {
	return c.Render(code, name, data)
}

func index(c echo.Context) error {
	return renderHTML(c, http.StatusOK, "index.html", nil)
}

type TemplateRenderer struct {
	templatesDir string
}

// Render renders the HTML template
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return template.Must(template.ParseFiles(filepath.Join(t.templatesDir, name))).Execute(w, data)
}

func main() {
	e := echo.New()

	e.Static("/static", "pages/asset")
	e.Renderer = &TemplateRenderer{
		templatesDir: "pages",
	}

	e.GET("/", index)

	e.GET("/users", crud.ListUsers)
	e.POST("/users", crud.CreateUser)

	e.Start(":8080")
}
