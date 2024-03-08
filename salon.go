package kursussalon

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

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
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

	if !UsernameExists(mongoenvkatalogfilm, dbname, User{Username: tokenusername}) {
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

	if !UsernameExists(mongoenvkatalogfilm, dbname, user) {
		response.Message = "Akun yang ingin dihapus tidak ditemukan"
		return ReturnStruct(response)
	}

	DeleteUser(mconn, collname, user)

	response.Status = true
	response.Message = "Berhasil hapus " + user.Username + " dari database"
	return ReturnStruct(response)
}

func UpdatePassword(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
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

	user.Nomor = auth.Nomor
	user.Name = user.Username
	user.Role = "user"
	HashPassword(user.Password)
	user.Password = hash
	EditUser(mconn, collname, user)

	response.Status = true
	response.Message = "Berhasil update password " + user.Username + " dari database"
	return ReturnStruct(response)

}

//salon

func FindSalon(mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	datafilm := FindallSalon(mconn, collname)
	var response Pesan
	response.Status = true
	response.Message = "Berhasil ambil data"
	response.Data = datafilm
	return ReturnStruct(response)
}

func AddedSalon(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var salon Salon
	err := json.NewDecoder(r.Body).Decode(&salon)

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

	if tokenname == "" || tokenusername == "" || tokenrole == "" {
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
	salon.Author = tokenusername
	InsertSalon(mconn, collname, salon)
	response.Status = true
	response.Message = "Berhasil input data"

	return ReturnStruct(response)
}

func UpdateSalonAPI(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var user Salon
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

	UpdatedSalon(mconn, collname, user)

	response.Status = true
	response.Message = "Berhasil update " + user.Name + " dari database"
	return ReturnStruct(response)
}

func DeleteSalon(publickeykatalogfilm, mongoenvkatalogfilm, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenvkatalogfilm, dbname)
	var salon Salon
	err := json.NewDecoder(r.Body).Decode(&salon)

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

	if tokenname == "" || tokenusername == "" || tokenrole == "" {
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
	DeletedSalon(mconn, collname, salon)
	response.Status = true
	response.Message = "Berhasil hapus data"

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
