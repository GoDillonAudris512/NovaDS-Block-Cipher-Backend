package models

type OFBRequest struct {
	TextBitArray	[]int	`json:"text-array"`
	KeyBitArray		[]int	`json:"key-array"`
	InitVector 		[]int	`json:"init-vector"`
	Encrypt			bool	`json:"encrypt"`
}

type OFBResponse struct {
	Success			bool	`json:"success"`
	ResultBitArray	[]int	`json:"result-array"`
}
