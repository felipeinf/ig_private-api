package server

import (
	"net/http"
)

//Router struct routing
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

//NewRouter create router
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

//FindHandler find path in rules
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]

	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}

//ServeHTTP http interface function
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, methodExist, exist := r.FindHandler(req.URL.Path, req.Method)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, req)
}
