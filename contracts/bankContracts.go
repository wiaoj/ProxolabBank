package contracts

type CreateBankRequest struct {
	Name string `json:"name"`
}

type BankResponse struct {
	Name      string
	Interests []BankInterestsResponse
}

type BankInterestsResponse struct {
	Interest     float32 `json:"interest"`
	TimeOptionID uint    `json:"time_option_id"`
	CreditTypeID uint    `json:"credit_type_id"`
}
