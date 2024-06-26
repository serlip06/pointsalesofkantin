package _714220023

import (
	"fmt"
	"testing"

	//"github.com/rogpeppe/go-internal/module"
	//"github.com/serlip06/pointsalesofkantin/model"
	"github.com/serlip06/pointsalesofkantin/module"

	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

func TestInsertPelanggan(t *testing.T) {
	nama := "askara dirgantara"
	phoneNumber := "0812-9098-8680"
	alamat := " perumahan dirgantara "
	email := []string{"akssa_88@gmail.com", "aska__89@gmail.com"}
	insertedID := module.InsertPelanggan(nama, phoneNumber, alamat, email)
	fmt.Println(insertedID)
}

// func TestGetPelangganByID(t *testing.T) {
// 	pelangganID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}
// 	pelanggan := module.GetPelangganByID(pelangganID)
// 	fmt.Println(pelanggan)
// }
func TestGetPelangganByID(t *testing.T) {
	pelangganID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
	if err != nil {
		fmt.Printf("error converting id to ObjectID: %v\n", err)
		return
	}

	// Asumsikan module.MongoConn telah diinisialisasi sebelumnya
	pelanggan, err := module.GetPelangganByID(pelangganID, module.MongoConn, "kantin_pelanggan")
	if err != nil {
		fmt.Printf("error calling GetPelangganByID: %v\n", err)
		return
	}
	fmt.Println(pelanggan)
}


func TestGetAllPelanggan(t *testing.T) {
	pelanggans := module.GetAllPelanggan()
	fmt.Println(pelanggans)
}

func TestInsertProduk(t *testing.T) {
	namaProduk := "lumpiah basah"
	deskripsi := "ayam bakar dengan berbagai bumbu  "
	harga := 16000
	insertedID := module.InsertProduk(namaProduk, deskripsi, harga)
	fmt.Println(insertedID)
}

func TestGetProdukByID(t *testing.T) {
	produkID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	produk := module.GetProdukByID(produkID)
	fmt.Println(produk)
}

func TestGetAllProduk(t *testing.T) {
	produks := module.GetAllProduk()
	fmt.Println(produks)
}

func TestInsertTransaksi(t *testing.T) {
	metodePembayaran := "bayar langsung"
	tanggalWaktu := "2021-11-23 22:00:00"
	insertedID := module.InsertTransaksi(metodePembayaran, tanggalWaktu)
	fmt.Println(insertedID)
}

func TestGetTransaksiByID(t *testing.T) {
	transaksiID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	transaksi := module.GetTransaksiByID(transaksiID)
	fmt.Println(transaksi)
}
func TestGetAllTransaksi(t *testing.T) {
	transaksis := module.GetAllTransaksi()
	fmt.Println(transaksis)
}

//testing delete pelanggan
func TestDeletePelangganByID(t *testing.T) {
	id := "663c6729918275d152c9d488" // ID data yang ingin dihapus id elisabet
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeletePelangganByID(objectID, module.MongoConn, "kantin_pelanggan")
	if err != nil {
		t.Fatalf("error calling DeletePelangganByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPelangganByID
	_, err = module.GetPelangganByID(objectID, module.MongoConn, "kantin_pelanggan")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

//insert data customer 
func TestInsertCustomer(t *testing.T) {
	nama := "Anindya Kirana"
	phoneNumber := "0856-2245-5522"
	alamat := "jl.sarijadi"
	email := []string{"kirana90@gmail.com", "krn_anindya@gmail.com"}
	namaProduk := "ayam geprek"
	deskripsi := "Ayam goreng disajikan dengan sambal pedas dan kerupuk"
	harga := 12000
	gambar := "https://i.pinimg.com/564x/d3/47/b0/d347b0132dcb98af18158cbebd533cc8.jpg"
	stok := "15"

	insertedID := module.InsertCustomer(nama, phoneNumber, alamat, email, namaProduk, deskripsi, harga, gambar, stok)
	fmt.Println(insertedID)
}

// get data customer by id
func TestCustomerFromID(t *testing.T) {
	id := "6682995cb6ea919536290321" 
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	profil, err := module.GetCustomerFromID(objectID, module.MongoConn, "customer")
	if err != nil {
		t.Fatalf("error calling GetCustomerFromID: %v", err)
	}
	fmt.Println(profil)
}

// function get all customer
func TestGetAllCustomer(t *testing.T) {
	customers := module.GetAllCustomer()
	fmt.Println(customers)
}

//function delete customer
func TestDeleteCustomerByID(t *testing.T) {
	id := "6682ac8719e5c29e437eac67" // ID data yang ingin dihapus id anindya kirana
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteCustomerByID(objectID, module.MongoConn, "customer")
	if err != nil {
		t.Fatalf("error calling DeleteCustomerByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPelangganByID
	_, err = module.GetCustomerFromID(objectID, module.MongoConn, "customer")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}