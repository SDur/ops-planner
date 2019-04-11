package ui

import (
	"bytes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	e.GET("/slack", getSlackHandler(m))

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

func getSlackHandler(m *model.Model) echo.HandlerFunc {
	return func(c echo.Context) error {
		webhook := "https://hooks.slack.com/services/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"
		resp, err := http.Post(webhook, "application/json", bytes.NewBufferString(`{"text":"Buy cheese and bread for breakfast."}`))

		if err != nil {
			c.Error(err)
		}
		return c.String(resp.StatusCode, "Message send")
	}
}
