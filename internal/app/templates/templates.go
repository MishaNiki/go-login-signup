package tepmplates

import "html/template"

// Tepmplates ...
//type Tepmplates map[string]*template.Template
// Tepmplates ...
type Tepmplates struct {
	Login  *template.Template
	SignUp *template.Template
	Forgod *template.Template
	Home   *template.Template
}

// New ...
func New(config *Config) *Tepmplates {
	return &Tepmplates{
		Login:  template.Must(template.ParseFiles(config.Login)),
		SignUp: template.Must(template.ParseFiles(config.SignUp)),
		Forgod: template.Must(template.ParseFiles(config.Forgod)),
		Home:   template.Must(template.ParseFiles(config.Home)),
	}
}
