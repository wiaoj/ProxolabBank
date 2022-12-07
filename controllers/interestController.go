package controllers

import (
	"bank-application/contracts"
	"bank-application/initializers"
	"bank-application/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateInterest(context *gin.Context) {
	var request contracts.CreateInterestRequest

	//418
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusTeapot, contracts.SingleResponse{
			Message: "geçersiz input tekrar deneyiniz",
			Item:    request,
		},
		)
		return
	}

	var creditType []models.CreditTypeTimeOption
	initializers.DB.Find(&creditType, request.CreditTypeID)

	isCreditTypeCorrect := false

	for i := 0; i < len(creditType); i++ {
		if creditType[i].TimeOptionID == int(request.TimeOptionID) {
			isCreditTypeCorrect = true
		}
	}

	if !isCreditTypeCorrect {
		context.JSON(http.StatusBadRequest, contracts.SingleResponse{
			Message: "faiz eklenemedi, bu kredi türüyle bu vadeyi ekleyemezsin",
			Item:    request,
		},
		)
		return
	}

	interest := models.Interest{
		BankID:       request.BankID,
		Interest:     request.Interest,
		TimeOptionID: request.TimeOptionID,
		CreditTypeID: request.CreditTypeID,
	}

	result := initializers.DB.Create(&interest)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, contracts.SingleResponse{
		Message: "interest eklendi",
		Item: contracts.InterestResponse{
			BankID:       interest.BankID,
			Interest:     interest.Interest,
			TimeOptionID: interest.TimeOptionID,
			CreditTypeID: interest.CreditTypeID,
		},
	},
	)
}

func DeleteInterest(context *gin.Context) {
	var interest models.Interest

	initializers.DB.Unscoped().Delete(&interest, context.Param("id"))

	context.JSON(http.StatusOK, gin.H{
		"message": "faiz kaldırıldı",
	})
}

func GetInterestsQuery(context *gin.Context) {
	var interests []*models.Interest
	database := *initializers.DB.Model(&interests)

	interestOrderType := context.Query("interestOrderType")

	if interestOrderType == "" {
		interestOrderType = "asc"
	}

	database.Preload("Bank").Preload("TimeOption").Preload("CreditType")
	if bankId, err := strconv.ParseUint(context.Query("bankId"), 10, 64); err == nil {
		database.Where(&models.Interest{BankID: uint(bankId)})
	}

	if timeOptionId, err := strconv.ParseUint(context.Query("timeOptionId"), 10, 64); err == nil {
		database.Where(&models.Interest{TimeOptionID: uint(timeOptionId)}).Find(&interests)
	} else if creditTypeId, err := strconv.ParseUint(context.Query("creditTypeId"), 10, 64); err == nil {
		database.Where(&models.Interest{CreditTypeID: uint(creditTypeId)}).Find(&interests)
	}

	database.Order("interest " + interestOrderType).Find(&interests)

	// initializers.DB.
	// 	Joins("JOIN banks ON banks.id = ?", uint(bankId)).
	// 	Joins("JOIN time_options ON time_options.id = ?", uint(timeOptionId)).
	// 	Joins("JOIN credit_types ON credit_types.id = ?", uint(creditTypeId)).
	// 	Order("interest " + interestOrderType).Find(&interests)

	var interestsResponse []contracts.InterestResponse

	for index := 0; index < len(interests); index++ {
		interestsResponse = append(interestsResponse, contracts.InterestResponse{
			Id:                    interests[index].ID,
			BankID:                interests[index].BankID,
			BankName:              interests[index].Bank.Name,
			Interest:              interests[index].Interest,
			CreditTypeID:          interests[index].CreditTypeID,
			CreditTypeDescription: interests[index].CreditType.Description,
			TimeOptionID:          interests[index].TimeOptionID,
			TimeOptionDescription: interests[index].TimeOption.Description,
		})
	}

	context.JSON(http.StatusOK, contracts.MultipleResponse{
		Message: "faizler " + interestOrderType + " şeklinde sıralanıp getirildi",
		Items:   interestsResponse,
	})
}
func GetAllInterest(context *gin.Context) {
	var interests []models.Interest

	initializers.DB.Preload("Bank").Preload("TimeOption").Preload("CreditType").Order("interest asc").Find(&interests)

	if interests[0].ID == 0 {
		context.JSON(http.StatusNotFound, contracts.MultipleResponse{
			Message: "herhangi bir faiz bulunamadı",
			Items:   []any{},
		})
		context.Abort()
		return
	}

	var interestsResponse []contracts.InterestResponse

	for index := 0; index < len(interests); index++ {
		interestsResponse = append(interestsResponse, contracts.InterestResponse{
			BankID:                interests[index].BankID,
			BankName:              interests[index].Bank.Name,
			Interest:              interests[index].Interest,
			CreditTypeID:          interests[index].CreditTypeID,
			CreditTypeDescription: interests[index].CreditType.Description,
			TimeOptionID:          interests[index].TimeOptionID,
			TimeOptionDescription: interests[index].TimeOption.Description,
		})
	}

	context.JSON(http.StatusOK, contracts.MultipleResponse{
		Message: "faizler düşükten yükseğe doğru sıralanıp getirildi",
		Items:   interestsResponse,
	})
}
