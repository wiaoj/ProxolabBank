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
	id := context.Param("id")

	initializers.DB.Model(&models.Bank{}).Preload("Interests").Find(&bank, id)

	if bank.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "banka bulunamadı",
		})
		return
	}

	var bankResponse contracts.BankResponse

	bankResponse.Name = bank.Name

	for index := 0; index < len(bank.Interests); index++ {
		bankResponse.Interests = append(bankResponse.Interests, contracts.BankInterestsResponse{
			Interest:     bank.Interests[index].Interest,
			CreditTypeID: bank.Interests[index].CreditTypeID,
			TimeOptionID: bank.Interests[index].TimeOptionID,
		})
	}

	context.JSON(http.StatusOK, contracts.SingleResponse{
		Message: id + "'ye ait banka detayları getirildi",
		Item:    bankResponse,
	})
}

func DeleteBank(context *gin.Context) {
	var bank models.Bank

	initializers.DB.Unscoped().Delete(&bank, context.Param("id"))

	context.JSON(http.StatusOK, gin.H{
		"message": "silme işlemi başarılı",
	})
}
