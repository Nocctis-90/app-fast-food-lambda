// repositories/cliente_repository.go
package repositorio

import (
	"github.com/nocctis-90/app-fastfood-lambda/app/dominio"
)

type RepositorioCliente interface {
	SalvarOuAtualizarCliente(cliente *dominio.Cliente) error
	BuscarClientePorID(idCliente string) (*dominio.Cliente, error)
}
