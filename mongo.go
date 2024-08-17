package kursussalon

import (
	"context"
	"fmt"
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

// find all sertificate
func FindallCertificate(mconn *mongo.Database, collname string) []Certificate {
	certificate := atdb.GetAllDoc[[]Certificate](mconn, collname)
	return certificate
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
func InsertQuestionAndAnswerVidieo(mconn *mongo.Database, collname string, dataquestion VidioQuestion) interface{} {
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
func InsertAccessControl(mconn *mongo.Database, collname string, access AccessControl) error {
	collection := mconn.Collection(collname)

	// Create the access control record
	accessRecord := bson.M{
		"username":   access.Username,
		"content_id": access.ContentID,
		"has_access": access.HasAccess,
	}

	// Insert the record into the database
	_, err := collection.InsertOne(context.Background(), accessRecord)
	if err != nil {
		return fmt.Errorf("error inserting access control: %v", err)
	}

	return nil
}
func CheckUserAccess(mconn *mongo.Database, username string, contentID int) bool {
	collection := mconn.Collection("access_control")
	filter := bson.M{"username": username, "content_id": contentID}
	var result AccessControl
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false // No access control record found
		}
		// Handle other errors
		return false
	}
	return result.HasAccess
}
func FindVideoByID(mdb *mongo.Database, collname string, contentID int) (VidioQuestion, error) {
	collection := mdb.Collection(collname)

	// Create a filter to find the document where contentID is part of the ContentID array
	filter := bson.M{"content_id": bson.M{"$in": []int{contentID}}}

	// Find the document
	var video VidioQuestion
	err := collection.FindOne(context.Background(), filter).Decode(&video)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return VidioQuestion{}, fmt.Errorf("video not found")
		}
		return VidioQuestion{}, fmt.Errorf("error retrieving video data: %v", err)
	}

	return video, nil
}
