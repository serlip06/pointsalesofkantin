package _714220023

import (
	"fmt"
	"testing"

	"context"
	//"github.com/rogpeppe/go-internal/module"
	"github.com/serlip06/pointsalesofkantin/model"
	"github.com/serlip06/pointsalesofkantin/module"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
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
	nama_produk := "Cireng"
	deskripsi := "Cireng pedas yang enak"
	gambar := "https://asset.kompas.com/crops/AN59bpBBYFwpIg7Xm8LFx16ldpk=/0x0:698x465/1200x800/data/photo/2024/04/12/6618d397934e0.jpg"
	harga := 5000
	stok := 15
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
	id := "678737e16677cb3541cb0ba6" // idnya bolu
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

// test untuk masuk keranjang
// insert to cartitme (keranjang)
func TestInsertDataCartItemFunc(t *testing.T) {
	db, err := module.MongoConnectdatabase("kantin") // Nama database yang benar
	assert.NoError(t, err)

	// Gunakan ID produk yang valid dan sesuaikan dengan format ObjectID
	idProduk, err := primitive.ObjectIDFromHex("673c90cd715120ed663eb984") // ID produk: ayam bakar
	assert.NoError(t, err)

	// Gunakan ID pengguna yang valid
	idUser, err := primitive.ObjectIDFromHex("678f3dec6c07fa5fb07d8e3a") // Contoh ID pengguna serli
	assert.NoError(t, err)

	// Pastikan ID produk tersebut ada di database
	collection := db.Collection("produk")
	var product model.Produk
	err = collection.FindOne(context.TODO(), bson.M{"_id": idProduk}).Decode(&product)
	assert.NoError(t, err)

	// Lanjutkan dengan pengujian InsertDataCartItemFunc
	quantity := 2
	result, err := module.InsertDataCartItemFunc(db, idProduk, idUser, quantity)

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
	assert.NoError(t, err)                                           // Tambahkan ini untuk menangani error

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
	db, err := module.MongoConnectdatabase("kantin") // Nama database yang benar
	assert.NoError(t, err)

	id := "6770f80a419da98516ba7db1" // ID item keranjang: ikan bakar
	objectID, err := primitive.ObjectIDFromHex(id)
	assert.NoError(t, err)

	// Panggil fungsi DeleteCartItemByID dengan semua parameter yang diperlukan
	err = module.DeleteCartItemByID(objectID, db, db, "cart_items")
	assert.NoError(t, err)

	// Verifikasi bahwa data telah dihapus dengan pengecekan menggunakan GetCartItemFromID
	_, err = module.GetCartItemFromID(objectID, db, "cart_items")
	assert.Error(t, err) // Harus menghasilkan error karena data sudah dihapus
}


// test untuk pesanan

// test untuk login
func TestRegisterHandler(t *testing.T) {
	// Setup test database
	db := module.MongoConnectdb("kantin")

	// Test case input
	req := model.RegisterRequest{
		Username: "testuser",
		Password: "testpassword",
		Role:     "customer",
	}

	// Call the signupHandler function
	message, err := module.RegisterHandler(req, db)

	// Test if there were no errors
	if err != nil {
		t.Errorf("Error in signupHandler: %v", err)
	}

	// Test if the message is as expected
	expectedMessage := "Registration submitted, waiting for admin approval"
	if message != expectedMessage {
		t.Errorf("Expected message: %s, got: %s", expectedMessage, message)
	}

	// Verify if the user was inserted into the database
	collection := db.Collection("pending_registrations")
	var result model.PendingRegistration
	err = collection.FindOne(context.TODO(), bson.M{"username": req.Username}).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to find user in the database: %v", err)
	}

	// Check if the user data is correct
	if result.Username != req.Username {
		t.Errorf("Expected username: %s, got: %s", req.Username, result.Username)
	}

	// Verify if the password is correctly hashed
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password))
	if err != nil {
		t.Fatalf("Password hash mismatch: %v", err)
	}

	// Print confirmation message
	fmt.Printf("User %s successfully registered and saved to the database.\n", req.Username)
}

