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
	//"golang.org/x/crypto/bcrypt"
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
	db, err := module.MongoConnectdatabase("kantin") // Pastikan database sudah benar
	assert.NoError(t, err)

	idProduk, err := primitive.ObjectIDFromHex("673c90cd715120ed663eb984") // ID produk
	assert.NoError(t, err)

	idUser, err := primitive.ObjectIDFromHex("678f3dec6c07fa5fb07d8e3a") // ID pengguna
	assert.NoError(t, err)

	// Pastikan ID produk tersedia di database
	collection := db.Collection("produk")
	var product model.Produk
	err = collection.FindOne(context.TODO(), bson.M{"_id": idProduk}).Decode(&product)
	assert.NoError(t, err)

	// Insert ke keranjang
	quantity := 2
	result, err := module.InsertDataCartItemFunc(db, idProduk, idUser, quantity)

	// Verifikasi hasil
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

// //get cartitemfromid
func TestGetCartItemFromID(t *testing.T) {
	db, err := module.MongoConnectdatabase("kantin")
	assert.NoError(t, err)

	id, err := primitive.ObjectIDFromHex("676fb8a26214462f56e4965d") // Ganti dengan ID yang valid
	assert.NoError(t, err)

	cartItem, err := module.GetCartItemFromID(id, db, "cart_items")

	// Verifikasi hasil
	assert.NoError(t, err)
	assert.NotNil(t, cartItem)
	assert.Equal(t, id, cartItem.IDCartItem) // Sesuaikan dengan struct
}

// //get all cartitem
func TestGetAllCartItems(t *testing.T) {
	cartitems := module.GetAllCartItems()
	// Menggunakan %v untuk mencetak slice dengan format default
	fmt.Printf("%v\n", cartitems)
}

// deletecartitem
func TestDeleteCartItemFromID(t *testing.T) {
	db, err := module.MongoConnectdatabase("kantin")
	assert.NoError(t, err)

	id := "6770f80a419da98516ba7db1" // ID item keranjang
	objectID, err := primitive.ObjectIDFromHex(id)
	assert.NoError(t, err)

	idUser, err := primitive.ObjectIDFromHex("678f3dec6c07fa5fb07d8e3a") // ID pengguna
	assert.NoError(t, err)

	// Delete item dari keranjang
	err = module.DeleteCartItemByID(objectID, idUser, db, "cart_items")
	assert.NoError(t, err)

	// Pastikan data sudah terhapus
	_, err = module.GetCartItemFromID(objectID, db, "cart_items")
	assert.Error(t, err) // Harus error karena data sudah dihapus
}


// test untuk chekout dari keranjang 

func TestCheckoutFromCart(t *testing.T) {
	// Koneksi ke database
	db, err := module.MongoConnectdatabase("kantin")
	assert.NoError(t, err)

	// Gunakan ID user dari database yang sudah ada
	idUser, err := primitive.ObjectIDFromHex("677bc9a25f0ee6d09631e07e") // ID user dari database sebastian
	assert.NoError(t, err)

	// Simpan item ke dalam cart sebagai data awal
	collectionCart := db.Collection("cart_items")

	item1 := model.CartItem{
		IDCartItem: primitive.NewObjectID(),
		IDUser:     idUser,
		IDProduk:   primitive.NewObjectID(),
		Harga:      15000,
		Quantity:   2,
		SubTotal:   30000,
		IsSelected: true, // Item harus dipilih agar masuk ke transaksi
	}
	_, err = collectionCart.InsertOne(context.TODO(), item1)
	assert.NoError(t, err)

	// Panggil fungsi CheckoutFromCart
	metodePembayaran := "Transfer Bank"
	buktiPembayaran := "bukti_transfer.jpg"
	alamat := "Jl. Kantin No. 123"

	transaksiID, err := module.CheckoutFromCart(db, idUser, metodePembayaran, buktiPembayaran, alamat)

	// Pastikan tidak ada error dan transaksi ID valid
	assert.NoError(t, err)
	assert.NotEqual(t, primitive.NilObjectID, transaksiID)

	// Pastikan transaksi masuk ke database
	collectionTransaksi := db.Collection("kantin_transaksi")
	var transaksi model.Transaksi
	err = collectionTransaksi.FindOne(context.TODO(), bson.M{"_id": transaksiID}).Decode(&transaksi)
	assert.NoError(t, err)

	// Pastikan item dalam transaksi benar
	assert.Equal(t, idUser, transaksi.IDUser)
	assert.Equal(t, metodePembayaran, transaksi.MetodePembayaran)
	assert.Equal(t, 30000, transaksi.TotalHarga)
	assert.Equal(t, "Pending", transaksi.Status)

	// Pastikan item cart sudah dihapus
	count, err := collectionCart.CountDocuments(context.TODO(), bson.M{"id_user": idUser})
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)
}

