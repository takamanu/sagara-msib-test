package repositories

import (
	"database/sql"
	"fmt"
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

	query := `
		INSERT INTO baju (nama, brand, warna, ukuran, harga, stok)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = ibr.db.Exec(
		query,
		baju.Nama,
		baju.Brand,
		baju.Warna,
		baju.Ukuran,
		baju.Harga,
		baju.Stok,
	)
	if err != nil {
		return fmt.Errorf("failed to create baju : %v", err.Error())
	}

	return err
}
