package moneymovement

import "github.com/gmessias/api-go-money/internal/core/domain"

type service struct{}

func New() *service {
	return &service{}
}

func (service *service) CreateMovement(cash domain.Cash, category string) (domain.Cash, error) {
	// Rebendo alguns campos preenchidos

	// Verificar categoria

	// Consultar categorias que existem

	// Atribuir ao modelo

	return domain.Cash{}, nil
}
