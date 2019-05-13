package main

import (
	"html/template"
	"net/http"
	"io"
	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type ServiceInfo struct {
	Title string
}

var serviceInfo = ServiceInfo {
	"TEST",
}

func main() {

  t := &Template{
    templates: template.Must(template.ParseGlob("views/*.html")),
  }

  e := echo.New()

  e.Renderer = t

  e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello World")
    })

  e.GET("/pg1", func(c echo.Context) error {
    // テンプレートに渡す値

    data := struct {
      ServiceInfo
      Content_a string
    } {
      ServiceInfo: serviceInfo,
      Content_a: "TEST",
    }
    return c.Render(http.StatusOK, "pg1", data)
  }) 
	e.Logger.Fatal(e.Start(":1323"))
}

//func main() {
//	e := echo.New()
//
//	e.Renderer = t
//
//	e.GET("/",func(c echo.Context) error {
//		return c.String(http.StatusOK,"Hello, World!")
//	})
//
//	e.Logger.Fatal(e.Start(":1323"))
//}
