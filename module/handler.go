package module

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/serlip06/pointsalesofkantin/model"
	

	//"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"encoding/json"
	//"golang.org/x/crypto/bcrypt"
)


func MongoConnectdb(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

//function untuk meyimpan data registrasi
func SavePendingRegistration(registration model.PendingRegistration, db*mongo.Database) error {
    collection := db.Collection("pending_registrations")
    _, err := collection.InsertOne(context.Background(), registration)
    return err
}

//memindahkan data pending ke data users(function untuk ACC)
func ApproveRegistration(id string, db *mongo.Database) (model.PendingRegistration, model.User, error) {
	// function yang dipake untuk mindahil data progress ke colekcion pengguna 
	collectionPending := db.Collection("pending_registrations")
	collectionUsers := db.Collection("users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.PendingRegistration{}, model.User{}, err
	}

	var pending model.PendingRegistration
	err = collectionPending.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&pending)
	if err != nil {
		return model.PendingRegistration{}, model.User{}, err
	}

	user := model.User{
		Username:  pending.Username,
		Password:  pending.Password,
		Role:      pending.Role,
		CreatedAt: time.Now(),
	}

	_, err = collectionUsers.InsertOne(context.Background(), user)
	if err != nil {
		return model.PendingRegistration{}, model.User{}, err
	}

	_, err = collectionPending.DeleteOne(context.Background(), bson.M{"_id": objID})
	return pending, user, err
}

// Register handler
func RegisterHandler(req model.RegisterRequest, db *mongo.Database) (string, error) {
	// Simpan password mentah tanpa di-hash
	password := req.Password

	// Membuat objek registrasi
	registration := model.PendingRegistration{
		Username:    req.Username,
		Password:    password,  // Simpan password dalam bentuk asli
		Role:        req.Role,
		SubmittedAt: time.Now(),
	}

	// Panggil fungsi untuk menyimpan data pengguna
	err := SavePendingRegistration(registration, db)
	if err != nil {
		return "", fmt.Errorf("failed to save registration: %v", err)
	}

	// Return success message
	return "Registration submitted, waiting for admin approval", nil
}


// function untuk memanggil data di colecction pending_registration dan user 
// GetAllUsers retrieves all user data from the users collection
func GetAllUsers(db *mongo.Database) ([]model.User, error) {
	collection := db.Collection("users")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []model.User
	for cursor.Next(context.Background()) {
		var user model.User
		var raw bson.M
		if err := cursor.Decode(&raw); err != nil {
			return nil, err
		}

		// Extract ObjectID and convert to string
		if objID, ok := raw["_id"].(primitive.ObjectID); ok {
			raw["_id"] = objID.Hex()
		}

		// Map raw data to User struct
		data, err := bson.Marshal(raw)
		if err != nil {
			return nil, err
		}
		if err := bson.Unmarshal(data, &user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Fungsi untuk mendapatkan pengguna berdasarkan ID
func GetUserByID(userID string, db *mongo.Database) (*model.User, error) {
	collection := db.Collection("users")
	var user model.User

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %v", err)
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user dengan ID %s tidak ditemukan", userID)
		}
		return nil, err
	}

	return &user, nil
}

func GetUserFromID(_id primitive.ObjectID, db *mongo.Database, col string) (user model.User, errs error) {
	User := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := User.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, fmt.Errorf("no data found for ID %s", _id)
		}
		return user, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return user, nil
}

// GetAllPendingRegistrations retrieves all pending registrations from the pending_registrations collection
func GetAllPendingRegistrations(db *mongo.Database) ([]model.PendingRegistration, error) {
	collection := db.Collection("pending_registrations")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var registrations []model.PendingRegistration
	for cursor.Next(context.Background()) {
		var registration model.PendingRegistration
		var raw bson.M
		if err := cursor.Decode(&raw); err != nil {
			return nil, err
		}

		// Extract ObjectID and convert to string
		if objID, ok := raw["_id"].(primitive.ObjectID); ok {
			raw["_id"] = objID.Hex()
		}

		// Map raw data to PendingRegistration struct
		data, err := bson.Marshal(raw)
		if err != nil {
			return nil, err
		}
		if err := bson.Unmarshal(data, &registration); err != nil {
			return nil, err
		}

		registrations = append(registrations, registration)
	}

	return registrations, nil
}