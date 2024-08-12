package kursussalon

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ID           string `json:"id" bson:"id"`
	Name         string `json:"name" bson:"name"`
	Author       string `json:"author" bson:"author"`
	Salon1       string `json:"salon1" bson:"salon1"`
	Salon2       string `json:"salon2" bson:"salon2"`
	Salon3       string `json:"salon3" bson:"salon3"`
	Salon4       string `json:"salon4" bson:"salon4"`
	Salon5       string `json:"salon5" bson:"salon5"`
	Salon6       string `json:"salon6" bson:"salon6"`
	Salon7       string `json:"salon7" bson:"salon7"`
	Salon8       string `json:"salon8" bson:"salon8"`
	Salon9       string `json:"salon9" bson:"salon9"`
	Salon10      string `json:"salon10" bson:"salon10"`
	Salon11      string `json:"salon11" bson:"salon11"`
	Salon12      string `json:"salon12" bson:"salon12"`
	Status       bool   `json:"status" bson:"status"`
	Image        string `json:"image" bson:"image"`
	Nomor_claims string `json:"nomor_claims" bson:"nomor_claims"`
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

type Response struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type QuestionAndAnswer struct {
	ID            int      `json:"id" bson:"id"`
	Question      string   `json:"question" bson:"question"`
	Answers       []string `json:"answers" bson:"answers"`
	CorrectAnswer string   `json:"correct_answer" bson:"correct_answer"`
	Status        bool     `json:"status" bson:"status"`
}

type VidioQuestion struct {
	ID            int      `json:"id" bson:"id"`
	Question      string   `json:"question" bson:"question"`
	Answers       []string `json:"answers" bson:"answers"`
	CorrectAnswer string   `json:"correct_answer" bson:"correct_answer"`
	Status        bool     `json:"status" bson:"status"`
	Salon         []string `json:"salon" bson:"salon"`
}

type History struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Salon       []Salon            `json:"salon" bson:"salon"`
	Certificate []Certificate      `json:"certificate" bson:"certificate"`
}

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"id"`
	Nomorid     int                `json:"nomorid" bson:"nomorid"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       int                `json:"price" bson:"price"`
	Stock       int                `json:"stock" bson:"stock"`
	Size        string             `json:"size" bson:"size"`
	Image       string             `json:"image" bson:"image"`
	Status      string             `json:"status" bson:"status"`
}

type Productt struct {
	ID          string `json:"id" bson:"id"`
	Nomorid     int    `json:"nomorid" bson:"nomorid"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Price       int    `json:"price" bson:"price"`
	Stock       int    `json:"stock" bson:"stock"`
	Size        string `json:"size" bson:"size"`
	Image       string `json:"image" bson:"image"`
	Status      bool   `json:"status" bson:"status"`
}

type Content struct {
	ID          int    `json:"id" bson:"id" `
	Content     string `json:"content" bson:"content"`
	Image       string `json:"image" bson:"image"`
	Description string `json:"description" bson:"description"`
	Status      bool   `json:"status" bson:"status"`
}

type Comment struct {
	ID        int    `json:"id" bson:"id"`
	Username  string `json:"username" bson:"username"`
	Answer    string `json:"comment" bson:"comment"`
	Questions string `json:"questions" bson:"questions"`
	Tanggal   string `json:"tanggal" bson:"tanggal"`
	Status    bool   `json:"status" bson:"status"`
}

type Blog struct {
	ID                int    `json:"id" bson:"id"`
	Content           string `json:"content_one" bson:"content_one"`
	Content_two       string `json:"content_two" bson:"content_two"`
	Image             string `json:"image" bson:"image"`
	Title             string `json:"title" bson:"title"`
	Title_two         string `json:"title_two" bson:"title_two"`
	Description       string `json:"description" bson:"description"`
	Description_twoo  string `json:"description_two" bson:"description_two"`
	Description_three string `json:"description_3" bson:"description_3"`
	Status            bool   `json:"status" bson:"status"`
}
