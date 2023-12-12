package katalogfilm

import (
	"encoding/json"
	"net/http"
	"os"
)

func ReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

//------------------------------------------------------------------- User

func Authorization(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response CredentialUser
	var auth User
	response.Status = false

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}

	tokenname := DecodeGetName(os.Getenv(publickeykatalogfilm), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)

	auth.Username = tokenusername

	if tokenname == "" || tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !usernameExists(mongoenvkatalogfilm, dbname, auth) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	response.Message = "Berhasil decode token"
	response.Status = true
	response.Data.Name = tokenname
	response.Data.Username = tokenusername
	response.Data.Role = tokenrole

	return ReturnStruct(response)
}

func Registrasi(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	if usernameExists(mongoenvkatalogfilm, dbname, user) {
		response.Message = "Username telah dipakai"
		return ReturnStruct(response)
	}

	hash, hashErr := HashPassword(user.Password)
	if hashErr != nil {
		response.Message = "Gagal hash password: " + hashErr.Error()
		return ReturnStruct(response)
	}

	user.Password = hash

	InsertUser(mconn, collname, user)
	response.Status = true
	response.Message = "Berhasil input data"

	return ReturnStruct(response)
}

func Login(privatekeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	if !usernameExists(mongoenvkatalogfilm, dbname, user) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if !IsPasswordValid(mconn, collname, user) {
		response.Message = "Password Salah"
		return ReturnStruct(response)
	}

	auth := FindUser(mconn, collname, user)

	tokenstring, tokenerr := Encode(auth.Name, auth.Username, auth.Role, os.Getenv(privatekeykatalogfilm))
	if tokenerr != nil {
		response.Message = "Gagal encode token: " + tokenerr.Error()
		return ReturnStruct(response)
	}

	response.Status = true
	response.Message = "Berhasil login"
	response.Token = tokenstring

	return ReturnStruct(response)
}

func AmbilSemuaUser(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}

	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !usernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}

	datauser := GetAllUser(mconn, collname)
	return ReturnStruct(datauser)
}

func UpdateUser(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}

	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !usernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}

	if user.Username == "" {
		response.Message = "Parameter dari function ini adalah username"
		return ReturnStruct(response)
	}

	if !usernameExists(mongoenvkatalogfilm, dbname, user) {
		response.Message = "Akun yang ingin diedit tidak ditemukan"
		return ReturnStruct(response)
	}

	if user.Password != "" {
		hash, hashErr := HashPassword(user.Password)
		if hashErr != nil {
			response.Message = "Gagal Hash Password: " + hashErr.Error()
			return ReturnStruct(response)
		}
		user.Password = hash
	} else {
		olduser := FindUser(mconn, collname, user)
		user.Password = olduser.Password
	}

	EditUser(mconn, collname, user)

	response.Status = true
	response.Message = "Berhasil update " + user.Username + " dari database"
	return ReturnStruct(response)
}

func HapusUser(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}

	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !usernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}

	if user.Username == "" {
		response.Message = "Parameter dari function ini adalah username"
		return ReturnStruct(response)
	}

	if !usernameExists(mongoenvkatalogfilm, dbname, user) {
		response.Message = "Akun yang ingin dihapus tidak ditemukan"
		return ReturnStruct(response)
	}

	DeleteUser(mconn, collname, user)

	response.Status = true
	response.Message = "Berhasil hapus " + user.Username + " dari database"
	return ReturnStruct(response)
}

//------------------------------------------------------------------- Film

func TambahFilm(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var film Film
	err := json.NewDecoder(r.Body).Decode(&film)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}

	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !usernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}

	if film.ID == "" {
		response.Message = "ID dibutuhkan untuk membuat datafilm"
		return ReturnStruct(response)
	}

	if idFilmExists(mongoenvkatalogfilm, dbname, film) {
		response.Message = "ID yang ingin digunakan telah digunakan"
		return ReturnStruct(response)
	}

	InsertFilm(mconn, collname, film)
	response.Status = true
	response.Message = "Berhasil input data"

	return ReturnStruct(response)
}

func AmbilSemuaFilm(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}

	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !usernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		if tokenrole != "user" {
			response.Message = "Anda tidak memiliki akses"
			return ReturnStruct(response)
		}
	}

	datafilm := GetAllFilm(mconn, collname)
	return ReturnStruct(datafilm)
}

func AmbilSatuFilm(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var film Film
	err := json.NewDecoder(r.Body).Decode(&film)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}

	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !usernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		if tokenrole != "user" {
			response.Message = "Anda tidak memiliki akses"
			return ReturnStruct(response)
		}
	}

	if film.ID == "" {
		response.Message = "ID dibutuhkan untuk memanggil datafilm"
		return ReturnStruct(response)
	}

	if !idFilmExists(mongoenvkatalogfilm, dbname, film) {
		response.Message = "Film tidak ditemukan"
		return ReturnStruct(response)
	}

	datafilm := FindFilm(mconn, collname, film)
	return ReturnStruct(datafilm)
}

func UpdateFilm(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var film Film
	err := json.NewDecoder(r.Body).Decode(&film)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}

	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !usernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}

	if film.ID == "" {
		response.Message = "ID dibutuhkan untuk update datafilm"
		return ReturnStruct(response)
	}

	if !idFilmExists(mongoenvkatalogfilm, dbname, film) {
		response.Message = "Film tidak ditemukan"
		return ReturnStruct(response)
	}

	EditFilm(mconn, collname, film)
	response.Status = true
	response.Message = "Berhasil update data"

	return ReturnStruct(response)
}

func HapusFilm(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var film Film
	err := json.NewDecoder(r.Body).Decode(&film)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}

	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)

	if tokenusername == "" || tokenrole == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !usernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}

	if film.ID == "" {
		response.Message = "ID dibutuhkan untuk hapus datafilm"
		return ReturnStruct(response)
	}

	if !idFilmExists(mongoenvkatalogfilm, dbname, film) {
		response.Message = "Film tidak ditemukan"
		return ReturnStruct(response)
	}

	DeleteFilm(mconn, collname, film)
	response.Status = true
	response.Message = "Berhasil hapus data"

	return ReturnStruct(response)
}
