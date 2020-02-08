package templates

import "html/template"

// Templates ...
//type Tepmplates map[string]*template.Template
// Tepmplates ...
type Templates struct {
	Login  *template.Template
	SignUp *template.Template
	Forgod *template.Template
	Home   *template.Template
}

// New ...
func New(config *Config) (*Templates, error) {

	login, err := template.ParseFiles(config.Login)
	if err != nil {
		return nil, err
	}

	signup, err := template.ParseFiles(config.Login)
	if err != nil {
		return nil, err
	}

	forgod, err := template.ParseFiles(config.Login)
	if err != nil {
		return nil, err
	}

	home, err := template.ParseFiles(config.Login)
	if err != nil {
		return nil, err
	}

	return &Templates{
		Login:  login,
		SignUp: signup,
		Forgod: forgod,
		Home:   home,
	}, nil
}
