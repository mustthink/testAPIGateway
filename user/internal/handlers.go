package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id       int
	Username string
	Email    string
	Dob      time.Time
	Age      int
	Number   string
}

type Service struct {
	errorLog *log.Logger
	DB       *sql.DB
	url      *string
	secret   *string
}

func NewService(errorLog *log.Logger, db *sql.DB, url *string, s *string) *Service {
	return &Service{
		errorLog: errorLog,
		DB:       db,
		url:      url,
		secret:   s,
	}
}

func (s *Service) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/user/profile", s.userInformation)
	mux.HandleFunc("/microservice/name", s.nameService)

	return mux
}

func fromNormal(dec, secret string) string {
	key := []byte(secret)

	h := hmac.New(sha256.New, key)
	h.Write([]byte(dec))

	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func (s *Service) userInformation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		s.errorLog.Println(w, http.StatusMethodNotAllowed)
		return
	}

	uname := r.Header.Get("Username")
	u := fromNormal(uname, *s.secret)

	us := &User{}
	stmt := `select * from users where Username = $1`
	err := s.DB.QueryRow(stmt, u).Scan(&us.Id, &us.Username, &us.Email, &us.Dob, &us.Age, &us.Number)
	if err != nil {
		s.errorLog.Println(err)
	}
	us.Username = uname

	jsondata, err := json.MarshalIndent(us, "", "    ")
	if err != nil {
		s.errorLog.Println(err)
	}

	_, err = w.Write(jsondata)
	if err != nil {
		s.errorLog.Println(err)
	}
}

func (s *Service) nameService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		s.errorLog.Println(w, http.StatusMethodNotAllowed)
		return
	}

	_, err := w.Write([]byte("User microservice"))
	if err != nil {
		s.errorLog.Println(err)
	}
}
