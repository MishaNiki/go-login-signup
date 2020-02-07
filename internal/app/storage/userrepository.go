package storage

import "github.com/MishaNiki/go-login-signup/internal/app/model"

// UserRepository ...
type UserRepository struct {
	stor *Storage
}

// Create создание пользователя в базе данных
func (r *UserRepository) Create(u *model.User) (*model.User, error) {

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.stor.db.QueryRow(
		"INSERT INTO \"shGologin\".\"objUsers\" (email, pass) Values($1, $2) RETURNING id_user",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	u := &model.User{}
	if err := r.stor.db.QueryRow(
		"SELECT id_user, email, pass FROM \"shGologin\".\"objUsers\" WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword); err != nil {
		return nil, err
	}

	return u, nil
}
