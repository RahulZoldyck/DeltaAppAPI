package controllers

import (
	"github.com/revel/revel"
	"encoding/json"
	"hash/fnv"
	"strconv"
	"github.com/RahulZoldyck/DeltaAppAPI/app/models"
	"time"
)


type App struct {
	*revel.Controller
}

func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}

func (c App) Index() revel.Result {
	return c.Render()
}

type LoginRequest struct {
	RollNo int64
	Password string
}

type LoginResponse struct {
	Token uint32
	Success bool
	StatusCode int
}

func (c App) Login() revel.Result {
	var request LoginRequest
    err := json.NewDecoder(c.Request.Body).Decode(&request)
	currTime := time.Now().UnixNano() / int64(time.Millisecond)
	tokenstring := strconv.FormatInt(currTime,10)+ strconv.FormatInt(request.RollNo,10)
	s := LoginResponse {hash(tokenstring),true,200}
	js,err := json.Marshal(s)
	if(err!=nil) {
		panic(err)
	}
	return c.RenderJson(string(js))
}

type PeopleInDeanResponse struct {
	PeopleInRoom []models.PeopleOfDelta
	PeopleHasKey []models.PeopleOfDelta
}

func (c App) PeopleInDean() revel.Result {
	pid:=GetPeopleInDean()
	phk:=GetPeopleHasKey()
	pidr:=PeopleInDeanResponse{pid,phk}
	js,err := json.Marshal(pidr)
	if(err!=nil) {
		panic(err)
	}
	return c.RenderJson(string(js))
}