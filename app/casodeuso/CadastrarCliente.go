package casodeuso

import (
	"github.com/nocctis-90/app-fastfood-lambda/app/apresentacao"
	"github.com/nocctis-90/app-fastfood-lambda/app/dominio"
)

// CadastrarCliente é a interface que define o caso de uso de cadastro de cliente
type CadastrarCliente interface {
	CadastrarCliente(inputCliente apresentacao.ClienteDTO) (*dominio.Cliente, error)
}
