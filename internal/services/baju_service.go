package services

import (
	"context"
	"log"
	"sagara-msib-test/internal/entities"
	"sagara-msib-test/internal/repositories"
)

type BajuServices interface {
	CreateBaju(ctx context.Context, baju entities.Baju) error
	GetBajuByID(ctx context.Context, id int) (entities.Baju, error)
	GetAllBaju(ctx context.Context) ([]entities.Baju, error)
	UpdateBaju(ctx context.Context, id int, baju entities.Baju) error
	DeleteBaju(ctx context.Context, id int) error
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

func (ibs *bajuServices) CreateBaju(ctx context.Context, baju entities.Baju) (err error) {
	log.Printf("[LOG][Service] Nama Baju Request : %v\n", baju.Nama)
	err = ibs.bajuRepo.Create(ctx, baju)

	if err != nil {
		return err
	}

	return err
}

func (s *bajuServices) GetBajuByID(ctx context.Context, id int) (entities.Baju, error) {
	return s.bajuRepo.GetByID(ctx, id)
}

func (s *bajuServices) GetAllBaju(ctx context.Context) ([]entities.Baju, error) {
	return s.bajuRepo.GetAll(ctx)
}

func (s *bajuServices) UpdateBaju(ctx context.Context, id int, baju entities.Baju) error {
	return s.bajuRepo.Update(ctx, id, baju)
}

func (s *bajuServices) DeleteBaju(ctx context.Context, id int) error {
	return s.bajuRepo.Delete(ctx, id)
}
