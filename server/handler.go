package server

import (
	"fmt"
	"log"
	"net/http"

	m "../models"
	s "../scrap"
)

//HandleHome handle home route
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the API end point")
}

//HandleInfo info endpoint
func HandleInfo(w http.ResponseWriter, r *http.Request) {
	body := &m.InfoBody{}
	GetHTTPBody(w, r, body)

	// Get instagram info
	user, err := s.GetUsernameInfo(body.Username)
	if err != nil {
		log.Printf("Error scraping user %s. Error: %s\n", body.Username, err)
		SendResponse(false, "Internal Server Error", 500, http.StatusInternalServerError, w)
		return
	}

	SendResponseData(true, "OK", user, 200, http.StatusOK, w)
	return
}

//HandleValidate validate endpoint
func HandleValidate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Validate")
}
