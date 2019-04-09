package ui

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"

	"github.com/SDur/ops-planner/model"
)

type Config struct {
	Assets http.FileSystem
}

func Start(cfg Config, m *model.Model) {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", indexHandler(m))
	e.GET("/members", getMembersHandler(m))
	e.PUT("/members", putMembersHandler(m))
	e.DELETE("/members", deleteMembersHandler(m))

	e.GET("/sprints", getSprintsHandler(m))
	e.POST("/sprints", postSprintsHandler(m))
	e.PUT("/sprints", putSprintsHandler(m))

	e.File("/js/app.jsx", "assets/js/app.jsx")
	e.File("/js/style.css", "assets/js/style.css")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

const (
	cdnReact           = "https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"
	cdnReactDom        = "https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"
	cdnBabelStandalone = "https://cdnjs.cloudflare.com/ajax/libs/babel-standalone/6.24.0/babel.min.js"
	cdnAxios           = "https://cdnjs.cloudflare.com/ajax/libs/axios/0.16.1/axios.min.js"
)

const indexHTML = `
<!DOCTYPE HTML>
<html>
  <head>
    <meta charset="utf-8">
	<link rel="stylesheet" type="text/css" href="/js/style.css" media="screen" />
    <title>OPSer van de dag</title>
  </head>
  <body>
    <div id='root'></div>
    <script src="` + cdnReact + `"></script>
    <script src="` + cdnReactDom + `"></script>
    <script src="` + cdnBabelStandalone + `"></script>
    <script src="` + cdnAxios + `"></script>
    <script src="/js/app.jsx" type="text/babel"></script>
  </body>
</html>
`

func indexHandler(m *model.Model) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.HTML(http.StatusOK, indexHTML)
	}
}

func postSprintsHandler(m *model.Model) echo.HandlerFunc {
	return func(c echo.Context) error {
		s := new(model.Sprint)
		if err := c.Bind(s); err != nil {
			c.Error(err)
		}
		log.Println("Received update for sprint: ")
		log.Println(s)
		if err := m.SaveSprint(s); err != nil {
			c.Error(err)
		}
		return c.NoContent(204)
	}
}

func putSprintsHandler(m *model.Model) echo.HandlerFunc {
	return func(c echo.Context) error {
		s := new(model.Sprint)
		if err := c.Bind(s); err != nil {
			c.Error(err)
		}
		for i := 0; i <= 9; i += 1 {
			if s.Days[i] == 0 {
				s.Days[i] = -1
			}
		}
		log.Println("Received new sprint: ")
		log.Println(s)
		if err := m.AddSprint(s); err != nil {
			c.Error(err)
		}
		return c.NoContent(204)
	}
}

func getSprintsHandler(m *model.Model) echo.HandlerFunc {
	return func(c echo.Context) error {
		sprint, err := m.CurrentSprint()
		if err != nil {
			log.Println(err)
			c.Error(err)
		}
		return c.JSON(http.StatusOK, sprint)
	}
}
