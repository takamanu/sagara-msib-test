package services

import (
	"sagara-msib-test/internal/entities"
	"sagara-msib-test/internal/repositories"
)

type InventoryBajuServices interface {
	CreateBaju(baju *entities.Baju) (err error)
}

type inventoryBajuServices struct {
	bajuRepo repositories.InventoryBajuRepository
}

func NewInventoryBajuService(r repositories.InventoryBajuRepository) (ibs InventoryBajuServices) {
	ibs = &inventoryBajuServices{
		bajuRepo: r,
	}

	return ibs
}

func (ibs *inventoryBajuServices) CreateBaju(baju *entities.Baju) (err error) {
	err = ibs.bajuRepo.Save(baju)
	if err != nil {
		return err
	}

	return err
}
