package auth

import (
	"database/sql"
	"log"
	"net/http"
)

type Service struct {
	errorLog *log.Logger
	DB       *sql.DB
	url      *string
}

func NewService(errorLog *log.Logger, db *sql.DB, url *string) *Service {
	return &Service{
		errorLog: errorLog,
		DB:       db,
		url:      url,
	}
}

func (s *Service) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth", s.Verify)

	return mux
}

func (s *Service) Verify(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		s.errorLog.Println(w, http.StatusMethodNotAllowed)
		return
	}
	i := 0
	u := r.Header.Get("Username")
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
