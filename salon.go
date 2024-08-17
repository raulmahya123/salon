package kursussalon

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
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
	tokennomor := DecodeGetNomor(os.Getenv(publickeykatalogfilm), header)
	auth.Username = tokenusername

	if tokenname == "" || tokenusername == "" || tokenrole == "" || tokennomor == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, auth) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	response.Message = "Berhasil decode token"
	response.Status = true
	response.Data.Name = tokenname
	response.Data.Username = tokenusername
	response.Data.Role = tokenrole
	response.Data.Nomor = tokennomor

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

	if UsernameExists(mongoenvkatalogfilm, dbname, user) {
		response.Message = "Username telah dipakai"
		return ReturnStruct(response)
	}

	hash, hashErr := HashPassword(user.Password)
	if hashErr != nil {
		response.Message = "Gagal hash password: " + hashErr.Error()
		return ReturnStruct(response)
	}

	//generate nomor random
	user.Nomor = GenerateRandomNumber()
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

	if !UsernameExists(mongoenvkatalogfilm, dbname, user) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if !IsPasswordValid(mconn, collname, user) {
		response.Message = "Password Salah"
		return ReturnStruct(response)
	}

	auth := FindUser(mconn, collname, user)

	tokenstring, tokenerr := Encode(auth.Name, auth.Username, auth.Role, auth.Nomor, os.Getenv(privatekeykatalogfilm))
	if tokenerr != nil {
		response.Message = "Gagal encode token: " + tokenerr.Error()
		return ReturnStruct(response)
	}

	response.Status = true
	response.Message = "Berhasil login"
	response.Token = tokenstring
	response.Role = auth.Role
	response.Nomor = auth.Nomor

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

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
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

func UpdateUser(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	if user.Username == "" {
		response.Message = "Parameter dari function ini adalah username"
		return ReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, user) {
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

func HapusUser(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	DeleteUser(mconn, collname, user)

	response.Status = true
	response.Message = "Berhasil hapus " + user.Username + " dari database"
	return ReturnStruct(response)
}

func UpdatePassword(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	auth := FindUser(mconn, collname, user)

	if auth.Username == "" {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, user) {
		response.Message = "Akun yang ingin diedit tidak ditemukan"
		return ReturnStruct(response)
	}
	findpassword := FindPassword(mconn, collname, user)
	if auth.Password == findpassword.Password {
		response.Message = "Password sama dengan yang lama"
		return ReturnStruct(response)
	}
	hash, hashErr := HashPassword(user.Password)
	if hashErr != nil {
		response.Message = "Gagal hash password: " + hashErr.Error()
		return ReturnStruct(response)
	}

	user.Name = user.Username
	user.Role = "user"
	HashPassword(user.Password)
	user.Password = hash
	EditUser(mconn, collname, user)

	response.Status = true
	response.Message = "Berhasil update password " + user.Username + " dari database"
	return ReturnStruct(response)

}

// certificate salon
func FindCertificate(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	datafilm := FindallCertificate(mconn, collname)
	var response Pesan
	response.Status = true
	response.Message = "Berhasil ambil data"
	response.Data = datafilm
	return ReturnStruct(response)
}

func AddedCertificate(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var certificate Certificate
	err := json.NewDecoder(r.Body).Decode(&certificate)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}
	tokenname := DecodeGetName(os.Getenv(publickeykatalogfilm), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)
	tokennomor := DecodeGetNomor(os.Getenv(publickeykatalogfilm), header)
	if tokenname == "" || tokenusername == "" || tokenrole == "" || tokennomor == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "user" && tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}
	certificate.Nomor = tokennomor
	InsertCertificate(mconn, collname, certificate)
	response.Status = true
	response.Message = "Berhasil input data"

	return ReturnStruct(response)
}

func DeleteCertificate(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var certificate Certificate
	err := json.NewDecoder(r.Body).Decode(&certificate)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}
	tokenname := DecodeGetName(os.Getenv(publickeykatalogfilm), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)
	tokennomor := DecodeGetNomor(os.Getenv(publickeykatalogfilm), header)
	if tokenname == "" || tokenusername == "" || tokenrole == "" || tokennomor == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}
	DeletedCertificate(mconn, collname, certificate)
	response.Status = true
	response.Message = "Berhasil hapus data"

	return ReturnStruct(response)
}

func AddedBlog(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var blog Blog
	err := json.NewDecoder(r.Body).Decode(&blog)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}
	tokenname := DecodeGetName(os.Getenv(publickeykatalogfilm), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)
	tokennomor := DecodeGetNomor(os.Getenv(publickeykatalogfilm), header)
	if tokenname == "" || tokenusername == "" || tokenrole == "" || tokennomor == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "user" && tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}
	InsertBlog(mconn, collname, blog)
	response.Status = true
	response.Message = "Berhasil input data"

	return ReturnStruct(response)
}

func FindBlog(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	datablog := FindallBlog(mconn, collname)
	var response Pesan
	response.Status = true
	response.Message = "Berhasil ambil data"
	response.Data = datablog
	return ReturnStruct(response)
}

