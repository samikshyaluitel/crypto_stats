package model

// Symbols for currencies
type Symbol struct {
	ID          string `json:"id"`
	FeeCurrency string `json:"feeCurrency"`
	Bid         string `json:"bid"`
}
