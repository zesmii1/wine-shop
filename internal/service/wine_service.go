package service

import "wine-shop/internal/models"

type WineRepository interface {
	GetAll() ([]models.Wine, error)
	GetById(id uint) (*models.Wine, error)
	Create(wine *models.Wine) error
	Update(id uint, wineEdit *models.WineEdit) error
	Delete(id uint) error
}

type WineService struct {
	repo WineRepository
}

func NewWineService(repo WineRepository) *WineService {
	return &WineService{repo: repo}
}

func (s *WineService) GetAllWines() ([]models.Wine, error) {
	return s.repo.GetAll()
}

func (s *WineService) GetWineById(id uint) (*models.Wine, error) {
	return s.repo.GetById(id)
}

func (s *WineService) CreateWine(name string, year int, price float64, desc string) (*models.Wine, error) {
	wine := &models.Wine{
		Name:        name,
		Year:        year,
		Price:       price,
		Description: desc,
	}
	err := s.repo.Create(wine)
	return wine, err
}

func (s *WineService) UpdateWine(id uint, edit *models.WineEdit) (*models.Wine, error) {
	err := s.repo.Update(id, edit)
	if err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *WineService) DeleteWine(id uint) error {
	return s.repo.Delete(id)
}