// UJICOBA APPROVE NYA
func ApproveRegistration(t *testing.T) {
	// Setup test database
	db := module.MongoConnectdb("kantin")

	// Tentukan ID yang sudah ada di koleksi pending_registrations
	// Misalnya ID sudah diketahui sebelumnya
	existingID := "677a8fff42740fa6xxxxx" // Ganti dengan ID yang sesuai

	// Mengonversi ID string menjadi ObjectID
	objectID, err := primitive.ObjectIDFromHex(existingID)
	if err != nil {
		t.Fatalf("Failed to convert ID to ObjectID: %v", err)
	}

	// Ambil data dari koleksi pending_registrations dengan ID yang ada
	collection := db.Collection("pending_registrations")
	var result model.PendingRegistration
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to find user in pending_registrations collection: %v", err)
	}

	// Panggil ConfirmRegistration dengan ID yang sudah ada
	_, user, err := module.ApproveRegistration(existingID, db)
	if err != nil {
		t.Fatalf("Error in ApproveRegistration: %v", err)
	}

	// Verifikasi data yang dipindahkan ke collection pengguna
	if user.Username != result.Username {
		t.Errorf("Expected username: %s, got: %s", result.Username, user.Username)
	}
	if user.Role != result.Role {
		t.Errorf("Expected role: %s, got: %s", result.Role, user.Role)
	}

	// Verifikasi data yang dihapus dari unverified_users
	var deletedUser model.PendingRegistration
	err = collection.FindOne(context.TODO(), bson.M{"_id": result.ID}).Decode(&deletedUser)
	if err == nil {
		t.Errorf("Expected user to be deleted, but found user in pending_registrations: %v", deletedUser)
	}

	// Verifikasi data ada di koleksi pengguna
	collectionUsers := db.Collection("users")
	var foundUser model.User
	err = collectionUsers.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&foundUser)
	if err != nil {
		t.Fatalf("Failed to find user in users collection: %v", err)
	}
	if foundUser.Username != user.Username {
		t.Errorf("Expected username in user: %s, got: %s", user.Username, foundUser.Username)
	}

	// Print confirmation message
	fmt.Printf("User %s confirmed and moved to users collection.\n", user.Username)
}

//function test untuk get all pending dan user
// function test untuk get all pending
// func TestGetAllPendingRegistrations(t *testing.T) {
// 	pendingRegistrations := module.GetAllPendingRegistrations()
// 	fmt.Printf("%v\n", pendingRegistrations)
// }

// function test untuk get all users
// func TestGetAllUsers(t *testing.T) {
// 	// Menyambung ke database
// 	db, err := GetDatabase()
// 	if err != nil {
// 		t.Fatalf("Failed to connect to database: %v", err)
// 	}

// 	// Memanggil fungsi GetAllUsers dan menangkap hasilnya
// 	users, err := module.GetAllUsers(db)
// 	if err != nil {
// 		t.Fatalf("Error getting all users: %v", err)
// 	}

// 	// Menampilkan hasil
// 	fmt.Printf("%v\n", users)
// }

// test untuk transaksi
func TestInsertTransaksi(t *testing.T) {
	// Mock data untuk pengujian
	idUser, err := primitive.ObjectIDFromHex("6784d0ce0e8e100dae5a9921") // Ganti dengan ID yang sesuai
	if err != nil {
		t.Errorf("Error converting ObjectID: %v\n", err)
		return
	}

	username := "Serli" // Ganti dengan username yang sesuai
	metodePembayaran := "Bayar Langsung"
	buktiPembayaran := "https://i.pinimg.com/736x/95/f8/f0/95f8f07eaf103282dbd9518ab8175931.jpg" //link  gambar bukti pembayaran 
	status := "pending"
	alamat := "batujajar"

	// Mock data untuk item keranjang
	items := []model.CartItem{
		{
			IDCartItem:  primitive.NewObjectID(),
			IDProduk:    primitive.NewObjectID(),
			Nama_Produk: "Thai Tea",
			Harga:       5000,
			Quantity:    3,
			SubTotal:    15000,
			Gambar:      "https://th.bing.com/th/id/OIP.AfTYyVvuiEzo20poD4EnsgHaHa?rs=1&pid=ImgDetMain",
		},
		{
			IDCartItem:  primitive.NewObjectID(),
			IDProduk:    primitive.NewObjectID(),
			Nama_Produk: "Mie Pedas",
			Harga:       15000,
			Quantity:    1,
			SubTotal:    15000,
			Gambar:      "https://i.pinimg.com/564x/6c/9c/fb/6c9cfbda40f0d15572fb59e4ad30965e.jpg",
		},
	}

	// Panggil fungsi InsertTransaksi
	insertedID, err := module.InsertTransaksi(idUser, username, items, metodePembayaran, buktiPembayaran, status, alamat)
	if err != nil {
		t.Errorf("Error inserting transaksi: %v", err)
		return
	}
	fmt.Printf("Inserted Transaksi ID: %v\n", insertedID)

	fmt.Printf("Inserted Transaksi ID: %v\n", insertedID)
}

