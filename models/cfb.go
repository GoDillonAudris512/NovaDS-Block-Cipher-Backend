package models

type CFBRequest struct {
	TextBitArray	[]int	`json:"text-array"`
	KeyBitArray		[]int	`json:"key-array"`
	InitVector 		[]int	`json:"init-vector"`
	Encrypt			bool	`json:"encrypt"`
}

type CFBResponse struct {
	Success			bool	`json:"success"`
	ResultBitArray	[]int	`json:"result-array"`
}