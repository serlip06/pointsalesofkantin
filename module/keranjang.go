package module

import (
	"context"
	"errors"
	"fmt"
	"github.com/serlip06/pointsalesofkantin/model" // Mengimpor package model dengan benar
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Fungsi untuk menghubungkan ke MongoDB
func  MongoConnectdatabase(dbname string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
		return nil, err
	}
	return client.Database(dbname), nil
}

// Insert item ke keranjang 
func InsertDataCartItem(dbname, collection string, doc interface{}) (interface{}, error) {
	insertResult, err :=  MongoConnectdatabase(dbname)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	result, err := insertResult.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
		return nil, err
	}
	return result.InsertedID, nil
}

// Fungsi untuk menambahkan item ke keranjang
func InsertDataCartItemFunc(db *mongo.Database, idProduk primitive.ObjectID,idUser primitive.ObjectID, quantity int) (interface{}, error) {
	// Mengambil data produk berdasarkan IDProduk
	product, err := GetProduksFromID(idProduk, db, "produk") // Menggunakan fungsi yang kamu buat
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %v", err)
	}

	// Menghitung subtotal
	subTotal := product.Harga * quantity

	// Buat item keranjang
	var cartItem model.CartItem
	cartItem.IDProduk = idProduk
	cartItem.IDUser = idUser 
	cartItem.Harga = product.Harga
	cartItem.Quantity = quantity
	cartItem.SubTotal = subTotal
	cartItem.Nama_Produk = product.Nama_Produk // Menyimpan nama produk jika diperlukan
	cartItem.Gambar = product.Gambar// meyimpan data gambar 

	// Menyimpan item ke dalam keranjang
	result, err := InsertDataCartItem("kantin", "cart_items", cartItem)
	if err != nil {
		return nil, fmt.Errorf("failed to insert cart item: %v", err)
	}

	return result, nil
}

// Get cart item by ID 
func GetCartItemFromID(_id primitive.ObjectID, db *mongo.Database, col string) (model.CartItem, error) {
    cartItem := model.CartItem{}
    filter := bson.M{"_id": _id}
    
    fmt.Println("Looking for _id:", _id.Hex())  // Log ID yang dicari
    
    err := db.Collection(col).FindOne(context.TODO(), filter).Decode(&cartItem)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return cartItem, fmt.Errorf("no data found for ID %s", _id.Hex())
        }
        return cartItem, fmt.Errorf("error retrieving data for ID %s: %s", _id.Hex(), err.Error())
    }
    
    fmt.Println("Found Cart Item:", cartItem)  // Log data yang ditemukan
    
    return cartItem, nil
}

// Get all cart items
func GetAllCartItems() (cartitems []model.CartItem) {
	// Menangani nilai error yang dikembalikan oleh MongoConnectdatabase
	db, err := MongoConnectdatabase("kantin")
	if err != nil {
		fmt.Printf("GetAllCartItem: %v\n", err)
		return nil
	}


	// Mengakses collection dari database yang berhasil terhubung
	collection := db.Collection("cart_items")

	// Melakukan query ke collection
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllCartItem: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())

	// Menyimpan hasil query ke dalam slice cartitems
	for cursor.Next(context.Background()) {
		var CartItem model.CartItem
		if err := cursor.Decode(&CartItem); err != nil {
			fmt.Printf("GetAllCartItem: %v\n", err)
			continue
		}
		cartitems = append(cartitems, CartItem)
	}

	// Menangani jika ada error selama iterasi cursor
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllCartItem: %v\n", err)
	}

	// Mengembalikan hasil
	return cartitems
}

// Update cart item
func UpdateCartItem(db *mongo.Database, col string, id primitive.ObjectID,idUser, nama string, harga int, quantity int, gambar string) error {
	filter := bson.M{"_id": id, "id_user" : idUser}// fitler untuk user
	update := bson.M{
		"$set": bson.M{
			"nama":      nama,
			"harga":     harga,
			"quantity":  quantity,
			"sub_total": harga * quantity, // Menghitung ulang subtotal
			"gambar":    gambar,    //mengupdate data gambar

		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateCartItem: %v\n", err)
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}
	return nil
}

// Delete cart item
func DeleteCartItemByID(_id primitive.ObjectID,idUser, db *mongo.Database, col string) error {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id, "id_user" : idUser}// filter untuk user 

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

//logika untuk chartitem jaga jaga aja 
func InsertOrUpdateCartItem(db *mongo.Database, idProduk, idUser primitive.ObjectID, quantity int) (interface{}, error) {
	collection := db.Collection("cart_items")

	// Cek apakah produk dengan IDProduk dan IDUser yang sama sudah ada di keranjang
	filter := bson.M{"id_produk": idProduk, "id_user": idUser}
	var existingItem model.CartItem
	err := collection.FindOne(context.TODO(), filter).Decode(&existingItem)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Jika tidak ada, tambahkan sebagai item baru
			product, err := GetProduksFromID(idProduk, db, "produk")
			if err != nil {
				return nil, fmt.Errorf("failed to get product: %v", err)
			}

			// Hitung subtotal
			subTotal := product.Harga * quantity

			// Buat item keranjang baru
			newItem := model.CartItem{
				IDCartItem:  primitive.NewObjectID(),
				IDProduk:    idProduk,
				IDUser:      idUser, // Tambahkan IDUser
				Nama_Produk: product.Nama_Produk,
				Harga:       product.Harga,
				Quantity:    quantity,
				SubTotal:    subTotal,
				Gambar:      product.Gambar,
			}

			// Simpan ke database
			result, err := collection.InsertOne(context.TODO(), newItem)
			if err != nil {
				return nil, fmt.Errorf("failed to insert cart item: %v", err)
			}

			return result.InsertedID, nil
		}
		return nil, fmt.Errorf("error finding cart item: %v", err)
	}

	// Jika item sudah ada, update quantity dan subtotal
	newQuantity := existingItem.Quantity + quantity
	newSubTotal := newQuantity * existingItem.Harga

	update := bson.M{
		"$set": bson.M{
			"quantity":  newQuantity,
			"sub_total": newSubTotal,
		},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update cart item: %v", err)
	}

	return existingItem.IDCartItem, nil
}



// ini filter untuk kategorinya 
// func GetAllCartItems(kategori string) ([]model.CartItem, error) {
// 	collection, err :=  MongoConnectdatabase("kantin")
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to database: %v", err)
// 	}

// 	// Filter berdasarkan kategori jika diberikan
// 	var filter bson.M
// 	if kategori != "" {
// 		filter = bson.M{"kategori": kategori}
// 	} else {
// 		filter = bson.M{}
// 	}

// 	cursor, err := collection.Collection("cart_items").Find(context.TODO(), filter)
// 	if err != nil {
// 		return nil, fmt.Errorf("GetAllCartItems: %v", err)
// 	}
// 	defer cursor.Close(context.TODO())

// 	var cartItems []model.CartItem
// 	for cursor.Next(context.Background()) {
// 		var item model.CartItem
// 		if err := cursor.Decode(&item); err != nil {
// 			fmt.Printf("GetAllCartItems: %v\n", err)
// 			continue
// 		}
// 		cartItems = append(cartItems, item)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil, fmt.Errorf("GetAllCartItems: %v", err)
// 	}
// 	return cartItems, nil
// }
