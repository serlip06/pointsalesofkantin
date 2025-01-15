package module

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/serlip06/pointsalesofkantin/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnectDb(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
		return nil
	}
	return client.Database(dbname)
}

func InsertData(dbname, collection string, doc interface{}) interface{} {
	insertResult, err := MongoConnect(dbname).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
		return nil
	}
	return insertResult.InsertedID
}

// insert data
func InsertDataProduk(nama_produk string, deskripsi string, harga int, gambar string, stok int, kategori string) (interface{}, error) {
	// Validasi kategori
	if kategori != "Makanan" && kategori != "Minuman" {
		return nil, errors.New("kategori harus berupa 'Makanan' atau 'Minuman'")
	}

	var produk model.Produk
	produk.IDProduk = primitive.NewObjectID()
	produk.Nama_Produk = nama_produk
	produk.Deskripsi = deskripsi
	produk.Harga = harga
	produk.Gambar = gambar
	produk.Stok = stok
	produk.Kategori = kategori
	produk.CreatedAt = time.Now() // Atur CreatedAt di properti produk

	// Simpan produk ke database
	result := InsertData("kantin", "produk", produk)
	if result == nil {
		return nil, errors.New("gagal menyimpan produk")
	}

	return result, nil
}


// memanggil data by id
func GetProduksFromID(_id primitive.ObjectID, db *mongo.Database, col string) (produk model.Produk, errs error) {
	Produk := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := Produk.FindOne(context.TODO(), filter).Decode(&produk)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return produk, fmt.Errorf("no data found for ID %s", _id)
		}
		return produk, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return produk, nil
}

// memanggil semua data produk
func GetAllProduks(kategori string) (produks []model.Produk, err error) {
	collection := MongoConnect("kantin").Collection("produk")

	// Filter berdasarkan kategori jika diberikan
	var filter bson.M
	if kategori != "" {
		filter = bson.M{"kategori": kategori}
	} else {
		filter = bson.M{}
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("GetAllProduks: %v", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.Background()) {
		var produk model.Produk
		if err := cursor.Decode(&produk); err != nil {
			fmt.Printf("GetAllProduks: %v\n", err)
			continue
		}
		// Validasi tambahan (opsional)
		if produk.CreatedAt.IsZero() {
			produk.CreatedAt = time.Now() // Default ke waktu sekarang jika kosong
		}
		produks = append(produks, produk)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("GetAllProduks: %v", err)
	}
	return produks, nil
}

// update data produk
func UpdateProduks(db *mongo.Database, col string, id primitive.ObjectID, Nama_Produk string, Deskripsi string, Harga int, Gambar string, Stok int, Kategori string) (err error) {
	// Validasi kategori
	if Kategori != "Makanan" && Kategori != "Minuman" {
		return errors.New("kategori harus berupa 'Makanan' atau 'Minuman'")
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama_produk": Nama_Produk,
			"deskripsi":   Deskripsi,
			"harga":       Harga,
			"gambar":      Gambar,
			"stok":        Stok,
			"kategori":    Kategori, // Tambahkan kategori
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateProduks: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}
	return nil
}

// delete data produk
func DeleteProduksByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
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
