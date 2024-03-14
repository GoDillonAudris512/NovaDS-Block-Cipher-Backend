package models

type ECBRequest struct {
	TextBitArray	[]int	`json:"text-array"`
	KeyBitArray		[]int	`json:"key-array"`
	Encrypt			bool	`json:"encrypt"`
}

type ECBResponse struct {
	Success			bool	`json:"success"`
	ResultBitArray	[]int	`json:"result-array"`
}
