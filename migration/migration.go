package migration

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

}

func Migration() {
	if tables, _ := initializers.DB.Migrator().GetTables(); len(tables) == 0 {
		//x := initializers.DB.Migrator().HasTable("users")
		Migrate()
	}
}

const databaseName = "public"

const upSchema = "Create SCHEMA public"
const dropSchema = "DROP SCHEMA IF EXISTS public CASCADE;"

const createUser = `-- public.users definition

-- Drop table

-- DROP TABLE users;

CREATE TABLE users (
	id bigserial NOT NULL PRIMARY KEY,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	username text NULL,
	first_name text NULL,
	last_name text NULL,
	age int2 NULL,
	email text NULL,
	phone_number text NULL,
	password_hash bytea NULL
	--CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);`

const createClaims = `-- public.claims definition

-- Drop table

-- DROP TABLE claims;

CREATE TABLE claims (
	id bigserial NOT NULL PRIMARY KEY,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" text NULL,
	"level" int8 NULL
	--CONSTRAINT claims_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_claims_deleted_at ON claims (deleted_at timestamptz_ops);`

const createUsersClaims = `-- public.users_claims definition

-- Drop table

-- DROP TABLE users_claims;

CREATE TABLE users_claims (
	claim_id bigserial NOT NULL,
	user_id bigserial NOT NULL
	--CONSTRAINT users_claims_pkey NULL
);


-- public.users_claims foreign keys

ALTER TABLE public.users_claims ADD CONSTRAINT fk_users_claims_claim FOREIGN KEY (claim_id) REFERENCES claims(id);
ALTER TABLE public.users_claims ADD CONSTRAINT fk_users_claims_user FOREIGN KEY (user_id) REFERENCES users(id);`

const createBanks = `-- public.banks definition

-- Drop table

-- DROP TABLE banks;

CREATE TABLE banks (
	id bigserial NOT NULL PRIMARY KEY,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" text NOT NULL
);
CREATE INDEX idx_banks_deleted_at ON banks (deleted_at timestamptz_ops);`

const createInterests = `-- public.interests definition

-- Drop table

-- DROP TABLE interests;

CREATE TABLE interests (
	id bigserial NOT NULL PRIMARY KEY,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	bank_id bigserial NOT NULL,
	interest numeric NOT NULL,
	time_option_id bigserial NOT NULL,
	credit_type_id bigserial NOT NULL
	--CONSTRAINT interests_pkey null
);
ALTER TABLE public.interests ADD CONSTRAINT fk_banks FOREIGN KEY (bank_id) REFERENCES banks(id);
ALTER TABLE public.interests ADD CONSTRAINT fk_time_option FOREIGN KEY (time_option_id) REFERENCES time_options(id);
ALTER TABLE public.interests ADD CONSTRAINT fk_credit_type FOREIGN KEY (credit_type_id) REFERENCES credit_types(id);
`

const createCreditTypes = `-- public.credit_types definition

-- Drop table

-- DROP TABLE credit_types;

CREATE TABLE credit_types (
	id bigserial NOT NULL PRIMARY KEY,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	description text NULL
);
CREATE INDEX idx_credit_types_deleted_at ON credit_types (deleted_at timestamptz_ops);`

const createTimeOptions = `-- public.time_options definition

-- Drop table

-- DROP TABLE time_options;

CREATE TABLE time_options (
	id bigserial NOT NULL PRIMARY KEY,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	description text NULL
);
CREATE INDEX idx_time_options_deleted_at ON time_options (deleted_at timestamptz_ops);`

const createCreditTypeTimeOptions = `-- public.credit_type_time_options definition

-- Drop table

-- DROP TABLE credit_type_time_options;

CREATE TABLE credit_type_time_options (
	time_option_id bigserial NOT NULL,
	credit_type_id bigserial NOT NULL
	--CONSTRAINT credit_type_time_options_pkey null
);


-- public.credit_type_time_options foreign keys

ALTER TABLE public.credit_type_time_options ADD CONSTRAINT fk_credit_type_time_options_credit_type FOREIGN KEY (credit_type_id) REFERENCES credit_types(id);
ALTER TABLE public.credit_type_time_options ADD CONSTRAINT fk_credit_type_time_options_time_option FOREIGN KEY (time_option_id) REFERENCES time_options(id);`

func Migrate() {
	if err := database.Exec(dropSchema).Error; err != nil {
		panic(err)
	}

	if err := database.Exec(upSchema).Error; err != nil {
		panic(err)
	}

	database.Exec(createUser)
	database.Exec(createClaims)
	database.Exec(createUsersClaims)

	database.Exec(createBanks)

	database.Exec(createCreditTypes)
	database.Exec(createTimeOptions)
	database.Exec(createCreditTypeTimeOptions)

	database.Exec(createInterests)
	// database.AutoMigrate(&models.User{})
	// database.AutoMigrate(&models.Claim{})
	// database.AutoMigrate(&models.UsersClaims{})

	// database.AutoMigrate(&models.Bank{})
	// database.AutoMigrate(&models.Interest{})

	// database.AutoMigrate(&models.CreditType{})
	// database.AutoMigrate(&models.TimeOption{})
	// database.AutoMigrate(&models.CreditTypeTimeOption{})

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
