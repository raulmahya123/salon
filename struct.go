package kursussalon

import (
	"time"
)

type Payload struct {
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Nomor    string    `json:"nomor"`
	Exp      time.Time `json:"exp"`
	Iat      time.Time `json:"iat"`
	Nbf      time.Time `json:"nbf"`
}

type User struct {
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
	Nomor    string `json:"nomor" bson:"nomor"`
}

type CredentialUser struct {
	Status bool `json:"status" bson:"status"`
	Data   struct {
		Name     string `json:"name" bson:"name"`
		Username string `json:"username" bson:"username"`
		Role     string `json:"role" bson:"role"`
		Nomor    string `json:"nomor" bson:"nomor"`
	} `json:"data" bson:"data"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Pesan struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data,omitempty" bson:"data,omitempty"`
	Role    string      `json:"role,omitempty" bson:"role,omitempty"`
	Token   string      `json:"token,omitempty" bson:"token,omitempty"`
	Nomor   string      `json:"nomor,omitempty" bson:"nomor,omitempty"`
}
type Salon struct {
	ID      string `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Author  string `json:"author" bson:"author"`
	Salon1  string `json:"salon1" bson:"salon1"`
	Salon2  string `json:"salon2" bson:"salon2"`
	Salon3  string `json:"salon3" bson:"salon3"`
	Salon4  string `json:"salon4" bson:"salon4"`
	Salon5  string `json:"salon5" bson:"salon5"`
	Salon6  string `json:"salon6" bson:"salon6"`
	Salon7  string `json:"salon7" bson:"salon7"`
	Salon8  string `json:"salon8" bson:"salon8"`
	Salon9  string `json:"salon9" bson:"salon9"`
	Salon10 string `json:"salon10" bson:"salon10"`
	Salon11 string `json:"salon11" bson:"salon11"`
	Salon12 string `json:"salon12" bson:"salon12"`
	Status  bool   `json:"status" bson:"status"`
}

type Certificate struct {
	Nama             string `json:"nama" bson:"nama"`
	Nomorcertificate string `json:"nomorcertificate" bson:"nomorcertificate"`
	Tanggal          string `json:"tanggal" bson:"tanggal"`
	Expired          string `json:"expired" bson:"expired"`
	Jurusan          string `json:"jurusan" bson:"jurusan"`
	Status           bool   `json:"status" bson:"status"`
	Nomor            string `json:"nomor" bson:"nomor"`
	Ttd              string `json:"ttd" bson:"ttd"`
}

type Blog struct {
	ID      string `json:"id" bson:"id"`
	Title   string `json:"title" bson:"title"`
	Author  string `json:"author" bson:"author"`
	Content string `json:"content" bson:"content"`
	Status  bool   `json:"status" bson:"status"`
	Image   string `json:"image" bson:"image"`
}

type Response struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}
