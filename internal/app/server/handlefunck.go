package server

import (
	"net/http"
	"time"
)

// handleLodin ...
func (serv *Server) handleLogin() http.HandlerFunc {

	// предоставляем доступ к серверу
	return func(w http.ResponseWriter, r *http.Request) {

		//  смотрим есть ли у пользователя куки
		c, err := r.Cookie("session")
		logdedIn := (err != http.ErrNoCookie)
		if logdedIn {
			// получение id пользователя из сессии или получении ошибки если кук нет
			_, err := serv.session.FindUserID(c.Value)

			if err != nil {
				// если ключа кук нет в сессии то
				c.Expires = time.Now().AddDate(0, 0, -1) // удаляем куки
				http.SetCookie(w, c)
				http.Redirect(w, r, "/login", http.StatusFound) // запрашиваем логин
			}

			// если куки есть и они валидны
			// то id парсим из базы данных пользователя
			// заносим нового пользователя  в структуру пользователей в  Server.treeUsers
			// указатель на нового пльзователя заносим в Contex
			// редиректим на главную страницу

			http.Redirect(w, r, "/home", http.StatusFound)
		}
	}
}

// handleSignup ...
func (serv *Server) handleSignup() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// handleForgod ...
func (serv *Server) handleForgod() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// handleHome ...
func (serv Server) handleHome() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
