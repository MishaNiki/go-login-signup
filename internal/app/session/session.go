package session

type sessionData struct {
	IDUser int
}

// Session ...
type Session map[string]*sessionData

// New ...
func New() *Session {
	s := make(Session)
	return &s
}
