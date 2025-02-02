package module

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/serlip06/pointsalesofkantin/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

// Fungsi untuk koneksi ke database MongoDB
func MongoConnectDBase(dbname string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnectDBase: %v\n", err)
		return nil, err
	}
	return client.Database(dbname), nil
}

func InsertTransaksiToDatabase(dbName, collectionName string, data interface{}) (interface{}, error) {
	// Mendapatkan koneksi database MongoDB
	collection, err := MongoConnectDBase(dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Memasukkan data ke dalam koleksi yang sesuai
	result, err := collection.Collection(collectionName).InsertOne(context.TODO(), data)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %w", err)
	}

	// Mengembalikan ID yang dihasilkan setelah berhasil insert
	return result.InsertedID, nil
}

// Fungsi untuk menambahkan transaksi baru
func InsertTransaksi(db *mongo.Database, idUser primitive.ObjectID, idCartItems []primitive.ObjectID, metodePembayaran, buktiPembayaran, status, alamat string) (interface{}, error) {
	// Ambil item dari cart berdasarkan ID yang dipilih
	collectionCart := db.Collection("cart_items")
	filter := bson.M{"_id": bson.M{"$in": idCartItems}}

	cursor, err := collectionCart.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve selected cart items: %v", err)
	}
	defer cursor.Close(context.TODO())

	// Simpan item dalam slice
	var items []model.CartItem
	for cursor.Next(context.TODO()) {
		var item model.CartItem
		if err := cursor.Decode(&item); err != nil {
			return nil, fmt.Errorf("error decoding cart item: %v", err)
		}
		items = append(items, item)
	}

	// Pastikan ada item yang valid
	if len(items) == 0 {
		return nil, fmt.Errorf("no valid cart items found for checkout")
	}

	// Hitung total harga transaksi
	totalHarga := calculateTotalHarga(items)

	// Buat data transaksi
	transaksi := model.Transaksi{
		IDTransaksi:      primitive.NewObjectID(),
		IDUser:           idUser,
		IDCartItem:       idCartItems,
		MetodePembayaran: metodePembayaran,
		TotalHarga:       totalHarga,
		BuktiPembayaran:  buktiPembayaran,
		Alamat:           alamat,
		CreatedAt:        time.Now(),
		Status:           status,
	}

	// Simpan transaksi ke koleksi "kantin_transaksi"
	collectionTransaksi := db.Collection("kantin_transaksi")
	result, err := collectionTransaksi.InsertOne(context.TODO(), transaksi)
	if err != nil {
		return nil, fmt.Errorf("failed to insert transaction: %v", err)
	}

	// Hapus hanya item yang telah di-checkout dari cart
	_, err = collectionCart.DeleteMany(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to clear selected cart items: %v", err)
	}

	return result.InsertedID, nil
}

// handler transaksi untuk request
func CheckoutHandler(c *gin.Context) {
	var req model.TransaksiRequest

	// Bind request JSON ke struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil koneksi database
	db, err := MongoConnectDBase("kantin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	// Panggil fungsi transaksi dengan data dari request
	result, err := InsertTransaksi(db, req.IDUser, req.IDCartItems, req.MetodePembayaran, req.BuktiPembayaran, req.Status, req.Alamat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Response jika berhasil
	c.JSON(http.StatusOK, gin.H{"message": "Transaction successful", "transaction_id": result})
}

// Fungsi untuk menghitung total harga
func calculateTotalHarga(items []model.CartItem) int {
	total := 0
	for _, item := range items {
		total += item.SubTotal
	}
	return total
}

// Fungsi untuk mendapatkan transaksi berdasarkan ID
func GetTransaksiByID(transaksiID primitive.ObjectID) (model.Transaksi, error) {
	var transaksi model.Transaksi
	collection, err := MongoConnectDBase("kantin")
	if err != nil {
		return transaksi, err
	}

	filter := bson.M{"_id": transaksiID}
	err = collection.Collection("kantin_transaksi").FindOne(context.TODO(), filter).Decode(&transaksi)
	if err != nil {
		fmt.Printf("GetTransaksiByID: %v\n", err)
		return transaksi, err
	}
	return transaksi, nil
}

func GetAllTransaksiByIDUser(userID string, database *mongo.Database) ([]model.Transaksi, error) {
	var transaksis []model.Transaksi

	collection := database.Collection("kantin_transaksi")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Mengkonversi userID menjadi ObjectID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %v", err)
	}

	filter := bson.M{"id_user": objID}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var transaksi model.Transaksi
		if err := cursor.Decode(&transaksi); err != nil {
			return nil, err
		}
		transaksis = append(transaksis, transaksi)
	}

	return transaksis, nil
}

