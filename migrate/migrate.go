package main

import (
	"bank-application/initializers"
	"bank-application/models"
	"bank-application/utils"

	"gorm.io/gorm"
)

var database gorm.DB

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	database = *initializers.DB
}

func main() {
	Migrate()
}

func Migrate() {
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Claim{})
	database.AutoMigrate(&models.UsersClaims{})

	database.AutoMigrate(&models.Bank{})
	database.AutoMigrate(&models.Interest{})

	database.AutoMigrate(&models.CreditType{})
	database.AutoMigrate(&models.TimeOption{})
	database.AutoMigrate(&models.CreditTypeTimeOption{})

	seedClaims()
	seedAdmin()

	seedBank()
	seedCreditType()
	seedTimeOption()

	seedCreditTypeTimeOption()
}
func seedAdmin() {
	// var admin models.User
	// if admin := database.Where("username = ?", "proxolab").First(&admin); admin != nil {
	// 	return
	// }

	var adminClaim models.Claim
	database.Where("name = ?", "admin").First(&adminClaim)

	passwordHash, _ := utils.HashPassword("proxolab")
	adminUser := models.User{
		Username:     "proxolab",
		Email:        "proxolab",
		PasswordHash: passwordHash,
	}

	database.Create(&adminUser)
	database.Create(&models.UsersClaims{
		ClaimID: adminClaim.ID,
		UserID:  adminUser.ID,
	})

}

func seedClaims() {
	var claims []models.Claim

	claims = append(claims,
		models.Claim{
			Name:  "admin",
			Level: models.AdminClaimLevel,
		},
	)

	database.Create(&claims)
}

func seedBank() {
	var banks []models.Bank

	banks = append(banks,
		models.Bank{
			Name: "Banka 1",
		},
		models.Bank{
			Name: "Banka 2",
		})

	database.Create(&banks)
}

func seedCreditType() {
	var creditType []models.CreditType

	creditType = append(creditType,
		models.CreditType{
			Description: "Konut Kredisi",
		},
		models.CreditType{
			Description: "Tüketici Kredisi",
		},
		models.CreditType{
			Description: "Mevduat Kredisi",
		})

	database.Create(&creditType)
}

func seedTimeOption() {

	var timeOption []models.TimeOption

	timeOption = append(timeOption,
		models.TimeOption{
			Description: "3 Ay",
		},
		models.TimeOption{
			Description: "6 Ay",
		},
		models.TimeOption{
			Description: "12 Ay",
		},
		models.TimeOption{
			Description: "24 Ay",
		},
		models.TimeOption{
			Description: "36 Ay",
		},
		models.TimeOption{
			Description: "5 yıl",
		},
		models.TimeOption{
			Description: "10 yıl",
		})

	database.Create(&timeOption)
}

func seedCreditTypeTimeOption() {
	var creditTypeTimeOption []models.CreditTypeTimeOption

	creditTypeTimeOption = append(creditTypeTimeOption,
		models.CreditTypeTimeOption{
			CreditTypeID: 1,
			TimeOptionID: 6,
		},
		models.CreditTypeTimeOption{
			CreditTypeID: 1,
			TimeOptionID: 7,
		},
		models.CreditTypeTimeOption{
			CreditTypeID: 2,
			TimeOptionID: 3,
		},
		models.CreditTypeTimeOption{
			CreditTypeID: 2,
			TimeOptionID: 4,
		},
		models.CreditTypeTimeOption{
			CreditTypeID: 2,
			TimeOptionID: 5,
		},
		models.CreditTypeTimeOption{
			CreditTypeID: 3,
			TimeOptionID: 1,
		},
		models.CreditTypeTimeOption{
			CreditTypeID: 3,
			TimeOptionID: 2,
		},
		models.CreditTypeTimeOption{
			CreditTypeID: 3,
			TimeOptionID: 3,
		})

	database.Create(&creditTypeTimeOption)
}
