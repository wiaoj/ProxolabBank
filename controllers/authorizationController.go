package controllers

import (
	"bank-application/contracts"
	"bank-application/initializers"
	"bank-application/models"
	"bank-application/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var request contracts.UserRegisterRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusTeapot, contracts.SingleResponse{
			Message: "geçersiz input tekrar deneyiniz",
			Item:    request,
		},
		)
		return
	}

	var user models.User

	if initializers.DB.Where("username = ?", request.Username).First(&user); user.Username != "" {
		context.JSON(http.StatusBadRequest, contracts.SingleResponse{
			Message: "böyle bir kullanıcı zaten mevcut",
		})
		context.Abort()
		return
	}

	user = models.User{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Username:    request.Username,
		Age:         request.Age,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}

	passwordHash, hashError := utils.HashPassword(request.Password)

	if hashError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": hashError.Error()})
		context.Abort()
		return
	}

	user.PasswordHash = passwordHash

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		context.Abort()
		return
	}

	// token, jwtError := utils.GenerateJWT(user)

	// if jwtError != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": jwtError.Error()})
	// 	context.Abort()
	// 	return
	// }

	context.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func Login(context *gin.Context) {
	var request contracts.UserLoginRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusTeapot, contracts.SingleResponse{
			Message: "geçersiz input tekrar deneyiniz",
			Item:    request,
		},
		)
		return
	}

	var user models.User
	initializers.DB.Where("username = ?", request.Username).First(&user)

	if user.Username == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "kullanıcı bulunamadı"})
		context.Abort()
		return
	}

	chechPassword := utils.CheckPassword(user.PasswordHash, request.Password)

	if chechPassword != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": chechPassword.Error()})
		context.Abort()
		return
	}

	token, jwtError := utils.GenerateJWT(user)

	if jwtError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": jwtError.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// func GetAll(context *gin.Context) {
// 	var users []models.User

// 	initializers.DB.Find(&users)
// }

// func GetById(context *gin.Context) {
// 	var user models.User

// 	id := context.Param("id")

// 	initializers.DB.Find(&user, id)
// }

// func Update(context *gin.Context) {
// 	var user models.User

// 	id := context.Param("id")

// 	//update model
// 	context.Bind(&user)
// 	initializers.DB.First(&user, id)

// 	var x models.User
// 	initializers.DB.Model(&x).Updates(models.User{
// 		Username: user.Username, //bla bla
// 	})
// }

// func Delete(context *gin.Context) {
// 	id := context.Param("id")
// 	initializers.DB.Delete(&models.User{}, id)
// }
