package handler

import (
	"apptastic/dashboard/model"
	"apptastic/dashboard/session"
	"database/sql"
	"encoding/json"
	"net/http"
)

func ProfileHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session := session.Instance(r)

		user, _ := model.FindUserByID(db, session.Values["ID"].(int))

		data, err := json.Marshal(user)

		if err != nil {
			panic(err.Error())
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
}
