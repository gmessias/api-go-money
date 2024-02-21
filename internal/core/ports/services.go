package ports

import "github.com/gmessias/api-go-money/internal/core/domain"

type MoneyMovementService interface {
	CreateMovement(cash domain.Cash, category string) (domain.Cash, error)
}
