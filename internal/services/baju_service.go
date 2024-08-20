package services

import (
	"log"
	"sagara-msib-test/internal/entities"
	"sagara-msib-test/internal/repositories"
)

type BajuServices interface {
	CreateBaju(baju entities.Baju) error
	GetBajuByID(id int) (entities.Baju, error)
	GetAllBaju() ([]entities.Baju, error)
	UpdateBaju(id int, baju entities.Baju) error
	DeleteBaju(id int) error
}

type bajuServices struct {
	bajuRepo repositories.BajuRepository
}

func NewInventoryBajuService(r repositories.BajuRepository) (ibs BajuServices) {
	ibs = &bajuServices{
		bajuRepo: r,
	}

	return ibs
}

func (ibs *bajuServices) CreateBaju(baju entities.Baju) (err error) {
	log.Printf("[LOG][Service] Nama Baju Request : %v\n", baju.Nama)
	err = ibs.bajuRepo.Create(baju)

	if err != nil {
		return err
	}

	return err
}

func (s *bajuServices) GetBajuByID(id int) (entities.Baju, error) {
	return s.bajuRepo.GetByID(id)
}

func (s *bajuServices) GetAllBaju() ([]entities.Baju, error) {
	return s.bajuRepo.GetAll()
}

func (s *bajuServices) UpdateBaju(id int, baju entities.Baju) error {
	return s.bajuRepo.Update(id, baju)
}

func (s *bajuServices) DeleteBaju(id int) error {
	return s.bajuRepo.Delete(id)
}
