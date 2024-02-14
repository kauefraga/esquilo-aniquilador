package database

import (
	"github.com/kauefraga/esquilo-aniquilador/internal/domain"
)

var Clientes map[int]domain.Cliente
var Transacoes map[int]domain.Transacao

func InsereClientesIniciais() {
	Clientes[1] = domain.Cliente{
		ID:     1,
		Nome:   "o barato sai caro",
		Limite: 1000 * 100,
		Saldo:  0,
	}
	Clientes[2] = domain.Cliente{
		ID:     2,
		Nome:   "zan corp ltda",
		Limite: 800 * 100,
		Saldo:  0,
	}
	Clientes[3] = domain.Cliente{
		ID:     3,
		Nome:   "les cruders",
		Limite: 10000 * 100,
		Saldo:  0,
	}
	Clientes[4] = domain.Cliente{
		ID:     4,
		Nome:   "padaria joia de cocaia",
		Limite: 100000 * 100,
		Saldo:  0,
	}
	Clientes[5] = domain.Cliente{
		ID:     5,
		Nome:   "kid mais",
		Limite: 5000 * 100,
		Saldo:  0,
	}
}
