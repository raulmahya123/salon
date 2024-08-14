package main

import (
	"fmt"
	"log"
	"net/http"

	kursussalon "github.com/raulmahya123/salon"
)

func AuthorizationAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, kursussalon.Authorization("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "user", r))
}

func FindProductAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, kursussalon.FindProduct("mongoenv", "petapedia", "product", r))
}

func RegistrasiAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, kursussalon.Registrasi("mongoenvkatalogfilm", "katalogfilm", "user", r))
}

func LoginAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, kursussalon.Login("privatekeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "user", r))
}

func AmbilSemuaUserAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, kursussalon.AmbilSemuaUser("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "user", r))
}

func UpdateUserAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, kursussalon.UpdateUser("mongoenvkatalogfilm", "katalogfilm", "user", r))
}

func HapusUserAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, kursussalon.HapusUser("mongoenvkatalogfilm", "katalogfilm", "user", r))
}

func UpdatedPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, kursussalon.UpdatePassword("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "user", r))
}

func CreateSalonAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, kursussalon.AddedSalon("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "salon", r))
}

func GetallSalonAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, kursussalon.FindSalon("mongoenvkatalogfilm", "katalogfilm", "salon", r))
}

func CreateCertificateAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, kursussalon.AddedCertificate("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "certificate", r))
}

func UpdateSalonAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
	}
	fmt.Fprintf(w, kursussalon.UpdateSalonAPI("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "salon", r))
}

func DeleteSalonAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
	}
	fmt.Fprintf(w, kursussalon.DeleteSalon("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "salon", r))
}

func GetallCertificateAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, kursussalon.FindCertificate("mongoenvkatalogfilm", "katalogfilm", "certificate", r))
}

func CreateBlogAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	fmt.Fprintf(w, kursussalon.AddedBlog("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "blog", r))
}

func FindallBlogAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
	}
	fmt.Fprintf(w, kursussalon.FindBlog("mongoenvkatalogfilm", "katalogfilm", "blog", r))
}
func DeleteCertificateAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
	}
	fmt.Fprintf(w, kursussalon.DeleteCertificate("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "certificate", r))
}

func DeleteBlogAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
	}
	fmt.Fprintf(w, kursussalon.DeletedBlog("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "blog", r))
}

func UpdateBlogAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
	}
	fmt.Fprintf(w, kursussalon.UpdateBlog("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "blog", r))
}

func AddedQuestionAndAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
	}
	fmt.Fprintf(w, kursussalon.AddedQuestionAndAnswer("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "questionandanswer", r))
}

func AddedAnswerQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
	}
	fmt.Fprintf(w, kursussalon.CekAnswer("mongoenvkatalogfilm", "katalogfilm", "questionandanswer", r))
}

func updatedAnsweQurstion(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
	}
	fmt.Fprintf(w, kursussalon.UpdatedAnswer("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "questionandanswer", r))
}

func GetallQuestionAndAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")

	}
	fmt.Fprintf(w, kursussalon.GetFindAll("mongoenvkatalogfilm", "katalogfilm", "questionandanswer", r))
}

func DeleteQuestionAndAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")

	}
	fmt.Fprintf(w, kursussalon.DeleteAnswer("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "questionandanswer", r))
}

func ClaimsSalon(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")

	}
	fmt.Fprintf(w, kursussalon.ClaimsSalon("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "claims", r))
}

func ClaimsCertificate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")

	}
	fmt.Fprintf(w, kursussalon.GetClaimsSalon("publickeykatalogfilm", "mongoenvkatalogfilm", "katalogfilm", "claims", r))
}

func handlerRequests() {
	http.HandleFunc("/AuthorizationAPI", AuthorizationAPI)
	http.HandleFunc("/RegistrasiAPI", RegistrasiAPI)
	http.HandleFunc("/LoginAPI", LoginAPI)
	http.HandleFunc("/AmbilSemuaUserAPI", AmbilSemuaUserAPI)
	http.HandleFunc("/UpdateUserAPI", UpdateUserAPI)
	http.HandleFunc("/HapusUserAPI", HapusUserAPI)
	http.HandleFunc("/UpdatedPassword", UpdatedPassword)
	http.HandleFunc("/CreateSalonAPI", CreateSalonAPI)
	http.HandleFunc("/GetallSalonAPI", GetallSalonAPI)
	http.HandleFunc("/UpdateSalonAPI", UpdateSalonAPI)
	http.HandleFunc("/DeleteSalonAPI", DeleteSalonAPI)
	http.HandleFunc("/CreateCertificateAPI", CreateCertificateAPI)
	http.HandleFunc("/GetallCertificateAPI", GetallCertificateAPI)
	http.HandleFunc("/DeleteCertificateAPI", DeleteCertificateAPI)
	http.HandleFunc("/CreateBlogAPI", CreateBlogAPI)
	http.HandleFunc("/FindallBlogAPI", FindallBlogAPI)
	http.HandleFunc("/DeleteBlogAPI", DeleteBlogAPI)
	http.HandleFunc("/UpdateBlogAPI", UpdateBlogAPI)
	http.HandleFunc("/AddedQuestionAndAnswer", AddedQuestionAndAnswer)
	http.HandleFunc("/GetQuestionAndAnswer", AddedAnswerQuestion)
	http.HandleFunc("/UpdatedAnswer", updatedAnsweQurstion)
	http.HandleFunc("/GetallQuestionAndAnswer", GetallQuestionAndAnswer)
	http.HandleFunc("/DeleteQuestionAndAnswer", DeleteQuestionAndAnswer)
	http.HandleFunc("/ClaimsSalon", ClaimsSalon)
	http.HandleFunc("/FindProductAPI", FindProductAPI)

	http.HandleFunc("/ClaimsCertificate", ClaimsCertificate)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handlerRequests()
}
