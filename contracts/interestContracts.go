package contracts

type CreateInterestRequest struct {
	BankID       uint    `json:"bank_id"`
	Interest     float32 `json:"interest"`
	TimeOptionID uint    `json:"time_option_id"`
	CreditTypeID uint    `json:"credit_type_id"`
}

type InterestResponse struct {
	BankID                uint    `json:"bank_id"`
	BankName              string  `json:"bank_name"`
	Interest              float32 `json:"interest"`
	TimeOptionID          uint    `json:"time_option_id"`
	TimeOptionDescription string  `json:"time_option_description"`
	CreditTypeID          uint    `json:"credit_type_id"`
	CreditTypeDescription string  `json:"credit_type_description"`
}
