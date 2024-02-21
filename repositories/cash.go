package repositories

import (
	"github.com/gmessias/api-go-money/database"
	"github.com/gmessias/api-go-money/internal/core/domain"
	"gorm.io/gorm"
)

type CashRepositoryImpl struct{}

func (cr *CashRepositoryImpl) GetAllCash(cashList *[]domain.Cash) *gorm.DB {
	result := database.DB.Find(cashList)

	return result
}

func (cr *CashRepositoryImpl) GetCashPerId(cash *domain.Cash, id string) *gorm.DB {
	result := database.DB.First(cash, id)

	return result
}

func (cr *CashRepositoryImpl) CreateCash(cash domain.Cash) *gorm.DB {
	result := database.DB.Create(&cash)

	return result
}

func (cr *CashRepositoryImpl) UpdateCash() *gorm.DB {

}

func (cr *CashRepositoryImpl) DeleteCash(cash domain.Cash, id string) *gorm.DB {
	result := database.DB.Delete(&cash, id)

	return result
}
