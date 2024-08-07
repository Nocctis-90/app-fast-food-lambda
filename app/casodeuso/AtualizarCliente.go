package casodeuso

import (
	"github.com/nocctis-90/app-fastfood-lambda/app/dominio"
)

type AtualizarCliente interface {
	AtualizarCliente(inputCliente *dominio.Cliente) (*dominio.Cliente, error)
}