// Fungsi untuk mendapatkan semua transaksi

// get all transaksi by user yang di filter berdasarkan data terbaru  yang masuk

func GetAllTransaksi() ([]model.Transaksi, error) {
	var transaksis []model.Transaksi
	collection, err := MongoConnectDBase("kantin")
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Collection("kantin_transaksi").Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllTransaksi: %v\n", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var transaksi model.Transaksi
		if err := cursor.Decode(&transaksi); err != nil {
			fmt.Printf("GetAllTransaksi: %v\n", err)
			continue
		}
		transaksis = append(transaksis, transaksi)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllTransaksi: %v\n", err)
		return nil, err
	}
	return transaksis, nil
}

// update dan delete

// update transaksi
func UpdateTransaksi(db *mongo.Database, id primitive.ObjectID, transaksi model.Transaksi) error {
	// Filter transaksi berdasarkan ID
	filter := bson.M{"_id": id}

	// Data yang akan diupdate
	update := bson.M{
		"$set": bson.M{
			"id_user":           transaksi.IDUser,
			"id_cartitem":       transaksi.IDCartItem,
			"metode_pembayaran": transaksi.MetodePembayaran,
			"bukti_pembayaran":  transaksi.BuktiPembayaran,
			"status":            transaksi.Status,
			"alamat":            transaksi.Alamat,
			"created_at":        transaksi.CreatedAt, // Jika tanggal diperbarui juga
		},
	}

	// Lakukan update
	result, err := db.Collection("kantin_transaksi").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update transaction: %w", err)
	}

	// Jika tidak ada data yang diperbarui
	if result.ModifiedCount == 0 {
		return fmt.Errorf("no transaction updated")
	}

	return nil
}

// detele transaksi
// Fungsi untuk menghapus transaksi berdasarkan ID
func DeleteTransaksiByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	Produk := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := Produk.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// cadangan

// func InsertTransaksi(db *mongo.Database, idUser primitive.ObjectID, metodePembayaran string, buktiPembayaran string, status string, alamat string) (interface{}, error) {
// 	// Ambil semua item dari keranjang user
// 	collectionCart := db.Collection("cart_items")
// 	filter := bson.M{"id_user": idUser}
// 	cursor, err := collectionCart.Find(context.TODO(), filter)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to retrieve cart items: %v", err)
// 	}
// 	defer cursor.Close(context.TODO())

// 	// Inisialisasi slice untuk menyimpan item transaksi
// 	var items []model.CartItem
// 	var idCartItems []primitive.ObjectID // Simpan IDCartItem dari item dalam keranjang

// 	for cursor.Next(context.TODO()) {
// 		var item model.CartItem
// 		if err := cursor.Decode(&item); err != nil {
// 			return nil, fmt.Errorf("error decoding cart item: %v", err)
// 		}
// 		items = append(items, item)
// 		idCartItems = append(idCartItems, item.IDCartItem) // Simpan hanya IDCartItem
// 	}

// 	// Pastikan ada item dalam keranjang
// 	if len(items) == 0 {
// 		return nil, fmt.Errorf("cart is empty, cannot proceed with transaction")
// 	}

// 	// Hitung total harga menggunakan fungsi calculateTotalHarga
// 	totalHarga := calculateTotalHarga(items)

