package session

import (
	"errors"

	"github.com/MishaNiki/go-login-signup/internal/app/utils"
)

type sessionData struct {
	IDUser int
}

// Session ...
type Session struct {
	data map[string]*sessionData
}

// New ...
func New() *Session {
	return &Session{
		data: make(map[string]*sessionData),
	}
}

// InitUser ...
func (s *Session) InitUser(userID int) string {
	sessionID := utils.GenerateKeySession()
	data := &sessionData{IDUser: userID}
	s.data[sessionID] = data
	return sessionID
}

// ErrorFindID ...
var ErrorFindID = errors.New("key is not session")

// FindUserID ...
func (s *Session) FindUserID(key string) (int, error) {
	data, ok := s.data[key]
	if !ok {
		return 0, ErrorFindID
	}

	return data.IDUser, nil
}

// RemoveUser remove user from session
func (s *Session) RemoveUser(key string) {
	delete(s.data, key)
}
