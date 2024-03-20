package models

type CounterRequest struct {
	TextBitArray	[]int	`json:"text-array"`
	KeyBitArray		[]int	`json:"key-array"`
	Encrypt			bool	`json:"encrypt"`
}

type CounterResponse struct {
	Success			bool	`json:"success"`
	ResultBitArray	[]int	`json:"result-array"`
}
