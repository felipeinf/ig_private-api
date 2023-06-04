package server

import "net/http"

//Server struct
type Server struct {
	port   string
	router *Router
}

//NewServer build server
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

//Handle handle function
func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]

	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.rules[path][method] = handler
}

//AddMiddleware add interceptors
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		f = middleware(f)
	}

	return f
}

//Listen serve
func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)

	if err != nil {
		return err
	}
	return nil
}
