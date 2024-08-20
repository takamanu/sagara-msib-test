package repositories

import (
	"context"
	"gorm.io/gorm"
	"sagara-msib-test/internal/entities"
)

type BajuRepository interface {
	Create(ctx context.Context, baju entities.Baju) error
	GetByID(ctx context.Context, id int) (entities.Baju, error)
	GetAll(ctx context.Context) ([]entities.Baju, error)
	Update(ctx context.Context, id int, baju entities.Baju) error
	Delete(ctx context.Context, id int) error
}

type bajuRepository struct {
	db *gorm.DB
}

func NewInventoryBajuRepository(db *gorm.DB) (ibr BajuRepository) {
	ibr = &bajuRepository{
		db: db,
	}
	return ibr
}

func (br *bajuRepository) Create(ctx context.Context, baju entities.Baju) (err error) {
	err = br.db.WithContext(ctx).Create(baju).Error
	return err
}

func (br *bajuRepository) GetByID(ctx context.Context, id int) (baju entities.Baju, err error) {
	err = br.db.WithContext(ctx).First(&baju, id).Error
	return baju, err
}

func (br *bajuRepository) GetAll(ctx context.Context) (bajuList []entities.Baju, err error) {
	err = br.db.WithContext(ctx).Find(&bajuList).Error
	return bajuList, err
}

func (br *bajuRepository) Update(ctx context.Context, id int, updatedBaju entities.Baju) (err error) {
	var (
		baju entities.Baju
	)

	if err = br.db.WithContext(ctx).First(&baju, id).Error; err != nil {
		return err
	}

	baju.Nama = updatedBaju.Nama
	baju.Brand = updatedBaju.Brand
	baju.Warna = updatedBaju.Warna
	baju.Ukuran = updatedBaju.Ukuran
	baju.Harga = updatedBaju.Harga
	baju.Stok = updatedBaju.Stok

	err = br.db.WithContext(ctx).Save(&baju).Error

	return err
}

func (br *bajuRepository) Delete(ctx context.Context, id int) (err error) {
	err = br.db.WithContext(ctx).Delete(&entities.Baju{}, id).Error

	return err
}