//register
func TestRegisterHandler(t *testing.T) {
	// Setup test database
	db := module.MongoConnectdb("kantin")

	// Test case input
	req := model.RegisterRequest{
		Username: "cello",
		Password: "ce123",
		Role:     "customer",
	}

	// Call the RegisterHandler function
	message, err := module.RegisterHandler(req, db)

	// Test if there were no errors
	if err != nil {
		t.Errorf("Error in RegisterHandler: %v", err)
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

	// Verify if the password is correctly stored (in plain text)
	if result.Password != req.Password {
		t.Fatalf("Expected password: %s, got: %s", req.Password, result.Password)
	}

	// Print confirmation message
	fmt.Printf("User %s successfully registered and saved to the database.\n", req.Username)
}

// UJICOBA APPROVE NYA
func TestApproveRegistration(t *testing.T) {
	// Setup test database
	db := module.MongoConnectdb("kantin")

	// Tentukan ID yang sudah ada di koleksi pending_registrations
	// Misalnya ID sudah diketahui sebelumnya
	existingID := "67aa36c6c0f544aeabdd877f" // Ganti dengan ID yang sesuai

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
	// Setup koneksi database
	db, err := module.MongoConnectDBase("kantin")
	assert.NoError(t, err)

	// ID User yang akan melakukan transaksi
	idUser, err := primitive.ObjectIDFromHex("67852acf2b6bacb779d85e66") // ID asti
	assert.NoError(t, err)

	// Ambil semua item dari cart milik user dari database
	cartCollection := db.Collection("cart_items")
	cursor, err := cartCollection.Find(context.TODO(), bson.M{"id_user": idUser})
	assert.NoError(t, err)
	defer cursor.Close(context.TODO())

	// Simpan ID cart items yang ditemukan
	var idCartItems []primitive.ObjectID
	for cursor.Next(context.TODO()) {
		var cartItem struct {
			ID primitive.ObjectID `bson:"_id"`
		}
		err := cursor.Decode(&cartItem)
		assert.NoError(t, err)
		idCartItems = append(idCartItems, cartItem.ID)
	}

	// Pastikan ada data cart items yang ditemukan
	assert.NotEmpty(t, idCartItems, "Tidak ada cart item yang ditemukan untuk user ini")

	// Parameter transaksi
	metodePembayaran := "Bayar Langsung"
	buktiPembayaran := "https://i.pinimg.com/736x/95/f8/f0/95f8f07eaf103282dbd9518ab8175931.jpg"
	status := "pending"
	alamat := "batujajar"

	// Panggil fungsi InsertTransaksi
	insertedID, err := module.InsertTransaksi(db, idUser, idCartItems, metodePembayaran, buktiPembayaran, status, alamat)
	assert.NoError(t, err)
	assert.NotNil(t, insertedID)

	// Pastikan transaksi masuk ke koleksi "kantin_transaksi"
	transaksiCollection := db.Collection("kantin_transaksi")
	var transaksi struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	err = transaksiCollection.FindOne(context.TODO(), bson.M{"_id": insertedID}).Decode(&transaksi)
	assert.NoError(t, err)

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
	id := "677bc9a25f0ee6d09631e07e" // ID dalam format string
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

// Unit Test get transaksi by userid 

func TestGetAllTransaksiByIDUser(t *testing.T) {
	// Gunakan ID user yang valid
	userID := "67852acf2b6bacb779d85e66" // Ganti dengan IDUser punya asti
	
	// Menggunakan koneksi yang sudah didefinisikan di config.go
	db := module.MongoConn.Client().Database("kantin") // Menggunakan database "kantin" yang sudah ada dalam konfigurasi

	// Memanggil fungsi untuk mengambil transaksi berdasarkan userID
	transaksis, err := module.GetAllTransaksiByIDUser(userID, db)
	if err != nil {
		t.Errorf("Error retrieving transactions: %v", err)
		return
	}

	if len(transaksis) == 0 {
		t.Errorf("No transactions found for user ID: %v", userID)
	} else {
		fmt.Printf("Transactions for user ID %v: %+v\n", userID, transaksis)
	}
}







//cadangan
// func TestInsertTransaksi(t *testing.T) {
//     db, err := module.MongoConnectdatabase("kantin") // Nama database yang benar
//     assert.NoError(t, err)

//     // Mock ID User yang valid
//     idUser, err := primitive.ObjectIDFromHex("678cfa4f8c2198269380a729") // Contoh ID pengguna serli
//     assert.NoError(t, err)

//     // Parameter transaksi
//     metodePembayaran := "Bayar Langsung"
//     buktiPembayaran := "https://i.pinimg.com/736x/95/f8/f0/95f8f07eaf103282dbd9518ab8175931.jpg"
//     status := "pending"
//     alamat := "batujajar"

//     // Panggil fungsi InsertTransaksi dengan menyesuaikan jumlah argumen
//     insertedID, err := module.InsertTransaksi(db, idUser, metodePembayaran, buktiPembayaran, status, alamat)
//     if err != nil {
//         t.Errorf("Error inserting transaksi: %v", err)
//         return
//     }

//     // Verifikasi hasil
//     assert.NoError(t, err)
//     assert.NotNil(t, insertedID)

//     // Pastikan ID transaksi yang dihasilkan valid
//     if insertedID == nil {
//         t.Errorf("Expected a valid inserted ID, but got nil")
//     }

//     fmt.Printf("Inserted Transaksi ID: %v\n", insertedID)
// }







// func TestGetAllTransaksiByIDUser(t *testing.T) {
// 	// ID User dalam format string
// 	id := "678cfa508c2198269380a72a"

// 	// Dapatkan database yang sudah terkoneksi
// 	database, err := MongoConnectDBase("kantin") // Fungsi untuk menghubungkan ke database
// 	if err != nil {
// 		t.Fatalf("error connecting to database: %v", err)
// 	}

// 	// Memanggil fungsi GetAllTransaksiByIDUser dengan argumen string dan database
// 	transaksis, err := module.GetAllTransaksiByIDUser(id, database) // Gunakan id (string) dan database
// 	if err != nil {
// 		t.Fatalf("error calling GetAllTransaksiByIDUser: %v", err)
// 	}

// 	// Memeriksa apakah transaksi ditemukan
// 	if len(transaksis) == 0 {
// 		t.Fatalf("no transactions found for user ID: %v", id)
// 	}

// 	// Iterasi melalui setiap transaksi
// 	for _, transaksi := range transaksis {
// 		// Menampilkan detail transaksi
// 		t.Logf("Transaksi ID: %s, Total Harga: %d, Metode Pembayaran: %s, Status: %s, Created At: %v\n",
// 			transaksi.IDTransaksi.Hex(), transaksi.TotalHarga, transaksi.MetodePembayaran, transaksi.Status, transaksi.CreatedAt)

// 		// Iterasi melalui item dalam transaksi (CartItem)
// 		for _, item := range transaksi.Items {
// 			t.Logf("  - Nama Produk: %s, Harga: %d, Quantity: %d, Sub Total: %d\n",
// 				item.Nama_Produk, item.Harga, item.Quantity, item.SubTotal)
// 		}
// 	}
// }





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

// func TestInsertTransaksi(t *testing.T) {
// 	// Mock data untuk pengujian
// 	idUser, err := primitive.ObjectIDFromHex("6784d0ce0e8e100dae5a9921") // Ganti dengan ID yang sesuai
// 	if err != nil {
// 		t.Errorf("Error converting ObjectID: %v\n", err)
// 		return

// 	}

// 	username := "Serli" // Ganti dengan username yang sesuai
// 	metodePembayaran := "Bayar Langsung"
// 	buktiPembayaran := "https://i.pinimg.com/736x/95/f8/f0/95f8f07eaf103282dbd9518ab8175931.jpg" //link  gambar bukti pembayaran 
// 	status := "pending"
// 	alamat := "batujajar"

// 	// Mock data untuk item keranjang
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
// 		{
// 			IDCartItem:  primitive.NewObjectID(),
// 			IDProduk:    primitive.NewObjectID(),
// 			Nama_Produk: "Mie Pedas",
// 			Harga:       15000,
// 			Quantity:    1,
// 			SubTotal:    15000,
// 			Gambar:      "https://i.pinimg.com/564x/6c/9c/fb/6c9cfbda40f0d15572fb59e4ad30965e.jpg",
// 		},
// 	}

// 	// Panggil fungsi InsertTransaksi
// 	insertedID, err := module.InsertTransaksi(idUser, username, items, metodePembayaran, buktiPembayaran, status, alamat)
// 	if err != nil {
// 		t.Errorf("Error inserting transaksi: %v", err)
// 		return
// 	}
// 	fmt.Printf("Inserted Transaksi ID: %v\n", insertedID)

// 	fmt.Printf("Inserted Transaksi ID: %v\n", insertedID)
// }