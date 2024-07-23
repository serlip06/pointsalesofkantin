package module

import (
	"context"
	"errors"
	"fmt"

	"github.com/serlip06/pointsalesofkantin/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"os"
	//"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database)  {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
		return nil
	}
	return client.Database(dbname)
}

func InsertOneDoc(dbname, collection string, doc interface{}) interface{} {
	insertResult, err := MongoConnect(dbname).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
		return nil
	}
	return insertResult.InsertedID
}

func InsertPelanggan(nama string, phoneNumber string, alamat string, email []string) interface{} {
	var pelanggan  model.Pelanggan
	pelanggan.ID = primitive.NewObjectID()
	pelanggan.Nama = nama
	pelanggan.Phone_number = phoneNumber
	pelanggan.Alamat = alamat
	pelanggan.Email = email
	return InsertOneDoc("kantin", "kantin_pelanggan", pelanggan)
}

func GetPelangganByID(pelangganID primitive.ObjectID, db *mongo.Database, collectionName string) (pelanggan model.Pelanggan, err error) {
	collection := db.Collection(collectionName)
	filter := bson.M{"_id": pelangganID}
	err = collection.FindOne(context.TODO(), filter).Decode(&pelanggan)
	if err != nil {
		fmt.Printf("GetPelangganByID: %v\n", err)
	}
	return pelanggan, err
}


func GetAllPelanggan() (pelanggans [] model.Pelanggan) {
	collection := MongoConnect("kantin").Collection("kantin_pelanggan")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllPelanggan: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var pelanggan model.Pelanggan
		if err := cursor.Decode(&pelanggan); err != nil {
			fmt.Printf("GetAllPelanggan: %v\n", err)
			continue
		}
		pelanggans = append(pelanggans, pelanggan)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllPelanggan: %v\n", err)
	}
	return pelanggans
}

// func InsertProduk(namaProduk string, deskripsi string, harga int) interface{} {
// 	var produk model.Produk
// 	produk.ID = primitive.NewObjectID()
// 	produk.Nama_Produk = namaProduk
// 	produk.Deskripsi = deskripsi
// 	produk.Harga = harga
// 	return InsertOneDoc("kantin", "Menu_produk", produk)
// }

func GetProdukByID(produkID primitive.ObjectID) (produk model.Produk) {
	collection := MongoConnect("kantin").Collection("Menu_produk")
	filter := bson.M{"_id": produkID}
	err := collection.FindOne(context.TODO(), filter).Decode(&produk)
	if err != nil {
		fmt.Printf("GetProdukByID: %v\n", err)
	}
	return produk
}

func GetAllProduk() (produks [] model.Produk) {
	collection := MongoConnect("kantin").Collection("Menu_produk")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllProduk: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var produk model.Produk
		if err := cursor.Decode(&produk); err != nil {
			fmt.Printf("GetAllProduk: %v\n", err)
			continue
		}
		produks = append(produks, produk)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllProduk: %v\n", err)
	}
	return produks
}

func InsertTransaksi(metodePembayaran string, tanggalWaktu string) interface{} {
	var transaksi model.Transaksi
	transaksi.ID = primitive.NewObjectID()
	transaksi.Metode_Pembayaran = metodePembayaran
	transaksi.Tanggal_Waktu = tanggalWaktu
	return InsertOneDoc("kantin", "kantin_transaksi", transaksi)
}

func GetTransaksiByID(transaksiID primitive.ObjectID) (transaksi model.Transaksi) {
	collection := MongoConnect("kantin").Collection("kantin_transaksi")
	filter := bson.M{"_id": transaksiID}
	err := collection.FindOne(context.TODO(), filter).Decode(&transaksi)
	if err != nil {
		fmt.Printf("GetTransaksiByID: %v\n", err)
	}
	return transaksi
}

func GetAllTransaksi() (transaksis [] model.Transaksi) {
	collection := MongoConnect("kantin").Collection("kantin_transaksi")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllTransaksi: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var transaksi model.Transaksi
		if err := cursor.Decode(&transaksi); err != nil {
			fmt.Printf("GetAllTransaksi: %v\n", err)
			continue
		}
		transaksis = append(transaksis, transaksi)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllTransaksi: %v\n", err)
	}
	return transaksis
}

