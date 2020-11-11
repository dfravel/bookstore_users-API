package app

import "../controllers/ping/"

func mapUrls() {
	router.GET("/ping", ping.Ping)
}
