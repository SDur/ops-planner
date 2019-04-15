package cron

import (
	"fmt"
	"github.com/SDur/ops-planner/model"
	"github.com/SDur/ops-planner/ui"
	"github.com/robfig/cron"
	"log"
	"time"
)

func StartCron(m *model.Model) {
	c := cron.New()
	err := c.AddFunc("0 0 6 * * MON,TUE,WED,THU,FRI", func() {
		sendOpser(m)
	})
	if err != nil {
		log.Fatal("Couldnt init cron job")
	}
	c.Start()
	log.Println("Cron job started: ")
	log.Println(c.Entries()[0])
}

func sendOpser(m *model.Model) {
	fmt.Println("*** Running cron job, who is the opser?")
	member, e := m.GetMemberForDate(time.Now())
	if e != nil {
		log.Println("Something went wrong in the cron job getting the opser of today")
		log.Println(e.Error())
	}
	e = ui.SendSlackMessage(member)
	if e != nil {
		log.Println("Something went wrong in the cron job sending the opser of today to slack")
		log.Println(e.Error())
	}
}
