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
	nama := "Aksara Dirgantara "
	phoneNumber := "08577908763"
	alamat := "Jalan Dirgantara No. 50"
	email := []string{"Aksara90@example.com", "Askara67@example.com"}
	insertedID := module.InsertPelanggan(nama, phoneNumber, alamat, email)
	fmt.Println(insertedID)
}

func TestGetPelangganByID(t *testing.T) {
	pelangganID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	pelanggan := module.GetPelangganByID(pelangganID)
	fmt.Println(pelanggan)
}
func TestGetAllPelanggan(t *testing.T) {
	pelanggans := module.GetAllPelanggan()
	fmt.Println(pelanggans)
}

func TestInsertProduk(t *testing.T) {
	namaProduk := "Martabak Telor"
	deskripsi := "Martabak Asin melipah akan toping dagingnya"
	harga := 25000
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
	metodePembayaran := "Cash On Delivery"
	tanggalWaktu := "2024-03-28 10:00:00"
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
