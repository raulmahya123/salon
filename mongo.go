package kursussalon

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
