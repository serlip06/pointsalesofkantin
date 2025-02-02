package module

import (
	"context"
	"errors"
	"fmt"
	"github.com/serlip06/pointsalesofkantin/model" // Mengimpor package model dengan benar
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Fungsi untuk menghubungkan ke MongoDB
func MongoConnectdatabase(dbname string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
		return nil, err
	}
	return client.Database(dbname), nil
}

// Insert item ke keranjang
func InsertDataCartItem(dbname, collection string, doc interface{}) (interface{}, error) {
	insertResult, err := MongoConnectdatabase(dbname)
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
func InsertDataCartItemFunc(db *mongo.Database, idProduk primitive.ObjectID, idUser primitive.ObjectID, quantity int) (interface{}, error) {
	product, err := GetProduksFromID(idProduk, db, "produk")
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %v", err)
	}

	subTotal := product.Harga * quantity

	var cartItem model.CartItem
	cartItem.IDCartItem = primitive.NewObjectID()
	cartItem.IDProduk = idProduk
	cartItem.IDUser = idUser
	cartItem.Harga = product.Harga
	cartItem.Quantity = quantity
	cartItem.SubTotal = subTotal
	cartItem.Nama_Produk = product.Nama_Produk
	cartItem.Gambar = product.Gambar
	cartItem.IsSelected = false // Defaultnya tidak dicentang

	result, err := InsertDataCartItem("kantin", "cart_items", cartItem)
	if err != nil {
		return nil, fmt.Errorf("failed to insert cart item: %v", err)
	}

	return result, nil
}

// fungtion update bersasarkkan item yang di pilih
func UpdateCartItemSelection(db *mongo.Database, idCartItems []primitive.ObjectID, isSelected bool) error {
	collection := db.Collection("cart_items")
	filter := bson.M{"_id": bson.M{"$in": idCartItems}}

	update := bson.M{
		"$set": bson.M{
			"is_selected": isSelected,
		},
	}

	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update cart items: %v", err)
	}

	return nil
}

// funstion untuk checkout dari keranjang
func CheckoutFromCart(db *mongo.Database, idUser primitive.ObjectID, metodePembayaran, buktiPembayaran, alamat string) (primitive.ObjectID, error) {
	collectionCart := db.Collection("cart_items")
	filter := bson.M{"id_user": idUser, "is_selected": true}

	cursor, err := collectionCart.Find(context.TODO(), filter)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to get selected cart items: %v", err)
	}
	defer cursor.Close(context.TODO())

	var items []model.CartItem
	var idCartItems []primitive.ObjectID
	totalHarga := 0

	for cursor.Next(context.TODO()) {
		var item model.CartItem
		if err := cursor.Decode(&item); err != nil {
			return primitive.NilObjectID, fmt.Errorf("error decoding cart item: %v", err)
		}
		items = append(items, item)
		idCartItems = append(idCartItems, item.IDCartItem)
		totalHarga += item.SubTotal
	}

	// Pastikan ada item yang dipilih
	if len(items) == 0 {
		return primitive.NilObjectID, errors.New("no selected items to checkout")
	}

	// Buat transaksi baru
	transaksi := model.Transaksi{
		IDTransaksi:      primitive.NewObjectID(),
		IDUser:           idUser,
		IDCartItem:       idCartItems,
		MetodePembayaran: metodePembayaran,
		TotalHarga:       totalHarga,
		BuktiPembayaran:  buktiPembayaran,
		Alamat:           alamat,
		CreatedAt:        time.Now(),
		Status:           "Pending",
	}

	collectionTransaksi := db.Collection("kantin_transaksi")
	_, err = collectionTransaksi.InsertOne(context.TODO(), transaksi)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to create transaction: %v", err)
	}

	// Hapus item dari cart setelah transaksi berhasil
	_, err = collectionCart.DeleteMany(context.TODO(), filter)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to remove items from cart: %v", err)
	}

	return transaksi.IDTransaksi, nil
}

// Get cart item by ID
func GetCartItemFromID(_id primitive.ObjectID, db *mongo.Database, col string) (model.CartItem, error) {
	cartItem := model.CartItem{}
	filter := bson.M{"_id": _id}

	fmt.Println("Looking for _id:", _id.Hex()) // Log ID yang dicari

	err := db.Collection(col).FindOne(context.TODO(), filter).Decode(&cartItem)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return cartItem, fmt.Errorf("no data found for ID %s", _id.Hex())
		}
		return cartItem, fmt.Errorf("error retrieving data for ID %s: %s", _id.Hex(), err.Error())
	}

	fmt.Println("Found Cart Item:", cartItem) // Log data yang ditemukan

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
func UpdateCartItem(db *mongo.Database, col string, id primitive.ObjectID, idUser, nama string, harga int, quantity int, gambar string) error {
	filter := bson.M{"_id": id, "id_user": idUser} // fitler untuk user
	update := bson.M{
		"$set": bson.M{
			"nama":      nama,
			"harga":     harga,
			"quantity":  quantity,
			"sub_total": harga * quantity, // Menghitung ulang subtotal
			"gambar":    gambar,           //mengupdate data gambar

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
func DeleteCartItemByID(_id primitive.ObjectID, idUser primitive.ObjectID, db *mongo.Database, col string) error {
	collection := db.Collection("cart_items")
	filter := bson.M{"_id": _id, "id_user": idUser}// _id dari chartitem

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting cart item: %v", err)
	}

	if result.DeletedCount == 0 {
		return errors.New("cart item not found or does not belong to user")
	}

	return nil
}

// logika untuk chartitem jaga jaga aja
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
