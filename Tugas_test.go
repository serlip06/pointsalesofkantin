package _714220023

import (
	"fmt"
	"testing"

	"context"
	//"github.com/rogpeppe/go-internal/module"
	"github.com/serlip06/pointsalesofkantin/module"
	"github.com/serlip06/pointsalesofkantin/model"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

// insert data customer
func TestInsertCustomer(t *testing.T) {
	nama := "Tom Holland"
	phoneNumber := "0856-2245-5522"
	alamat := "jl.dago pakar"
	email := []string{"tomhllnd@gmail.com", "holandtom22@gmail.com"}

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

// function delete customer
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

// test untuk produknya
func TestInsertDataProduk(t *testing.T) {
	nama_produk := "Pangsit pedas"
	deskripsi := "Pangsit pedas yang enak"
	gambar := "https://i.pinimg.com/736x/b2/db/a4/b2dba44f2dd0ad242e0575fd17ad94c7.jpg"
	harga := 10000
	stok := 10
	kategori := "Makanan"

	insertedID, err := module.InsertDataProduk(nama_produk, deskripsi, harga, gambar, stok, kategori)
	if err != nil {
		t.Fatalf("error inserting data produk: %v", err)
	}
	fmt.Printf("Inserted ID: %v\n", insertedID)
}

// test buat delete data produknya
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

// function get all produks
func TestGetAllProduks(t *testing.T) {
	kategori := "Makanan" // Ubah sesuai kebutuhan: "Makanan" atau "Minuman"
	produks, err := module.GetAllProduks(kategori)
	if err != nil {
		t.Fatalf("error calling GetAllProduks: %v", err)
	}
	fmt.Println("Produk yang ditemukan:", produks)
}

// get produk by id
func TestProduksFromID(t *testing.T) {
	id := "674b290af85cc65485585875" // idnya bolu 
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

//test untuk masuk keranjang
//insert to cartitme (keranjang)
func TestInsertDataCartItemFunc(t *testing.T) {
    db, err := module.MongoConnectdatabase("kantin") // Nama database yang benar
    assert.NoError(t, err)

    // Gunakan ID produk yang valid dan sesuaikan dengan format ObjectID
    idProduk, err := primitive.ObjectIDFromHex("673c90cd715120ed663eb984") // id produk : ikan bakar
    if err != nil {
        t.Fatalf("Invalid ObjectID: %v", err)
    }

    // Pastikan ID tersebut ada di database
    collection := db.Collection("produk")
    var product model.Produk
    err = collection.FindOne(context.TODO(), bson.M{"_id": idProduk}).Decode(&product)
    if err != nil {
        t.Fatalf("Product with ID %v not found: %v", idProduk, err)
    }

    // Lanjutkan dengan pengujian InsertDataCartItemFunc
    quantity := 2
    result, err := module.InsertDataCartItemFunc(db, idProduk, quantity)

    // Verifikasi hasil
    assert.NoError(t, err)
    assert.NotNil(t, result)
}

// //get cartitemfromid
func TestGetCartItemFromID(t *testing.T) {
    // Setup database dan koleksi
    db, err := module.MongoConnectdatabase("kantin")
    assert.NoError(t, err)

    // ID cart item yang valid, misalnya ID yang sudah ada di database
    id, err := primitive.ObjectIDFromHex("676fb8a26214462f56e4965d") // Ganti dengan ID yang valid
    assert.NoError(t, err) // Tambahkan ini untuk menangani error

    // Panggil fungsi GetCartItemFromID
    cartItem, err := module.GetCartItemFromID(id, db, "cart_items")

    // Verifikasi hasil
    assert.NoError(t, err)
    assert.NotNil(t, cartItem)
    assert.Equal(t, id, cartItem.IDCartItem) // Sesuaikan atribut dengan nama di struct Anda
    // assert.Equal(t, "NamaProduk yang diharapkan", cartItem.NamaProduk) // Tambahkan verifikasi field lain jika perlu
}

// //get all cartitem
func TestGetAllCartItems(t *testing.T) {
	cartitems := module.GetAllCartItems()
	// Menggunakan %v untuk mencetak slice dengan format default
	fmt.Printf("%v\n", cartitems)
}

// deletecartitem
func TestDeleteCartItemFromID(t *testing.T) {
	id := "6770f80a419da98516ba7db1" //id yang akan dihapus ini pake id yang ikan bakar
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	err = module.DeleteCartItemByID(objectID, module.MongoConn, "cart_items")
	if err != nil {
		t.Fatalf("error calling DeleteCartItemFromID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan Getprodukfrom id
	_, err = module.GetCartItemFromID(objectID, module.MongoConn, "cart_items")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

// test untuk pesanan 

