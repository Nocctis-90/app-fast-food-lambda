package casodeuso

import (
	"github.com/nocctis-90/app-fastfood-lambda/app/dominio"
)

type AutenticarUsuario interface {
	AutenticarClienteAnonimo() (string, error)
	AutenticarCliente(cliente *dominio.Cliente) (string, error)
}
