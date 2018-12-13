package auth

import (
	"apptastic/dashboard/model"
	"apptastic/dashboard/session"
	"database/sql"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func MustAuthorize(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Auth check")

		sess := session.Instance(r)
		userID := sess.Values["ID"]

		if userID == nil {

			fmt.Println("Auth Error: 401 Unauthorized")

			fmt.Println("A", r.Header.Get("Content-Type"))

			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)

			return

			//todo redirect
		}

		next.ServeHTTP(w, r)

		fmt.Println("Auth Check Finished")
	})
}

//GetUser returns the authenticated user
func GetUser(r *http.Request, db *sql.DB) (*model.User, error) {

	sess := session.Instance(r)

	userID := sess.Values["ID"].(int)

	user, err := model.FindUserByID(db, userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login() {}

func Logout() {}
