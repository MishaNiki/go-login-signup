package server

import (
	//"fmt"
	//"html/template"
	"net/http"

	"github.com/MishaNiki/go-login-signup/internal/app/mail"
	"github.com/MishaNiki/go-login-signup/internal/app/storage"
)

// Server ...
type Server struct {
	config *Config
	router *http.ServeMux
	mail   *mail.Mail
	stor   *storage.Storage
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

	serv.configureRouter() // конфигурация роутера

	// конфигурация и подключение к бд
	if err := serv.configureStorage(); err != nil {
		return err
	}

	// конфигурация и подключение к почтовому ящику с которого будет произведена рассылка
	if err := serv.configureMail(); err != nil {
		return err
	}

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
	serv.router.Handle("/", stdAPIHandler)
}

//  configureStorage ...
func (serv *Server) configureStorage() error {
	stor := storage.New(serv.config.Storage)
	if err := stor.Open(); err != nil {
		return err
	}

	serv.stor = stor
	return nil
}

// configureMail ...
func (serv *Server) configureMail() error {

	mail := mail.New(serv.config.Mail)
	if err := mail.Ping(); err != nil {
		return err
	}

	serv.mail = mail
	return nil
}
