package domain

// Limite e saldo em centavos
type Cliente struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Limite int    `json:"limite"`
	Saldo  int    `json:"saldo"`
}
