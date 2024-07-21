package module

import (
	"context"
	"errors"
	"fmt"

	"github.com/serlip06/pointsalesofkantin/model"
	"go.mongodb.org/mongo-driver/bson"
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
//inserr data 
func InsertDataProduk(nama_produk string, deskripsi string, harga int, gambar string, stok int) interface{} {
    var produk model.Produk
    produk.IDProduk = primitive.NewObjectID()
    produk.Nama_Produk = nama_produk
    produk.Deskripsi = deskripsi
    produk.Harga = harga
    produk.Gambar = gambar
    produk.Stok = stok
    return InsertData("kantin", "produk", produk)
}
//memanggil data by id
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
//memanggil semua data produk
func GetAllProduks() (produks [] model.Produk) {
	collection := MongoConnect("kantin").Collection("produk")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllProduks: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var produk model.Produk
		if err := cursor.Decode(&produk); err != nil {
			fmt.Printf("GetAllProduks: %v\n", err)
			continue
		}
		produks = append(produks, produk )
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllProduks: %v\n", err)
	}
	return produks
}
//update data produk 
func UpdateProduks(db *mongo.Database, col string, id primitive.ObjectID, Nama_Produk string, Deskripsi string, Harga int, Gambar string, Stok int) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama_produk": Nama_Produk,
			"deskripsi":   Deskripsi,
			"harga":       Harga,
			"gambar":      Gambar,
			"stok":        Stok,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateProduks: %v\n", err)
		return 
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified ID")
		return 
	}
	return nil
}
//delete data produk
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