func UpdateBlog(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var blog Blog
	err := json.NewDecoder(r.Body).Decode(&blog)

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

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}

	UpdatedBlog(mconn, collname, blog)

	response.Status = true
	response.Message = "Berhasil update " + blog.Title + " dari database"
	return ReturnStruct(response)
}

func DeletedBlog(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var blog Blog
	err := json.NewDecoder(r.Body).Decode(&blog)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}
	tokenname := DecodeGetName(os.Getenv(publickeykatalogfilm), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)
	tokennomor := DecodeGetNomor(os.Getenv(publickeykatalogfilm), header)
	if tokenname == "" || tokenusername == "" || tokenrole == "" || tokennomor == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}
	DeleteBlog(mconn, collname, blog)
	response.Status = true
	response.Message = "Berhasil hapus data"
	return ReturnStruct(response)
}

// commment
func AddedComment(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}
	tokenname := DecodeGetName(os.Getenv(publickeykatalogfilm), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)
	tokennomor := DecodeGetNomor(os.Getenv(publickeykatalogfilm), header)
	if tokenname == "" || tokenusername == "" || tokenrole == "" || tokennomor == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "user" && tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}
	InsertComment(mconn, collname, comment)
	response.Status = true
	response.Message = "Berhasil input data"

	return ReturnStruct(response)
}

func FindComment(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	datafilm := FindallComment(mconn, collname)
	var response Pesan
	response.Status = true
	response.Message = "Berhasil ambil data"
	response.Data = datafilm
	return ReturnStruct(response)
}

func UpdateComment(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)

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

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}

	UpdatedComment(mconn, collname, comment)

	response.Status = true
	response.Message = "Berhasil update " + comment.Answer + " dari database"
	return ReturnStruct(response)
}

func DeletedComment(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}
	tokenname := DecodeGetName(os.Getenv(publickeykatalogfilm), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)
	tokennomor := DecodeGetNomor(os.Getenv(publickeykatalogfilm), header)
	if tokenname == "" || tokenusername == "" || tokenrole == "" || tokennomor == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}
	DeleteComment(mconn, collname, comment)
	response.Status = true
	response.Message = "Berhasil hapus data"
	return ReturnStruct(response)
}

// question

func AddedQuestionAndAnswer(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var qas []QuestionAndAnswer
	err := json.NewDecoder(r.Body).Decode(&qas)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}
	tokenname := DecodeGetName(os.Getenv(publickeykatalogfilm), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)
	tokennomor := DecodeGetNomor(os.Getenv(publickeykatalogfilm), header)
	log.Println("tokenname", tokenname)
	if tokenname == "" || tokenusername == "" || tokenrole == "" || tokennomor == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "user" && tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}

	for _, qa := range qas {
		// Validate each answer
		if !checkAnswer(qa.Answers, qa.CorrectAnswer) {
			response.Message = "Jawaban salah"
			return ReturnStruct(response)
		}
	}

	// If all answers are correct, insert all questions and answers into the database
	for _, qa := range qas {
		InsertQuestionAndAnswer(mconn, collname, qa)
	}
	// author ambil dari token name
	response.Status = true
	response.Message = "Berhasil input data"

	return ReturnStruct(response)
}

func checkAnswer(answers []string, correctAnswer string) bool {
	for _, answer := range answers {
		if answer == correctAnswer {
			return true
		}
	}
	return false
}

func GetQuestionAndAnswer(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	dataqa := FindallQuestionAndAnswer(mconn, collname)
	var response Pesan
	response.Status = false

	// Check each question's answer
	// Check each question's answer
	allIncorrect := true
	for _, qa := range dataqa {
		// Validate each answer
		if checkAnswer(qa.Answers, qa.CorrectAnswer) {
			allIncorrect = false
			break
		}
	}

	// If all answers are incorrect, set the response accordingly
	if allIncorrect {
		response.Status = true
		response.Message = "Berhasil ambil data"
		response.Data = "Semua jawaban salah"
	} else {
		// Check the correctness of answers only if there is at least one correct answer
		// Check each question's answer
		for _, qa := range dataqa {
			// Validate if the correct answer is present in the answers slice
			if checkAnswer(qa.Answers, qa.CorrectAnswer) {
				response.Status = true
				response.Message = "Berhasil ambil data"
				response.Data = "Benar"
				return ReturnStruct(response)
			}
		}

		// If no correct answer is found, return a message indicating that the question is incorrect
		response.Message = "Soal salah"
	}

	return ReturnStruct(response)
}

func checkAnswers(answers []string, correctAnswers []string) bool {
	// Check if all correct answers are present in the provided answers
	for _, correctAnswer := range correctAnswers {
		found := false
		for _, answer := range answers {
			if answer == correctAnswer {
				found = true
				break
			}
		}
		// If any correct answer is not found, return false
		if !found {
			return false
		}
	}
	return true
}

