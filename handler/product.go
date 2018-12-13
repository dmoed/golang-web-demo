package handler

import (
	"apptastic/dashboard/model"
	"apptastic/dashboard/view"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func ProductHandler(db *sql.DB, v *view.View) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var p = 1
		var l = 5
		var err error
		q := r.URL.Query()

		limit := q.Get("limit")
		page := q.Get("page")

		if page != "" {

			p, _ = strconv.Atoi(page)

			if err != nil {
				panic(err.Error())
			}
		}

		if limit != "" {

			l, _ = strconv.Atoi(limit)

			if err != nil {
				panic(err.Error())
			}
		}

		//Get all products
		products, err := model.GetAllByPage(db, p, l)

		if err != nil {
			panic(err.Error())
		}

		data, err := json.Marshal(map[string]interface{}{
			"payload": products,
		})

		if err != nil {
			panic(err.Error())
		}

		v.RenderJSON(w, data)
	})
}
