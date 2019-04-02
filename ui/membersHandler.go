package ui

import (
	"encoding/json"
	"fmt"
	"github.com/SDur/ops-planner/model"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

func getMembersHandler(m *model.Model) echo.HandlerFunc {
	return func(c echo.Context) error {
		members, err := m.Members()
		if err != nil {
			c.Error(err)
		}
		return c.JSON(http.StatusOK, members)
	}
}

func putMembersHandler(m *model.Model) echo.HandlerFunc {
	return func(c echo.Context) error {
		firstname := c.QueryParam("firstname")
		lastname := c.QueryParam("lastname")

		log.Println("Received new member: " + firstname + " " + lastname)
		newMember := &model.Member{
			Id:        0,
			Firstname: firstname,
			Lastname:  lastname}

		err := m.AddMember(newMember)
		if err != nil {
			c.Error(err)
		}
		return c.String(http.StatusOK, "member added")
	}
}

func deleteMembersHandler(m *model.Model) echo.HandlerFunc {
	return func(c echo.Context) error {
		idString := c.QueryParam("id")

		id, erri := strconv.Atoi(idString)
		if erri != nil {
			log.Println("Error parsing url param to int")
			c.Error(erri)
		}
		err := m.RemoveMember(id)
		if err != nil {
			log.Println("Something went wrong")
			c.Error(err)
		}
		return c.String(http.StatusOK, "member deleted")
	}
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
