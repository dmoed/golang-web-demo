package handler

import (
	"apptastic/dashboard/auth"
	"apptastic/dashboard/model"
	"apptastic/dashboard/session"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

func LoginHandler(db *sql.DB) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		formErrors := []string{}
		redirectTarget := "/dashboard/"
		username := ""

		sess := session.Instance(r)

		//check POST
		if r.Method == "POST" {

			username, plainPassword := r.FormValue("_username"), r.FormValue("_password")

			if user, err := model.FindUserByEmail(db, username); err != nil {

				fmt.Println("error", err)

				formErrors = append(formErrors, "Credentials invalid")

				panic("invalid creds")

			} else {

				fmt.Println("user found", user)

				if auth.CheckPasswordHash(plainPassword, user.PasswordHash) {

					fmt.Println("auth OK")

					sess.Values["ID"] = user.ID
					err := sessions.Save(r, w)

					if err != nil {
						fmt.Println(err.Error())
					}

					http.Redirect(w, r, redirectTarget, 302)
				}

				formErrors = append(formErrors, "Credentials invalid")
			}
		}

		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, map[string]interface{}{
			"errors":       formErrors,
			"lastUsername": username,
		})
	})
}