// 	// Buat data transaksi baru
// 	transaksi := model.Transaksi{
// 		IDTransaksi:      primitive.NewObjectID(),
// 		IDUser:           idUser,
// 		IDCartItem:       idCartItems, // Simpan semua IDCartItem dalam bentuk slice
// 		MetodePembayaran: metodePembayaran,
// 		TotalHarga:       totalHarga,
// 		BuktiPembayaran:  buktiPembayaran,
// 		Alamat:           alamat,
// 		CreatedAt:        time.Now(),
// 		Status:           status,
// 	}

// 	// Simpan transaksi ke koleksi "kantin_transaksi"
// 	collectionTransaksi := db.Collection("kantin_transaksi")
// 	result, err := collectionTransaksi.InsertOne(context.TODO(), transaksi)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to insert transaction: %v", err)
// 	}

// 	// Hapus item dari keranjang setelah transaksi berhasil
// 	_, err = collectionCart.DeleteMany(context.TODO(), filter)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to clear cart: %v", err)
// 	}

// 	return result.InsertedID, nil
// }

// func InsertTransaksiToDatabase(collectionName string, collectionType string, data interface{}) (interface{}, error) {
// 	// Dapatkan koneksi database
// 	collection, err := MongoConnectDBase("kantin")
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to database: %w", err)
// 	}

// 	// Memasukkan data ke dalam koleksi "kantin_transaksi"
// 	result, err := collection.Collection(collectionName).InsertOne(context.TODO(), data)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to insert document: %w", err)
// 	}

//		return result.InsertedID, nil
//	}

// func calculateTotalHarga(items []model.CartItem) int {
// 	total := 0
// 	for _, item := range items {
// 		total += item.SubTotal
// 	}
// 	return total
// }

// func InsertTransaksi(idUser primitive.ObjectID, username string, items []model.CartItem, metodePembayaran string) interface{} {
// 	// Inisialisasi data transaksi
// 	var transaksi model.Transaksi
// 	transaksi.IDTransaksi = primitive.NewObjectID()
// 	transaksi.IDUser = idUser
// 	transaksi.Username = username
// 	transaksi.Items = items

// 	// Menghitung total harga
// 	totalHarga := 0
// 	for _, item := range items {
// 		totalHarga += item.SubTotal
// 	}
// 	transaksi.TotalHarga = totalHarga
// 	transaksi.CreatedAt = time.Now()
// 	transaksi.MetodePembayaran = metodePembayaran

// 	// Memasukkan ke database
// 	return InsertOneDoc("kantin", "kantin_transaksi", transaksi)
// }

//cadangan
// func GetAllTransaksiByIDUser(idUser string, db *mongo.Database) ([]model.CartItem, error) {
// 	// Konversi idUser ke ObjectID
// 	objectID, err := primitive.ObjectIDFromHex(idUser)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid id_user: %v", err)
// 	}

// 	// Buat koleksi transaksi
// 	collection := db.Collection("kantin_transaksi")

// 	// Filter berdasarkan id_user
// 	filter := bson.M{"id_user": objectID}

// 	// Cari semua transaksi dengan id_user tertentu
// 	cursor, err := collection.Find(context.TODO(), filter)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch transactions: %v", err)
// 	}
// 	defer cursor.Close(context.TODO())

// 	var items []model.CartItem
// 	for cursor.Next(context.TODO()) {
// 		var transaksi model.Transaksi
// 		if err := cursor.Decode(&transaksi); err != nil {
// 			fmt.Printf("failed to decode transaksi: %v\n", err)
// 			continue
// 		}
// 		// Tambahkan semua items dari transaksi ini ke hasil akhir
// 		items = append(items, transaksi.Items...)
// 	}

// 	// Cek error setelah iterasi
// 	if err := cursor.Err(); err != nil {
// 		return nil, fmt.Errorf("cursor error: %v", err)
// 	}

// 	return items, nil
// }
