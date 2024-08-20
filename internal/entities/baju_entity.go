package entities

import "gorm.io/gorm"

type Baju struct {
	gorm.Model
	ID     int     `json:"id"`
	Nama   string  `json:"nama"`
	Brand  string  `json:"brand"`
	Ukuran string  `json:"ukuran"`
	Stok   int     `json:"stok"`
	Harga  float64 `json:"harga"`
	Warna  string  `json:"warna"`
}
