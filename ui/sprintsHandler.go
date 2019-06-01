package ui

import (
	"github.com/SDur/ops-planner/model"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

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

func getLatestSprintsHandler(m *model.Model) echo.HandlerFunc {
	return func(c echo.Context) error {
		sprint, err := m.LatestSprint()
		if err != nil {
			log.Println(err)
			c.Error(err)
		}
		return c.JSON(http.StatusOK, sprint)
	}
}

func getSprintsHandler(m *model.Model) echo.HandlerFunc {
	return func(c echo.Context) error {
		sprints, err := m.Sprints()
		if err != nil {
			log.Println(err)
			c.Error(err)
		}
		return c.JSON(http.StatusOK, sprints)
	}
}

func deleteSprintHandler(m *model.Model) echo.HandlerFunc {
	return func(c echo.Context) error {
		idString := c.QueryParam("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			log.Println("Error parsing url param to int")
			c.Error(err)
		}
		err = m.RemoveSprint(id)
		if err != nil {
			log.Println("Something went wrong")
			c.Error(err)
		}
		return c.String(http.StatusOK, "sprint deleted")
	}
}
