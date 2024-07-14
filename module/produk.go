package module

import (
	"context"
	"fmt"

	"github.com/serlip06/pointsalesofkantin/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnectDb(dbname string) (db *mongo.Database)  {
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

func InsertDataProduk(adminID string, nama_produk string, deskripsi string, harga int, gambar string, stok int) interface{} {
    var produk model.Produk
    produk.IDProduk = primitive.NewObjectID()
    produk.AdminID = primitive.NewObjectID() // Assuming adminID is an ObjectID, change this if adminID is different
    produk.Nama_Produk = nama_produk
    produk.Deskripsi = deskripsi
    produk.Harga = harga
    produk.Gambar = gambar
    produk.Stok = stok
    return InsertData("kantin", "produk", produk)
}