package auth

import (
	"database/sql"
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
	u := r.Header.Get("Username")
	u = "Test5"
	us := User{}
	stmt := `select * from users where Username = $1`

	err := s.DB.QueryRow(stmt, u).Scan(&us.Id, &us.Username, &us.Email, &us.Dob, &us.Age, &us.Number)
	if err != nil {
		if err != sql.ErrNoRows {
			s.errorLog.Println(err)
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
	w.WriteHeader(http.StatusOK)
}
