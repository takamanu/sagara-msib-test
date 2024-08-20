package repositories

import (
	"database/sql"
	"sagara-msib-test/internal/entities"
)

type BajuRepository interface {
	Create(baju entities.Baju) error
	GetByID(id int) (entities.Baju, error)
	GetAll() ([]entities.Baju, error)
	Update(id int, baju entities.Baju) error
	Delete(id int) error
}

type bajuRepository struct {
	db *sql.DB
}

func NewInventoryBajuRepository(db *sql.DB) (ibr BajuRepository) {
	ibr = &bajuRepository{
		db: db,
	}
	return ibr
}

func (ibr *bajuRepository) Create(baju entities.Baju) error {
	_, err := ibr.db.Exec(`INSERT INTO baju (warna, ukuran, harga, stok, nama, brand) VALUES ($1, $2, $3, $4, $5, $6)`,
		baju.Warna, baju.Ukuran, baju.Harga, baju.Stok, baju.Nama, baju.Brand)
	return err
}

func (ibr *bajuRepository) GetByID(id int) (entities.Baju, error) {
	row := ibr.db.QueryRow(`SELECT id, warna, ukuran, harga, stok, nama, brand FROM baju WHERE id = $1`, id)
	var baju entities.Baju
	err := row.Scan(&baju.ID, &baju.Warna, &baju.Ukuran, &baju.Harga, &baju.Stok, &baju.Nama, &baju.Brand)
	if err != nil {
		return entities.Baju{}, err
	}
	return baju, nil
}

func (ibr *bajuRepository) GetAll() ([]entities.Baju, error) {
	rows, err := ibr.db.Query(`SELECT id, warna, ukuran, harga, stok, nama, brand FROM baju`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bajuList []entities.Baju
	for rows.Next() {
		var baju entities.Baju
		if err := rows.Scan(&baju.ID, &baju.Warna, &baju.Ukuran, &baju.Harga, &baju.Stok, &baju.Nama, &baju.Brand); err != nil {
			return nil, err
		}
		bajuList = append(bajuList, baju)
	}
	return bajuList, nil
}

func (r *bajuRepository) Update(id int, baju entities.Baju) error {
	_, err := r.db.Exec(`UPDATE baju SET warna = $1, ukuran = $2, harga = $3, stok = $4, nama = $5, brand = $6 WHERE id = $7`,
		baju.Warna, baju.Ukuran, baju.Harga, baju.Stok, baju.Nama, baju.Brand, id)
	return err
}

func (r *bajuRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM baju WHERE id = $1`, id)
	return err
}
