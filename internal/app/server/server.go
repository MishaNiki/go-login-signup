package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	templLogin  *template.Template
	templSignup *template.Template
)

// Server ...
type Server struct {
	config *Config
	router *http.ServeMux // главный роутер
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
	templLogin = template.Must(template.ParseFiles("web/template/login.html"))
	templSignup = template.Must(template.ParseFiles("web/template/signup.html"))
	serv.configureRouter()
	return http.ListenAndServe(serv.config.BindPort, serv.router)
}

func (serv *Server) configureRouter() {
	ApiHandler := httprouter.New()      // кастомный роутер
	StdAPiHandler := http.NewServeMux() // стандартный роутер

	StdAPiHandler.HandleFunc("/", Login)
	StdAPiHandler.HandleFunc("/profile", Profile)
	StdAPiHandler.HandleFunc("/signup", Signup)
	ApiHandler.GET("/hello/:name", Hello)

	serv.router.Handle("/", StdAPiHandler)
	serv.router.Handle("/hello/", ApiHandler)
}

// Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logEr := ""
		templLogin.Execute(w,
			struct {
				LogError string
			}{
				logEr,
			})
		return
	}
	http.Redirect(w, r, "/profile", http.StatusFound)
}

// Signup ...
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logEr := ""
		templSignup.Execute(w,
			struct {
				LogError string
			}{
				logEr,
			})
		return
	}
}

// Profile ...
func Profile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

// Hello ...
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
