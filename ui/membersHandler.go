package ui

import (
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
