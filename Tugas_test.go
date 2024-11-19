package _714220023

import (
	"fmt"
	"testing"

	//"github.com/rogpeppe/go-internal/module"
	"github.com/serlip06/pointsalesofkantin/module"

	 "go.mongodb.org/mongo-driver/bson/primitive"
	
)

// func TestInsertPelanggan(t *testing.T) {
// 	nama := "askara dirgantara"
// 	phoneNumber := "0812-9098-8680"
// 	alamat := " perumahan dirgantara "
// 	email := []string{"akssa_88@gmail.com", "aska__89@gmail.com"}
// 	insertedID := module.InsertPelanggan(nama, phoneNumber, alamat, email)
// 	fmt.Println(insertedID)
// }

// // func TestGetPelangganByID(t *testing.T) {
// // 	pelangganID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
// // 	if err != nil {
// // 		fmt.Printf("Error: %v\n", err)
// // 		return
// // 	}
// // 	pelanggan := module.GetPelangganByID(pelangganID)
// // 	fmt.Println(pelanggan)
// // }
// func TestGetPelangganByID(t *testing.T) {
// 	pelangganID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
// 	if err != nil {
// 		fmt.Printf("error converting id to ObjectID: %v\n", err)
// 		return
// 	}

// 	// Asumsikan module.MongoConn telah diinisialisasi sebelumnya
// 	pelanggan, err := module.GetPelangganByID(pelangganID, module.MongoConn, "kantin_pelanggan")
// 	if err != nil {
// 		fmt.Printf("error calling GetPelangganByID: %v\n", err)
// 		return
// 	}
// 	fmt.Println(pelanggan)
// }


// func TestGetAllPelanggan(t *testing.T) {
// 	pelanggans := module.GetAllPelanggan()
// 	fmt.Println(pelanggans)
// }

// func TestInsertProduk(t *testing.T) {
// 	namaProduk := "lumpiah basah"
// 	deskripsi := "ayam bakar dengan berbagai bumbu  "
// 	harga := 16000
// 	insertedID := module.InsertProduk(namaProduk, deskripsi, harga)
// 	fmt.Println(insertedID)
// }

// func TestGetProdukByID(t *testing.T) {
// 	produkID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}
// 	produk := module.GetProdukByID(produkID)
// 	fmt.Println(produk)
// }

// func TestGetAllProduk(t *testing.T) {
// 	produks := module.GetAllProduk()
// 	fmt.Println(produks)
// }

// func TestInsertTransaksi(t *testing.T) {
// 	metodePembayaran := "bayar langsung"
// 	tanggalWaktu := "2021-11-23 22:00:00"
// 	insertedID := module.InsertTransaksi(metodePembayaran, tanggalWaktu)
// 	fmt.Println(insertedID)
// }

// func TestGetTransaksiByID(t *testing.T) {
// 	transaksiID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}
// 	transaksi := module.GetTransaksiByID(transaksiID)
// 	fmt.Println(transaksi)
// }
// func TestGetAllTransaksi(t *testing.T) {
// 	transaksis := module.GetAllTransaksi()
// 	fmt.Println(transaksis)
// }

// //testing delete pelanggan
// func TestDeletePelangganByID(t *testing.T) {
// 	id := "663c6729918275d152c9d488" // ID data yang ingin dihapus id elisabet
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		t.Fatalf("error converting id to ObjectID: %v", err)
// 	}

// 	err = module.DeletePelangganByID(objectID, module.MongoConn, "kantin_pelanggan")
// 	if err != nil {
// 		t.Fatalf("error calling DeletePelangganByID: %v", err)
// 	}

// 	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPelangganByID
// 	_, err = module.GetPelangganByID(objectID, module.MongoConn, "kantin_pelanggan")
// 	if err == nil {
// 		t.Fatalf("expected data to be deleted, but it still exists")
// 	}
// }

//insert data customer 
func TestInsertCustomer(t *testing.T) {
	nama := "Tom Holland"
	phoneNumber := "0856-2245-5522"
	alamat := "jl.dago pakar"
	email := []string{"tomhllnd@gmail.com","holandtom22@gmail.com"}
	
	insertedID := module.InsertCustomer(nama, phoneNumber, alamat, email)
	fmt.Println(insertedID)
}

// get data customer by id
func TestCustomerFromID(t *testing.T) {
	id := "673c9e05adbfb49a59ab07c1" 
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
// func TestGetCustomerByID(t *testing.T) {
// 	customerID, err := primitive.ObjectIDFromHex("6682995cb6ea919536290321")
// 	if err != nil {
// 		fmt.Printf("error converting id to ObjectID: %v\n", err)
// 		return
// 	}

// 	// Asumsikan module.MongoConn telah diinisialisasi sebelumnya
// 	customer, err := module.GetCustomerByID(customerID, module.MongoConn, "customer")
// 	if err != nil {
// 		fmt.Printf("error calling GetCustomerByID: %v\n", err)
// 		return
// 	}
// 	fmt.Println(customer)
// }

// //insertbarang



// // Mock InsertOneDoc function for testing purposes
// func InsertOneDoc(db string, collection string, doc interface{}) interface{} {
// 	// Mocked implementation, you can add logic to validate input or return a mock result
// 	return doc
// }

// func InsertBarang(namaProduk string, deskripsi string, harga int, gambar string, stok int) interface{} {
// 	// Implementation of InsertBarang function
// 	// You can add your logic here
// 	return nil
// }

// func TestInsertBarang(t *testing.T) {
// 	// Data untuk pengujian
// 	namaProduk := "ayam geprek"
// 	deskripsi := "Ayam goreng disajikan dengan sambal pedas dan kerupuk"
// 	harga := 12000
// 	gambar := "https://i.pinimg.com/564x/d3/47/b0/d347b0132dcb98af18158cbebd533cc8.jpg"
// 	stok := 15



// 	fmt.Println(barang.ID_barang)
// }

func TestInsertDataProduk(t *testing.T) {
	nama_produk := "Ikan bakar "
	deskripsi := "Ikan bakar yang enak "
	gambar := "https://i.pinimg.com/736x/a9/c1/86/a9c186a2a4c0202f2bef835b3fbf5327.jpg"
	harga := 25000
	stok := 15
	insertedID := module.InsertDataProduk(nama_produk, deskripsi, harga, gambar, stok)
	fmt.Println(insertedID)
}

//test buat delete data produknya 
func TestDeleteProduksByID(t *testing.T) {
	id := "6693f0a667dbd67851b4f04c" // ID data yang ingin dihapus id ayam bakar
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteCustomerByID(objectID, module.MongoConn, "produk")
	if err != nil {
		t.Fatalf("error calling DeleteCustomerByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan Getprodukfrom id 
	_, err = module.GetProduksFromID(objectID, module.MongoConn, "produk")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}
//function get all produks
func  TestGetAllProduks(t *testing.T) {
	produks := module.GetAllProduks()
	fmt.Println(produks)
}
//get produk by id 
func TestProduksFromID(t *testing.T) {
	id := "669cbbf56733fd391c96bb86" // idnya mie pedas
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	profil, err := module.GetProduksFromID(objectID, module.MongoConn, "produk") // Perbaiki pemanggilan fungsi
	if err != nil {
		t.Fatalf("error calling GetProduksFromID: %v", err)
	}
	fmt.Println(profil)
}
