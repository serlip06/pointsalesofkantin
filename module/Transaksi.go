package module

import (
	"context"
	"fmt"
	"time"

	"github.com/serlip06/pointsalesofkantin/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func InsertTransaksi(idUser primitive.ObjectID, username string, items []model.CartItem, metodePembayaran string, buktiPembayaran string, status string, alamat string) (interface{}, error) {
	// Validasi Items tidak boleh kosong
	if len(items) == 0 {
		return nil, fmt.Errorf("items cannot be empty")
	}

	// Hitung total harga
	calculatedTotal := calculateTotalHarga(items)

	// Buat objek transaksi
	var transaksi model.Transaksi
	transaksi.IDTransaksi = primitive.NewObjectID() // Generate ID transaksi baru
	transaksi.IDUser = idUser                       // Masukkan ID user
	transaksi.Username = username                   // Masukkan username user
	transaksi.Items = items                         // Masukkan data cart item
	transaksi.TotalHarga = calculatedTotal          // Set total harga
	transaksi.MetodePembayaran = metodePembayaran   // Masukkan metode pembayaran
	transaksi.CreatedAt = time.Now()                // Masukkan timestamp transaksi
	transaksi.Buktipembayaran = buktiPembayaran     // Masukkan buktipembayaran
	transaksi.Status = status                       // Masukkan status dari transaksi
	transaksi.Alamat = alamat                       // mesukkan alamat dari pengguna

	// Validasi Total Harga (cek ulang jika diperlukan)
	if transaksi.TotalHarga != calculatedTotal {
		return nil, fmt.Errorf("total price mismatch: expected %d, got %d", calculatedTotal, transaksi.TotalHarga)
	}

	// Masukkan transaksi ke database
	result, err := InsertTransaksiToDatabase("kantin", "kantin_transaksi", transaksi)
	if err != nil {
		return nil, fmt.Errorf("failed to insert transaction: %w", err)
	}

	return result, nil
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
func UpdateTransaksi(db *mongo.Database, col string, id primitive.ObjectID, transaksi model.Transaksi) (err error) {
	// Membuat filter untuk mencocokkan ID transaksi yang akan diupdate
	filter := bson.M{"_id": id}

	// Menyiapkan data update
	update := bson.M{
		"$set": bson.M{
			"id_user":           transaksi.IDUser,           // Sesuai dengan `bson:"id_user"`
			"username":          transaksi.Username,         // Sesuai dengan `bson:"username"`
			"items":             transaksi.Items,            // Sesuai dengan `bson:"items"`
			"total_harga":       transaksi.TotalHarga,       // Sesuai dengan `bson:"total_harga"`
			"metode_pembayaran": transaksi.MetodePembayaran, // Sesuai dengan `bson:"metode_pembayaran"`
			"created_at":        transaksi.CreatedAt,        // Sesuai dengan `bson:"created_at"`
			"bukti_pembayaran":  transaksi.Buktipembayaran,  // Sesuai dengan `bson:"bukti_pembayaran"`
			"status":            transaksi.Status,
			"alamat":            transaksi.Alamat,
		},
	}

	// Melakukan update pada koleksi transaksi
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateTransaksi: %v\n", err)
		return
	}

	// Mengecek jika tidak ada data yang diubah
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("no data has been changed with the specified ID")
		return
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