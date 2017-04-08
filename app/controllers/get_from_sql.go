package controllers

import(
	"github.com/RahulZoldyck/DeltaAppAPI/app/models"
)

func GetPeopleInDean() []models.PeopleOfDelta{
	n1:= models.PeopleOfDelta{"RahulZoldyck",110114070}
	pid := []models.PeopleOfDelta{n1}
	return pid
}

func GetPeopleHasKey() []models.PeopleOfDelta{
	n1:= models.PeopleOfDelta{"RahulZoldyck",110114070}
	phk := []models.PeopleOfDelta{n1}
	return phk
}

func SqlAuth(rollno int64,password string) bool {
	return true
}

func GetNotifications() {
	n1:= models.Notifications {"First Notifications","RahulZoldyck",time.Now().UTC().Format("2006-01-02T15:04:05-0700")}
	notify := []models.Notifications {n1}
	return notify
}

func GetScrumTalks() {
	n1:= models.ScrumTalks {"First ScrumTalks","RahulZoldyck",time.Now().UTC().Format("2006-01-02T15:04:05-0700")}
	st := []models.ScrumTalks {n1}
	return st
}