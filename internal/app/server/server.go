package server

import (
	"fmt"
	"html/template"
	"net/http"
)

// Server ...
type Server struct {
	config *Config
	router *http.ServeMux
}

// New ...
func New(config *Config) *Server {
	return &Server{
		config: config,
		router: http.NewServeMux(),
	}
}

// Start ...
func (serv *Server) Start() error {

	serv.configureRouter()

	return http.ListenAndServe(serv.config.BindPort, serv.router)
}

//
// configureRouter ...
// присваивание к url функций обработки
func (serv *Server) configureRouter() {
	// Создание роутеров
	stdAPIHandler := http.NewServeMux()

	// Создание обработчиков url для роутеров
	stdAPIHandler.HandleFunc("/", Login)
	stdAPIHandler.HandleFunc("/signup", Signup)
	stdAPIHandler.HandleFunc("/profile", Profile)

	// слив всех роутеров в один главный
	serv.router.Handle("/", stdAPiHandler)
}
