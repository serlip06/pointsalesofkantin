package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Produk struct {
	IDProduk    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Produk string             `bson:"nama_produk,omitempty" json:"nama_produk,omitempty"`
	Deskripsi   string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Harga       int                `bson:"harga,omitempty" json:"harga,omitempty"`
	Gambar      string             `bson:"gambar,omitempty" json:"gambar,omitempty"`
	Stok        int                `bson:"stok,omitempty" json:"stok,omitempty"`
	Kategori    string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

// struct transaksi
type Transaksi struct {
	IDTransaksi      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`         // ID unik transaksi
	IDUser           primitive.ObjectID `bson:"id_user" json:"id_user"`                     // Referensi ke ID pengguna
	IDCartItem       []primitive.ObjectID `bson:"id_cartitem" json:"id_cartitem"`           // Referensi ke ID CartItem
	TotalHarga       int                `bson:"total_harga" json:"total_harga"`             // Total harga (dihitung otomatis)
	CreatedAt        time.Time          `bson:"created_at" json:"created_at"`               // Waktu transaksi
	MetodePembayaran string             `bson:"metode_pembayaran" json:"metode_pembayaran"` // Metode pembayaran
	BuktiPembayaran  string             `bson:"bukti_pembayaran" json:"bukti_pembayaran"`   // Bukti pembayaran
	Status           string             `bson:"status" json:"status"`                       // Status transaksi
	Alamat           string             `bson:"alamat,omitempty" json:"alamat,omitempty"`   // Alamat pengiriman
}

//transaksi request 
type TransaksiRequest struct {
	IDUser           primitive.ObjectID   `json:"id_user" binding:"required"`
	IDCartItems      []primitive.ObjectID `json:"id_cartitems" binding:"required"`
	MetodePembayaran string               `json:"metode_pembayaran" binding:"required"`
	BuktiPembayaran  string               `json:"bukti_pembayaran"`
	Status           string               `json:"status" binding:"required"`
	Alamat           string               `json:"alamat"`
}

type Customer struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Email        []string           `bson:"email,omitempty" json:"email"`
}

// struct untuk item dalam keranjang
type CartItem struct {
	IDCartItem  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`             // ID unik untuk item keranjang
	IDProduk    primitive.ObjectID `bson:"id_produk,omitempty" json:"id_produk,omitempty"` // Referensi ke ID Produk
	IDUser      primitive.ObjectID `bson:"id_user" json:"id_user"`
	Nama_Produk string             `bson:"nama_produk,omitempty" json:"nama_produk,omitempty"` //nama untuk produknya
	Harga       int                `bson:"harga,omitempty" json:"harga,omitempty"`             // Harga produk pada saat dimasukkan ke keranjang
	Quantity    int                `bson:"quantity,omitempty" json:"quantity,omitempty"`       // Jumlah produk dalam keranjang
	SubTotal    int                `bson:"sub_total,omitempty" json:"sub_total,omitempty"`     // Total harga (Harga * Quantity)
	Gambar      string             `bson:"gambar,omitempty" json:"gambar,omitempty"`           // Gambar produk
	IsSelected  bool               `bson:"is_selected,omitempty" json:"is_selected,omitempty"` // Tambahkan flag ini
	
}

type Pesanan struct {
	IDPesanan    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID           primitive.ObjectID `bson:"id_customer,omitempty" json:"id,omitempty"` // idnya yang ada di customer
	Produk       []Produk           `bson:"produk,omitempty" json:"produk,omitempty"`
	TotalHarga   int                `bson:"total_harga,omitempty" json:"total_harga,omitempty"`
	Status       string             `bson:"status,omitempty" json:"status,omitempty"` // Contoh: "Menunggu Konfirmasi", "Dikonfirmasi", "Selesai"
	TanggalPesan string             `bson:"tanggal_pesan,omitempty" json:"tanggal_pesan,omitempty"`
}

// stuct untuk login register
type User struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"` // ID unik dari MongoDB
	Username  string    `bson:"username" json:"username"`          // Username pengguna
	Password  string    `bson:"password" json:"password"`          // Password dalam bentuk hash
	Role      string    `bson:"role" json:"role"`                  // Peran pengguna (admin, customer, kasir, operator)
	CreatedAt time.Time `bson:"created_at" json:"created_at"`      // Waktu pembuatan akun
}

type PendingRegistration struct {
	ID          string    `bson:"_id,omitempty" json:"id,omitempty"` // ID unik dari MongoDB
	Username    string    `bson:"username" json:"username"`          // Username pengguna
	Password    string    `bson:"password" json:"password"`          // Password dalam bentuk hash
	Role        string    `bson:"role" json:"role"`                  // Peran pengguna (customer, kasir, operator)
	SubmittedAt time.Time `bson:"submitted_at" json:"submitted_at"`  // Waktu registrasi
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"` // Username pengguna
	Password string `json:"password"` // Password pengguna
}

type Response struct {
	Status  string `json:"status"`  // Status operasi (success, error)
	Message string `json:"message"` // Pesan deskripsi
	// Data    any    `json:"data"`    // Data tambahan (opsional)
}

// notifikasi

// Struct untuk notifikasi produk terbaru
// type NewProductNotification struct {
// 	IDProduk    string `json:"id_produk"`   // ID unik produk
// 	Nama_Produk string `json:"nama_produk"` // Nama produk baru
// 	Kategori    string `json:"kategori"`    // Kategori produk
// 	AddedBy     string `json:"added_by"`    // Admin/Operator yang menambahkan produk
// 	Timestamp   string `json:"timestamp"`   // Waktu produk ditambahkan
// 	Message     string `json:"message"`     // Pesan notifikasi
// }

// Struct untuk notifikasi stok menipis
// type LowStockNotification struct {
// 	IDProduk     string `json:"id_produk"`     // ID unik produk
// 	Nama_Produk  string `json:"nama_produk"`   // Nama produk
// 	CurrentStock int    `json:"current_stock"` // Jumlah stok saat ini
// 	Threshold    int    `json:"threshold"`     // Ambang batas stok
// 	NotifiedAt   string `json:"notified_at"`   // Waktu notifikasi dibuat
// 	Message      string `json:"message"`       // Pesan notifikasi
// }

// Struct untuk notifikasi stok habis
// type OutOfStockNotification struct {
// 	IDProduk    string `json:"id_produk"`   // ID unik produk
// 	Nama_Produk string `json:"nama_produk"` // Nama produk baru
// 	NotifiedAt  string `json:"notified_at"` // Waktu notifikasi dibuat
// 	Message     string `json:"message"`     // Pesan notifikasi
// }
// type Pelanggan struct {
// 	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
// 	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
// 	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
// 	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
// 	Email        []string           `bson:"email,omitempty" json:"email,omitempty"`
// }
