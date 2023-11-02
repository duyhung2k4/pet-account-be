package main

import (
	"account/config"
	_ "account/docs"
	"account/router"
	"account/utils"
	"log"
	"net/http"
	"time"
)

// @title account
// @version 1.0
// @description account

// @host localhost:18888
// @BasePath /api/v1
func main() {
	server := &http.Server{
		Addr:         "localhost:" + config.GetAppPort(),
		Handler:      router.Router(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	utils.Cronjob()

	if err := server.ListenAndServe(); err != nil {
		log.Println("Error run server: ", err)
	}
}
