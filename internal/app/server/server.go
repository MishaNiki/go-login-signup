package server

import (
	//"fmt"
	//"html/template"
	"net/http"

	"github.com/MishaNiki/go-login-signup/internal/app/mail"
	"github.com/MishaNiki/go-login-signup/internal/app/session"
	"github.com/MishaNiki/go-login-signup/internal/app/storage"
	"github.com/MishaNiki/go-login-signup/internal/app/templates"
	"github.com/MishaNiki/go-login-signup/internal/app/treeusers"
)

// Server ...
type Server struct {
	config   *Config
	router   *http.ServeMux
	mail     *mail.Mail
	stor     *storage.Storage
	templ    *templates.Templates
	session  *session.Session
	treeUser *treeusers.TreeUsers
}

// New ...
func New(config *Config) *Server {
	return &Server{
		config:  config,
		router:  http.NewServeMux(),
		session: session.New(),
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

	// конфигуранция шаблонов страниц перед началом работы сервера
	if err := serv.configureTemplates(); err != nil {
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
	stdAPIHandler.HandleFunc("/", serv.handleLogin())
	stdAPIHandler.HandleFunc("/signup", serv.handleSignup())
	stdAPIHandler.HandleFunc("/forgod", serv.handleForgod())
	stdAPIHandler.HandleFunc("/home", serv.handleHome())

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

// configureTemplates ...
func (serv *Server) configureTemplates() error {

	templ, err := templates.New(serv.config.Templates)
	if err != nil {
		return err
	}
	serv.templ = templ
	return nil
}
