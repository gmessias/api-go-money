package ports

import (
	"github.com/gmessias/api-go-money/internal/core/domain"
	"gorm.io/gorm"
)

type CashRepository interface {
	GetAllCash(cashList *[]domain.Cash) *gorm.DB
	GetCashPerId(cash *domain.Cash, id string) *gorm.DB
	CreateCash(cash *domain.Cash) *gorm.DB
	UpdateCash(cash *domain.Cash) *gorm.DB
	DeleteCash(cash *domain.Cash, id string) *gorm.DB
}

type CategoryRepository interface {
	GetAllCategories()
	GetCategoryPerId()
	CreateCategory()
	UpdateCategory()
	DeleteCategory()
}