// update function
func UpdatePelanggan(db *mongo.Database, col string, id primitive.ObjectID, nama string, phonenumber string, alamat string, email []string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama":         nama,
			"phone_number": phonenumber,
			"alamat":       alamat,
			"email":        email,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdatePelanggan: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified ID")
		return
	}
	return nil
}
//function delete 
func DeletePelangganByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	Pelanggan := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := Pelanggan.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

//function untuk bagian customer 
func InsertCustomer(nama string, phoneNumber string, alamat string, email string, namaProduk string, deskripsi string, harga int, gambar string, stok string) interface{} {
	var customer model.Customer
	customer.ID = primitive.NewObjectID()
	customer.Nama = nama
	customer.Phone_number = phoneNumber
	customer.Alamat = alamat
	customer.Email = email
	customer.Nama_Produk = namaProduk
	customer.Deskripsi = deskripsi
	customer.Harga = harga
	customer.Gambar = gambar
	customer.Stok = stok
	return InsertOneDoc("kantin", "customer", customer)
}
func GetCustomerFromID(_id primitive.ObjectID, db *mongo.Database, col string) (customer model.Customer, errs error) {
	Customer := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := Customer.FindOne(context.TODO(), filter).Decode(&customer)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return customer, fmt.Errorf("no data found for ID %s", _id)
		}
		return customer, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return customer, nil
}

func GetCustomerByID(customerID primitive.ObjectID, db *mongo.Database, collectionName string) (customer model.Customer, err error) {
	collection := db.Collection(collectionName)
	filter := bson.M{"_id": customerID}
	err = collection.FindOne(context.TODO(), filter).Decode(&customer)
	if err != nil {
		fmt.Printf("GetCutomerByID: %v\n", err)
	}
	return customer, err
}

func GetAllCustomer() (customers [] model.Customer) {
	collection := MongoConnect("kantin").Collection("customer")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllCustomer: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var customer model.Customer
		if err := cursor.Decode(&customer); err != nil {
			fmt.Printf("GetAllCustomer: %v\n", err)
			continue
		}
		customers = append(customers, customer )
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllPelanggan: %v\n", err)
	}
	return customers
}

// function update dan delete untuk data customer 
//function update 
func UpdateCustomer(db *mongo.Database, col string, id primitive.ObjectID, nama string, phoneNumber string, alamat string, email string, namaProduk string, deskripsi string, harga int, gambar string, stok string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama":          nama,
			"phone_number":  phoneNumber,
			"alamat":        alamat,
			"email":         email,
			"nama_produk":   namaProduk,
			"deskripsi":     deskripsi,
			"harga":         harga,
			"gambar":        gambar,
			"stok":          stok,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateCustomer: %v\n", err)
		return 
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified ID")
		return 
	}
	return nil
}

// function delete
func DeleteCustomerByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	Customer := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := Customer.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

//insert barang\
func InsertBarang(namaProduk string, deskripsi string, harga int, gambar string, stok int) interface{} {
	var barang model.Barang
	barang.ID_barang = primitive.NewObjectID()
	barang.Nama_Produk = namaProduk
	barang.Deskripsi = deskripsi
	barang.Harga = harga
	barang.Gambar = gambar
	barang.Stok = stok
	return InsertOneDoc("kantin", "barang", barang)
}

// func InsertDataProduk(adminID string, nama_produk string, deskripsi string, harga int, gambar string, stok int) interface{} {
//     var produk model.Produk
//     produk.IDProduk = primitive.NewObjectID()
//     produk.AdminID = primitive.NewObjectID() // Assuming adminID is an ObjectID, change this if adminID is different
//     produk.Nama_Produk = nama_produk
//     produk.Deskripsi = deskripsi
//     produk.Harga = harga
//     produk.Gambar = gambar
//     produk.Stok = stok
//     return InsertOneDoc("kantin", "produk", produk)
// }