func CekAnswer(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var data []QuestionAndAnswer
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	// Iterate over each question and check its answers
	for _, userQuestion := range data {
		// Retrieve the corresponding question and its correct answers from the database
		question := CheckAnswerdb(mconn, collname, userQuestion)

		// Check if any of the user's answers match the correct answer for this question
		correct := false
		for _, userAnswer := range userQuestion.Answers {
			if userAnswer == string(question.CorrectAnswer) {
				correct = true
				break
			}
		}

		// If none of the user's answers match the correct answer, set response accordingly and return
		if !correct {
			response.Message = "Jawaban salah"
			return ReturnStruct(response)
		}
	}

	// If at least one user answer matched a correct answer for each question, set response accordingly
	response.Status = true
	response.Message = "Jawaban benar"
	return ReturnStruct(response)
}

func UpdatedAnswer(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var user []QuestionAndAnswer
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

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}

	for _, qa := range user {
		UpdatedAnswerdb(mconn, collname, qa)
	}

	response.Status = true
	response.Message = "Berhasil update " + " dari database"
	return ReturnStruct(response)
}

// find answerd and video
func AddedQuestionAnswerAndVidieo(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var qas []VidioQuestion
	err := json.NewDecoder(r.Body).Decode(&qas)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	header := r.Header.Get("token")
	if header == "" {
		response.Message = "Header login tidak ditemukan"
		return ReturnStruct(response)
	}
	tokenname := DecodeGetName(os.Getenv(publickeykatalogfilm), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)
	tokennomor := DecodeGetNomor(os.Getenv(publickeykatalogfilm), header)
	log.Println("tokenname", tokenname)
	if tokenname == "" || tokenusername == "" || tokenrole == "" || tokennomor == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "user" && tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}

	for _, qa := range qas {
		// Validate each answer
		if !checkAnswer(qa.Answers, qa.CorrectAnswer) {
			response.Message = "Jawaban salah"
			return ReturnStruct(response)
		}
	}

	// If all answers are correct, insert all questions and answers into the database
	for _, qa := range qas {
		InsertQuestionAndAnswerVidieo(mconn, collname, qa)
	}
	// author ambil dari token name
	response.Status = true
	response.Message = "Berhasil input data"

	return ReturnStruct(response)
}
func GetFindAll(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	datafilm := FindallQuestionAndAnswer(mconn, collname)
	var response Pesan
	response.Status = true
	response.Message = "Berhasil ambil data"
	response.Data = datafilm
	return ReturnStruct(response)
}

func DeleteAnswer(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var user QuestionAndAnswer
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
	tokenname := DecodeGetName(os.Getenv(publickeykatalogfilm), header)
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)
	tokenrole := DecodeGetRole(os.Getenv(publickeykatalogfilm), header)
	tokennomor := DecodeGetNomor(os.Getenv(publickeykatalogfilm), header)
	if tokenname == "" || tokenusername == "" || tokenrole == "" || tokennomor == "" {
		response.Message = "Hasil decode tidak ditemukan"
		return ReturnStruct(response)
	}

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
		response.Message = "Akun tidak ditemukan"
		return ReturnStruct(response)
	}

	if tokenrole != "admin" {
		response.Message = "Anda tidak memiliki akses"
		return ReturnStruct(response)
	}
	DeleteAnswerdb(mconn, collname, user)
	response.Status = true
	response.Message = "Berhasil hapus data"
	return ReturnStruct(response)
}

func GrantAccess(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false

	// Establish MongoDB connection
	mconn := SetConnection(mongoenvkatalogfilm, dbname)

	// Parse JSON request body to AccessControl struct
	var access AccessControl
	err := json.NewDecoder(r.Body).Decode(&access)
	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	// Ensure ContentID array is not empty
	if len(access.ContentID) == 0 {
		response.Message = "No content IDs provided"
		return ReturnStruct(response)
	}

	// Insert access control entries into the database
	for _, contentID := range access.ContentID {
		entry := AccessControl{
			Username:  access.Username,
			ContentID: []int{contentID}, // Single content ID in array
			HasAccess: access.HasAccess,
		}

		err = InsertAccessControl(mconn, collname, entry)
		if err != nil {
			response.Message = "Error inserting access control: " + err.Error()
			return ReturnStruct(response)
		}
	}

	response.Status = true
	response.Message = "Akses diberikan"
	return ReturnStruct(response)
}

func GetVideoWithAccessCheck(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false

	// Establish MongoDB connection
	mconn := SetConnection(mongoenvkatalogfilm, dbname)

	// Extract token from header and decode username
	header := r.Header.Get("token")
	tokenusername := DecodeGetUsername(os.Getenv(publickeykatalogfilm), header)

	if tokenusername == "" {
		response.Message = "Token username tidak ditemukan"
		return ReturnStruct(response)
	}

	// Extract content ID from request parameters
	contentIDStr := r.URL.Query().Get("content_id")
	contentID, err := strconv.Atoi(contentIDStr)
	if err != nil {
		response.Message = "Invalid content ID format"
		return ReturnStruct(response)
	}

	// Retrieve video data from the database
	videoData, err := FindVideoByID(mconn, collname, contentID)
	if err != nil {
		response.Message = "Error retrieving video data: " + err.Error()
		return ReturnStruct(response)
	}

	response.Status = true
	response.Message = "Berhasil mengambil data video"
	response.Data = videoData

	return ReturnStruct(response)
}
