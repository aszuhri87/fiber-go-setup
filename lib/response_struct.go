package lib

import (
	"net/http"
	"time"
)

type ThingResponse struct {
	UUID    string    `json:"uuid"`
	Name    string    `json:"name"`
	Value   string    `json:"value"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func GetThing(w http.ResponseWriter, r *http.Request) {}

func DeleteThing(w http.ResponseWriter, r *http.Request) {}

type ThingsResponse struct {
	Total  int             `json:"total"`
	Page   int             `json:"page"`
	Limit  int             `json:"limit"`
	Things []ThingResponse `json:"things"`
}

func ListThings(w http.ResponseWriter, r *http.Request) {}

type ResponseStatus struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
