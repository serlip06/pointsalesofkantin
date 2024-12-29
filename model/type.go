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

// struct untuk item dalam keranjang
type CartItem struct {
	IDCartItem primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`               // ID unik untuk item keranjang
	IDProduk   primitive.ObjectID `bson:"id_produk,omitempty" json:"id_produk,omitempty"`   // Referensi ke ID Produk
	NamaProduk string             `bson:"NamaProduk,omitempty" json:"NamaProduk,omitempty"` //nama untuk produknya
	Harga      int                `bson:"harga,omitempty" json:"harga,omitempty"`           // Harga produk pada saat dimasukkan ke keranjang
	Quantity   int                `bson:"quantity,omitempty" json:"quantity,omitempty"`     // Jumlah produk dalam keranjang
	SubTotal   int                `bson:"sub_total,omitempty" json:"sub_total,omitempty"`   // Total harga (Harga * Quantity)
}

type Pesanan struct {
	IDPesanan    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID           primitive.ObjectID `bson:"id_customer,omitempty" json:"id,omitempty"` // idnya yang ada di customer
	Produk       []Produk           `bson:"produk,omitempty" json:"produk,omitempty"`
	TotalHarga   int                `bson:"total_harga,omitempty" json:"total_harga,omitempty"`
	Status       string             `bson:"status,omitempty" json:"status,omitempty"` // Contoh: "Menunggu Konfirmasi", "Dikonfirmasi", "Selesai"
	TanggalPesan string             `bson:"tanggal_pesan,omitempty" json:"tanggal_pesan,omitempty"`
}
