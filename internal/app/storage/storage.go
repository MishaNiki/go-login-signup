package storage

import (
	"database/sql"

	_ "github.com/lib/pq" // driver postgresql
)

// Storage ...
type Storage struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open ...
func (stor *Storage) Open() error {
	db, err := sql.Open("postgres", stor.config.DataBaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	stor.db = db
	return nil
}

// Close ...
func (stor *Storage) Close() {
	stor.db.Close()
}

// User ...
// templete : storage.User().Create()
func (stor *Storage) User() *UserRepository {
	if stor.userRepository != nil {
		return stor.userRepository
	}

	stor.userRepository = &UserRepository{
		stor: stor,
	}

	return stor.userRepository
}
