package controllers

import (
	"bank-application/migration"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMigration(context *gin.Context) {

	status := migration.Migration()

	context.JSON(http.StatusCreated, gin.H{
		"message": status,
	})
}

func DeleteDatabase(context *gin.Context) {

	status := migration.DeleteDatabase()

	context.JSON(http.StatusCreated, gin.H{
		"message": status,
	})
}
