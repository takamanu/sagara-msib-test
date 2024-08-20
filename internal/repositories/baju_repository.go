package repositories

import (
	"database/sql"
	"log"
	"sagara-msib-test/internal/entities"
)

type InventoryBajuRepository interface {
	Save(baju *entities.Baju) (err error)
}

type inventoryBajuRepository struct {
	db *sql.DB
}

func NewInventoryBajuRepository(db *sql.DB) (ibr InventoryBajuRepository) {
	ibr = &inventoryBajuRepository{
		db: db,
	}
	return ibr
}

func (ibr *inventoryBajuRepository) Save(baju *entities.Baju) (err error) {
	log.Printf("[LOG][Repository] Nama Baju Request : %v\n", baju.Nama)

	return err
}
