package repository

import (
	"gorm.io/gorm"
	"wine-shop/internal/models"
)

type WineRepositoryImpl struct {
	db *gorm.DB
}

func NewWineRepository(db *gorm.DB) *WineRepositoryImpl {
	return &WineRepositoryImpl{db: db}
}

func (r *WineRepositoryImpl) GetAll() ([]models.Wine, error) {
	var wines []models.Wine
	err := r.db.Find(&wines).Error
	return wines, err
}

func (r *WineRepositoryImpl) GetById(id uint) (*models.Wine, error) {
	var wine models.Wine
	err := r.db.First(&wine, id).Error
	return &wine, err
}

func (r *WineRepositoryImpl) Create(wine *models.Wine) error {
	return r.db.Create(wine).Error
}

func (r *WineRepositoryImpl) Update(id uint, wineEdit *models.WineEdit) error {
	return r.db.Model(&models.Wine{}).Where("id = ?", id).Updates(wineEdit).Error
}

func (r *WineRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Wine{}, id).Error
}
