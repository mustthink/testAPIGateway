package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"
)

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

	mux.HandleFunc("/auth", s.Verify)

	return mux
}

func fromNormal(dec, secret string) string {
	key := []byte(secret)

	h := hmac.New(sha256.New, key)
	h.Write([]byte(dec))

	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func (s *Service) Verify(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		s.errorLog.Println(w, http.StatusMethodNotAllowed)
		return
	}
	i := 0
	u := r.Header.Get("Username")
	u = fromNormal(u, *s.secret)

	stmt := `select count(*) from users where Username = $1`
	err := s.DB.QueryRow(stmt, u).Scan(&i)
	if err != nil {
		s.errorLog.Println(err)
	}

	if i != 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
