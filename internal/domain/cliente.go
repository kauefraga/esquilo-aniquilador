package domain

type Cliente struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Limite int    `json:"limite"`
	Saldo  int    `json:"saldo"`
}
