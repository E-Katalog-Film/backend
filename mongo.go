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

func usernameExists(mongoenvkatalogfilm, dbname string, userdata User) bool {
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

func idFilmExists(mongoenvkatalogfilm, dbname string, datafilm Film) bool {
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
