package main

import (
	"backend/database"
	"github.com/gin-gonic/gin"
)

func main(){
	db := &database.Database{}
	db.Connect()
	defer db.Connection.Close()

	router := gin.Default()

	router.Run(":8000")
}
