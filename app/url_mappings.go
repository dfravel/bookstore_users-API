package app

import (
	"bookstore_users-API/controllers/ping"
	"bookstore_users-API/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/user", user.GetUser)
}
