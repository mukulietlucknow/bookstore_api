package app

import (
	"github.com/bookstore_api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
