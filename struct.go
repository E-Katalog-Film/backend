package katalogfilm

import (
	"time"
)

type Payload struct {
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Exp      time.Time `json:"exp"`
	Iat      time.Time `json:"iat"`
	Nbf      time.Time `json:"nbf"`
}

type User struct {
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type CredentialUser struct {
	Status bool `json:"status" bson:"status"`
	Data   struct {
		Name     string `json:"name" bson:"name"`
		Username string `json:"username" bson:"username"`
		Role     string `json:"role" bson:"role"`
	} `json:"data" bson:"data"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Pesan struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data,omitempty" bson:"data,omitempty"`
	Role    string      `json:"role,omitempty" bson:"role,omitempty"`
	Token   string      `json:"token,omitempty" bson:"token,omitempty"`
}

type Film struct {
	ID        string `json:"id" bson:"id"`
	Judul     string `json:"judul" bson:"judul"`
	Image     string `json:"image" bson:"image"`
	Tanggal   string `json:"tanggal" bson:"tanggal"`
	Genre     string `json:"genre" bson:"genre"`
	Sinopsis  string `json:"sinopsis" bson:"sinopsis"`
	Penulis   string `json:"penulis" bson:"penulis"`
	Sutradara string `json:"sutradara" bson:"sutradara"`
	Aktor     string `json:"aktor" bson:"aktor"`
}

type Komentar struct {
	ID       string `json:"id" bson:"id"`
	ID_Film  string `json:"id_film" bson:"id_film"`
	Name     string `json:"name" bson:"name"`
	Tanggal  string `json:"tanggal" bson:"tanggal"`
	Komentar string `json:"komentar" bson:"komentar"`
}

type Rating struct {
	ID       string `json:"id" bson:"id"`
	ID_Film  string `json:"id_film" bson:"id_film"`
	Username string `json:"username" bson:"username"`
	Rating   int    `json:"rating" bson:"rating"`
	Kualitas string `json:"kualitas" bson:"kualitas"`
	Note     string `json:"note" bson:"note"`
	Tanggal  string `json:"tanggal" bson:"tanggal"`
}
