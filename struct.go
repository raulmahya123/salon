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
	Name      string `json:"name" bson:"name"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Role      string `json:"role" bson:"role"`
	Nomor     string `json:"nomor" bson:"nomor"`
	Ispaid    bool   `json:"ispaid" bson:"ispaid"`       // Indicates whether the user has paid
	Hasaccess bool   `json:"hasaccess" bson:"hasaccess"` // Indicates whether access has been granted by admin
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
type PesanAnswer struct {
	Status         bool             `json:"status"`
	Message        string           `json:"message"`
	CorrectCount   int              `json:"correct_count"`
	IncorrectCount int              `json:"incorrect_count"`
	Details        []QuestionDetail `json:"details"`
	UserDetails    User             `json:"user_details,omitempty"`
	Certificate    string           `json:"certificate,omitempty"`
	Nomor          string           `json:"nomor,omitempty"`
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
	Username      string   `json:"username"`
}

type QuestionDetail struct {
	QuestionID string `json:"question_id"`
	IsCorrect  bool   `json:"is_correct"`
}
type AccessControl struct {
	Username  string `json:"username" bson:"username"`
	ContentID []int  `json:"content_id" bson:"content_id"`
	HasAccess bool   `json:"has_access" bson:"has_access"`
}

type VidioQuestion struct {
	ID            int      `json:"id" bson:"id"`
	Question      string   `json:"question" bson:"question"`
	Answers       []string `json:"answers" bson:"answers"`
	CorrectAnswer string   `json:"correct_answer" bson:"correct_answer"`
	Status        bool     `json:"status" bson:"status"`
	Video         []string `json:"video" bson:"video"`
	ContentID     []int    `json:"content_id" bson:"content_id"` // Added to link with AccessControl
	Username      string   `json:"username"`
	Nomor         string   `json:"nomor" bson:"nomor"`
}

type Content struct {
	ID          int    `json:"id" bson:"id"`
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
