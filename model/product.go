package model

type Product struct {
	Deskripsi string `json:"id" validate:"required"`
	Img       string `json:"img" validate:"required"`
	Harga     int    `json:"harga" validate:"required"`
	Jenis     string `json:"jenis" validate:"required"`
	Nama      string `json:"nama" validate:"required"`
	Rating    int    `json:"rating"`
}
