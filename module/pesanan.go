package module

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/serlip06/pointsalesofkantin/model" // Mengimpor package model dengan benar
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

// Fungsi untuk menghubungkan ke MongoDB
func MongoConnectDatabase(dbname string) (db *mongo.Database, err error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
		return nil, err
	}
	return client.Database(dbname), nil
}

// Fungsi untuk menyisipkan pesanan ke MongoDB
func InsertPesanan(client *mongo.Client, id_customer primitive.ObjectID, produk []model.Produk, total_harga int, status string) (interface{}, error) {

	// Validasi status pesanan (misalnya, hanya "Menunggu Konfirmasi", "Dikonfirmasi", atau "Selesai" yang diperbolehkan)
	if status != "Menunggu Konfirmasi" && status != "Dikonfirmasi" && status != "Selesai" {
		return nil, errors.New("status harus berupa 'Menunggu Konfirmasi', 'Dikonfirmasi', atau 'Selesai'")
	}

	// Menyiapkan data pesanan
	var pesanan model.Pesanan
	pesanan.IDPesanan = primitive.NewObjectID() // Generate ID unik untuk pesanan baru
	pesanan.ID = id_customer                    // ID Customer yang membuat pesanan
	pesanan.Produk = produk                      // Daftar produk yang dipesan
	pesanan.TotalHarga = total_harga            // Total harga pesanan
	pesanan.Status = status                      // Status pesanan
	// Format tanggal menjadi string
	pesanan.TanggalPesan = time.Now().Format("2006-01-02 15:04:05") // Format: "YYYY-MM-DD HH:MM:SS"

	// Mengakses koleksi Pesanan di MongoDB
	collection := client.Database("kantin").Collection("pesanan")

	// Menyisipkan pesanan baru ke dalam koleksi
	result, err := collection.InsertOne(context.TODO(), pesanan)
	if err != nil {
		return nil, err
	}

	return result, nil
}


// Memanggil data pesanan berdasarkan ID
func GetPesananByID(_id primitive.ObjectID, db *mongo.Database, col string) (pesanan model.Pesanan, errs error) {
	// Mengakses koleksi pesanan
	Pesanan := db.Collection(col)

	// Filter pencarian berdasarkan ID
	filter := bson.M{"_id": _id}
	err := Pesanan.FindOne(context.TODO(), filter).Decode(&pesanan)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return pesanan, fmt.Errorf("no data found for ID %s", _id)
		}
		return pesanan, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return pesanan, nil
}

// Memanggil semua data pesanan berdasarkan ID customer
func GetAllPesananByID(idCustomer primitive.ObjectID, db *mongo.Database, col string) ([]model.Pesanan, error) {
	var pesananList []model.Pesanan
	Pesanan := db.Collection(col)

	// Filter berdasarkan ID customer
	filter := bson.M{"id_customer": idCustomer}

	cursor, err := Pesanan.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("GetAllPesananByID: %v", err)
	}
	defer cursor.Close(context.TODO())

	// Menyusun hasil pencarian ke dalam slice
	for cursor.Next(context.TODO()) {
		var pesanan model.Pesanan
		if err := cursor.Decode(&pesanan); err != nil {
			fmt.Printf("GetAllPesananByID: %v\n", err)
			continue
		}
		pesananList = append(pesananList, pesanan)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("GetAllPesananByID: %v", err)
	}
	return pesananList, nil
}

// Fungsi untuk memperbarui data pesanan
func UpdatePesanan(db *mongo.Database, col string, id primitive.ObjectID, status string, produk []model.Produk, totalHarga int) error {
	// Validasi status pesanan
	if status != "Menunggu Konfirmasi" && status != "Dikonfirmasi" && status != "Selesai" {
		return errors.New("status harus berupa 'Menunggu Konfirmasi', 'Dikonfirmasi', atau 'Selesai'")
	}

	// Menentukan filter berdasarkan ID pesanan
	filter := bson.M{"_id": id}

	// Data update yang akan diterapkan
	update := bson.M{
		"$set": bson.M{
			"status":       status,         // Update status pesanan
			"produk":       produk,         // Update produk yang dipesan
			"total_harga":  totalHarga,     // Update total harga
			"tanggal_pesan": time.Now().Format("2006-01-02 15:04:05"), // Update tanggal pesanan
		},
	}

	// Melakukan update pada koleksi pesanan
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdatePesanan: %v\n", err)
		return err
	}

	// Memastikan bahwa ada data yang diubah
	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}

	return nil
}

// Menghapus pesanan berdasarkan ID
func DeletePesananByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	Pesanan := db.Collection(col)

	// Menentukan filter berdasarkan ID pesanan
	filter := bson.M{"_id": _id}

	// Melakukan penghapusan pesanan
	result, err := Pesanan.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	// Memastikan ada data yang terhapus
	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
