package katalogfilm

import (
	"fmt"
	"testing"
)

var privatekeykatalogfilm = ""
var publickeykatalogfilm = ""
var encode = ""

func TestGeneratePaseto(t *testing.T) {
	privateKey, publicKey := GenerateKey()
	fmt.Println("Private Key: " + privateKey)
	fmt.Println("Public Key: " + publicKey)
}

func TestEncode(t *testing.T) {
	name := "Test Nama"
	username := "Test Username"
	role := "Test Role"

	tokenstring, err := Encode(name, username, role, privatekeykatalogfilm)
	fmt.Println("error : ", err)
	fmt.Println("token : ", tokenstring)
}

func TestDecode(t *testing.T) {
	pay, err := Decode(publickeykatalogfilm, encode)
	name := DecodeGetName(publickeykatalogfilm, encode)
	username := DecodeGetUsername(publickeykatalogfilm, encode)
	role := DecodeGetRole(publickeykatalogfilm, encode)

	fmt.Println("name :", name)
	fmt.Println("username :", username)
	fmt.Println("role :", role)
	fmt.Println("err : ", err)
	fmt.Println("payload : ", pay)
}

func TestRegistrasi(t *testing.T) {
	mconn := SetConnection("mongoenvkatalogfilm", "katalogfilm")
	var user User
	user.Name = "Ibrohim"
	user.Username = "Ibrohim"
	user.Password = "Ibrohim"
	user.Role = "admin"
	hash, hashErr := HashPassword(user.Password)
	if hashErr != nil {
		fmt.Println(hashErr)
	}
	user.Password = hash
	InsertUser(mconn, "user", user)

	fmt.Println("Berhasil insert data user")
}

func TestGetAllUser(t *testing.T) {
	mconn := SetConnection("mongoenvkatalogfilm", "katalogfilm")
	datauser := GetAllUser(mconn, "user")

	fmt.Println(datauser)
}

func TestFindUser(t *testing.T) {
	mconn := SetConnection("mongoenvkatalogfilm", "katalogfilm")
	var user User
	user.Username = "Ibrohim"
	datauser := FindUser(mconn, "user", user)

	fmt.Println(datauser)
}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("mongoenvkatalogfilm", "katalogfilm")
	var user User
	user.Username = "Ibrohim"
	user.Password = "Ibrohim"
	datauser := IsPasswordValid(mconn, "user", user)

	fmt.Println(datauser)
}

func TestUsernameExists(t *testing.T) {
	var user User
	user.Username = "Ibrohim"
	datauser := UsernameExists("mongoenvkatalogfilm", "katalogfilm", user)

	fmt.Println(datauser)
}

func TestEditUser(t *testing.T) {
	mconn := SetConnection("mongoenvkatalogfilm", "katalogfilm")
	var user User
	user.Name = "Ibrohim"
	user.Username = "Ibrohim"
	datauser := FindUser(mconn, "user", user)
	user.Password = datauser.Password
	user.Role = "admin"
	EditUser(mconn, "user", user)

	fmt.Println("Berhasil edit data user")
}

func TestDeleteUser(t *testing.T) {
	mconn := SetConnection("mongoenvkatalogfilm", "katalogfilm")
	var user User
	user.Username = "xxx"
	DeleteUser(mconn, "user", user)

	fmt.Println("Berhasil hapus data user")
}
