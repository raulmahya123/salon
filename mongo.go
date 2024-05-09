package kursussalon

import (
	"context"
	"log"
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func FindPassword(mconn *mongo.Database, collname string, userdata User) User {
	filter := bson.M{"password": userdata.Password}
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

//forgot password

func ForgotPassword(mconn *mongo.Database, collname string, userdata User) interface{} {
	filter := bson.M{"password": userdata.Password}
	return atdb.ReplaceOneDoc(mconn, collname, filter, userdata)
}

// find all salon
func FindallSalon(mconn *mongo.Database, collname string) []Salon {
	salon := atdb.GetAllDoc[[]Salon](mconn, collname)
	return salon
}

// find all sertificate
func FindallCertificate(mconn *mongo.Database, collname string) []Certificate {
	certificate := atdb.GetAllDoc[[]Certificate](mconn, collname)
	return certificate
}

// insert salon

func InsertSalon(mconn *mongo.Database, collname string, datasalon Salon) interface{} {
	return atdb.InsertOneDoc(mconn, collname, datasalon)
}

// insert certificate

func InsertCertificate(mconn *mongo.Database, collname string, datacertificate Certificate) interface{} {
	return atdb.InsertOneDoc(mconn, collname, datacertificate)
}
func CreateResponse(status bool, message string, data interface{}) Response {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return response
}

func UpdatedSalon(mconn *mongo.Database, collname string, datasalon Salon) interface{} {
	filter := bson.M{"id": datasalon.ID}
	return atdb.ReplaceOneDoc(mconn, collname, filter, datasalon)
}

func DeletedSalon(mconn *mongo.Database, collname string, datasalon Salon) interface{} {
	filter := bson.M{"id": datasalon.ID}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

func InsertBlog(mconn *mongo.Database, collname string, datablog Blog) interface{} {
	return atdb.InsertOneDoc(mconn, collname, datablog)
}

func FindallBlog(mconn *mongo.Database, collname string) []Blog {
	blog := atdb.GetAllDoc[[]Blog](mconn, collname)
	return blog
}

func DeletedCertificate(mconn *mongo.Database, collname string, datacertificate Certificate) interface{} {
	filter := bson.M{"nomor": datacertificate.Nomor}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

func DeleteBlog(mconn *mongo.Database, collname string, datablog Blog) interface{} {
	filter := bson.M{"id": datablog.ID}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

func UpdatedBlog(mconn *mongo.Database, collname string, datablog Blog) interface{} {
	filter := bson.M{"id": datablog.ID}
	return atdb.ReplaceOneDoc(mconn, collname, filter, datablog)
}

func InsertQuestionAndAnswer(mconn *mongo.Database, collname string, dataquestion QuestionAndAnswer) interface{} {
	return atdb.InsertOneDoc(mconn, collname, dataquestion)
}

func CheckAnswerdb(mconn *mongo.Database, collname string, dataquestion QuestionAndAnswer) QuestionAndAnswer {
	filter := bson.M{"correct_answer": dataquestion.CorrectAnswer}
	return atdb.GetOneDoc[QuestionAndAnswer](mconn, collname, filter)
}

func FindallQuestionAndAnswer(mconn *mongo.Database, collname string) []QuestionAndAnswer {
	question := atdb.GetAllDoc[[]QuestionAndAnswer](mconn, collname)
	return question
}

func UpdatedAnswerdb(mconn *mongo.Database, collname string, dataquestion QuestionAndAnswer) interface{} {
	filter := bson.M{"id": dataquestion.ID}
	return atdb.ReplaceOneDoc(mconn, collname, filter, dataquestion)
}

func DeleteAnswerdb(mconn *mongo.Database, collname string, dataquestion QuestionAndAnswer) interface{} {
	filter := bson.M{"id": dataquestion.ID}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

func FindallHistorySalon(mconn *mongo.Database, collname string) []History {
	history := atdb.GetAllDoc[[]History](mconn, collname)
	return history
}

func ClaimsSalondb(mconn *mongo.Database, collname string, datahistory History) interface{} {
	return atdb.InsertOneDoc(mconn, collname, datahistory)
}

func GetAllClaimsSalon(mconn *mongo.Database, collname string) []History {
	history := atdb.GetAllDoc[[]History](mconn, collname)
	return history
}
func ClaimsExists(mconn *mongo.Database, collname string, salonName string) bool {
	filter := bson.M{"salon.name": salonName}
	var history History
	err := mconn.Collection(collname).FindOne(context.Background(), filter).Decode(&history)
	return err == nil
}

// Fungsi untuk mendapatkan data dari database
func RetrieveDataFromDatabase(mconn *mongo.Database, collname string) []History {
	// Tentukan filter kosong jika Anda ingin mengambil semua dokumen
	filter := bson.M{}

	// Tentukan opsi untuk kueri, misalnya batasan atau urutan
	options := options.Find()

	// Lakukan kueri ke koleksi yang diberikan dengan filter dan opsi yang ditentukan
	cursor, err := mconn.Collection(collname).Find(context.Background(), filter, options)
	if err != nil {
		// Handle kesalahan jika ada
		// Contoh: log pesan kesalahan atau kembalikan data kosong
		log.Println("Error querying database:", err)
		return []History{}
	}
	defer cursor.Close(context.Background())

	// Iterasi melalui hasil kueri
	var data []History
	for cursor.Next(context.Background()) {
		var history History
		if err := cursor.Decode(&history); err != nil {
			// Handle kesalahan jika ada saat mem-parsing hasil kueri
			// Contoh: log pesan kesalahan atau lanjutkan ke dokumen berikutnya
			log.Println("Error decoding document:", err)
			continue
		}
		// Tambahkan data ke slice data
		data = append(data, history)
	}

	if err := cursor.Err(); err != nil {
		// Handle kesalahan jika ada saat iterasi melalui hasil kueri
		// Contoh: log pesan kesalahan atau kembalikan data kosong
		log.Println("Error iterating over query results:", err)
		return []History{}
	}

	// Kembalikan data yang ditemukan
	return data
}

func InsertContent(mconn *mongo.Database, collname string, datacontent Content) interface{} {
	return atdb.InsertOneDoc(mconn, collname, datacontent)
}

func FindallContent(mconn *mongo.Database, collname string) []Content {
	content := atdb.GetAllDoc[[]Content](mconn, collname)
	return content
}

func UpdatedContent(mconn *mongo.Database, collname string, datacontent Content) interface{} {
	filter := bson.M{"id": datacontent.ID}
	return atdb.ReplaceOneDoc(mconn, collname, filter, datacontent)
}

func DeleteContent(mconn *mongo.Database, collname string, datacontent Content) interface{} {
	filter := bson.M{"id": datacontent.ID}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

func InsertComment(mconn *mongo.Database, collname string, datacomment Comment) interface{} {
	return atdb.InsertOneDoc(mconn, collname, datacomment)
}

func FindallComment(mconn *mongo.Database, collname string) []Comment {
	comment := atdb.GetAllDoc[[]Comment](mconn, collname)
	return comment
}

func UpdatedComment(mconn *mongo.Database, collname string, datacomment Comment) interface{} {
	filter := bson.M{"id": datacomment.ID}
	return atdb.ReplaceOneDoc(mconn, collname, filter, datacomment)
}

func DeleteComment(mconn *mongo.Database, collname string, datacomment Comment) interface{} {
	filter := bson.M{"id": datacomment.ID}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

func InsertProduct(mconn *mongo.Database, collname string, dataproduct Product) interface{} {
	return atdb.InsertOneDoc(mconn, collname, dataproduct)
}

func FindallProduct(mconn *mongo.Database, collname string) []Product {
	product := atdb.GetAllDoc[[]Product](mconn, collname)
	return product
}

func UpdatedProduct(mconn *mongo.Database, collname string, dataproduct Product) interface{} {
	filter := bson.M{"id": dataproduct.ID}
	return atdb.ReplaceOneDoc(mconn, collname, filter, dataproduct)
}

func DeleteProductt(mconn *mongo.Database, collname string, dataproduct Product) interface{} {
	filter := bson.M{"id": dataproduct.ID}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}
