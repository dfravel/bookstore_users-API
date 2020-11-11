package app

import (
	"bookstore_users-API/controllers/ping"
	"bookstore_users-API/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users", user.GetUser)
	router.GET("/users/:user_id", user.GetUserByID)
}
