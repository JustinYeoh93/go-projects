package services

import (
	"encoding/json"
	"example/src/domains"
	"example/src/repositories"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent domains.Event
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Kindly enter data with the event title and description only in order to update")
		return
	}

	err = json.Unmarshal(reqBody, &newEvent)

	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	err = repositories.SaveEvent(&newEvent)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&newEvent)
}

func GetOneEventById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	ev := repositories.FindOneEventById(id)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&ev)

}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	evs := repositories.FindAllEvents()
	json.NewEncoder(w).Encode(&evs)
}
