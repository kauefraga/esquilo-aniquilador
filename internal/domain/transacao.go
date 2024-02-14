package domain

import "time"

type Transacao struct {
	ID          int       `json:"id"`
	Valor       int       `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	ClienteId   int       `json:"cliente_id"`
	RealizadaEm time.Time `json:"realizada_em"`
}

type TransacaoPayload struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type TransacaoResponse struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}
