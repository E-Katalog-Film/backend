package katalogfilm

import (
	"context"
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(mongoenvkatalogfilm, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(mongoenvkatalogfilm),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

//------------------------------------------------------------------- User

// Create

func InsertUser(mconn *mongo.Database, collname string, datauser User) interface{} {
	return atdb.InsertOneDoc(mconn, collname, datauser)
}

// Read

func GetAllUser(mconn *mongo.Database, collname string) []User {
	user := atdb.GetAllDoc[[]User](mconn, collname)
	return user
}

func FindUser(mconn *mongo.Database, collname string, userdata User) User {
	filter := bson.M{"username": userdata.Username}
	return atdb.GetOneDoc[User](mconn, collname, filter)
}

func IsPasswordValid(mconn *mongo.Database, collname string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, collname, filter)
	hashChecker := CheckPasswordHash(userdata.Password, res.Password)
	return hashChecker
}

func UsernameExists(mongoenvkatalogfilm, dbname string, userdata User) bool {
	mconn := SetConnection(mongoenvkatalogfilm, dbname).Collection("user")
	filter := bson.M{"username": userdata.Username}

	var user User
	err := mconn.FindOne(context.Background(), filter).Decode(&user)
	return err == nil
}

// Update

func EditUser(mconn *mongo.Database, collname string, datauser User) interface{} {
	filter := bson.M{"username": datauser.Username}
	return atdb.ReplaceOneDoc(mconn, collname, filter, datauser)
}

// Delete

func DeleteUser(mconn *mongo.Database, collname string, userdata User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

//------------------------------------------------------------------- Film

// Create

func InsertFilm(mconn *mongo.Database, collname string, datafilm Film) interface{} {
	return atdb.InsertOneDoc(mconn, collname, datafilm)
}

// Read

func GetAllFilm(mconn *mongo.Database, collname string) []Film {
	film := atdb.GetAllDoc[[]Film](mconn, collname)
	return film
}

func FindFilm(mconn *mongo.Database, collname string, datafilm Film) Film {
	filter := bson.M{"id": datafilm.ID}
	return atdb.GetOneDoc[Film](mconn, collname, filter)
}

func IdFilmExists(mongoenvkatalogfilm, dbname string, datafilm Film) bool {
	mconn := SetConnection(mongoenvkatalogfilm, dbname).Collection("film")
	filter := bson.M{"id": datafilm.ID}

	var film Film
	err := mconn.FindOne(context.Background(), filter).Decode(&film)
	return err == nil
}

// Update

func EditFilm(mconn *mongo.Database, collname string, datafilm Film) interface{} {
	filter := bson.M{"id": datafilm.ID}
	return atdb.ReplaceOneDoc(mconn, collname, filter, datafilm)
}

// Delete

func DeleteFilm(mconn *mongo.Database, collname string, datafilm Film) interface{} {
	filter := bson.M{"id": datafilm.ID}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

//------------------------------------------------------------------- Komentar

// Create

func InsertKomentar(mconn *mongo.Database, collname string, datakomentar Komentar) interface{} {
	return atdb.InsertOneDoc(mconn, collname, datakomentar)
}

// Read

func GetAllKomentar(mconn *mongo.Database, collname string) []Komentar {
	film := atdb.GetAllDoc[[]Komentar](mconn, collname)
	return film
}

func FindKomentar(mconn *mongo.Database, collname string, datakomentar Komentar) Komentar {
	filter := bson.M{"id": datakomentar.ID}
	return atdb.GetOneDoc[Komentar](mconn, collname, filter)
}

func IdKomentarExists(mongoenvkatalogfilm, dbname string, datakomentar Komentar) bool {
	mconn := SetConnection(mongoenvkatalogfilm, dbname).Collection("komentar")
	filter := bson.M{"id": datakomentar.ID}

	var komentar Komentar
	err := mconn.FindOne(context.Background(), filter).Decode(&komentar)
	return err == nil
}

// Update

func EditKomentar(mconn *mongo.Database, collname string, datakomentar Komentar) interface{} {
	filter := bson.M{"id": datakomentar.ID}
	return atdb.ReplaceOneDoc(mconn, collname, filter, datakomentar)
}

// Delete

func DeleteKomentar(mconn *mongo.Database, collname string, datakomentar Komentar) interface{} {
	filter := bson.M{"id": datakomentar.ID}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

//------------------------------------------------------------------- Rating

// Create

func InsertRating(mconn *mongo.Database, collname string, datarating Rating) interface{} {
	return atdb.InsertOneDoc(mconn, collname, datarating)
}

// Read

func GetAllRating(mconn *mongo.Database, collname string) []Rating {
	film := atdb.GetAllDoc[[]Rating](mconn, collname)
	return film
}

func FindRating(mconn *mongo.Database, collname string, datarating Rating) Rating {
	filter := bson.M{"id": datarating.ID}
	return atdb.GetOneDoc[Rating](mconn, collname, filter)
}

func IdRatingExists(mongoenvkatalogfilm, dbname string, datarating Rating) bool {
	mconn := SetConnection(mongoenvkatalogfilm, dbname).Collection("rating")
	filter := bson.M{"id": datarating.ID}

	var rating Rating
	err := mconn.FindOne(context.Background(), filter).Decode(&rating)
	return err == nil
}

func RatingFilmExists(mongoenvkatalogfilm, dbname string, datarating Rating) bool {
	mconn := SetConnection(mongoenvkatalogfilm, dbname).Collection("rating")
	filter := bson.M{"id_film": datarating.ID_Film, "username": datarating.Username}

	var rating Rating
	err := mconn.FindOne(context.Background(), filter).Decode(&rating)
	return err == nil
}

// Update

func EditRating(mconn *mongo.Database, collname string, datarating Rating) interface{} {
	filter := bson.M{"id": datarating.ID}
	return atdb.ReplaceOneDoc(mconn, collname, filter, datarating)
}

// Delete

func DeleteRating(mconn *mongo.Database, collname string, datarating Rating) interface{} {
	filter := bson.M{"id": datarating.ID}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}
