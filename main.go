package main

import (
	"dot/api"
	"dot/core/config"
	"dot/core/db"
)

func init() {
	db.ConnectDB()
}

func main() {
	e := config.SetupEcho()

	e.GET("/account/:id", api.GetAccount)
	e.POST("/account", api.CreateAccount)
	e.PUT("/account/:id", api.UpdateAccount)
	e.PATCH("/account/:id", api.UpdateAccount)
	e.DELETE("/account/:id", api.DeleteAccount)

	e.Logger.Fatal(e.Start(":8080"))
}
