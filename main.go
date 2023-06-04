package main

import (
	s "./server"
)

func main() {
	server := s.NewServer(":3000")
	server.Handle("GET", "/api", server.AddMiddleware(s.HandleHome, s.Logging()))
	server.Handle("POST", "/info", server.AddMiddleware(s.HandleInfo, s.Logging()))
	server.Handle("POST", "/validate", server.AddMiddleware(s.HandleValidate, s.Logging()))
	server.Listen()
}
