package controllers

import (
	"bank-application/contracts"
	"bank-application/initializers"
	"bank-application/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBank(context *gin.Context) {
	var request contracts.CreateBankRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusTeapot, contracts.SingleResponse{
			Message: "geçersiz input tekrar deneyiniz",
			Item:    request,
		},
		)
		return
	}

	bank := models.Bank{
		Name: request.Name,
	}

	result := initializers.DB.Create(&bank)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"bank": bank,
	})
}

func GetAllBanks(context *gin.Context) {
	var banks []models.Bank

	initializers.DB.Model(&models.Bank{}).Preload("Interests").Find(&banks)

	context.JSON(http.StatusOK, gin.H{
		"banks": banks,
	})
}

func GetByIdBank(context *gin.Context) {
	var bank models.Bank

	initializers.DB.Model(&models.Bank{}).Preload("Interests").Find(&bank, context.Param("id"))

	if bank.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "banka bulunamadı",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"bank": bank,
	})
}

func DeleteBank(context *gin.Context) {
	var bank models.Bank

	initializers.DB.Delete(&bank, context.Param("id"))

	context.JSON(http.StatusOK, gin.H{
		"message": "silme işlemi başarılı",
	})
}
