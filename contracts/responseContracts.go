package contracts

type SingleResponse struct {
	Message string      `json:"message"`
	Item    interface{} `json:"item"`
}

type MultipleResponse struct {
	Message string `json:"message"`
	Items   any    `json:"items"`
}
