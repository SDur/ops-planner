package ui

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/SDur/ops-planner/model"
)

type Config struct {
	Assets http.FileSystem
}

func Start(cfg Config, m *model.Model, listener net.Listener) {

	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16}

	http.Handle("/", indexHandler(m))
	http.Handle("/members", membersHandler(m))
	http.Handle("/js/", http.FileServer(cfg.Assets))

	go server.Serve(listener)
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

func indexHandler(m *model.Model) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, indexHTML)
	})
}

func membersHandler(m *model.Model) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			members, err := m.Members()
			if err != nil {
				http.Error(w, "This is an error", http.StatusBadRequest)
				return
			}

			js, err := json.Marshal(members)
			if err != nil {
				http.Error(w, "This is an error", http.StatusBadRequest)
				return
			}

			fmt.Fprintf(w, string(js))
		case "POST":
			firstnames, ok := r.URL.Query()["firstname"]
			lastnames, ok := r.URL.Query()["lastname"]

			if !ok || len(firstnames[0]) < 1 || len(lastnames[0]) < 0 {
				log.Println("Url Params are incomplete or missing")
				return
			}

			// Query()["key"] will return an array of items,
			// we only want the single item.
			firstname := firstnames[0]
			lastname := lastnames[0]

			log.Println("Received new member: " + firstname + " " + lastname)
			newMember := &model.Member{
				Id:        0,
				Firstname: firstname,
				Lastname:  lastname}

			err := m.AddMember(newMember)
			if err != nil {
				log.Println("Something went wrong")
			}
		case "DELETE":
			ids, ok := r.URL.Query()["id"]
			if !ok || len(ids[0]) < 1 {
				log.Println("Url Params are incomplete or missing")
				return
			}
			id, ierr := strconv.Atoi(ids[0])
			if ierr != nil {
				log.Println("Error parsing url param to int")
			}
			err := m.RemoveMember(id)
			if err != nil {
				log.Println("Something went wrong")
			}
		}
	})
}
