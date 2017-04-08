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
	pid := []models.PeopleOfDelta{n1}
	return pid
}