// transakasi get byid and get all

func TestGetTransaksiByID(t *testing.T) {
	// Masukkan ID transaksi yang sudah ada (misalnya dari ID yang Anda tentukan)
	id := "678c9ffa0bcc1355f48e2782" // ID yang sudah ada
	transaksiID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	// Panggil fungsi GetTransaksiByID
	transaksi, err := module.GetTransaksiByID(transaksiID)
	if err != nil {
		t.Errorf("Error retrieving transaksi: %v", err)
		return
	}

	// Output transaksi yang berhasil ditemukan
	fmt.Printf("Retrieved Transaksi: %+v\n", transaksi)
}


func TestGetAllTransaksi(t *testing.T) {
	// Panggil fungsi GetAllTransaksi
	transaksis, err := module.GetAllTransaksi()
	if err != nil {
		t.Errorf("Error retrieving all transaksi: %v", err)
		return
	}
	if len(transaksis) == 0 {
		t.Errorf("No transaksi retrieved")
		return
	}
	for _, transaksi := range transaksis {
		fmt.Printf("Transaksi: %+v\n", transaksi)
	}
}

// delete transaksi byid 
func TestDeleteTransaksiByID(t *testing.T) {
	id := "678ca4e99a64a4a6842cc7c0" //id transaksi punya nicky
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	err = module.DeleteTransaksiByID(objectID, module.MongoConn, "kantin_transaksi")
	if err != nil {
		t.Fatalf("error calling DeleteTransaksiByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetTransaksiByID
	_, err = module.GetTransaksiByID(objectID)
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestGetUserByID(t *testing.T) {
	id := "678d202875bdf2c478072312" // ID dalam format string
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	// Memanggil GetUserByID dengan parameter yang sesuai
	profil, err := module.GetUserByID(objectID.Hex(), module.MongoConn) // objectID.Hex() mengonversi ke string
	if err != nil {
		t.Fatalf("error calling GetUserByID: %v", err)
	}

	fmt.Println(profil)
}


// func TestGetTransaksiByID(t *testing.T) {
// 	// Insert transaksi terlebih dahulu untuk mendapatkan ID
// 	idUser := primitive.NewObjectID()
// 	username := "test_user"
// 	metodePembayaran := "Bayar Langsung"

// 	items := []model.CartItem{
// 		{
// 			IDCartItem:  primitive.NewObjectID(),
// 			IDProduk:    primitive.NewObjectID(),
// 			Nama_Produk: "Thai Tea",
// 			Harga:       5000,
// 			Quantity:    3,
// 			SubTotal:    15000,
// 			Gambar:      "https://th.bing.com/th/id/OIP.AfTYyVvuiEzo20poD4EnsgHaHa?rs=1&pid=ImgDetMain",
// 		},
// 	}

// 	insertedID, err := module.InsertTransaksi(idUser, username, items, metodePembayaran)
// 	if err != nil {
// 		t.Errorf("Error inserting transaksi: %v", err)
// 		return
// 	}

// 	// Ambil ID transaksi yang baru saja dibuat
// 	transaksiID, ok := insertedID.(primitive.ObjectID)
// 	if !ok {
// 		t.Errorf("Inserted ID is not a valid ObjectID")
// 		return
// 	}

// 	// Panggil fungsi GetTransaksiByID
// 	transaksi, err := module.GetTransaksiByID(transaksiID)
// 	if err != nil {
// 		t.Errorf("Error retrieving transaksi: %v", err)
// 		return
// 	}
// 	fmt.Printf("Retrieved Transaksi: %+v\n", transaksi)
// }