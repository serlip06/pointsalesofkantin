package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pelanggan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Email        []string           `bson:"email,omitempty" json:"email,omitempty"`
}

type Produk struct {
	IDProduk    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Produk string             `bson:"nama_produk,omitempty" json:"nama_produk,omitempty"`
	Deskripsi   string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Harga       int                `bson:"harga,omitempty" json:"harga,omitempty"`
	Gambar      string             `bson:"gambar,omitempty" json:"gambar,omitempty"`
	Stok        int                `bson:"stok,omitempty" json:"stok,omitempty"`
	Kategori    string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
}

type Transaksi struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Metode_Pembayaran string             `bson:"metode_pembayaran,omitempty" json:"metode_pembayaran,omitempty"`
	Tanggal_Waktu     string             `bson:"tanggal_waktu,omitempty" json:"tanggal_waktu,omitempty"`
}

type Customer struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Email        []string           `bson:"email,omitempty" json:"email"`
}

type Barang struct {
	ID_barang   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Produk string             `bson:"nama_produk,omitempty" json:"nama_produk,omitempty"`
	Deskripsi   string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Harga       int                `bson:"harga,omitempty" json:"harga,omitempty"`
	Gambar      string             `bson:"gambar,omitempty" json:"gambar,omitempty"`
	Stok        int                `bson:"stok,omitempty" json:"stok,omitempty"`
}
