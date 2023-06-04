package server

import "net/http"

//Middleware interceptor
type Middleware func(http.HandlerFunc) http.HandlerFunc
