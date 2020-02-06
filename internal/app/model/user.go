package model

// модель пользователя
type User struct {
	ID            int
	Email         string
	EncryptedPass string
}
