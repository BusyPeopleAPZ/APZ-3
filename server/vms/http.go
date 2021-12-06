package vms

import (
	"encoding/json"
	"github.com/BusyPeople/APZ-3/server/tools"
	"log"
	"net/http"
)

type VMS struct {
	Id int64 `json:"id"`
	Status bool `json:"status"`
}

func MainLoader(cont *Container) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			loadConnInfo(responseWriter, cont)
		} else if request.Method == "PUT" {
			loadUpdateVMS(request, responseWriter, cont)
		} else {
			responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func loadConnInfo(responseWriter http.ResponseWriter, cont *Container) {
	result, fault := cont.ConnInfo()
	if fault == nil {
		tools.WriteJsonOk(responseWriter, result)
	} else {
		log.Printf("%s", fault)
		tools.WriteJsonInternalError(responseWriter, "sth went wrong...")
	}
}

func loadUpdateVMS(request *http.Request, responseWriter http.ResponseWriter, cont *Container) {
	var vms VMS
	if fault := json.NewDecoder(request.Body).Decode(&vms); fault != nil {
		log.Printf("%s", fault)
		tools.WriteJsonBadRequest(responseWriter, "fault to decode json")
		return
	}

	if fault := cont.UpdateVMS(vms.Id, vms.Status); fault == nil {
		tools.WriteJsonOkReplacement(responseWriter)
	} else {
		return
	}
}
