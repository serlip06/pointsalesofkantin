package module

import(
	"context"
	//"errors"
	"fmt"
	"time"
	"github.com/serlip06/pointsalesofkantin/model"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func  MongoConnectDB(dbname string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
		return nil, err
	}
	return client.Database(dbname), nil
}

func CreateNewProductNotificationFromProduk(produk model.Produk, addedBy string) model.NewProductNotification {
	return model.NewProductNotification{
		IDProduk:   produk.IDProduk.Hex(), // Konversi ObjectID ke string
		Nama_Produk: produk.Nama_Produk,
		Kategori :    produk.Kategori,
		AddedBy:     addedBy,
		Timestamp:   time.Now().Format(time.RFC3339),
		Message:     fmt.Sprintf("Produk baru telah ditambahkan: %s di kategori %s.", produk.Nama_Produk, produk.Kategori),
	}
}

func CreateLowStockNotificationFromProduk(produk model.Produk, threshold int) model.LowStockNotification {
    return model.LowStockNotification{
        IDProduk:    produk.IDProduk.Hex(),              // Menggunakan IDProduk dari produk
        Nama_Produk: produk.Nama_Produk,                  // Menggunakan Nama_Produk dari produk
        CurrentStock: produk.Stok,                        // Menggunakan Stok dari produk
        Threshold:    threshold,                           // Ambang batas stok yang diberikan
        NotifiedAt:   time.Now().Format(time.RFC3339),     // Waktu notifikasi dibuat
        Message:      fmt.Sprintf("Stok %s menipis (%d item). Segera lakukan restock.", produk .Nama_Produk,produk.Stok), // Pesan notifikasi
    }
}

// Fungsi untuk membuat notifikasi stok habis
func CreateOutOfStockNotificationFromProduk(produk model.Produk) model.OutOfStockNotification {
    return model.OutOfStockNotification{
        IDProduk:   produk.IDProduk.Hex(),               // Menggunakan IDProduk dari produk
        Nama_Produk: produk.Nama_Produk,                  // Menggunakan Nama_Produk dari produk
        NotifiedAt:  time.Now().Format(time.RFC3339),      // Waktu notifikasi dibuat
        Message:     fmt.Sprintf("Stok %s telah habis. Produk tidak tersedia.", produk.Nama_Produk), // Pesan notifikasi
    }
}
