package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"
)

//CheckAuht middleware
func CheckAuht() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Authentication")

			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

//Logging middleware
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			defer func() {
				log.Println(r.Method, r.URL.Path, time.Since(start))
			}()

			f(w, r)
		}
	}
}

func GetHTTPBody(res http.ResponseWriter, req *http.Request, body interface{}) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	CheckErr(err, res)

	err = json.Unmarshal(bodyBytes, body)
	CheckErr(err, res)
}
